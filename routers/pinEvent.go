package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/pinEvent"
)
/**
前端路由
*/
func pinEventRouter() {
	beego.Router("/", &pinEvent.EventController{},"*:Index")
}