package dbType

// 米游社Cookie 表 表结构  ========== Account ==============
type MiYouSheCookieTable struct {
	Email  string `bson:"Email"`  // 用户 Email
	UserID string `bson:"UserID"` // 用户 ID

	CreateTime    int64  `bson:"CreateTime"`    // 创建时间
	CreateTimeStr string `bson:"CreateTimeStr"` // 创建时间

	MiYouSheCookie string `bson:"MiYouSheCookie"` // Cookie
}
