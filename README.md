# UserCenter

AItrade 的数据中心,用来进行用户账户的数据管理

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
