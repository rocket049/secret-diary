#!/bin/sh
nano secret-diary-deb/DEBIAN/control
cp -ru secret-diary-linux_amd64/* secret-diary-deb/opt/secret-diary/
VERSION="$(fgrep Version secret-diary-deb/DEBIAN/control |awk '{print $2}')"
DEBNAME="secret-diary-${VERSION}_amd64.deb"
dpkg -b secret-diary-deb $DEBNAME
cp $DEBNAME ~/src/pkg2appimage/debs/secret-diary_amd64.deb
cd ~/src/pkg2appimage/
./pkg2appimage recipes/secret-diary.yml
mv out/Secret-Diary-amd64.deb.glibc2.14-x86_64.AppImage ~/go/src/secret-diary/deploy/Secret-Diary-${VERSION}-x86_64.AppImage
