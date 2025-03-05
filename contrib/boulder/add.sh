#!/usr/bin/bash
# boulder/add.sh : Adds Boulder to your latest go venv

# Pathing
CONTRIB=$PWD$(dirname $(dirname $0) | cut -d"." -f2 )
echo $CONTRIB
ROOT=$(dirname $CONTRIB )
echo $ROOT
SRC="${ROOT}/src"
echo $SRC
SHIM="${ROOT}/shim"
echo $SHIM
MYSHIM="$SHIM/$(ls -t $SHIM | grep go-shim )"
echo $MYSHIM

zGOROOT=$GOROOT
zGOBIN=$GOBIN
xPATH=$PATH

export GOROOT="$MYSHIM"
export GOBIN="$MYSHIM/bin"
export PATH="$GOBIN:$PATH"

cd $SRC
# Install Boulder's sources into our src 
if [ -d "boulder" ];then
    cd go/
    git stash
    git pull origin main
else 
    git clone https://github.com/letsencrypt/boulder
    cd boulder/ 
fi 

# Git Commits
GIT_COMMIT=$(git rev-parse --short HEAD)

# Apply our Patch
git apply ../../patches/boulder/p1.patch

# Build Sources
make all

# Backups
cd $SHIM
cp -a $MYSHIM $SHIM/go-boulder-$GIT_COMMIT
tar czf go-boulder-$GIT_COMMIT.tar.gz go-boulder-$GIT_COMMIT
tar tf go-boulder-$GIT_COMMIT.tar.gz 

export GOROOT="$zGOROOT"
export GOBIN="$zGOBIN"
export PATH="$xPATH"