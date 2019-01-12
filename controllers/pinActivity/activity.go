package pinActivity
import (
	"groupRegistration/controllers"
	"groupRegistration/lib"
	"groupRegistration/service/pinActivityService"
)

type ActivityController struct {
	controllers.BaseController
}
/**
手机端活动列表页
*/
func (this *ActivityController) ActivityList() {
	pageNum ,_:= this.GetInt("pageNum",1)
	pageSize ,_:= this.GetInt("pageSize",10)
	offset := lib.PageInit(pageNum,pageSize)
	filters := make([]interface{}, 0)
	filters = append(filters,"status",1)
	status,code,msg,res := pinActivityService.GetInfoList(pageNum,pageSize,offset,filters...)
	this.Rsps(status,code,msg,res)
}

