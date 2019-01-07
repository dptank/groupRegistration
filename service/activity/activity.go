package activity

import (
	"groupRegistration/models"
	"encoding/json"
	"time"
	"github.com/astaxie/beego/validation"
	"groupRegistration/lib"
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
	row["stock"] = info.Stock
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
	var ob models.PinActivity
	info,err := ob.GetActivityInfo(id)
	if err != nil {
		status := false
		code := 500
		msg := "信息不存在"
		res := ""
		return status ,code, msg,res
	}
	info.Status = st
	res :=  info.UpdateActivityInfo()
	if res != nil {
		status := false
		code := 500
		msg := err.Error()
		res := ""
		return status ,code, msg,res
	}
	status := true
	code := 200
	msg := "修改成功"
	result := ""
	return status ,code, msg,result
}
/**
获取信息列表
*/
func GetInfoList(pageNum int,pageSize int,offset int,filters ...interface{}) ( bool,int,string,interface{}){
	var ob models.PinActivity
	result ,total :=ob.GetActivityList(pageNum,offset,pageSize,filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		//fmt.Println(v.StartTime)
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["joinCount"] = v.JoinCount
		row["ownerPrice"] = v.OwnerPrice
		row["memberPrice"] = v.MemberPrice
		row["ownerPrice"] = v.OwnerPrice
		row["priceType"] = v.PriceType
		row["startTime"] = time.Unix(v.StartTime, 0).Format("2006-01-02 15:04:05")//beego.Date(time.Unix(v.StartTime, 0), "Y-m-d H:i:s")
		row["endTime"] = time.Unix(v.EndTime, 0).Format("2006-01-02 15:04:05")//beego.Date(time.Unix(v.EndTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["img"] = v.Img
		row["stock"] = v.Stock
		list[k] = row
	}
	data := lib.PageUtil(int(total),pageNum,pageSize,list)
	status := true
	code := 200
	msg := "操作成功"
	return status ,code, msg,data
}