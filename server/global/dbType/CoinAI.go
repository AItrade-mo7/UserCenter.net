package dbType

type ApiKeyType struct {
	Name       string `bson:"Name"`
	ApiKey     string `bson:"ApiKey"`
	SecretKey  string `bson:"SecretKey"`
	Passphrase string `bson:"Passphrase"`
}

type AppEnv struct {
	Name       string       `bson:"Name"`
	Type       string       `bson:"Type"` // 服务类型
	Version    string       `bson:"Version"`
	IP         string       `bson:"IP"`
	Port       string       `bson:"Port"`
	ServeID    string       `bson:"ServeID"`
	UserID     string       `bson:"UserID"`
	ApiKeyList []ApiKeyType `bson:"ApiKeyList"`
}
