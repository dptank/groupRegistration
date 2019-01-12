package pinActivity
import (
	"groupRegistration/controllers"
	"groupRegistration/service/pinActivityService"
	"github.com/astaxie/beego/validation"
)

type EventController struct {
	controllers.BaseController
}
/**
发起拼团
*/
func (this *EventController) AddPinEvent() {
	pinActivityId ,_ := this.GetInt("pinActivityId")
	userId,_ := this.GetInt("userId")
	valid := validation.Validation{}
	valid.Required(userId, "userId")
	valid.Required(pinActivityId, "pinActivityId")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			this.Rsps(false,500,err.Key+" "+err.Message,"")
			//log.Println(err.Key, err.Message)
		}
	}
	status,code,msg,res := pinActivityService.AddPinEvent(pinActivityId,userId)
	this.Rsps(status,code,msg,res)
}
