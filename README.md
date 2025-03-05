# shim-go
A shim for Go/Go Projects 

## Prep
Because of certain limitations in Almalinux 9, we'll use Ubuntu LTS 24.04.
1) Git Clone this repo ( `git clone https://github.com/melinko2003/shim-go` )
2) Move into the `shim-go` directory ( `cd shim-go` ) 
3) Run the contributed `contrib/shim/build.sh` script that will build a Go Shim for you. ( `./contrib/shim/build.sh` ) 
4) Wait 10 minutes, and then check in the `shim/` directory
5) Activate the shim via `source shim/shim-go-$your-hash/bin/activate`

Setup Example:
```bash
@melinko2003 ➜ /workspaces (main) $ git clone https://github.com/melinko2003/shim-go
@melinko2003 ➜ /workspaces (main) $ cd shim-go
@melinko2003 ➜ /workspaces/shim-go (main) $ ./contrib/shim/build.sh
< ... Wait 10 Minutes ... >
@melinko2003 ➜ /workspaces/shim-go (main) $ ls shim/
go-shim-0cfc641420
```
Activate the Go/Shim venv:
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ source /workspaces/shim-go/shim/go-shim-0cfc641420/bin/activate


░░░░░░░░      ░░░░      ░░░░░░  ░░      ░░░  ░░░░  ░░        ░░  ░░░░  ░░░░░░░░  ░░░░  ░░        ░░   ░░░  ░░  ░░░░  ░░░░░░░
▒▒▒▒▒▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒▒▒  ▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒▒▒▒  ▒▒▒▒▒   ▒▒   ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒  ▒▒▒▒▒▒▒▒    ▒▒  ▒▒  ▒▒▒▒  ▒▒▒▒▒▒▒
▓▓▓▓▓▓▓  ▓▓▓   ▓▓  ▓▓▓▓  ▓▓▓  ▓▓▓▓      ▓▓▓        ▓▓▓▓▓  ▓▓▓▓▓        ▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓▓      ▓▓▓▓  ▓  ▓  ▓▓▓  ▓▓  ▓▓▓▓▓▓▓▓
███████  ████  ██  ████  ██  ██████████  ██  ████  █████  █████  █  █  ██████████    ████  ████████  ██    ████    █████████
████████      ████      ██  ██████      ███  ████  ██        ██  ████  ███████████  █████        ██  ███   █████  ██████████
                                                                                                                            

[GO/VENV] ➜ GOROOT @ /workspaces/shim-go/shim/go-shim-0cfc641420
[GO/VENV] ➜ GOBIN @ /workspaces/shim-go/shim/go-shim-0cfc641420/bin
[GO/VENV] ➜ Version @ go version devel go1.25-0cfc641420 Tue Mar 4 12:43:35 2025 -0800 linux/amd64
(0cfc641420) @melinko2003 ➜ /workspaces/shim-go (main) 
```
To Deactivate the Shim
```bash
(0cfc641420) @melinko2003 ➜ /workspaces/shim-go (main) $ deactivate
@melinko2003 ➜ /workspaces/shim-go (main) $ 
```
## Test Patch Feature
```bash
(0cfc641420) @melinko2003 ➜ /workspaces/shim-go (main) $ cd tests/go/
(0cfc641420) @melinko2003 ➜ /workspaces/shim-go/tests/go (main) $ go run rsa_root.go 
Successfully created SHA3256WithRSA Root -> SHA3_256WithRSA.pem
Successfully created SHA3384WithRSA Root -> SHA3_384WithRSA.pem
Successfully created SHA3512WithRSA Root -> SHA3_512WithRSA.pem
(0cfc641420) @melinko2003 ➜ /workspaces/shim-go/tests/go (main) $ go run ecdsa_root.go 
Successfully created ECDSAWithSHA3_256 Root -> ECDSAWithSHA3_256.pem
Successfully created ECDSAWithSHA3_384 Root -> ECDSAWithSHA3_384.pem
Successfully created ECDSAWithSHA3_512 Root -> ECDSAWithSHA3_512.pem
```
## Add Boulder to your venv
Will install boulder into your latest `shim/go-shim`, and make a `shim/go-boulder` shim. 
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ ./contrib/boulder/add.sh
```
## Extending Examples
We have a Soft CA example [guide](EXTRA.md).
