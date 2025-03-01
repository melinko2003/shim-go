# Prep
Grab `Go`'s source from somewhere
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ cd src/
@melinko2003 ➜ /workspaces/shim-go/src (main) $ git clone https://github.com/melinko2003/go
Cloning into 'go'...
remote: Enumerating objects: 615180, done.
remote: Counting objects: 100% (28/28), done.
remote: Compressing objects: 100% (18/18), done.
remote: Total 615180 (delta 15), reused 10 (delta 10), pack-reused 615152 (from 3)
Receiving objects: 100% (615180/615180), 364.77 MiB | 34.34 MiB/s, done.
Resolving deltas: 100% (483873/483873), done.
Updating files: 100% (14189/14189), done.
```
Update to a modern `Go`: https://go.dev/wiki/Ubuntu
```bash
@melinko2003 ➜ /workspaces/shim-go (main) $ sudo add-apt-repository ppa:longsleep/golang-backports
 Golang 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.20, 1.21, 1.22 and 1.23 PPA for Ubuntu

See https://github.com/longsleep/golang-deb for the sources for the packages in this PPA and for reporting issues / change requests.
 More info: https://launchpad.net/~longsleep/+archive/ubuntu/golang-backports
Press [ENTER] to continue or Ctrl-c to cancel adding it.

Get:1 https://packages.microsoft.com/repos/microsoft-ubuntu-focal-prod focal InRelease [3632 B]
Get:3 https://packages.microsoft.com/repos/microsoft-ubuntu-focal-prod focal/main all Packages [2938 B]                                                                                                                   
Get:4 https://packages.microsoft.com/repos/microsoft-ubuntu-focal-prod focal/main amd64 Packages [336 kB]                                                                                                                 
Get:5 https://dl.yarnpkg.com/debian stable InRelease [17.1 kB]                                                                                                                                                            
Get:6 https://repo.anaconda.com/pkgs/misc/debrepo/conda stable InRelease [3961 B]                                                                                                               
Get:7 http://archive.ubuntu.com/ubuntu focal InRelease [265 kB]                                                                         
Get:8 http://security.ubuntu.com/ubuntu focal-security InRelease [128 kB]                 
Get:2 https://packagecloud.io/github/git-lfs/ubuntu focal InRelease [29.1 kB]                                  
Get:9 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal InRelease [18.3 kB]                                  
Get:10 https://dl.yarnpkg.com/debian stable/main all Packages [11.8 kB]                  
Get:11 https://dl.yarnpkg.com/debian stable/main amd64 Packages [11.8 kB]                                                               
Get:12 https://repo.anaconda.com/pkgs/misc/debrepo/conda stable/main amd64 Packages [4557 B]                                                                   
Get:13 https://packagecloud.io/github/git-lfs/ubuntu focal/main amd64 Packages [4064 B]                                                
Get:14 http://archive.ubuntu.com/ubuntu focal-updates InRelease [128 kB]
Get:15 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal/main amd64 Packages [7176 B]
Get:16 http://archive.ubuntu.com/ubuntu focal-backports InRelease [128 kB]    
Get:17 http://security.ubuntu.com/ubuntu focal-security/universe amd64 Packages [1301 kB]
Get:18 http://archive.ubuntu.com/ubuntu focal/main amd64 Packages [1275 kB]
Get:19 http://security.ubuntu.com/ubuntu focal-security/multiverse amd64 Packages [30.9 kB]
Get:20 http://security.ubuntu.com/ubuntu focal-security/main amd64 Packages [4252 kB]
Get:21 http://archive.ubuntu.com/ubuntu focal/universe amd64 Packages [11.3 MB]   
Get:22 http://security.ubuntu.com/ubuntu focal-security/restricted amd64 Packages [4389 kB]    
Get:23 http://archive.ubuntu.com/ubuntu focal/restricted amd64 Packages [33.4 kB]         
Get:24 http://archive.ubuntu.com/ubuntu focal/multiverse amd64 Packages [177 kB]
Get:25 http://archive.ubuntu.com/ubuntu focal-updates/multiverse amd64 Packages [34.6 kB]
Get:26 http://archive.ubuntu.com/ubuntu focal-updates/main amd64 Packages [4730 kB]
Get:27 http://archive.ubuntu.com/ubuntu focal-updates/universe amd64 Packages [1593 kB]
Get:28 http://archive.ubuntu.com/ubuntu focal-updates/restricted amd64 Packages [4577 kB]
Get:29 http://archive.ubuntu.com/ubuntu focal-backports/universe amd64 Packages [28.6 kB]
Get:30 http://archive.ubuntu.com/ubuntu focal-backports/main amd64 Packages [55.2 kB]
Fetched 34.9 MB in 3s (12.6 MB/s)                       
Reading package lists... Done
@melinko2003 ➜ /workspaces/shim-go (main) $ sudo apt install golang-go
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  golang-1.23-go golang-1.23-src golang-src
Suggested packages:
  bzr | brz mercurial subversion
The following NEW packages will be installed:
  golang-1.23-go golang-1.23-src golang-go golang-src
0 upgraded, 4 newly installed, 0 to remove and 71 not upgraded.
Need to get 46.2 MB of archives.
After this operation, 242 MB of additional disk space will be used.
Do you want to continue? [Y/n] Y
Get:1 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal/main amd64 golang-1.23-src all 1.23.6-1longsleep1+focal [19.2 MB]
Get:2 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal/main amd64 golang-1.23-go amd64 1.23.6-1longsleep1+focal [27.0 MB]
Get:3 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal/main amd64 golang-src all 2:1.23~0longsleep1 [6184 B]
Get:4 http://ppa.launchpad.net/longsleep/golang-backports/ubuntu focal/main amd64 golang-go amd64 2:1.23~0longsleep1 [45.4 kB]
Fetched 46.2 MB in 3s (13.4 MB/s)         
Selecting previously unselected package golang-1.23-src.
(Reading database ... 70107 files and directories currently installed.)
Preparing to unpack .../golang-1.23-src_1.23.6-1longsleep1+focal_all.deb ...
Unpacking golang-1.23-src (1.23.6-1longsleep1+focal) ...
Selecting previously unselected package golang-1.23-go.
Preparing to unpack .../golang-1.23-go_1.23.6-1longsleep1+focal_amd64.deb ...
Unpacking golang-1.23-go (1.23.6-1longsleep1+focal) ...
Selecting previously unselected package golang-src.
Preparing to unpack .../golang-src_2%3a1.23~0longsleep1_all.deb ...
Unpacking golang-src (2:1.23~0longsleep1) ...
Selecting previously unselected package golang-go:amd64.
Preparing to unpack .../golang-go_2%3a1.23~0longsleep1_amd64.deb ...
Unpacking golang-go:amd64 (2:1.23~0longsleep1) ...
Setting up golang-1.23-src (1.23.6-1longsleep1+focal) ...
Setting up golang-src (2:1.23~0longsleep1) ...
Setting up golang-1.23-go (1.23.6-1longsleep1+focal) ...
Setting up golang-go:amd64 (2:1.23~0longsleep1) ...
Processing triggers for man-db (2.9.1-1) ...
```
