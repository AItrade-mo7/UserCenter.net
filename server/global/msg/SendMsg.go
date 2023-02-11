package msg

// import (
// 	"fmt"

// 	"UserCenter.net/server/global/config"
// 	"github.com/EasyGolang/goTools/mEncrypt"
// 	"github.com/EasyGolang/goTools/mJson"
// 	"github.com/EasyGolang/goTools/mTask"
// 	jsoniter "github.com/json-iterator/go"
// )

// func SendMsg() {
// 	Data := NewTask()

// 	resData, resErr := NewReq(ReqOpt{
// 		Origin: config.SysEnv.MessageBaseUrl,
// 		Path:   "/api/public/InsertTaskQueue",
// 		Method: "POST",
// 		Data:   Data,
// 	})

// 	fmt.Println(mJson.JsonFormat(resData), resErr)
// }

// func NewTask() map[string]any {
// 	NewTaskData := mTask.TaskType{
// 		TaskID:   mEncrypt.GetUUID(),
// 		TaskType: "SendEmail",
// 		Source:   "",
// 	}

// 	jsonStr := mJson.ToJson(NewTaskData)

// 	var returnData map[string]any

// 	jsoniter.Unmarshal(jsonStr, &returnData)

// 	return returnData
// }
