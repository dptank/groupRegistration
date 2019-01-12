package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/pinActivity"
)
/**
前端活动路由
*/
func pinActivityRouter() {
	//活动列表
	beego.Router("/activty/list", &pinActivity.ActivityController{},"POST:ActivityList")
	//拼团
	beego.Router("/activity/event/add", &pinActivity.EventController{},"*:AddPinEvent")
	//打卡
	beego.Router("/activity/click/index", &pinActivity.ClickController{},"*:Index")
	//参团
	beego.Router("/activity/join/index", &pinActivity.JoinController{},"*:Index")
}