package main

import (
	"found/common"
	_ "found/routers"
	"github.com/astaxie/beego"
)

func init() {
	common.MySQLInit()
}

func main() {
	//beego.ErrorHandler("404",page_not_found)
	beego.Run()
}

