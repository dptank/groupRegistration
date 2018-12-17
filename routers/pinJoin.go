package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/pinJoin"
)
/**
前端路由
*/
func pinJoinRouter() {
	beego.Router("/", &pinJoin.JoinController{},"*:Index")
}