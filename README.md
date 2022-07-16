# UserCenter

AITrade 的数据中心,用来进行用户账户的数据管理

## sys_env.yaml

```yaml
# IP
LocalIP: '50.18.29.218'
# Mongodb
MongoAddress: 'trade.mo7.cc:17017'
MongoUserName: 'mo7'
MongoPassword: 'asdasd55555'

# 运行模式
RunMod: 1
```

## go work

```go

go 1.18

use (
./
)

replace (
github.com/EasyGolang/goTools => /root/EasyGolang/goTools
)

```
