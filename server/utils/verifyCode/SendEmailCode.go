package verifyCode

/*
	type SendCodeParam struct {
		Email        string
		Action       string
		SecurityCode string
	}

	func CheckAndSendCode(opt SendCodeParam) (resErr error) {
		resErr = nil

		isEmail := mVerify.IsEmail(opt.Email)
		if !isEmail {
			resErr = fmt.Errorf("email 格式不正确 %+v", opt.Email)
			return
		}

		db := mMongo.New(mMongo.Opt{
			UserName: config.SysEnv.MongoUserName,
			Password: config.SysEnv.MongoPassword,
			Address:  config.SysEnv.MongoAddress,
			DBName:   "AItrade",
		}).Connect().Collection("EmailCode")

		err := db.Ping()
		if err != nil {
			db.Close()
			resErr = fmt.Errorf("EmailCode 数据库连接错误 %+v", err)
			global.LogErr(resErr)
			return
		}

		// 先进行查找

		var result dbType.EmailCodeTable
		findOptions := options.FindOne()
		findOptions.SetSort(bson.D{
			{
				Key:   "_id",
				Value: -1,
			},
		})

		FK := bson.D{{
			Key:   "Email",
			Value: opt.Email,
		}}

		db.Table.FindOne(db.Ctx, FK, findOptions).Decode(&result)

		if len(result.Email) < 1 {
			// 没找到数据,发送验证码并存储
			vCode, err := SendCode(SendCodeParam{
				Email:        opt.Email,
				Action:       opt.Action,
				SecurityCode: opt.SecurityCode,
			})
			if err != nil {
				db.Close()
				resErr = err
				return
			}

			var WK dbType.EmailCodeTable
			WK.Email = opt.Email
			WK.Code = vCode
			WK.SendTime = mTime.GetUnixInt64()

			_, err = db.Table.InsertOne(db.Ctx, WK)
			if err != nil {
				db.Close()
				resErr = fmt.Errorf("EmailCode,插入数据失败 %+v", err)
				global.LogErr(resErr)
				return
			}
		} else {
			// 找到了数据,则对比时间

			sendTime := mStr.ToStr(result.SendTime)
			nowTime := mTime.GetUnix()
			subStr := mCount.Sub(nowTime, sendTime)

			// 时间不符合则不发送,直接返回频繁提示
			if mCount.Le(subStr, mCount.Mul(mTime.UnixTime.Minute, "10")) < 0 {
				db.Close()
				resErr = fmt.Errorf("刚刚的验证码还能再用哦~")
				return
			}

			// 时间符合,则发送验证码并更新
			vCode, err := SendCode(SendCodeParam{
				Email:        opt.Email,
				Action:       opt.Action,
				SecurityCode: opt.SecurityCode,
			})
			if err != nil {
				db.Close()
				resErr = err
				return
			}

			UK := bson.D{
				{
					Key: "$set",
					Value: bson.D{
						{
							Key:   "Code",
							Value: vCode,
						},
					},
				},
				{
					Key: "$set",
					Value: bson.D{
						{
							Key:   "SendTime",
							Value: mTime.GetUnixInt64(),
						},
					},
				},
			}

			_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
			if err != nil {
				db.Close()
				resErr = fmt.Errorf("数据库更新失败 %+v", err)
				global.LogErr(resErr)
				return
			}
		}

		db.Close()

		return
	}
func SendCode(opt SendCodeParam) (vCode string, resErr error) {
	vCode = ""
	resErr = nil

	SecurityCode := "trade.mo7.cc-System"

	if len(opt.SecurityCode) > 0 {
		SecurityCode = opt.SecurityCode
	}

	vCode = mVerify.NewCode()

	Email := global.Email(global.EmailOpt{
		To: []string{
			opt.Email,
		},
		Subject:  "请查收您的验证码",
		Template: tmpl.CodeEmail,
		SendData: tmpl.CodeParam{
			SysTime:      mTime.UnixFormat(""),
			Action:       opt.Action,
			VerifyCode:   vCode,
			SecurityCode: SecurityCode,
		},
	})

	resErr = Email.Send()

	return
}
*/
