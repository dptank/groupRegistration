package pinActivityService

import (
	"groupRegistration/models"
	"encoding/json"
	"time"
	"github.com/astaxie/beego/validation"
	"groupRegistration/lib"
)
type ActivityInfo struct {
	Id int
	Title string
	CountLimit int
	OwnerPrice int
	MemberPrice int
	PriceType int
	StartTime int64
	EndTime int64
	Status int
	Img string
	Stock int
}
/**
获取活动详情
*/
func GetInfo(id int) (res ActivityInfo){
	var ob models.PinActivity
	var result ActivityInfo
	info,err := ob.GetActivityInfo(id)
	if err != nil {
		return result
	}
	result.Id = info.Id
	result.Img = info.Img
	result.Title = info.Title
	result.CountLimit = info.CountLimit
	result.OwnerPrice = info.OwnerPrice
	result.MemberPrice = info.MemberPrice
	result.StartTime = info.StartTime
	result.EndTime = info.EndTime
	result.Stock = info.Stock
	result.Status = info.Status
	//fmt.Println(result)
	return result
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
		row["countLimit"] = v.CountLimit
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
/**
检查活动信息
*/
func CheckActivityInfo(res ActivityInfo) ( bool,string) {
	if res.Id == 0 {
		status := false
		msg := "活动不存在"
		return status , msg
	}
	if res.Status == 2{
		status := false
		msg := "活动以下线"
		return status ,msg
	}
	if res.Status == 0{
		status := false
		msg := "活动还没上线"
		return status ,msg
	}
	currentTime := time.Now().Unix()
	if res.StartTime>currentTime {
		status := false
		msg := "活动还没上线"
		return status ,msg
	}
	if res.EndTime<currentTime {
		status := false
		msg := "活动时间以过"
		return status ,msg
	}
	status := true
	msg := "活动在线"
	return status,msg
}