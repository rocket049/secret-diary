# 说明
`secret-diary` (安全日记本)是一个加密日记本，加密强度非常强，只能暴力破解数据，只要密码足够复杂，数据就足够安全。
编辑器也相当完善，支持富文本格式。
兼容 `linux/windows`。
支持语言：中文、英语

### ChangeLog

[At github](https://github.com/rocket049/secret-diary/releases)

[码云上的](https://gitee.com/rocket049/secret-diary/releases)

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
