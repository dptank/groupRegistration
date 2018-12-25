package routers

import (
	"github.com/astaxie/beego"
	"groupRegistration/controllers/admin"
)
/*
管理后台接口路由
*/
func pinAdminRouter() {
	beego.Router("/admin/activty/index", &admin.ActivityController{},"POST:Index")
	beego.Router("/admin/activty/info", &admin.ActivityController{},"POST:Info")
	beego.Router("/admin/activty/add", &admin.ActivityController{},"*:Add")
	beego.Router("/admin/activty/update", &admin.ActivityController{},"*:Update")
	beego.Router("/admin/activty/change/status", &admin.ActivityController{},"*:ChangStatus")
}