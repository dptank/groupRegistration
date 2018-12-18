package admin

import "github.com/astaxie/beego"

type baseController struct {
	beego.Controller
}
func (c *baseController) Rsp(status bool, str string) {
	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	c.ServeJSON()
	c.StopRun()
}