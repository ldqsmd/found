package main

import (
	"found/common"
	"found/controllers"
	_ "found/routers"
	"github.com/astaxie/beego"
)

func init() {
	common.MySQLInit()
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
