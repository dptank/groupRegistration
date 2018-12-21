package admin

import "github.com/astaxie/beego"

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type baseController struct {
	beego.Controller
}
//返回
func (c *baseController) Rsp(res bool, str string) {
	c.Data["json"] = &map[string]interface{}{"success": res, "info": str}
	c.ServeJSON()
	c.StopRun()
}
//ajax返回 列表
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