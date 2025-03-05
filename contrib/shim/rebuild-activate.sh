#!/usr/bin/bash

# Usage: 
# $ ./contrib/shim/rebuild-activate.sh "$(pwd)/shim/go-shim-0cfc641420" 0cfc641420
# Creating: 0cfc641420 @ /workspaces/go-shim/shim/go-shim-0cfc641420/bin/activate 
# $ source /workspaces/go-shim/shim/go-shim-0cfc641420/bin/activate

# ░░░░░░░░      ░░░░      ░░░░░░░░░░      ░░░  ░░░░  ░░        ░░  ░░░░  ░░░░░░░░  ░░░░  ░░        ░░   ░░░  ░░  ░░░░  ░░░░░░░
# ▒▒▒▒▒▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒▒▒▒  ▒▒▒▒▒   ▒▒   ▒▒▒▒▒▒▒▒  ▒▒▒▒  ▒▒  ▒▒▒▒▒▒▒▒    ▒▒  ▒▒  ▒▒▒▒  ▒▒▒▒▒▒▒
# ▓▓▓▓▓▓▓  ▓▓▓   ▓▓  ▓▓▓▓  ▓▓▓▓▓▓▓▓▓      ▓▓▓        ▓▓▓▓▓  ▓▓▓▓▓        ▓▓▓▓▓▓▓▓▓  ▓▓  ▓▓▓      ▓▓▓▓  ▓  ▓  ▓▓▓  ▓▓  ▓▓▓▓▓▓▓▓
# ███████  ████  ██  ████  ██████████████  ██  ████  █████  █████  █  █  ██████████    ████  ████████  ██    ████    █████████
# ████████      ████      ██████████      ███  ████  ██        ██  ████  ███████████  █████        ██  ███   █████  ██████████
                                                                                                                            
# [GO/VENV] ➜ GOROOT@/workspaces/go-shim/shim/go-shim-0cfc641420
# [GO/VENV] ➜ GOBIN@/workspaces/go-shim/shim/go-shim-0cfc641420/bin
# [GO/VENV] ➜ Version: go version devel go1.25-0cfc641420 Tue Mar 4 12:43:35 2025 -0800 linux/amd64
# (0cfc641420) @melinko2003 ➜ /workspaces/go-shim (main) $ 

# Venv Name
if [ -z "$1" ] ; then
    echo "Please provide a shim at least."
    exit 1
fi

cp contrib/shim/go-activate.sh "$1/bin/activate"
chmod +x "$1/bin/activate"

# Venv Name
if [ -z "$2" ] ; then
    # Easter Egg +
    sed -i "s,GO_VENV_P2,6D656C696E6B6F32303033,g" "$1/bin/activate"
else
    # Set your Githash
    sed -i "s,GO_VENV_P2,$2,g" "$1/bin/activate"
fi

echo "Creating: $2 @ $1/bin/activate "