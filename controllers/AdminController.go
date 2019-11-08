package controllers

import (
	"found/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type AdminController struct {
	BaseController
}

func (this *AdminController)AdminList() {

	var admin models.Admin
	list,total := admin.GetAdminList()
	this.Data["list"]  = list
	this.Data["total"] = total
	this.SetTpl("base/layout_page.html","admin/list.html")
}


func (this *AdminController)filterParams(admin models.Admin)  {
	//表单验证
	valid := validation.Validation{}

	if this.requestMethod == "AddAdmin"{
		valid.Required(admin.Account, "账号").Message("不能为空")
		valid.Required(admin.Password, "密码").Message("不能为空")
		valid.Required(admin.Email, "e-mail").Message("不能为空")

	}
	if this.actionName == "EditAdmin"{
		//编辑的时候 必填 adminId
		valid.Required(admin.Id, "管理员ID").Message("不能为空")
		if this.requestMethod == "POST"{
			valid.Required(admin.Account, "账号").Message("不能为空")
			valid.Required(admin.Email, "e-mail").Message("不能为空")
		}
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {

			if this.IsAjax(){
				this.ReturnJson(-1,err.Key+err.Message,nil)
			}else{
				this.Data["error"] = err.Key+err.Message
				this.Abort("404")
			}
		}
	}

}

//添加管理员
func (this *AdminController)AddAdmin() {

	switch this.requestMethod {
		case "GET":
			this.SetTpl("base/layout_page.html","admin/add.html")
		case "POST":
			var admin models.Admin
			if err := this.ParseForm(&admin); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(admin)

			if  err := admin.InsertOrUpdate() ; err != nil{
				this.ReturnJson(-1,err.Error(),nil)
			}else{
				this.ReturnJson(0,"添加成功",nil)
			}

	}
}

//编辑管理员
func (this *AdminController)EditAdmin() {

	switch this.requestMethod {
		case "GET":
			var  admin models.Admin
			if adminInfo :=  admin.GetAdminDetail(this.GetString("adminId")) ;adminInfo.Id ==0{
				this.ReturnJson(-2,"信息不存在",nil)
			}else{
				this.Data["adminInfo"] = adminInfo
				this.SetTpl("base/layout_page.html","admin/edit.html")
			}
	case "POST":
			var admin models.Admin
			if err := this.ParseForm(&admin); err != nil {
				this.ReturnJson(-1,err.Error(),nil)
			}
			//校验必填参数
			this.filterParams(admin)
			if err := admin.InsertOrUpdate() ;err != nil{
				this.ReturnJson(-2,err.Error(),nil)
			}else{
				this.ReturnJson(0,"修改成功",nil)
			}
	}
}

//软删除
func (this *AdminController)ForbidAdmin() {

	var admin models.Admin
	//校验必填参数
	if  adminId := this.GetString("adminId") ; adminId == ""{
		this.ReturnJson(-1,"管理员ID",nil)
	}else{
		intId,_ := strconv.Atoi(adminId)
		if err := admin.ForbidAdmin(intId) ;err != nil{
			this.ReturnJson(-1,err.Error(),nil)
		}else{
			this.ReturnJson(0,"禁用成功",nil)
		}
	}

}
