package apiType

type CreateCoinServerParam struct {
	OkxKeyID string `bson:"OkxKeyID"` // ApiKey 的 ID
	Port     string `bson:"Port"`     // 需要请求的服务的端口  Host + Port 必须唯一
	Note     string `bson:"Note"`     // 备注
	Password string `bson:"Password"`
}
