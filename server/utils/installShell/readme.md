# 创建服务

服务器结构：

AItrade：

9000
9001
9002
9003
9004
9005
9006
9007
9008
9009

读取目录, 来判定端口号是多少，关闭一个服务，则关闭一个端口号，并删除数据库

以 user 为中心，创建一个服务，则读取一个目录是否存在，关闭一个服务，则删除对应目录

9000 - 9009

目录存在，则表示该服务不可用，删除一个服务，则数据库对应一并删除

## 前端界面

其它服务器 & 本机

其它服务器，填写 ip 地址&端口号，测试是否连通

本机，则直接选择端口号即可，罗列出可用端口号即可

9001 - 9018

选择其他服务器则 ip 地址和端口号自己随便填写

## 需要实现的功能

读取某一个目录下的目录列表

user 表新增字段

```json
{
  "Server": {
    "IsLocal": true,
    "Host": "mo7.cc",
    "Port": "9003"
  }
}
```

## 打包安装的功能

提供安装步骤，用户自行部署即可，
