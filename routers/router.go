package routers
func init() {
	pinActivityRouter()//活动接口路由
	pinEventRouter()//参团事件接口路由
	pinClickRouter()//打卡接口路由
	pinJoinRouter()//参与接口路由
	pinAdminRouter()//管理后台接口路由
}
