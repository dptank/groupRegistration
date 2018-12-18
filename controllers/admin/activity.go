package admin
import (
	//"github.com/astaxie/beego"
	"groupRegistration/models"
	"encoding/json"
	"time"
	"fmt"
)

type ActivityController struct {
	baseController
}

func (this *ActivityController) Index() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 13
	c := make(map[string]bool)
	c["code"]=true

	this.Data["json"] = &m
	//this.Data["Email"] = "astaxie@gmail.com"
	this.ServeJSON()
	//c.TplName = "index.tpl"
}
func (this *ActivityController) Add() {
	var ob models.PinActivity
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	fmt.Println(err)
	if err != nil {
		status := false
		message := "数据格式错误"
		this.Rsp(status,message)
	}
	nowData := time.Now().Local()
	ob.CreatedAt = nowData
	ob.UpdatedAt = nowData
	fmt.Println(ob)
	res := ob.PinActivityAdd()
	if res != nil {
		status := false
		message := "数据格式错误"
		this.Rsp(status,message)
	}
	status := true
	message := "增加成功"
	this.Rsp(status,message)

	//json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	//res := ob.PinActivityAdd()
	
	//nowData := time.Now().Local()
	//pinActivity := new(models.PinActivity)
	//pinActivity.Title = this.GetString("title")
	//pinActivity.Join_count,_ = this.GetInt("joinCount")
	//pinActivity.Member_price,_ = this.GetInt("memberPrice")
	//pinActivity.Owner_price,_ = this.GetInt("ownerPrice")
	//pinActivity.Price_type,_ = this.GetInt("priceType")
	//pinActivity.Start_time,_ = this.GetInt64("startTime")
	//pinActivity.End_time,_ = this.GetInt64("endTime")
	//pinActivity.Img = this.GetString("img")
	//pinActivity.Created_at = nowData
	//pinActivity.Updated_at = nowData
	//err := pinActivity.PinActivityAdd()
	if err!=nil {
		status := false
		message := err.Error()
		this.Rsp(status,message)
	} else {
		status := true
		message := "增加成功"
		this.Rsp(status,message)
	}

}

func (this *ActivityController) Update() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 13

	this.Data["json"] = &m
	this.ServeJSON()
}


//func (c *ActivityController) Rsp(status bool, str string) {
//	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
//	c.ServeJSON()
//	c.StopRun()
//}