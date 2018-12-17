package pinEvent
import (
	"github.com/astaxie/beego"
)

type EventController struct {
	beego.Controller
}

func (c *EventController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
