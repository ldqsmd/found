package controllers

import (
	"found/models"
	"github.com/astaxie/beego/validation"
)

type NoticeController struct {
	BaseController
}

const (
	NoticeId      = "公告ID"
	NoticeTitle   = "公告标题"
	NoticeContent = "公告内容"
)

func (this *NoticeController) filterParams(mod models.Notice) {
	//表单验证
	valid := validation.Validation{}
	if this.actionName == "Edit" {
		valid.Required(mod.Id, NoticeId).Message(ErrNotEmpty)
	}
	valid.Required(mod.Title, NoticeTitle).Message(ErrNotEmpty)
	valid.Required(mod.Content, NoticeContent).Message(ErrNotEmpty)

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

func (this *NoticeController) List() {
	this.checAdminLogin()
	var obj models.Notice
	this.Data["list"], _ = obj.GetList()
	this.SetTpl(BaseLayoutPage, "notice/list.html")
}

//添加
func (this *NoticeController) Add() {
	this.checAdminLogin()
	switch this.requestMethod {
	case "GET":
		this.SetTpl(BaseLayoutPage, "notice/add.html")
	case "POST":
		var detail models.Notice
		if err := this.ParseForm(&detail); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(detail)
		if err := detail.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}

	}
}

//编辑
func (this *NoticeController) Edit() {
	this.checAdminLogin()
	switch this.requestMethod {
	case "GET":
		var detail models.Notice
		if data, err := detail.GetDetail(this.GetString("noticeId")); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		} else {
			this.Data["data"] = data
			this.SetTpl(BaseLayoutPage, "notice/edit.html")
		}
	case "POST":
		var detail models.Notice
		if err := this.ParseForm(&detail); err != nil {
			this.ReturnJson(FilterErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(detail)
		if err := detail.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}
	}
}

//删除
func (this *NoticeController) Del() {

	this.checAdminLogin()
	var obj models.Notice
	if err := this.ParseForm(&obj); err != nil {
		this.ReturnJson(BindErrCode, err.Error(), nil)
	}
	if obj.Id == 0 {
		this.ReturnJson(FilterErrCode, ItemId+ErrNotEmpty, nil)
	}
	if err := obj.Delete(); err != nil {
		this.ReturnJson(DelErrCode, err.Error(), nil)
	}
	this.ReturnJson(SuccessCode, SuccessMessage, nil)
}

//>>>>>>>>>>>>>>>>home

//列表
func (this *NoticeController) HomeList() {
	var obj models.Notice
	pageSize := 5
	page, _ := this.GetInt("page")
	count, list, _ := obj.HomeGetList(pageSize, pageSize*(page-1))
	this.Data["page"] = PageUtil(int(count), page, pageSize, list)
	this.SetTpl(HomeBaseLayout, HomeTplPath, "/notice/list.html")
}

//详情
func (this *NoticeController) HomeDetail() {
	if id := this.GetString("noticeId"); id == "" {
		this.Abort(NotFundCode)
	} else {
		var item models.Item
		var obj models.Notice
		data, err := obj.GetDetail(id)
		if err != nil {
			this.Abort(NotFundCode)
		}
		this.Data["data"] = data
		this.Data["latestList"], _ = item.GetNewItem(5)
		this.Data["latestNotice"], _ = obj.GetTheLatest(5)
		this.SetTpl(HomeBaseLayout, HomeTplPath, "/notice/detail.html")
	}
}
