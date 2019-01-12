package pinActivityService

import (
	"groupRegistration/models"
	"time"
)

/**
发起拼团
*/
func AddPinEvent(pinActivityId int,userId int) ( bool,int,string,interface{}) {
	activityInfo := GetInfo(pinActivityId)
	st ,message := CheckActivityInfo(activityInfo)
	if st!=true {
		code := 500
		res := ""
		return st ,code, message,res
	}
	//开团的数据
	var ob models.PinEvent
	ob.PinActivityId = pinActivityId
	ob.JoinCount = 1
	ob.CountLimit = activityInfo.CountLimit
	ob.EndTime = activityInfo.EndTime
	ob.Owner = userId
	ob.Status = 1
	nowData := time.Now().Local()
	ob.CreatedAt = nowData
	ob.UpdatedAt = nowData
	//入库
	addRes := ob.AddPinEvent()
	if addRes != nil {
		status := false
		code := 500
		msg := "开团失败"
		res := ""
		return status ,code, msg,res
	}
	status := true
	code := 200
	msg := "开团成功"
	res := ""
	return status ,code, msg,res
}