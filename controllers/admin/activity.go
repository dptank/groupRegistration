package admin
import (
	"groupRegistration/models"
	"time"
	"groupRegistration/lib"
	"github.com/astaxie/beego"
	"groupRegistration/service"
)

type ActivityController struct {
	baseController
}
/**
根据id获取详情
*/
func (this *ActivityController) Info() {
	id ,_:= this.GetInt("id")
	info := service.GetInfo(id)
	this.Rsps(true,200,"ok",info)
}
/**
添加活动信息
*/
func (this *ActivityController) Add() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := service.AddInfo(data)
	this.Rsps(status,code,msg,res)
}
/**
修改活动信息
*/
func (this *ActivityController) Update() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := service.UpdateInfo(data)
	this.Rsps(status,code,msg,res)
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
	var ob models.PinActivity
	pageNum ,_:= this.GetInt("pageNum",1)
	pageSize ,_:= this.GetInt("pageSize",10)
	offset := lib.PageInit(pageNum,pageSize)
	result ,total :=ob.GetActivityList(pageNum,offset,pageSize)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["joinCount"] = v.JoinCount
		row["ownerPrice"] = v.OwnerPrice
		row["memberPrice"] = v.MemberPrice
		row["ownerPrice"] = v.OwnerPrice
		row["priceType"] = v.PriceType
		row["startTime"] = beego.Date(time.Unix(v.StartTime, 0), "Y-m-d H:i:s")
		row["endTime"] = beego.Date(time.Unix(v.EndTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["img"] = v.Img
		list[k] = row
	}

	data := lib.PageUtil(int(total),pageNum,pageSize,list)

	this.Rsps(true,200,"修改成功",data)
}