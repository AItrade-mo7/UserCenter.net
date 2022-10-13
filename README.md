# UserCenter

AITrade 的数据中心,用来进行用户账户的数据管理

## sys_env.yaml

```yaml
# Mongodb
MongoAddress: "xxxx.xxxx.xxxx:79456"
MongoUserName: "sssss"
MongoPassword: "123456"
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
