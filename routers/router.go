package routers

import (
	"found/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//home
	beego.Router("/", &controllers.HomeController{}, "GET:Index")
	//item
	beego.Router("/home/item/list", &controllers.ItemController{}, "Get:HomeList")
	beego.Router("/home/item/detail", &controllers.ItemController{}, "Get:HomeItemDetail")
	beego.Router("/home/item/add", &controllers.ItemController{}, "*:HomeItemAdd")
	beego.Router("/home/user/item/del", &controllers.ItemController{}, "Post:HomeDel")
	beego.Router("/home/user/item/list", &controllers.ItemController{}, "*:UserItemList")
	beego.Router("/home/user/item/status", &controllers.ItemController{}, "Post:HomeChangeStatus")
	//login
	beego.Router("/login", &controllers.LoginController{}, "*:HomeLogin")
	beego.Router("/home/logout", &controllers.LoginController{}, "*:HomeLogout")
	beego.Router("/home/register", &controllers.LoginController{}, "*:HomeRegister")
	//user
	beego.Router("/home/user/info", &controllers.UserController{}, "*:UserInfo")

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	beego.Router("/admin/index", &controllers.ItemController{}, "GET:Index")
	//login
	beego.Router("/admin/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")

	//item
	beego.Router("/item/list", &controllers.ItemController{}, "GET:List")
	beego.Router("/item/add", &controllers.ItemController{}, "*:Add")
	beego.Router("/item/edit", &controllers.ItemController{}, "*:Edit")
	beego.Router("/item/del", &controllers.ItemController{}, "*:Del")
	beego.Router("/item/status", &controllers.ItemController{}, "POST:ChangeStatus")

	//admin
	beego.Router("/admin/list", &controllers.AdminController{}, "GET:AdminList")
	beego.Router("/admin/add", &controllers.AdminController{}, "*:AddAdmin")
	beego.Router("/admin/edit", &controllers.AdminController{}, "*:EditAdmin")
	beego.Router("/admin/forbid", &controllers.AdminController{}, "POST:ForbidAdmin")

	//user
	beego.Router("/user/list", &controllers.UserController{}, "GET:List")
	beego.Router("/user/add", &controllers.UserController{}, "*:Add")
	beego.Router("/user/edit", &controllers.UserController{}, "*:Edit")
	beego.Router("/user/forbid", &controllers.UserController{}, "POST:ForbidUser")

}
