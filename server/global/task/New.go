package task

import (
	"fmt"

	"UserCenter.net/server/global/config"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mRes"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

const Source = "UserCenter.net"

type NewTaskOpt struct {
	TaskType    string
	Content     map[string]any
	Description string
}

func NewTask(opt NewTaskOpt) error {
	CreateTime := mTime.GetUnixInt64()
	CreateTimeStr := mTime.UnixFormat(CreateTime)

	NewTaskData := mTask.TaskType{
		TaskID:        mEncrypt.GetUUID(),
		TaskType:      opt.TaskType,
		Source:        Source,
		Description:   opt.Description, // 任务描述
		CreateTime:    CreateTime,
		CreateTimeStr: CreateTimeStr,
		Content:       opt.Content,
	}

	jsonStr := mJson.ToJson(NewTaskData)

	var returnData map[string]any
	jsoniter.Unmarshal(jsonStr, &returnData)

	resData, err := NewReq(ReqOpt{
		Origin: config.SysEnv.MessageBaseUrl,
		Path:   "/api/public/InsertTaskQueue",
		Method: "POST",
		Data:   returnData,
	})
	if err != nil {
		return err
	}

	var resObj mRes.ResType
	jsoniter.Unmarshal(resData, &resObj)

	if resObj.Code != 0 {
		return fmt.Errorf(resObj.Msg)
	}

	return err
}
