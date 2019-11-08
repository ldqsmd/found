package controllers

import (
	"found/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)Login() {

	switch this.Ctx.Request.Method  {
		case "GET":
			this.TplName = "login/login.html"
		case "POST":
			var admin models.Admin

			if err := this.ParseForm(&admin); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//表单验证
			valid := validation.Validation{}
			valid.Required(admin.Account, "账号").Message("不能为空")
			valid.Required(admin.Password, "密码").Message("不能为空")

			if valid.HasErrors() {
				for _, err := range valid.Errors {
					this.ReturnJson(-1,err.Key+err.Message,nil)
				}
			}

			//判断用户名密码
			err :=  admin.GetAdminInfo()
			if err != nil{
				this.ReturnJson(-1001,"登录失败,账号或密码错误！",nil)
			}else if admin.Status == 1{
				this.ReturnJson(-1002,"账号已被禁用,请联系管理员！",nil)
			}
			//设置session信息
			this.SetSession("adminInfo",admin)
			this.ReturnJson(0,"登录成功",nil)
		}
}

func (this *LoginController)Logout() {
	this.DestroySession()
	this.Redirect(this.URLFor("LoginController.Login"),302)
	this.StopRun()
}

func (this *LoginController) ReturnJson(code int,message string,data interface{}) {

	var  json models.ReturnJson
	json.Code	 	= code
	json.Message 	= message
	json.Data 		= data
	this.Data["json"] 	= &json
	this.ServeJSON()
	this.StopRun()
}