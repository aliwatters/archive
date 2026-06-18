#!/bin/bash

HOST=dev.example.org
DIR=$(dirname "$0") 

echo "Setting up nginx for development on $HOST"
mkdir -p $DIR/ssl

# 1 - generate some private keys
if [ ! -f "$DIR/ssl/$HOST.key" ]; then
    $DIR/certify.sh $HOST
fi

# 2 - detect hostname
if grep -Fq "$HOST" /etc/hosts
then
    echo "done"
else
    echo ""
    echo "======================================================"
    echo "ACTION NEEDED: add '127.0.1.1 $HOST' to your '/etc/hosts'"
    echo "MacOS flush dns with:"
    echo "sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder"
    echo ""
fi
