package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
}
//返回值
func (self *BaseController) Rsps(res bool, code int, msg string, data interface{}) {
	out := make(map[string]interface{})
	errmsg := make(map[string]interface{})
	errmsg["code"] = code
	errmsg["msg"] = msg
	out["err"] = errmsg
	out["success"] = res
	out["data"] = data
	fmt.Println(out)
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}