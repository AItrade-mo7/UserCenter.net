package apiType

// UserInfo 的数据结构  ======== 对外展示 =======

type CreateOkxKeyParam struct {
	ApiKey     string `bson:"ApiKey"`     // ApiKey
	SecretKey  string `bson:"SecretKey"`  // SecretKey  密钥
	Passphrase string `bson:"Passphrase"` // Passphrase  密码
	IP         string `bson:"IP"`         // IP地址
	Name       string `bson:"Name"`       // 备注名,同一个账号下必须唯一
	Note       string `bson:"Note"`       // 备注
	Password   string `bson:"Password"`
}
