package apiType

// UserInfo 的数据结构  ======== 对外展示 =======

type UserInfo struct {
	UserID         string `bson:"UserID"`         // 用户 ID
	Email          string `bson:"Email"`          // 用户 Email
	Avatar         string `bson:"Avatar"`         // 用户头像
	NickName       string `bson:"NickName"`       // 用户昵称
	CreateTime     int64  `bson:"CreateTime"`     // 创建时间
	UpdateTime     int64  `bson:"UpdateTime"`     // 更新时间
	EntrapmentCode string `bson:"EntrapmentCode"` // 防钓鱼码
}

type UserEmailInfo struct {
	UserID    string   `bson:"UserID"`    // 用户 ID
	Email     string   `bson:"Email"`     // 用户主要的 Email
	UserEmail []string `bson:"UserEmail"` // 用户的 Email 列表
	Avatar    string   `bson:"Avatar"`    // 用户头像
	NickName  string   `bson:"NickName"`  // 用户昵称
}
