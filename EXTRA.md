# Extending the Shim

## Boulder Shim+
In the example below we install a ca and see if we can get a softhsm to produce keys, and certificates. 

### Building Boulder with Patches
Git clone main, apply patches, make, and done. Make sure you set your GOROOT, GOBIN, and PATH variables like when we tested Go's libraries.
```bash 
@melinko2003 ➜ /workspaces/shim-go (main) $ cd src
@melinko2003 ➜ /workspaces/shim-go/src (main) $ git clone https://github.com/letsencrypt/boulder
Cloning into 'boulder'...
remote: Enumerating objects: 78773, done.
remote: Counting objects: 100% (122/122), done.
remote: Compressing objects: 100% (77/77), done.
remote: Total 78773 (delta 62), reused 58 (delta 45), pack-reused 78651 (from 2)
Receiving objects: 100% (78773/78773), 54.93 MiB | 27.69 MiB/s, done.
Resolving deltas: 100% (50802/50802), done.
@melinko2003 ➜ /workspaces/shim-go/src (main) $ cd boulder/
@melinko2003 ➜ /workspaces/shim-go/src/boulder (main) $ git apply ../../patches/boulder/p1.patch
@melinko2003 ➜ /workspaces/shim-go/src/boulder (main) $ make all
echo bin/admin bin/boulder bin/ceremony bin/ct-test-srv
bin/admin bin/boulder bin/ceremony bin/ct-test-srv
GOBIN=/workspaces/shim-go/shim/go-shim-039b3ebeba/bin GO111MODULE=on go install -mod=vendor -ldflags "-X \"github.com/letsencrypt/boulder/core.BuildID= +10d9ef9a\" -X \"github.com/letsencrypt/boulder/core.BuildTime=Sat Mar  1 06:36:58 UTC 2025\" -X \"github.com/letsencrypt/boulder/core.BuildHost=codespace@codespaces-c4a42a\"" ./...
@melinko2003 ➜ /workspaces/shim-go/src/boulder (main) $ cd ../..
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -help
Usage of shim/go-shim-039b3ebeba/bin/ceremony:
  -config string
        Path to ceremony configuration file
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony
2025/03/01 06:43:16 --config is required
```
### Install softhsm2
Install Softhsm2, pkcs11 tool, and additional support via package management
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ sudo apt install softhsm2 opensc opensc-pkcs11
Reading package lists... Done
Building dependency tree       
Reading state information... Done
< ... continues ... >
```
### Configure softhsm2
Configure softhsm2, and store HSM tokens in shim
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ cp contrib/softhsm2/softhsm2.conf shim/go-shim-039b3ebeba/
# Create a softhsm2 token directory in our shim
@melinko2003 ➜ /workspaces/shim-go (main) $ mkdir shim/go-shim-039b3ebeba/tokens
# Update values
@melinko2003 ➜ /workspaces/shim-go (main) $ vi shim/go-shim-039b3ebeba/softhsm2.conf
# Use that config
@melinko2003 ➜ /workspaces/shim-go (main) $ export SOFTHSM2_CONF="$(pwd)/shim/go-shim-039b3ebeba/softhsm2.conf"
```
### Configure/Add Slots 
Setup softhsm2 slots via `softhsm2-util`. I just use two slots for the demo, for rsa keys, and ecdsa keys. 
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ softhsm2-util --init-token --slot 0 --label "RSA Tokens"
# Pin 1234
=== SO PIN (4-255 characters) ===
Please enter SO PIN: ****
Please reenter SO PIN: ****
# Pin 1234
=== User PIN (4-255 characters) ===
Please enter user PIN: ****
Please reenter user PIN: ****
The token has been initialized and is reassigned to slot 118006441
@melinko2003 ➜ /workspaces/shim-go (main) $ softhsm2-util --init-token --slot 1 --label "EC Tokens"
# Pin 1234
=== SO PIN (4-255 characters) ===
Please enter SO PIN: ****
Please reenter SO PIN: ****
# Pin 1234
=== User PIN (4-255 characters) ===
Please enter user PIN: ****
Please reenter user PIN: ****
The token has been initialized and is reassigned to slot 828895493
```
### Getting Slot Id's and Verifying slot creation
To get Slot Id's and Verifying slot creation, we use `pkcs11-tool`
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so -L
Available slots:
Slot 0 (0x17d90540): SoftHSM slot ID 0x17d90540
  token label        : EC Tokens
  token manufacturer : SoftHSM project
  token model        : SoftHSM v2
  token flags        : login required, rng, token initialized, PIN initialized, other flags=0x20
  hardware version   : 2.5
  firmware version   : 2.5
  serial num         : b963596d97d90540
  pin min/max        : 4/255
Slot 1 (0x7f3f1524): SoftHSM slot ID 0x7f3f1524
  token label        : RSA Tokens
  token manufacturer : SoftHSM project
  token model        : SoftHSM v2
  token flags        : login required, rng, token initialized, PIN initialized, other flags=0x20
  hardware version   : 2.5
  firmware version   : 2.5
  serial num         : 5ac0ba017f3f1524
  pin min/max        : 4/255
Slot 2 (0x2): SoftHSM slot ID 0x2
  token state:   uninitialized
```
You'll see these two lines, and the hex at the end is the slot id ( ie. slot when using boulder with softhsm2 )
```bash
Slot 0 (0x17d90540): SoftHSM slot ID 0x17d90540
Slot 1 (0x7f3f1524): SoftHSM slot ID 0x7f3f1524
```
### Create Demo Keys in there Slots
Create a Demo RSA key
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so \
>             --slot 0x708a2a9 \
>             --login \
>             --pin 1234 \
>             --keypairgen \
>             --key-type rsa:4096 \
>             --id 01 \
>             --label "example_rsa_keypair"
Key pair generated:
Private Key Object; RSA 
  label:      example_rsa_keypair
  ID:         01
  Usage:      decrypt, sign, unwrap
  Access:     sensitive, always sensitive, never extractable, local
Public Key Object; RSA 4096 bits
  label:      example_rsa_keypair
  ID:         01
  Usage:      encrypt, verify, wrap
  Access:     local
```
Create a Demo ECDSA key
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so \
>             --slot 0x3167f105 \
>             --login \
>             --pin 1234 \
>             --keypairgen \
>             --key-type EC:secp384r1 \
>             --id 02 \
>             --label "example_ecdsa_keypair"
Key pair generated:
Private Key Object; EC
  label:      example_root_e1
  ID:         02
  Usage:      decrypt, sign, unwrap, derive
  Access:     sensitive, always sensitive, never extractable, local
Public Key Object; EC  EC_POINT 384 bits
  EC_POINT:   0461044f14fdc30573f6755cd5fbce7423c6460cb8412854dd1bbdfe3bb6ced99c0cf6d04f2bac13ce3639be1e5f6c313f0c8d8cb431d555a78cb2c6512aa71153a975e373b1f4f9f1d2be81023503197e804c5aa56893f03f74ff0e4908a3a34e489e
  EC_PARAMS:  06052b81040022
  label:      example_root_e1
  ID:         02
  Usage:      encrypt, verify, wrap, derive
  Access:     local
```
Create ECDSA Certificates:
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -config=contrib/boulder/example_root_e1.yaml 
2025/03/01 07:56:34 Preparing root ceremony for bldr_exmple_root_e1.crt.pem
2025/03/01 07:56:34 Opened PKCS#11 session for slot 828895493
2025/03/01 07:56:34 Generating ECDSA key with curve P-384 and ID 60832e7f
2025/03/01 07:56:34     Encoded curve parameters for P-384: 06052B81040022
2025/03/01 07:56:34 Key generated
2025/03/01 07:56:34 Extracting public key
2025/03/01 07:56:34     X: 051AD255279A791F8CC4666469C0E8B594CBE74D341C674A91AF6407BF37B088F41E13DA2E06B159D027E76E7C87F051
2025/03/01 07:56:34     Y: A3C57D4A83940DEDE2941BB6E85B1292783C3937687526BB846880067280D75EE3651820E526A71E4CA896F7BD46A974
2025/03/01 07:56:34 Extracted public key
2025/03/01 07:56:34 Public key PEM:
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEBRrSVSeaeR+MxGZkacDotZTL5000HGdK
ka9kB783sIj0HhPaLgaxWdAn5258h/BRo8V9SoOUDe3ilBu26FsSkng8OTdodSa7
hGiABnKA117jZRgg5SanHkyolve9Rql0
-----END PUBLIC KEY-----

2025/03/01 07:56:34 Public key written to "bldr_exmple_root_e1.key.pem"
2025/03/01 07:56:34 Signed certificate PEM:
-----BEGIN CERTIFICATE-----
MIIB9DCCAXqgAwIBAgIQMHJh1mxzAqkshuOcdCW25DALBglghkgBZQMEAwswOzEL
MAkGA1UEBhMCVVMxEjAQBgNVBAoTCWdvb2QgZ3V5czEYMBYGA1UEAxMPQ0EgaW50
ZXJtZWRpYXRlMB4XDTIwMDEwMTEyMDAwMFoXDTQwMDEwMTEyMDAwMFowOzELMAkG
A1UEBhMCVVMxEjAQBgNVBAoTCWdvb2QgZ3V5czEYMBYGA1UEAxMPQ0EgaW50ZXJt
ZWRpYXRlMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEBRrSVSeaeR+MxGZkacDotZTL
5000HGdKka9kB783sIj0HhPaLgaxWdAn5258h/BRo8V9SoOUDe3ilBu26FsSkng8
OTdodSa7hGiABnKA117jZRgg5SanHkyolve9Rql0o0IwQDAOBgNVHQ8BAf8EBAMC
AQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUaviBLDNhOs8YV+3zzbmGz8lw
SVwwCwYJYIZIAWUDBAMLA2cAMGQCMEA7jvTIXOBmhH02uh/ojKJox1EbaT/3ZTeA
kam1bLtDXvrukZ4PPovewAMcJQsa5wIwe4XYp1TKl3CaOPSUme3I6JbVAKXzcS0I
v5Jo/bPMqWN2/j5hvis4HPYdIQganZlc
-----END CERTIFICATE-----
2025/03/01 07:56:34 Certificate written to "bldr_exmple_root_e1.crt.pem"
2025/03/01 07:56:34 Post issuance linting completed for bldr_exmple_root_e1.crt.pem
@melinko2003 ➜ /workspaces/shim-go (main) $ openssl x509 -in bldr_exmple_root_e1.crt.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            30:72:61:d6:6c:73:02:a9:2c:86:e3:9c:74:25:b6:e4
        Signature Algorithm: ecdsa_with_SHA3-384
        Issuer: C = US, O = good guys, CN = CA intermediate
        Validity
            Not Before: Jan  1 12:00:00 2020 GMT
            Not After : Jan  1 12:00:00 2040 GMT
        Subject: C = US, O = good guys, CN = CA intermediate
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (384 bit)
                pub:
                    04:05:1a:d2:55:27:9a:79:1f:8c:c4:66:64:69:c0:
                    e8:b5:94:cb:e7:4d:34:1c:67:4a:91:af:64:07:bf:
                    37:b0:88:f4:1e:13:da:2e:06:b1:59:d0:27:e7:6e:
                    7c:87:f0:51:a3:c5:7d:4a:83:94:0d:ed:e2:94:1b:
                    b6:e8:5b:12:92:78:3c:39:37:68:75:26:bb:84:68:
                    80:06:72:80:d7:5e:e3:65:18:20:e5:26:a7:1e:4c:
                    a8:96:f7:bd:46:a9:74
                ASN1 OID: secp384r1
                NIST CURVE: P-384
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                6A:F8:81:2C:33:61:3A:CF:18:57:ED:F3:CD:B9:86:CF:C9:70:49:5C
    Signature Algorithm: ecdsa_with_SHA3-384
    Signature Value:
        30:64:02:30:40:3b:8e:f4:c8:5c:e0:66:84:7d:36:ba:1f:e8:
        8c:a2:68:c7:51:1b:69:3f:f7:65:37:80:91:a9:b5:6c:bb:43:
        5e:fa:ee:91:9e:0f:3e:8b:de:c0:03:1c:25:0b:1a:e7:02:30:
        7b:85:d8:a7:54:ca:97:70:9a:38:f4:94:99:ed:c8:e8:96:d5:
        00:a5:f3:71:2d:08:bf:92:68:fd:b3:cc:a9:63:76:fe:3e:61:
        be:2b:38:1c:f6:1d:21:08:1a:9d:99:5c
@melinko2003 ➜ /workspaces/shim-go (main) $ 
```
Create RSA Certificates:
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-0a0e6af39b/bin/ceremony -config=contrib/boulder/example_root_r1.yaml
2025/03/03 21:52:52 Preparing root ceremony for bldr_exmple_root_r1.crt.pem
2025/03/03 21:52:52 Opened PKCS#11 session for slot 2134840612
2025/03/03 21:52:52 Generating RSA key with 4096 bit modulus and public exponent 65537 and ID 2c2c096a
2025/03/03 21:52:52     Encoded public exponent (65537) as: 010001
2025/03/03 21:52:52 Key generated
2025/03/03 21:52:52 Extracting public key
2025/03/03 21:52:52     Public exponent: 65537
2025/03/03 21:52:52     Modulus: (4096 bits) B96973FD9D0109109A149E38A289CCCD0B80BA34E055D836D8586DCCA30C616C9B3EA96DEB7C0965147CD857408BAA72BBA503C5FEA549B7A64AF7DE84025C9AAF34E63DDEC35CF48866CBABC2CDD2B101A08A2A1AC59B15D8B9BB4D4C8AD583E23C29D7205638C4B8B826EEC387C83A5A4BCD1ADE0F27EE051CA9B1E78A53B6DD48188B281E71F3EF19C400E921B6B380966A23AB5B09B917238B649C6AFB525EAC0C43CE406E640972DD9EB1106BD85E7E2DC83A7B599A261AD36B51310C68C1014ECF1601C7E91E14301670CDB17E10DF3A6CC53F0BA29FD874B1B829981FE0DBABAECCFE3D6E8307673FB9341F8733DAC3FF2296635726D49058428D48F7F13EFA5A610CB1DC58E7E325CFDFDC33C0BC4281C37FB5B426D0C5CC347FC4E0FB64DB792CAB269AE3A4348C58CA98C753885245E8804D59E53D762B13BA96AF7A5EDF3BDCFFA088CCAD91080C387E59F7F78F0643BA2553354A43FC8B87F9E6F01389AED230450BC73F25A1E64AA52433265871398F857B05ADDE66516A854FD117C15E6ECD846942C6E9316474E75C4CAD706F8D8FF89BCBBD2E5614103081650330201972AD505FFD6F6774F2E947668C2D9EB8F24E66A27E6C1D8DDEBA89EB3666ED5D3555F08AFA23C7B8C32EFFBFED0378D77A64B3F7640DD2FD1895F67987475304460F838E34B84F188BD6D93F242832C1483878A2BCD92443B133A7
2025/03/03 21:52:52 Extracted public key
2025/03/03 21:52:52 Public key PEM:
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAuWlz/Z0BCRCaFJ44oonM
zQuAujTgVdg22FhtzKMMYWybPqlt63wJZRR82FdAi6pyu6UDxf6lSbemSvfehAJc
mq805j3ew1z0iGbLq8LN0rEBoIoqGsWbFdi5u01MitWD4jwp1yBWOMS4uCbuw4fI
OlpLzRreDyfuBRypseeKU7bdSBiLKB5x8+8ZxADpIbazgJZqI6tbCbkXI4tknGr7
Ul6sDEPOQG5kCXLdnrEQa9hefi3IOntZmiYa02tRMQxowQFOzxYBx+keFDAWcM2x
fhDfOmzFPwuin9h0sbgpmB/g26uuzP49boMHZz+5NB+HM9rD/yKWY1cm1JBYQo1I
9/E++lphDLHcWOfjJc/f3DPAvEKBw3+1tCbQxcw0f8Tg+2TbeSyrJprjpDSMWMqY
x1OIUkXogE1Z5T12KxO6lq96Xt873P+giMytkQgMOH5Z9/ePBkO6JVM1SkP8i4f5
5vATia7SMEULxz8loeZKpSQzJlhxOY+FewWt3mZRaoVP0RfBXm7NhGlCxukxZHTn
XEytcG+Nj/iby70uVhQQMIFlAzAgGXKtUF/9b2d08ulHZowtnrjyTmaifmwdjd66
ies2Zu1dNVXwivojx7jDLv+/7QN413pks/dkDdL9GJX2eYdHUwRGD4OONLhPGIvW
2T8kKDLBSDh4orzZJEOxM6cCAwEAAQ==
-----END PUBLIC KEY-----

2025/03/03 21:52:52 Public key written to "bldr_exmple_root_r1.key.pem"
2025/03/03 21:52:53 Signed certificate PEM:
-----BEGIN CERTIFICATE-----
MIIFQzCCAyugAwIBAgIRAN7Jdy76Mg1+tLwpIbEdeTAwDQYJYIZIAWUDBAMOBQAw
OzELMAkGA1UEBhMCVVMxEjAQBgNVBAoTCWdvb2QgZ3V5czEYMBYGA1UEAxMPQ0Eg
aW50ZXJtZWRpYXRlMB4XDTIwMDEwMTEyMDAwMFoXDTQwMDEwMTEyMDAwMFowOzEL
MAkGA1UEBhMCVVMxEjAQBgNVBAoTCWdvb2QgZ3V5czEYMBYGA1UEAxMPQ0EgaW50
ZXJtZWRpYXRlMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAuWlz/Z0B
CRCaFJ44oonMzQuAujTgVdg22FhtzKMMYWybPqlt63wJZRR82FdAi6pyu6UDxf6l
SbemSvfehAJcmq805j3ew1z0iGbLq8LN0rEBoIoqGsWbFdi5u01MitWD4jwp1yBW
OMS4uCbuw4fIOlpLzRreDyfuBRypseeKU7bdSBiLKB5x8+8ZxADpIbazgJZqI6tb
CbkXI4tknGr7Ul6sDEPOQG5kCXLdnrEQa9hefi3IOntZmiYa02tRMQxowQFOzxYB
x+keFDAWcM2xfhDfOmzFPwuin9h0sbgpmB/g26uuzP49boMHZz+5NB+HM9rD/yKW
Y1cm1JBYQo1I9/E++lphDLHcWOfjJc/f3DPAvEKBw3+1tCbQxcw0f8Tg+2TbeSyr
JprjpDSMWMqYx1OIUkXogE1Z5T12KxO6lq96Xt873P+giMytkQgMOH5Z9/ePBkO6
JVM1SkP8i4f55vATia7SMEULxz8loeZKpSQzJlhxOY+FewWt3mZRaoVP0RfBXm7N
hGlCxukxZHTnXEytcG+Nj/iby70uVhQQMIFlAzAgGXKtUF/9b2d08ulHZowtnrjy
Tmaifmwdjd66ies2Zu1dNVXwivojx7jDLv+/7QN413pks/dkDdL9GJX2eYdHUwRG
D4OONLhPGIvW2T8kKDLBSDh4orzZJEOxM6cCAwEAAaNCMEAwDgYDVR0PAQH/BAQD
AgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFHCF9CqxDx/cphH3zgQ+2L69
oEJGMA0GCWCGSAFlAwQDDgUAA4ICAQBFjqtJq3SXEcyV1uaQkdS1na5TgewGWew5
VpuV6vQ8WijkRTAjmNZpZ0B3hl9gD6J9p/25/vS73cP6B6agyUd3ukj/Ih8c2eUp
TBtwyns+FOX3D4w/ixSpToyKywA4Tq9XLY3I2V4pVNyxFmFFEqx/A8W7pK99t0PH
pzd113ltQ4KkW+eqAZ5wlZ3FP/YH8FRtyv6JRAq+F4+To35nmdeC+uL49vVYwo1+
PWWZlY31bTGTjGZ6x4iQceWVIopJBwXVJwpRCuGQOeZvjPdEG7plXoP5g0HoJeTL
erWhFeP340VWLE567M3C4hxtCIIhFJ4nsU59QeVvwVLrx0FO1iIn8DH34TdeVHo6
XkwKjB30Zccb1f/ljBKcuqMm+dpLma0hy20lkpY++rbrmsJNLZpKJI/4+A2gwP0C
AZI5egUmP/RU5AYt6/5j1pK86+NfwSj6zA9oidEvk+cTEly0W0ovcMhJ5HDoBEIF
JGVAcQqiKej+mqzA2fn3M4UacTYHwMymNm+XKzRTlD/wPSnfCfcF20HcoNnkzuLg
+1aYvOOcmUAwOFiPdpxrjY2cEMW3DWuW7FTlPa3jE+iisbZeJL78Pvy7wzNLIr5t
QZFal8pTcRWEiv942SO/4h9xGPTNhvq9APTHy5/CPXn0rZ3f02F1POPxPQoaZSsN
LIE4kUDj9Q==
-----END CERTIFICATE-----
2025/03/03 21:52:53 Certificate written to "bldr_exmple_root_r1.crt.pem"
2025/03/03 21:52:53 Post issuance linting completed for bldr_exmple_root_r1.crt.pem
@melinko2003 ➜ /workspaces/shim-go (main) ➜ $ openssl x509 -in bldr_exmple_root_r1.crt.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            de:c9:77:2e:fa:32:0d:7e:b4:bc:29:21:b1:1d:79:30
        Signature Algorithm: RSA-SHA3-256
        Issuer: C = US, O = good guys, CN = CA intermediate
        Validity
            Not Before: Jan  1 12:00:00 2020 GMT
            Not After : Jan  1 12:00:00 2040 GMT
        Subject: C = US, O = good guys, CN = CA intermediate
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (4096 bit)
                Modulus:
                    00:b9:69:73:fd:9d:01:09:10:9a:14:9e:38:a2:89:
                    cc:cd:0b:80:ba:34:e0:55:d8:36:d8:58:6d:cc:a3:
                    0c:61:6c:9b:3e:a9:6d:eb:7c:09:65:14:7c:d8:57:
                    40:8b:aa:72:bb:a5:03:c5:fe:a5:49:b7:a6:4a:f7:
                    de:84:02:5c:9a:af:34:e6:3d:de:c3:5c:f4:88:66:
                    cb:ab:c2:cd:d2:b1:01:a0:8a:2a:1a:c5:9b:15:d8:
                    b9:bb:4d:4c:8a:d5:83:e2:3c:29:d7:20:56:38:c4:
                    b8:b8:26:ee:c3:87:c8:3a:5a:4b:cd:1a:de:0f:27:
                    ee:05:1c:a9:b1:e7:8a:53:b6:dd:48:18:8b:28:1e:
                    71:f3:ef:19:c4:00:e9:21:b6:b3:80:96:6a:23:ab:
                    5b:09:b9:17:23:8b:64:9c:6a:fb:52:5e:ac:0c:43:
                    ce:40:6e:64:09:72:dd:9e:b1:10:6b:d8:5e:7e:2d:
                    c8:3a:7b:59:9a:26:1a:d3:6b:51:31:0c:68:c1:01:
                    4e:cf:16:01:c7:e9:1e:14:30:16:70:cd:b1:7e:10:
                    df:3a:6c:c5:3f:0b:a2:9f:d8:74:b1:b8:29:98:1f:
                    e0:db:ab:ae:cc:fe:3d:6e:83:07:67:3f:b9:34:1f:
                    87:33:da:c3:ff:22:96:63:57:26:d4:90:58:42:8d:
                    48:f7:f1:3e:fa:5a:61:0c:b1:dc:58:e7:e3:25:cf:
                    df:dc:33:c0:bc:42:81:c3:7f:b5:b4:26:d0:c5:cc:
                    34:7f:c4:e0:fb:64:db:79:2c:ab:26:9a:e3:a4:34:
                    8c:58:ca:98:c7:53:88:52:45:e8:80:4d:59:e5:3d:
                    76:2b:13:ba:96:af:7a:5e:df:3b:dc:ff:a0:88:cc:
                    ad:91:08:0c:38:7e:59:f7:f7:8f:06:43:ba:25:53:
                    35:4a:43:fc:8b:87:f9:e6:f0:13:89:ae:d2:30:45:
                    0b:c7:3f:25:a1:e6:4a:a5:24:33:26:58:71:39:8f:
                    85:7b:05:ad:de:66:51:6a:85:4f:d1:17:c1:5e:6e:
                    cd:84:69:42:c6:e9:31:64:74:e7:5c:4c:ad:70:6f:
                    8d:8f:f8:9b:cb:bd:2e:56:14:10:30:81:65:03:30:
                    20:19:72:ad:50:5f:fd:6f:67:74:f2:e9:47:66:8c:
                    2d:9e:b8:f2:4e:66:a2:7e:6c:1d:8d:de:ba:89:eb:
                    36:66:ed:5d:35:55:f0:8a:fa:23:c7:b8:c3:2e:ff:
                    bf:ed:03:78:d7:7a:64:b3:f7:64:0d:d2:fd:18:95:
                    f6:79:87:47:53:04:46:0f:83:8e:34:b8:4f:18:8b:
                    d6:d9:3f:24:28:32:c1:48:38:78:a2:bc:d9:24:43:
                    b1:33:a7
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                70:85:F4:2A:B1:0F:1F:DC:A6:11:F7:CE:04:3E:D8:BE:BD:A0:42:46
    Signature Algorithm: RSA-SHA3-256
    Signature Value:
        45:8e:ab:49:ab:74:97:11:cc:95:d6:e6:90:91:d4:b5:9d:ae:
        53:81:ec:06:59:ec:39:56:9b:95:ea:f4:3c:5a:28:e4:45:30:
        23:98:d6:69:67:40:77:86:5f:60:0f:a2:7d:a7:fd:b9:fe:f4:
        bb:dd:c3:fa:07:a6:a0:c9:47:77:ba:48:ff:22:1f:1c:d9:e5:
        29:4c:1b:70:ca:7b:3e:14:e5:f7:0f:8c:3f:8b:14:a9:4e:8c:
        8a:cb:00:38:4e:af:57:2d:8d:c8:d9:5e:29:54:dc:b1:16:61:
        45:12:ac:7f:03:c5:bb:a4:af:7d:b7:43:c7:a7:37:75:d7:79:
        6d:43:82:a4:5b:e7:aa:01:9e:70:95:9d:c5:3f:f6:07:f0:54:
        6d:ca:fe:89:44:0a:be:17:8f:93:a3:7e:67:99:d7:82:fa:e2:
        f8:f6:f5:58:c2:8d:7e:3d:65:99:95:8d:f5:6d:31:93:8c:66:
        7a:c7:88:90:71:e5:95:22:8a:49:07:05:d5:27:0a:51:0a:e1:
        90:39:e6:6f:8c:f7:44:1b:ba:65:5e:83:f9:83:41:e8:25:e4:
        cb:7a:b5:a1:15:e3:f7:e3:45:56:2c:4e:7a:ec:cd:c2:e2:1c:
        6d:08:82:21:14:9e:27:b1:4e:7d:41:e5:6f:c1:52:eb:c7:41:
        4e:d6:22:27:f0:31:f7:e1:37:5e:54:7a:3a:5e:4c:0a:8c:1d:
        f4:65:c7:1b:d5:ff:e5:8c:12:9c:ba:a3:26:f9:da:4b:99:ad:
        21:cb:6d:25:92:96:3e:fa:b6:eb:9a:c2:4d:2d:9a:4a:24:8f:
        f8:f8:0d:a0:c0:fd:02:01:92:39:7a:05:26:3f:f4:54:e4:06:
        2d:eb:fe:63:d6:92:bc:eb:e3:5f:c1:28:fa:cc:0f:68:89:d1:
        2f:93:e7:13:12:5c:b4:5b:4a:2f:70:c8:49:e4:70:e8:04:42:
        05:24:65:40:71:0a:a2:29:e8:fe:9a:ac:c0:d9:f9:f7:33:85:
        1a:71:36:07:c0:cc:a6:36:6f:97:2b:34:53:94:3f:f0:3d:29:
        df:09:f7:05:db:41:dc:a0:d9:e4:ce:e2:e0:fb:56:98:bc:e3:
        9c:99:40:30:38:58:8f:76:9c:6b:8d:8d:9c:10:c5:b7:0d:6b:
        96:ec:54:e5:3d:ad:e3:13:e8:a2:b1:b6:5e:24:be:fc:3e:fc:
        bb:c3:33:4b:22:be:6d:41:91:5a:97:ca:53:71:15:84:8a:ff:
        78:d9:23:bf:e2:1f:71:18:f4:cd:86:fa:bd:00:f4:c7:cb:9f:
        c2:3d:79:f4:ad:9d:df:d3:61:75:3c:e3:f1:3d:0a:1a:65:2b:
        0d:2c:81:38:91:40:e3:f5