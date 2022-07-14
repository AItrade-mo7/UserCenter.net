package dbUser

import (
	"fmt"

	"DataCenter.net/server/global"
	"DataCenter.net/server/global/apiType"
	"DataCenter.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
)

func (dbObj *AccountType) CreateOkxKey(data apiType.CreateOkxKeyParam) (resErr error) {
	resErr = nil

	db := dbObj.DB

	OkxKeyID := mStr.Join(
		data.ApiKey, data.SecretKey, data.Passphrase,
	)

	okxKey := dbType.OkxKey{
		OkxKeyID:   mEncrypt.MD5(OkxKeyID),
		ApiKey:     data.ApiKey,
		Name:       data.Name,
		IP:         data.IP,
		SecretKey:  data.SecretKey,
		Passphrase: data.Passphrase,
		Note:       data.Note,
		CreateTime: mTime.GetUnixInt64(),
	}

	// 创建数组,并将数据库中的数组 copy
	size := len(dbObj.AccountData.OkxKeyList)
	okxKeyList := make([]dbType.OkxKey, size)
	copy(okxKeyList, dbObj.AccountData.OkxKeyList)

	ifThere := false
	for _, val := range okxKeyList {
		if val.ApiKey == okxKey.ApiKey {
			ifThere = true
			break
		}
	}
	if ifThere {
		resErr = fmt.Errorf("ApiKey 已存在")
		db.Close()
		return
	}

	okxKeyList = append(okxKeyList, okxKey)

	FK := bson.D{{
		Key:   "UserID",
		Value: dbObj.UserID,
	}}

	UK := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "OkxKeyList",
					Value: okxKeyList,
				},
			},
		},
	}

	_, err := db.Table.UpdateOne(db.Ctx, FK, UK)
	if err != nil {
		errStr := fmt.Errorf("数据库更新失败 %+v", err)
		global.LogErr(errStr)
		db.Close()
		return
	}

	dbObj.Update()

	return
}
