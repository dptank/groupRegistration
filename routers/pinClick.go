package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/pinClick"
)
/**
前端路由
*/
func pinClickRouter() {
	beego.Router("/", &pinClick.ClickController{},"*:Index")
}