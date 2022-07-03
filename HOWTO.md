# Linux上如何自行编译和部署本软件（go module）

### 一、安装go编译器

1. 从`https://go.dev`下载`go`编译器，解压缩到`/usr/local/`
2. 在`$HOME/.bashrc`中加入一行`export PATH=$PATH:/usr/local/go/bin`
3. 用命令`. $HOME/.bashrc`导入环境变量

然后可以用命令`go version`测试`go`编译器是否能运行。

接着要设置`GOPROXY`变量：`go env -w GOPROXY=https://goproxy.cn,direct`

### 二、下载本软件的源代码
运行命令：`git clone https://gitee.com/rocket049/secret-diary.git`

### 三、安装必须的共享库
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