package main

import (
	_ "shop/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/"] = "static"
	//beego.RunMode = "prod"
	beego.Run()
}

