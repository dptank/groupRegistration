package pinActivity
import (
	"github.com/astaxie/beego"
)

type JoinController struct {
	beego.Controller
}

func (c *JoinController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
