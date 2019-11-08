package controllers

import "fmt"

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.SetTpl()
}

func (this *HomeController)Page404() {
	this.Data["returnUrl"] =    this.GetString("returnUrl")

	fmt.Println( this.GetString("returnUrl"))
	this.TplName = "error/404.html"
}
