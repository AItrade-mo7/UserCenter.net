package dbUser

import (
	"UserCenter.net/server/global/dbType"
	"go.mongodb.org/mongo-driver/bson"
)

func (dbObj *AccountType) Update() {
	db := dbObj.DB
	var result dbType.AccountTable
	FK := bson.D{{
		Key:   "UserID",
		Value: dbObj.UserID,
	}}
	db.Table.FindOne(db.Ctx, FK).Decode(&result)
	dbObj.AccountData = result
}
