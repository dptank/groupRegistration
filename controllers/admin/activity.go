package admin
import (
	"groupRegistration/models"
	"encoding/json"
	"time"
	"fmt"
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
	if err != nil {
		status := false
		this.Rsp(status,err.Error())
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
		this.Rsp(status,err.Error())
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

func (this *ActivityController) Update() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 13

	this.Data["json"] = &m
	this.ServeJSON()
}
