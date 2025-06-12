# Linux上如何自行编译和部署本软件（go module）

### 一、安装go编译器

1. 从`https://go.dev`下载`go`编译器，解压缩到`/usr/local/`
2. 在`$HOME/.bashrc`中加入一行`export PATH=$PATH:/usr/local/go/bin`
3. 用命令`. $HOME/.bashrc`导入环境变量

然后可以用命令`go version`测试`go`编译器是否能运行。

接着要设置`GOPROXY`变量：`go env -w GOPROXY=https://goproxy.cn,direct`

### 二、下载本软件的源代码
运行命令：`git clone https://gitee.com/rocket049/secret-diary.git`

### 三、安装必须的共享库(来自 https://github.com/therecipe/qt)
```
ubuntu/debian:
sudo apt-get -y install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev
sudo apt-get --no-install-recommends install libqt*5-dev qt*5-dev qml-module-qtquick-* qt*5-doc-html

Fedora/RHEL/CentOS:
sudo yum -y groupinstall "C Development Tools and Libraries"
sudo yum -y install mesa-libGLU-devel gstreamer-plugins-base pulseaudio-libs-devel glib2-devel
sudo yum install qt5-* qt5-*-doc

openSUSE:
sudo zypper -n install -t pattern devel_basis
sudo zypper install --no-recommends libqt5-qt*-devel

Arch Linux:
sudo pacman -S base-devel
sudo pacman -S --needed qt5

Deepin Linux:
sudo apt-get -y install build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev
sudo apt install qtbase5-dev

```

### 四、用go编译器编译软件
在前面运行`git`命令的目录里，按顺序输入下面的命令：

```
cd secret-diary/
export QT_PKG_CONFIG=true
go mod tidy
go get -v -tags=no_env github.com/therecipe/qt/cmd/...
go install -v -tags=no_env github.com/therecipe/qt/cmd/...
go mod vendor
~/go/bin/qtdeploy build desktop
```

等待上面的命令运行结束，就可以得到编译好的程序，位置是:`deploy/linux/secret-diary`

#### 按上面的办法编译还是失败？因为那都是过去式了，下面介绍一下在ubuntu24.04上该怎么编译本软件。
1、`go get -v -tags=no_env github.com/therecipe/qt/cmd/...`这条命令现在不能用了，怎么办？按下面的步骤编译他们：
先执行下面的命令：
```
git clone https://github.com/therecipe/qt
cd github.com/therecipe/qt/cmd
```
`cmd`里面有`env  qtdeploy  qtminimal  qtmoc  qtrcc  qtsetup`这些目录，排除掉`env`，其他目录都`cd`进去分别编译，命令是：
```
go mod tidy
go install
```
2、增加两个环境变量：
```
export QT_API=5.13.0
export QT_PKG_CONFIG=true
```
3、到`secret-diary`目录下，执行下面3条命令：
```
go mod tidy
go mod vendor
~/go/bin/qtsetup generate
```
这时可能会报错，没关系，再执行一遍 `go mod vendor` 和 `~/go/bin/qtsetup generate`就可以成功了。

4、编译：

运行 `~/go/bin/qtdeploy build desktop`，输出文件在 `deploy/linux/`中。

5、如何使用自己编译的QT5.15.16？用下面这几步代替上面的第2步就可以了。

先到官网下载`qt-everywhere-opensource-src-5.15.16.tar.xz`，假设你已经编译成功，安装目录为`/home/app/qt5`，下一步你需要重建一个符合`github.com/therecipe/qt`标准的目录结构：
```
QT_DIR
    5.15.16
        gcc_64
```
具体做法如下：
```
sudo mkdir -p /home/app/qt5dir/5.15.16
sudo ln -s /home/app/qt5 /home/app/qt5dir/5.15.16/gcc_64
```
`gcc_64`就是指向`/home/app/qt5`的软链接，确保下一层就是`bin include lib plugins`等目录。

***注意：***如果你使用fcitx或fcitx5输入法，记得复制下面两个输入法插件，否则编译出来的程序不能使用输入法：
```
sudo cp usr/lib/x86_64-linux-gnu/qt5/plugins/platforminputcontexts/libfcitx5platforminputcontextplugin.so /home/app/qt5/plugins/platforminputcontexts/
sudo cp usr/lib/x86_64-linux-gnu/qt5/plugins/platforminputcontexts/libfcitxplatforminputcontextplugin.so /home/app/qt5/plugins/platforminputcontexts/
```

然后要设置对应的环境变量：
```
export QT_DIR=/home/app/qt5dir
export QT_VERSION=5.15.16
export QT_API=5.13.0
```
设置完成，可以继续第3步了。

### 五、部署软件
在`secret-diary`目录中依次运行下面的命令：

```
sudo mkdir /opt/secret-diary
sudo cp deploy/linux/secret-diary /opt/secret-diary/
sudo cp -r locale /opt/secret-diary/
sudo cp Sd.png /opt/secret-diary/
cp Secret-Diary.desktop ~/.local/share/applications/
```

这时你已经可以从开始菜单中点击启动`Secret-Diary`软件了。
