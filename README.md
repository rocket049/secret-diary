# 说明

`secret-diary` (安全日记本)是一个加密日记本，加密强度非常强，只能暴力破解数据，只要密码足够复杂，数据就足够安全。
编辑器也相当完善，支持富文本格式。

适用于场合：办公日记、个人日记。

兼容 `linux/windows`。
支持语言：中文、英语

### 百度下载
链接: https://pan.baidu.com/s/14Ltsh1WiuKhHgMA7KgA-dw 提取码: stdc

二维码下载：

![二维码下载](baidu.jpeg)

### ChangeLog

[At github](https://github.com/rocket049/secret-diary/releases)

[码云上的](https://gitee.com/rocket049/secret-diary/releases)

## 注册用户和删除用户
### 注册
在登陆界面上注册用户，方法是：
- 先输入用户名和密码
- 然后点击“注册”
- 接着在弹出的输入框中重复输入密码
- 点击确定

### 删除
1. 如何删除多余的用户？只需到数据存储目录中删除对应名字的目录。How to delete redundant users? Just go to the data storage directory and delete the directory with the corresponding name.
2. 数据存储目录在哪里？点击“帮助”菜单中的“备份”，会显示该目录路径。Where is the data storage directory? Click Backup on the Help menu to display the directory path.

# 加密算法说明

### 登陆验证算法
`sha256`

哈希值计算过程： 

1. 真实密码 = 用户密码 X 4
2. sha256( 真实密码 )

### 数据加密算法
`AES256`

数据加密密码为创建用户时从系统读取的32字节随机字符串。

“数据加密密码”被加密后存储在数据库中。

用于加密“数据加密密码”的密码形成算法：

1. 用户真实密码 = 用户密码 X 40
2. 真实密码 = sha256( 用户真实密码 )

## 支持作者 Support Author
全凭您的自愿！ Voluntary!

![支付宝 alipay](./pay/alipay.jpg)

![微信支付 wxpay](./pay/wxpay.png)
