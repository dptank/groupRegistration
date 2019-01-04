package service

import (
	"groupRegistration/models"
	"encoding/json"
	"time"
	"github.com/astaxie/beego/validation"
)
/**
获取活动详情
*/
func GetInfo(id int) (res interface{}){
	var ob models.PinActivity
	info,err := ob.GetActivityInfo(id)
	row := make(map[string]interface{})
	if err != nil {
		result := ""
		return result
	}
	row["id"] = info.Id
	row["img"] = info.Img
	row["title"] = info.Title
	row["joinCount"] = info.JoinCount
	row["ownerPrice"] = info.OwnerPrice
	row["memberPrice"] = info.MemberPrice
	row["startTime"] = info.StartTime
	row["endTime"] = info.EndTime
	row["status"] = info.Status
	return row
}
/**
添加活动信息
*/
func AddInfo(data []uint8) ( bool,int,string,interface{}) {
	var ob models.PinActivity
	err := json.Unmarshal(data, &ob)
	if err != nil {
		status := false
		code := 500
		msg := err.Error()
		res := ""
		return status ,code, msg,res
	}
	nowData := time.Now().Local()
	ob.CreatedAt = nowData
	ob.UpdatedAt = nowData
	addRes := ob.PinActivityAdd()
	if addRes != nil {
		status := false
		code := 500
		msg := err.Error()
		res := ""
		return status ,code, msg,res
	}
	status := true
	code := 200
	msg := "增加成功"
	res := ""
	return status ,code, msg,res
}
/**
活动信息修改
*/
func UpdateInfo(data []uint8)  (bool,int,string,interface{}) {
	var ob models.PinActivity
	err := json.Unmarshal(data, &ob)
	if err != nil {
		status := false
		code := 500
		msg := err.Error()
		res := ""
		return status ,code, msg,res
	}
	//检测
	valid := validation.Validation{}
	if v := valid.Required(ob.Id,"id").Message("不能为空！");!v.Ok {
		status := false
		code := 500
		msg := v.Error.Key + " " + v.Error.Message
		res := ""
		return status ,code, msg,res
	}
	result := ob.UpdateActivityInfo()
	if result != nil {
		status := false
		code := 500
		msg := err.Error()
		res := ""
		return status ,code, msg,res
	}
	status := true
	code := 200
	msg := "增加成功"
	res := ""
	return status ,code, msg,res
}
/**
修改状态
*/
func ChangeInfoStatus(id int,st int) ( bool,int,string,interface{}){
	status := true
	code := 200
	msg := "增加成功"
	res := ""
	return status ,code, msg,res
}