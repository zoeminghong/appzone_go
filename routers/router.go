package routers

import (
	"appzone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/getjson", &controllers.IndexController{})
	//beego.Router("/getjson",&controllers.MainController{},"get:Gett")
}
