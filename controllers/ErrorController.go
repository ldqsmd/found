package controllers

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	this.SetTpl(HomeBaseLayout, "home/404.html")
}
