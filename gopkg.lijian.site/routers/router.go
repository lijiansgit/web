package routers

import (
	"github.com/astaxie/beego"
	"gopkg.lijian.site/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})

	//pkgdoc
	beego.Router("/pkgdoc", &controllers.PkgDoc{})
}
