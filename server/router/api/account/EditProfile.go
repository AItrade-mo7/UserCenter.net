package account

import (
	"fmt"

	"UserCenter.net/server/global"
	"UserCenter.net/server/global/apiType"
	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"UserCenter.net/server/utils/taskPush"
	"UserCenter.net/sysPublic/dbType"
	"UserCenter.net/sysPublic/dbUser"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

type EditUserType struct {
	Avatar         string
	NickName       string
	EntrapmentCode string
	Password       string
	EmailCode      string
}

func EditProfile(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("编辑用户信息失败", "设备异常"))
	}

	var json EditUserType
	mFiber.Parser(c, &json)

	UserID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(err)))
	}
	defer UserDB.DB.Close()

	isName := mVerify.IsNickName(json.NickName)
	if !isName {
		return c.JSON(result.Fail.WithMsg("昵称不符合规范"))
	}

	// 记录老旧的邮箱

	NickName_edit := json.NickName != UserDB.Data.NickName
	Avatar_edit := len(json.Avatar) > 2 && json.Avatar != UserDB.Data.Avatar
	EntrapmentCode_edit := len(json.EntrapmentCode) > 2 && json.EntrapmentCode != UserDB.Data.EntrapmentCode

	FK := bson.D{{
		Key:   "UserID",
		Value: UserDB.UserID,
	}}
	UK := bson.D{}

	if NickName_edit {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "NickName",
					Value: json.NickName,
				},
			},
		})
	}

	if Avatar_edit {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "Avatar",
					Value: json.Avatar,
				},
			},
		})
	}

	// 需要验证老邮箱验证码 和 密码
	if EntrapmentCode_edit {
		// 密码验证
		err = UserDB.CheckPassword(json.Password)
		if err != nil {
			return c.JSON(result.ErrLogin.WithMsg(err))
		}
		// 验证码验证
		err := taskPush.CheckEmailCode(taskPush.CheckEmailCodeParam{
			Email: UserDB.Data.Email,
			Code:  json.EmailCode,
		})
		if err != nil {
			return c.JSON(result.ErrEmailCode.WithMsg(err))
		}

	}

	if EntrapmentCode_edit {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "EntrapmentCode",
					Value: json.EntrapmentCode,
				},
			},
		})
	}

	if len(UK) < 1 {
		return c.JSON(result.Fail.WithMsg("未作出任何更改"))
	}

	UK = append(UK, bson.E{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "UpdateTime",
				Value: mTime.GetUnixInt64(),
			},
		},
	})

	_, err = UserDB.DB.Table.UpdateOne(UserDB.DB.Ctx, FK, UK)
	if err != nil {
		errStr := fmt.Errorf("account.EditProfile %+v", err)
		global.LogErr(errStr)
		return c.JSON(result.ErrDB.WithData(mStr.ToStr(errStr)))
	}

	// 更新标识
	UserDB.Update()

	if EntrapmentCode_edit {

		// 生成 登录信息
		DeviceInfo := mVerify.DeviceToUA(c.Get("User-Agent"))
		IPInfoList := mVerify.GetIPS(c.IPs())
		var IPInfo mVerify.IPAddressType
		if len(IPInfoList) > 0 {
			IPInfo = IPInfoList[0]
		}
		nowTime := mTime.GetTime()
		LoginInfo := dbType.LoginSucceedType{
			UserID:         UserDB.Data.UserID,
			Email:          UserDB.Data.Email,
			BrowserName:    DeviceInfo.BrowserName,
			OsName:         DeviceInfo.OsName,
			Hostname:       IPInfo.Hostname,
			ISP:            IPInfo.ISP,
			Operators:      IPInfo.Operators,
			CreateTimeUnix: nowTime.TimeUnix,
			CreateTimeStr:  nowTime.TimeStr,
		}
		EmailCont := mStr.Join(
			"<br />",
			"新的防钓鱼码为: 【"+UserDB.Data.EntrapmentCode, "】<br />",
			"<br />",
			"时间: ", LoginInfo.CreateTimeStr, "<br />",
			"地区: ", LoginInfo.ISP, "<br />",
			"运营商: ", LoginInfo.Operators, "<br />",
			"系统: ", LoginInfo.OsName, "<br />",
			"设备: ", LoginInfo.BrowserName, "<br />",
			"IP: ", LoginInfo.Hostname, "<br />",
		)

		taskPush.SysEmail(taskPush.SysEmailOpt{
			To:             UserDB.Data.UserEmail,
			Subject:        "防钓鱼码修改",
			Title:          "防钓鱼码修改提醒",
			Message:        "您刚刚修改了自己的防钓鱼码",
			Content:        EmailCont,
			Description:    "防钓鱼码修改",
			EntrapmentCode: UserDB.Data.EntrapmentCode,
		})
	}

	var userinfoData apiType.UserInfo
	jsonStr := mJson.ToJson(UserDB.Data)
	jsoniter.Unmarshal(jsonStr, &userinfoData)

	taskPush.DelEmailCode(UserDB.Data.Email)

	return c.JSON(result.Succeed.With("修改成功", userinfoData))
}
