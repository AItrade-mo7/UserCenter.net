package apiType

// UserInfo 的数据结构  ======== 对外展示 =======

type UserInfo struct {
	UserID     string `bson:"UserID"`     // 用户 ID
	Email      string `bson:"Email"`      // 用户 Email
	Avatar     string `bson:"Avatar"`     // 用户头像
	NickName   string `bson:"NickName"`   // 用户昵称
	CreateTime int64  `bson:"CreateTime"` // 创建时间
	UpdateTime int64  `bson:"UpdateTime"` // 更新时间
}

type LoginSucceedType struct {
	UserID      string `bson:"UserID"`
	Email       string `bson:"Email"`
	BrowserName string `bson:"BrowserName"`
	OsName      string `bson:"OsName"`
	Hostname    string `bson:"Hostname"`
	ISP         string `bson:"ISP"`
	Operators   string `bson:"Operators"`
	Token       string `bson:"Token"`
}
