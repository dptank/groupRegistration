package admin
import (
	"groupRegistration/lib"
	"groupRegistration/controllers"
	"groupRegistration/service/pinActivityService"
)

type ActivityController struct {
	controllers.BaseController
}
/**
根据id获取详情
*/
func (this *ActivityController) Info() {
	id ,_:= this.GetInt("id")
	info := pinActivityService.GetInfo(id)
	if info.Id==0 {
		this.Rsps(true,200,"ok","")
	}
	row := make(map[string]interface{})
	row["id"] = info.Id
	row["img"] = info.Img
	row["title"] = info.Title
	row["countLimit"] = info.CountLimit
	row["ownerPrice"] = info.OwnerPrice
	row["memberPrice"] = info.MemberPrice
	row["startTime"] = info.StartTime
	row["endTime"] = info.EndTime
	row["status"] = info.Status
	row["stock"] = info.Stock
	this.Rsps(true,200,"ok",row)
}
/**
添加活动信息
*/
func (this *ActivityController) Add() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := pinActivityService.AddInfo(data)
	this.Rsps(status,code,msg,res)
}
/**
修改活动信息
*/
func (this *ActivityController) Update() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := pinActivityService.UpdateInfo(data)
	this.Rsps(status,code,msg,res)
}
/**
状态修改
*/
func (this *ActivityController) ChangStatus() {
	id ,_:= this.GetInt("id")
	st ,_:= this.GetInt("status")
	status,code,msg,res := pinActivityService.ChangeInfoStatus(id,st)
	this.Rsps(status,code,msg,res)
}
/**
活动list
*/
func (this *ActivityController) List(){
	pageNum ,_:= this.GetInt("pageNum",1)
	pageSize ,_:= this.GetInt("pageSize",10)
	offset := lib.PageInit(pageNum,pageSize)
	filters := make([]interface{}, 0)
	filters = append(filters,"status__in",[]int{0,1,2})
	status,code,msg,res := pinActivityService.GetInfoList(pageNum,pageSize,offset,filters...)
	this.Rsps(status,code,msg,res)
}