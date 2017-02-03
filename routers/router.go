package routers

import (
	"BetterClassViewer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
}
