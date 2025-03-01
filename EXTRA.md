# Extending the Shim
In the example below we install a ca and see if we can get a softhsm to produce keys, and certificates.
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
Install Softhsm2
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ sudo apt install softhsm2 opensc opensc-pkcs11
Reading package lists... Done
Building dependency tree       
Reading state information... Done
< ... continues ... >
```
Config softhsm2
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ cp contrib/softhsm2/softhsm2.conf shim/go-shim-039b3ebeba/
# Create a softhsm2 token directory in our shim
@melinko2003 ➜ /workspaces/shim-go (main) $ mkdir shim/go-shim-039b3ebeba/tokens
# Update values
@melinko2003 ➜ /workspaces/shim-go (main) $ vi shim/go-shim-039b3ebeba/softhsm2.conf
# Use that config
@melinko2003 ➜ /workspaces/shim-go (main) $ export SOFTHSM2_CONF="$(pwd)/shim/go-shim-039b3ebeba/softhsm2.conf"
```
Setup softhsm2
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ softhsm2-util --init-token --slot 0 --label "Init Token"
# Pin 1234
=== SO PIN (4-255 characters) ===
Please enter SO PIN: ****
Please reenter SO PIN: ****
# Pin 1234
=== User PIN (4-255 characters) ===
Please enter user PIN: ****
Please reenter user PIN: ****
The token has been initialized and is reassigned to slot 1532344672
@melinko2003 ➜ /workspaces/shim-go (main) $ softhsm2-util --init-token --slot 1 --label "RSA Tokens"
# Pin 1234
=== SO PIN (4-255 characters) ===
Please enter SO PIN: ****
Please reenter SO PIN: ****
# Pin 1234
=== User PIN (4-255 characters) ===
Please enter user PIN: ****
Please reenter user PIN: ****
The token has been initialized and is reassigned to slot 118006441
@melinko2003 ➜ /workspaces/shim-go (main) $ softhsm2-util --init-token --slot 2 --label "EC Tokens"
# Pin 1234
=== SO PIN (4-255 characters) ===
Please enter SO PIN: ****
Please reenter SO PIN: ****
# Pin 1234
=== User PIN (4-255 characters) ===
Please enter user PIN: ****
Please reenter user PIN: ****
The token has been initialized and is reassigned to slot 828895493
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so -L
Available slots:
Slot 0 (0x708a2a9): SoftHSM slot ID 0x708a2a9
  token label        : RSA Tokens
  token manufacturer : SoftHSM project
  token model        : SoftHSM v2
  token flags        : login required, rng, token initialized, PIN initialized, other flags=0x20
  hardware version   : 2.5
  firmware version   : 2.5
  serial num         : 78e289518708a2a9
  pin min/max        : 4/255
Slot 1 (0x3167f105): SoftHSM slot ID 0x3167f105
  token label        : EC Tokens
  token manufacturer : SoftHSM project
  token model        : SoftHSM v2
  token flags        : login required, rng, token initialized, PIN initialized, other flags=0x20
  hardware version   : 2.5
  firmware version   : 2.5
  serial num         : aa0416003167f105
  pin min/max        : 4/255
Slot 2 (0x5b55b960): SoftHSM slot ID 0x5b55b960
  token label        : Init Token
  token manufacturer : SoftHSM project
  token model        : SoftHSM v2
  token flags        : login required, rng, token initialized, PIN initialized, other flags=0x20
  hardware version   : 2.5
  firmware version   : 2.5
  serial num         : 4aa44023db55b960
  pin min/max        : 4/255
Slot 3 (0x3): SoftHSM slot ID 0x3
  token state:   uninitialized
```
Place RSA Keys into Slot 0
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so \
>             --slot 0x708a2a9 \
>             --login \
>             --pin 1234 \
>             --keypairgen \
>             --key-type rsa:4096 \
>             --id 01 \
>             --label "example_root_r1"
Key pair generated:
Private Key Object; RSA 
  label:      example_root_r1
  ID:         01
  Usage:      decrypt, sign, unwrap
  Access:     sensitive, always sensitive, never extractable, local
Public Key Object; RSA 4096 bits
  label:      example_root_r1
  ID:         01
  Usage:      encrypt, verify, wrap
  Access:     local
```
Place EC Keys into Slot 1
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ pkcs11-tool --module /usr/lib/softhsm/libsofthsm2.so \
>             --slot 0x3167f105 \
>             --login \
>             --pin 1234 \
>             --keypairgen \
>             --key-type EC:secp384r1 \
>             --id 02 \
>             --label "example_root_e1"
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
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -config=contrib/boulder/example_root_r1.yaml 
2025/03/01 08:03:13 Preparing root ceremony for bldr_exmple_root_r1.crt.pem
2025/03/01 08:03:13 Opened PKCS#11 session for slot 118006441
2025/03/01 08:03:13 Generating RSA key with 4096 bit modulus and public exponent 65537 and ID 3aacc0cd
2025/03/01 08:03:13     Encoded public exponent (65537) as: 010001
2025/03/01 08:03:13 Key generated
2025/03/01 08:03:13 Extracting public key
2025/03/01 08:03:13     Public exponent: 65537
2025/03/01 08:03:13     Modulus: (4096 bits) B2A08F82BADAF73EC031C37372E872778400D7D4D4E48CCFE7ADF70533B2655FF401091DDEA5D81ACF9FC65809C7C485195AD8F351AC25D4ECD8B51E863FD3F1431B61A93C92C679DA30D04C396B02DEB3857B28D20AA6F04B258D028FDEFD623B38A9FABEC2EC3047EF87AE76A5DAD38DA82859598428C87BAE481C0262589ACBDE0C0450E9D8774E7F35B97AAF8F93196D51EC0FDDD68903C8134E4B28933E75F7B7F6D8DF665329EAAB84E086A4CC323631E843C7943629E5BBD8E30891267B9BCF15C9810A1B8AF8461C9148395D0A86F37B57AB3E63669304920510DBE77D80DA2FDF0774CB52A53D7FE068A130EBBD1EC203F80C05C481F77E4A38DB9924D8F2553CAD2E85D61B88777279CC11FA562FB7B8823FCA41C68FF2170A14F2D00780D526B36E0E98C0D1D69289F9295FB48CD1A9F07EA71BCCADCE87D6AACF1D143187753C852EC75F48321DB4CCC3B5BD96181C00E37567F593E95CFDCD392A3A8F0F71C069C0438F042CD9A3D37EC8BC634DCA398CE6FD7644501F5AC9D8412D15FDB7C4AD0176A2FAFA58AB01A93FACA7A8E9822800DC966BCFD37FF9F1A838BDCAA3FA90B8A903B74B97ED501433F7241FB5AAE6823529F367A78AFCE672B5F541C2E907D0DA0A95F952F41C78402870214A88444554C55D7215637514BB3B95E804FC55DC9D76B5A483152EB0428AEFDCD6A79BBF551954621142CD8D
2025/03/01 08:03:13 Extracted public key
2025/03/01 08:03:13 Public key PEM:
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAsqCPgrra9z7AMcNzcuhy
d4QA19TU5IzP5633BTOyZV/0AQkd3qXYGs+fxlgJx8SFGVrY81GsJdTs2LUehj/T
8UMbYak8ksZ52jDQTDlrAt6zhXso0gqm8EsljQKP3v1iOzip+r7C7DBH74eudqXa
042oKFlZhCjIe65IHAJiWJrL3gwEUOnYd05/Nbl6r4+TGW1R7A/d1okDyBNOSyiT
PnX3t/bY32ZTKeqrhOCGpMwyNjHoQ8eUNinlu9jjCJEme5vPFcmBChuK+EYckUg5
XQqG83tXqz5jZpMEkgUQ2+d9gNov3wd0y1KlPX/gaKEw670ewgP4DAXEgfd+Sjjb
mSTY8lU8rS6F1huId3J5zBH6Vi+3uII/ykHGj/IXChTy0AeA1Sazbg6YwNHWkon5
KV+0jNGp8H6nG8ytzofWqs8dFDGHdTyFLsdfSDIdtMzDtb2WGBwA43Vn9ZPpXP3N
OSo6jw9xwGnAQ48ELNmj037IvGNNyjmM5v12RFAfWsnYQS0V/bfErQF2ovr6WKsB
qT+sp6jpgigA3JZrz9N/+fGoOL3Ko/qQuKkDt0uX7VAUM/ckH7Wq5oI1KfNnp4r8
5nK19UHC6QfQ2gqV+VL0HHhAKHAhSohERVTFXXIVY3UUuzuV6AT8VdyddrWkgxUu
sEKK79zWp5u/VRlUYhFCzY0CAwEAAQ==
-----END PUBLIC KEY-----

2025/03/01 08:03:13 Public key written to "bldr_exmple_root_r1.key.pem"
2025/03/01 08:03:16 root ceremony failed: failed to create certificate: unsupported hash function
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -config=contrib/boulder/example_root_r1.yaml 
2025/03/01 08:04:55 Preparing root ceremony for bldr_exmple_root_r1.crt.pem
2025/03/01 08:04:55 root ceremony failed: failed to validate config: outputs.public-key-path is "bldr_exmple_root_r1.key.pem", which already exists
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -config=contrib/boulder/example_root_r1.yaml 
2025/03/01 08:05:10 Preparing root ceremony for bldr_exmple_root_r1.crt.pem
2025/03/01 08:05:10 root ceremony failed: failed to validate config: outputs.public-key-path is "bldr_exmple_root_r1.key.pem", which already exists
@melinko2003 ➜ /workspaces/shim-go (main) $ rm -rf *.pem
@melinko2003 ➜ /workspaces/shim-go (main) $ shim/go-shim-039b3ebeba/bin/ceremony -config=contrib/boulder/example_root_r1.yaml 
2025/03/01 08:05:32 Preparing root ceremony for bldr_exmple_root_r1.crt.pem
2025/03/01 08:05:32 Opened PKCS#11 session for slot 118006441
2025/03/01 08:05:32 Generating RSA key with 4096 bit modulus and public exponent 65537 and ID c049706f
2025/03/01 08:05:32     Encoded public exponent (65537) as: 010001
2025/03/01 08:05:32 Key generated
2025/03/01 08:05:32 Extracting public key
2025/03/01 08:05:32     Public exponent: 65537
2025/03/01 08:05:32     Modulus: (4096 bits) EB87637C09BAA0F060370D148FC90CC03B59058A7C5C83CD64CC84E0EA28FEA16F193B0C5C24AB7E217C74517898CD2850C348E6AEE2C8C9279BFC980FA2F6B17926A5DF4E13625A82DFBA415ED63016C4E0F8376A7DC2908E494B2A6A3EBEE42B2FED8C211953E12F6FBB985C48B9FFB3B41B02CA122AFF63979A9D0C163C3063A2ABA739B07C17FA5450AFE1A27DA3BA7C5C1C228B3239F6B8D556423616117A6CB39A44D57CDF69D15D83B9D797C8D89C8D8F7A77C1FAEE501FD7E2832996B32726F2A8E0A3C4C5A7F38C4C7FC2AB6A554A718997462332E24EE6F1A76F92B5040F01A350EDF39193086290ADE12BB161B48B1248F790EF24457687DC187DD1228800320A05B652DEF3DF9521BEA4B88CAF1FA26C9C8780AC4A4B818199931298A66CE87E6ED2DE628721088413251D0C8739BFE309BCF99566B9B932723517C47BC312ADC6064770545B3DD172D52018DB50ED33B29984D281F3166D097D6CAF3CB963FB9DBC615480974DF7ED6CB5BAF2DA6FECB3BE7AE91BD6A55F77667151EF5C8DF4E5D323DED39568BBFF82B4915077303922CF3116888CCC2A09559D662493759DD0C66172B456509DCA3FCB9E517583BA16C180318C4014A1C07B7DDD9AFE6503B9787FEB5A8704D9EE6017FB193121440A07EF41C1A2F6788F1A00117CB9BA366EA2427AF42E8863A0F6C7444CC37E79D47B6FEC2446F11EA325
2025/03/01 08:05:32 Extracted public key
2025/03/01 08:05:32 Public key PEM:
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA64djfAm6oPBgNw0Uj8kM
wDtZBYp8XIPNZMyE4Ooo/qFvGTsMXCSrfiF8dFF4mM0oUMNI5q7iyMknm/yYD6L2
sXkmpd9OE2Jagt+6QV7WMBbE4Pg3an3CkI5JSypqPr7kKy/tjCEZU+Evb7uYXEi5
/7O0GwLKEir/Y5eanQwWPDBjoqunObB8F/pUUK/hon2junxcHCKLMjn2uNVWQjYW
EXpss5pE1XzfadFdg7nXl8jYnI2PenfB+u5QH9figymWsycm8qjgo8TFp/OMTH/C
q2pVSnGJl0YjMuJO5vGnb5K1BA8Bo1Dt85GTCGKQreErsWG0ixJI95DvJEV2h9wY
fdEiiAAyCgW2Ut7z35UhvqS4jK8fomych4CsSkuBgZmTEpimbOh+btLeYochCIQT
JR0Mhzm/4wm8+ZVmubkycjUXxHvDEq3GBkdwVFs90XLVIBjbUO0zspmE0oHzFm0J
fWyvPLlj+528YVSAl0337Wy1uvLab+yzvnrpG9alX3dmcVHvXI305dMj3tOVaLv/
grSRUHcwOSLPMRaIjMwqCVWdZiSTdZ3QxmFytFZQnco/y55RdYO6FsGAMYxAFKHA
e33dmv5lA7l4f+tahwTZ7mAX+xkxIUQKB+9BwaL2eI8aABF8ubo2bqJCevQuiGOg
9sdETMN+edR7b+wkRvEeoyUCAwEAAQ==
-----END PUBLIC KEY-----

2025/03/01 08:05:32 Public key written to "bldr_exmple_root_r1.key.pem"
2025/03/01 08:05:33 root ceremony failed: failed to create certificate: unsupported hash function
```
Softhsm only has partial SHA3 support so I suspect this is related.