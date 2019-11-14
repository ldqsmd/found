package controllers

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	this.SetTpl("", "home/404.html")
}

func (this *ErrorController) Error403() {
	this.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	this.SetTpl(BaseLayoutPage, "404.html")
}
