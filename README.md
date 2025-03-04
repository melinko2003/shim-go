# shim-go
A shim for Go/Go Projects 

## Prep
Because of certain limitations in Almalinux 9, we'll use Ubuntu LTS 24.04.
1) Git Clone this repo ( `git clone https://github.com/melinko2003/shim-go` )
2) Move into the `shim-go` directory ( `cd shim-go` ) 
3) Run the contributed `contrib/shim/build.sh` script that will build a Go Shim for you. ( `./contrib/shim/build.sh` ) 
4) Wait 10 minutes, and then check in the `shim/` directory
5) Activate the shim via `source shim/go-shim-$your-hash/bin/activate`

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
