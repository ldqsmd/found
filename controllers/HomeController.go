package controllers

import "fmt"

type HomeController struct {
	BaseController
}

const (
	Hmo = ""
)

func (this *HomeController) Index() {
	this.SetTpl(HomeBaseLayout, "home/index.html")
}

func (this *HomeController) Page404() {
	this.Data["returnUrl"] = this.GetString("returnUrl")

	fmt.Println(this.GetString("returnUrl"))
	this.TplName = "error/404.html"
}
