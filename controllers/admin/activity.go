package admin
import (
	"groupRegistration/models"
	"encoding/json"
	"time"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ActivityController struct {
	baseController
}
/**
根据id获取详情
*/
func (this *ActivityController) Info() {
	var ob models.PinActivity
	id ,_:= this.GetInt("id")
	info,err := ob.GetActivityInfo(id)
	//fmt.Println(info)
	if err != nil {
		status := true
		this.Rsps(status,200,"暂无数据","")
	}
	row := make(map[string]interface{})
	row["id"] = info.Id
	row["img"] = info.Img
	row["title"] = info.Title
	row["joinCount"] = info.JoinCount
	row["ownerPrice"] = info.OwnerPrice
	row["memberPrice"] = info.MemberPrice
	row["startTime"] = info.StartTime
	row["endTime"] = info.EndTime
	row["status"] = info.Status
	this.Rsps(true,200,"ok",row)
}
/**
添加活动信息
*/
func (this *ActivityController) Add() {
	var ob models.PinActivity
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		status := false
		this.Rsps(status,500,err.Error(),"")
	}
	nowData := time.Now().Local()
	ob.CreatedAt = nowData
	ob.UpdatedAt = nowData
	res := ob.PinActivityAdd()
	fmt.Println(res)
	if res != nil {
		status := false
		this.Rsps(status,500,res.Error(),"")
	}
	status := true
	this.Rsps(status,200,"增加成功","")
}
/**
修改活动信息
*/
func (this *ActivityController) Update() {
	var ob models.PinActivity
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		status := false
		this.Rsps(status,500,err.Error(),"")
	}
	//检测
	valid := validation.Validation{}
	if v := valid.Required(ob.Id,"id").Message("不能为空！");!v.Ok {
		status := false
		errInfo := v.Error.Key + " " + v.Error.Message
		this.Rsps(status,500,errInfo,"")
	}
	//fmt.Println(ob.Title)
	//info,_ := ob.GetActivityInfo(ob.Id)
	//if info == nil {
	//	status := false
	//	this.Rsps(status,500,"不存在该id","")
	//}
	//修改
	res := ob.UpdateActivityInfo()
	fmt.Println(res)
	if res != nil {
		status := false
		this.Rsps(status,500,res.Error(),"")
	}
	status := true
	this.Rsps(status,200,"修改成功","")
}
/**
状态修改
*/
func (this *ActivityController) ChangStatus() {
	var ob models.PinActivity
	id ,_:= this.GetInt("id")
	status ,_:= this.GetInt("status")
	info,err := ob.GetActivityInfo(id)
	if err != nil {
		res := false
		this.Rsps(res,500,"暂无数据","")
	}

	info.Status = status
	res :=  info.UpdateActivityInfo()
	if res != nil {
		status := false
		this.Rsps(status,500,res.Error(),"")
	}
	result := true
	this.Rsps(result,200,"修改成功","")
}
/**
活动list
*/
func (this *ActivityController) Index(){
	pageNum ,_:= this.GetInt("pageNum")
	pageSize ,_:= this.GetInt("pageSize")

	this.Rsps(true,200,"修改成功","")
}