package routers

import (
	"shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/home",&)
}
