package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/pinActivity"
)
/**
前端活动路由
*/
func pinActivityRouter() {
	beego.Router("/activty/list", &pinActivity.ActivityController{},"POST:ActivityList")
}