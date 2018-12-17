package pinActivity
import (
"github.com/astaxie/beego"
)

type ActivityController struct {
	beego.Controller
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
