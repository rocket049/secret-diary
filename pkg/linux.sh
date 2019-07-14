#!/bin/sh

~/go/bin/qtdeploy build

cd deploy
mkdir -p secret-diary-linux_amd64
cp -ru ../locale secret-diary-linux_amd64/
cp -u ../Sd.png secret-diary-linux_amd64/
cp -ru linux/* secret-diary-linux_amd64/
ln -sf ~/go/bin/qtdeploy/secret-diary-linux_amd64/secret-diary secret-diary
rm secret-diary-linux_amd64.zip
zip -r9 secret-diary-linux_amd64.zip secret-diary-linux_amd64/

sh mkdeb.sh
