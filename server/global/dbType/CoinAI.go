package dbType

type ApiKeyType struct {
	Name       string `bson:"Name"`
	ApiKey     string `bson:"ApiKey"`
	SecretKey  string `bson:"SecretKey"`
	Passphrase string `bson:"Passphrase"`
}

type AppEnv struct {
	Name       string       `bson:"Name"`
	Version    string       `bson:"Version"`
	Port       string       `bson:"Port"`
	IP         string       `bson:"IP"`
	ServeID    string       `bson:"ServeID"`
	UserID     string       `bson:"UserID"`
	ApiKeyList []ApiKeyType `bson:"ApiKeyList"`
}
