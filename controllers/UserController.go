package controllers

import (
	"found/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type UserController struct {
	BaseController
}

const (
	UserPkId      = "用户ID"
	UserAccount   = "账号"
	UserForbidden = "状态"
	UserPassword  = "密码"
	UserName      = "用户名"
	Email         = "e-mail"
	Phone         = "电话"
)

func (this *UserController) filterParams(mod models.User) {
	//表单验证
	valid := validation.Validation{}
	if this.actionName == "Edit" {
		valid.Required(mod.Id, UserPkId).Message(ErrNotEmpty)
	}
	if this.actionName == "Add" {
		valid.Required(mod.Password, UserPassword).Message(ErrNotEmpty)
	}
	valid.Required(mod.Name, UserName).Message(ErrNotEmpty)
	valid.Required(mod.Email, Email).Message(ErrNotEmpty)

	if valid.HasErrors() {
		for _, err := range valid.Errors {

			if this.IsAjax() {
				this.ReturnJson(FilterErrCode, err.Key+err.Message, nil)
			} else {
				this.Data["error"] = err.Key + err.Message
				this.Abort(NotFundCode)
			}
		}
	}
}

func (this *UserController) List() {

	var user models.User
	this.Data["list"], _ = user.GetList()
	this.SetTpl(BaseLayoutPage, "user/list.html")
}

//添加管理员
func (this *UserController) Add() {

	switch this.requestMethod {
	case "GET":
		this.SetTpl(BaseLayoutPage, "user/add.html")
	case "POST":
		var user models.User
		if err := this.ParseForm(&user); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(user)
		if err := user.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}

	}
}

//编辑管理员
func (this *UserController) Edit() {

	switch this.requestMethod {
	case "GET":
		var user models.User
		if userInfo, err := user.GetDetail(this.GetString("userId")); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		} else {
			this.Data["userInfo"] = userInfo
			this.SetTpl(BaseLayoutPage, "user/edit.html")
		}
	case "POST":
		var user models.User
		if err := this.ParseForm(&user); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(user)
		if err := user.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}
	}
}

//软删除
func (this *UserController) ForbidUser() {

	var user models.User
	forbidden := this.GetString("forbidden")
	//校验必填参数
	if forbidden == "" {
		this.ReturnJson(FilterErrCode, UserForbidden+ErrNotEmpty, nil)
	}
	userId := this.GetString("userId")
	if userId == "" {
		this.ReturnJson(FilterErrCode, UserPkId+ErrNotEmpty, nil)
	}
	intId, _ := strconv.Atoi(userId)
	intForbid, _ := strconv.Atoi(forbidden)
	if err := user.ForbidUser(intId, intForbid); err != nil {
		this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
	} else {
		this.ReturnJson(SuccessCode, SuccessMessage, nil)
	}

}

//>>>>>>>>>>>>>>>>home

//编辑管理员
func (this *UserController) UserInfo() {

	this.checkHomeLogin()
	switch this.requestMethod {
	case "GET":
		this.Data["userInfo"] = this.userInfo
		this.SetTpl(HomeBaseLayout, HomeTplPath, "/user/userinfo.html")
	case "POST":
		var user models.User
		if err := this.ParseForm(&user); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		}
		//校验必填参数
		valid := validation.Validation{}
		valid.Required(user.Id, UserPkId).Message(ErrNotEmpty)
		valid.Required(user.Name, UserName).Message(ErrNotEmpty)
		valid.Required(user.Email, Email).Message(ErrNotEmpty)
		if valid.HasErrors() {
			for _, err := range valid.Errors {

				if this.IsAjax() {
					this.ReturnJson(FilterErrCode, err.Key+err.Message, nil)
				} else {
					this.Data["error"] = err.Key + err.Message
					this.Abort(NotFundCode)
				}
			}
		}

		if err := user.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.SetSession("userInfo", user)
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}
	}
}
