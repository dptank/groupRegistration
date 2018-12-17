package pinClick
import (
	"github.com/astaxie/beego"
)

type ClickController struct {
	beego.Controller
}

func (c *ClickController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
