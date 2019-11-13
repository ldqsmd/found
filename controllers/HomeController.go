package controllers

import (
	"fmt"
	"found/models"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {

	var item models.Item
	this.Data["newList"], _ = item.GetNewItem(8)
	this.SetTpl(HomeBaseLayout, "home/index.html")
}

func (this *HomeController) Page404() {
	this.Data["returnUrl"] = this.GetString("returnUrl")

	fmt.Println(this.GetString("returnUrl"))
	this.TplName = "error/404.html"
}
