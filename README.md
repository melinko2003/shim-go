# shim-go
A shim for Go/Go Projects 

## Prep
Prep your environment with this [guide](PREP.md).

## Patch
Patch in when ever/what ever.
```bash
@melinko2003 ➜ /workspaces/shim-go/src/go (master) $ git apply ../../patches/go/p1.patch
```
## Build
1) Use Build [guide](BUILD.md). 
2) Afterwards define your shim 
    ```bash
    @melinko2003 ➜ .../shim-go/src/go/src (master) $ mkdir -p ../../../shim/go-shim-$(git rev-parse --short HEAD)
    @melinko2003 ➜ .../shim-go/src/go/src (master) $ cd ..
    @melinko2003 ➜ /workspaces/shim-go/src/go (master) $ cp -a bin pkg src misc lib ../../shim/go-shim-$(git rev-parse --short HEAD)/
    ```
## Verify version
```bash
@melinko2003 ➜ /workspaces/shim-go/src/go (master) $ cd ../../
@melinko2003 ➜ /workspaces/shim-go (main) $ ll shim/
total 16
drwxrwxrwx+ 3 codespace codespace 4096 Mar  1 06:03 ./
drwxrwxrwx+ 7 codespace root      4096 Mar  1 05:54 ../
-rw-rw-rw-  1 codespace codespace   27 Mar  1 05:55 README.md
drwxrwxrwx+ 7 codespace codespace 4096 Mar  1 06:07 go-shim-039b3ebeba/
@melinko2003 ➜ /workspaces/shim-go (main) $ export GOROOT="$(pwd)/shim/go-shim-039b3ebeba"
@melinko2003 ➜ /workspaces/shim-go (main) $ export GOBIN="$GOROOT/bin"
@melinko2003 ➜ /workspaces/shim-go (main) $ export PATH="$GOBIN:$PATH"
@melinko2003 ➜ /workspaces/shim-go (main) $ go version
go version devel go1.25-039b3ebeba Fri Feb 28 17:10:54 2025 -0800 linux/amd64
```
## Test Patch Feature
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ cd tests/go/
@melinko2003 ➜ /workspaces/shim-go/tests/go (main) $ go run rsa_root.go 
Successfully created SHA3256WithRSA Root -> SHA3_256WithRSA.pem
Successfully created SHA3384WithRSA Root -> SHA3_384WithRSA.pem
Successfully created SHA3512WithRSA Root -> SHA3_512WithRSA.pem
@melinko2003 ➜ /workspaces/shim-go/tests/go (main) $ go run ecdsa_root.go 
Successfully created ECDSAWithSHA3_256 Root -> ECDSAWithSHA3_256.pem
Successfully created ECDSAWithSHA3_384 Root -> ECDSAWithSHA3_384.pem
Successfully created ECDSAWithSHA3_512 Root -> ECDSAWithSHA3_512.pem
```
## Extending Examples
We have a Soft CA example [guide](EXTRA.md).
