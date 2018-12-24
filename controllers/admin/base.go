package admin

import "github.com/astaxie/beego"

type baseController struct {
	beego.Controller
}
//返回值
func (self *baseController) Rsps(res bool, code int, msg string, data interface{}) {
	out := make(map[string]interface{})
	errmsg := make(map[string]interface{})
	errmsg["code"] = code
	errmsg["msg"] = msg
	out["err"] = errmsg
	out["success"] = res
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}