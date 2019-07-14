#!/bin/sh

cp deploy/win32.syso .
~/go/bin/qtdeploy -docker build windows_32_static
rm win32.syso

cd deploy
mkdir -p secret-diary-win32
cp -ru ../locale secret-diary-win32/
cp -u ../Sd.ico secret-diary-win32/
cp -ru windows/* secret-diary-win32/
rm secret-diary-win32.zip
zip -r9 secret-diary-win32.zip secret-diary-win32/