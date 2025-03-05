#!/usr/bin/bash

# golang repo Install
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo sudo apt install golang-go -y

# Pathing
CONTRIB=$PWD$(dirname $(dirname $0) | cut -d"." -f2 )
echo $CONTRIB
ROOT=$(dirname $CONTRIB )
echo $ROOT
SRC="${ROOT}/src"
echo $SRC
SHIM="${ROOT}/shim"
echo $SHIM
GO_SRC="$SRC/go/src"
echo $GO_SRC
GO_BUILD_ROOT="$SRC/go"

# Install Go's sources into our src 
cd $SRC
if [ -d "go" ];then
    cd go/
    git stash
    git pull origin master
else 
    git clone https://github.com/golang/go
    cd go/ 
fi 

# Apply our Patch
git apply ../../patches/go/p1.patch

cd $GO_SRC
# Grab Go's short git commit
GIT_COMMIT=$(git rev-parse --short HEAD)
echo "Commit[$GIT_COMMIT] Current Path $(pwd)"
# This Takes like 10 minutes omg
./all.bash

# Build a Go-Shim
cd $SHIM
OURSHIM="$SHIM/go-shim-$GIT_COMMIT"
if [ -d $OURSHIM ]; then
    mv $OURSHIM $OURSHIM-previous 
fi
mkdir $OURSHIM

cd $GO_BUILD_ROOT
# This Part effectively builds the shim
cp -a bin pkg src misc lib $OURSHIM/
# adding a Psuedo Activate script to handle the exporting.
cp $CONTRIB/shim/golang-activate "$OURSHIM/bin/activate"
sed -i "s,GO_VENV_P2,$GIT_COMMIT,g" "$OURSHIM/bin/activate"
chmod +x "$OURSHIM/bin/activate"

echo "Creating: $GIT_COMMIT @ $OURSHIM "

# Backups
cd $SHIM
tar czf go-only-shim-$GIT_COMMIT.tar.gz go-shim-$GIT_COMMIT
tar tf go-only-shim-$GIT_COMMIT.tar.gz 