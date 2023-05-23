# UserCenter

AItrade 的数据中心,用来进行用户账户的数据管理


```bash
## 同步当前项目
npm run sync 

## 拉取公共模块
git clone git@github.com:AItrade-mo7/sysPublic.git

## 运行
npm run serve

## 编译打包
npm run build

## 部署
npm run deploy


```

## sys_env.yaml

```yaml
# Mongodb
MongoAddress: "test-mongo.mo7.cc:17017"
MongoUserName: "tester"
MongoPassword: "123456"
MessageBaseUrl: "http://test-msg.mo7.cc"
```

## go work

```go

go 1.20

use (
./
)

replace (
github.com/EasyGolang/goTools => /root/EasyGolang/goTools
)


```
