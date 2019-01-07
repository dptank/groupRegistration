package admin
import (
	"groupRegistration/lib"
	"groupRegistration/controllers"
	"groupRegistration/service/activity"
)

type ActivityController struct {
	controllers.BaseController
}
/**
根据id获取详情
*/
func (this *ActivityController) Info() {
	id ,_:= this.GetInt("id")
	info := activity.GetInfo(id)
	this.Rsps(true,200,"ok",info)
}
/**
添加活动信息
*/
func (this *ActivityController) Add() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := activity.AddInfo(data)
	this.Rsps(status,code,msg,res)
}
/**
修改活动信息
*/
func (this *ActivityController) Update() {
	data := this.Ctx.Input.RequestBody
	status,code,msg,res := activity.UpdateInfo(data)
	this.Rsps(status,code,msg,res)
}
/**
状态修改
*/
func (this *ActivityController) ChangStatus() {
	id ,_:= this.GetInt("id")
	st ,_:= this.GetInt("status")
	status,code,msg,res := activity.ChangeInfoStatus(id,st)
	this.Rsps(status,code,msg,res)
}
/**
活动list
*/
func (this *ActivityController) List(){
	pageNum ,_:= this.GetInt("pageNum",1)
	pageSize ,_:= this.GetInt("pageSize",10)
	offset := lib.PageInit(pageNum,pageSize)
	status,code,msg,res := activity.GetInfoList(pageNum,pageSize,offset)
	this.Rsps(status,code,msg,res)
}