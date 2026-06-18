#!/bin/bash -e
if [[ -z "$1" ]]; then cat <<EOF
certify - Creates and self-signs X.509 SSL/TLS certificates
          with the "subjectAltName" extension.
Usage: certify example.com [www.example.com] [mail.example.com] [...]
EOF
exit; fi

pushd $(dirname $0)

umask 066
KEYSIZE=2048
DAYS=3650
certname="$1"
prompt=no
# altnames="DNS:$1"
# while shift && (($#)); do altnames="$altnames,DNS:$1"; done
# echo -e "basicConstraints=critical,CA:true,pathlen:0\nsubjectAltName=$altnames" > extensions.ini

# https://github.com/openssl/openssl/issues/3536 -- certify.conf needed.
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -batch \
    -config certify.conf \
    -keyout ssl/$certname.key -out ssl/$certname.crt 

popd
