# secret-diary 在 ubuntu19.04 上运行的问题
secret-diary 运行需要 `libicu60` 的支持，ubuntu19.04 中没有 libicu60 只有 libicu63，导致运行失败。解决的办法是从 ubuntu18.04 的库中下载 libicu60 的 deb 包，然后手动安装进去。

Secret-diary needs the support of `libicu60` to run. There is no libicu60 but libicu63 in ubuntu19.04, which results in the failure of the operation. The solution is to download the DEB package of libicu60 from the library of ubuntu18.04 and install it manually.

安装命令(install)：
`sudo apt install ./libicu60_60.2-3ubuntu3_amd64.deb`

### 下载地址(Download address)：
1. [http://mirrors.163.com/ubuntu/pool/main/i/icu/libicu60_60.2-3ubuntu3_amd64.deb](http://mirrors.163.com/ubuntu/pool/main/i/icu/libicu60_60.2-3ubuntu3_amd64.deb)
2. [http://archive.ubuntu.com/ubuntu/pool/main/i/icu/libicu60_60.2-3ubuntu3_amd64.deb](http://archive.ubuntu.com/ubuntu/pool/main/i/icu/libicu60_60.2-3ubuntu3_amd64.deb)
