package controllers

import (
	"fmt"
	"found/models"
	"github.com/astaxie/beego/validation"
	"strconv"
)

const (
	ItemId     = "物品ID"
	UserId     = "用户ID"
	ItemName   = "物品名称"
	ItemType   = "物品类型"
	ItemTime   = "时间"
	ItemPlace  = "地点"
	ItemTitle  = "标题"
	ItemImage  = "物品图片"
	ItemStatus = "物品特点"
	ItemRemark = "详细描述"

	ActList = "list"
	ActAdd  = "add"
	ActEdit = "edit"

	UploadFileName = "itemPic"
)

type ItemController struct {
	BaseController
}

//参数过滤
func (this *ItemController) filterParams(params *models.Item) {

	//表单验证
	valid := validation.Validation{}
	if this.actionName == "Edit" {
		valid.Required(params.Id, ItemId).Message(ErrNotEmpty)
	}
	valid.Required(params.Title, ItemTitle).Message(ErrNotEmpty)
	valid.Required(params.UserId, UserId).Message(ErrNotEmpty)
	valid.Required(params.Name, ItemName).Message(ErrNotEmpty)
	valid.Required(this.GetString("type"), ItemType).Message(ErrNotEmpty)
	valid.Required(params.Time, ItemTime).Message(ErrNotEmpty)
	valid.Required(params.Place, ItemPlace).Message(ErrNotEmpty)
	valid.Required(params.Remark, ItemRemark).Message(ErrNotEmpty)

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			if this.IsAjax() {
				this.ReturnJson(FilterErrCode, err.Key+err.Message, nil)
			} else {
				this.Data["error"] = err.Key + err.Message
				this.Abort(ForbidCode)
			}
		}
	}
	if _, _, err := this.GetFile(UploadFileName); err == nil {
		filePath, err := this.UpFileTable(UploadFileName, UploadPicture)
		if err != nil {
			this.ReturnJson(FilterErrCode, UploadFileName+":"+err.Error(), nil)
		} else {
			params.Image = filePath
		}
	} else if this.actionName == "Add" || this.actionName == "HomeItemAdd" {
		this.ReturnJson(FilterErrCode, UploadFileName+":"+err.Error(), nil)
	}

}

func (this *ItemController) Index() {
	this.checAdminLogin()
	this.SetTpl("index.html")
}

//列表
func (this *ItemController) List() {
	this.checAdminLogin()
	if itemType := this.GetString("type"); itemType == "" {
		this.Abort(ForbidCode)
	} else {
		var item models.Item
		t, _ := strconv.Atoi(itemType)
		this.Data["list"], _ = item.GetList(t)
		this.GetContentByType(ActList, itemType)
		this.SetTpl(BaseLayoutPage, "item/list.html")
	}

}

//添加
func (this *ItemController) Add() {

	this.checAdminLogin()
	switch this.requestMethod {
	case "GET":
		if itemType := this.GetString("type"); itemType == "" {
			this.Abort(ForbidCode)
		} else {
			this.GetContentByType(ActAdd, itemType)
			this.SetTpl(BaseLayoutPage, "item/add.html")
		}
	case "POST":
		var item models.Item
		if err := this.ParseForm(&item); err != nil {
			this.ReturnJson(BindErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(&item)
		if err := item.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		}
		this.ReturnJson(SuccessCode, SuccessMessage, nil)
	}
}

//编辑
func (this *ItemController) Edit() {
	this.checAdminLogin()
	switch this.requestMethod {
	case "GET":
		if itemId := this.GetString("id"); itemId == "" {
			this.Abort(ForbidCode)
		} else {
			var item models.Item
			if itemInfo, err := item.GetDetail(itemId); err != nil || itemInfo.Id == 0 {
				this.Abort(ForbidCode)
			} else {
				this.GetContentByType(ActEdit, string(itemInfo.Type))
				this.Data["data"] = itemInfo
				this.SetTpl(BaseLayoutPage, "item/edit.html")
			}
		}
	case "POST":
		var item models.Item
		if err := this.ParseForm(&item); err != nil {
			this.ReturnJson(BindErrCode, err.Error(), nil)
		}
		//校验必填参数
		this.filterParams(&item)
		if err := item.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}
	}
}

func (this *ItemController) GetContentByType(act, itemType string) {
	t, _ := strconv.Atoi(itemType)

	this.Data["itemType"] = t

	switch act {
	case ActList:
		switch t {
		case 0:
			this.Data["addTile"] = "发布寻物启事"
			this.Data["happenTime"] = "丢失时间"
			this.Data["happenPlace"] = "丢失地点"
			this.Data["person"] = "失主"

		case 1:
			this.Data["addTile"] = "发布招领启事"
			this.Data["happenTime"] = "拾到时间"
			this.Data["happenPlace"] = "拾到地点"
			this.Data["person"] = "拾取人"
		default:
			this.ReturnJson(FilterErrCode, ItemType+"不正确", nil)
		}

	case ActEdit, ActAdd:
		switch t {
		case 0:
			if act == ActAdd {
				this.Data["actTitle"] = "发布寻物启事"
			} else {
				this.Data["actTitle"] = "编辑寻物启事"
			}
			this.Data["listTile"] = "寻物列表"
			this.Data["happenTime"] = "丢失时间"
			this.Data["userId"] = "失主用户ID"
			this.Data["happenPlace"] = "丢失地点"

		case 1:
			if act == ActEdit {
				this.Data["actTitle"] = "发布招领启事"
			} else {
				this.Data["actTitle"] = "编辑招领启事"
			}
			this.Data["listTile"] = "招领列表"
			this.Data["happenTime"] = "拾到时间"
			this.Data["userId"] = "拾到用户ID"
			this.Data["happenPlace"] = "拾到地点"
		default:
			this.ReturnJson(FilterErrCode, ItemType+"不正确", nil)
		}
	default:
		this.ReturnJson(FilterErrCode, "操作类型不正确", nil)
	}

}

//更改状态
func (this *ItemController) ChangeStatus() {
	this.checAdminLogin()
	var item models.Item
	if err := this.ParseForm(&item); err != nil {
		this.ReturnJson(BindErrCode, err.Error(), nil)
	}
	if item.Id == 0 {
		this.ReturnJson(FilterErrCode, ItemId+ErrNotEmpty, nil)
	}

	if string(item.Status) == "" {
		this.ReturnJson(FilterErrCode, ItemStatus+ErrNotEmpty, nil)
	}

	if err := item.UpdateStatus(item.Status); err != nil {
		this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
	} else {
		this.ReturnJson(SuccessCode, SuccessMessage, nil)
	}

}

//删除
func (this *ItemController) Del() {

	this.checAdminLogin()
	var item models.Item
	if err := this.ParseForm(&item); err != nil {
		this.ReturnJson(BindErrCode, err.Error(), nil)
	}
	if item.Id == 0 {
		this.ReturnJson(FilterErrCode, ItemId+ErrNotEmpty, nil)
	}
	if err := item.Delete(); err != nil {
		this.ReturnJson(DelErrCode, err.Error(), nil)
	}
	this.ReturnJson(SuccessCode, SuccessMessage, nil)
}

//========home===========
//列表
func (this *ItemController) HomeList() {
	if itemType := this.GetString("type"); itemType == "" {
		this.Abort(NotFundCode)
	} else {
		var item models.Item
		pageSize := 5
		itype, _ := strconv.Atoi(itemType)
		page, _ := this.GetInt("page")
		count, itemList, _ := item.HomeGetList(itype, pageSize, pageSize*(page-1))
		this.Data["listTen"], _ = item.GetNewItem(10)
		this.Data["page"] = PageUtil(int(count), page, pageSize, itemList)
		this.GetContentByType(ActList, itemType)
		this.SetTpl(HomeBaseLayout, HomeTplPath, "/item/list.html")
	}
}

//列表
func (this *ItemController) UserItemList() {
	this.checkHomeLogin()
	var item models.Item
	pageSize := 5
	page, _ := this.GetInt("page")
	count, itemList, _ := item.GetUserItemByUserId(this.userInfo.Id, pageSize, pageSize*(page-1))
	fmt.Printf("%#v \n", itemList)
	this.Data["page"] = PageUtil(int(count), page, pageSize, itemList)
	this.SetTpl(HomeBaseLayout, HomeTplPath, "/item/userItem.html")

}

//详情
func (this *ItemController) HomeItemDetail() {

	if itemId := this.GetString("itemId"); itemId == "" {
		this.Abort(NotFundCode)
	} else {
		var item models.Item
		var user models.User
		data, err := item.GetDetail(itemId)
		if err != nil {
			this.Abort(NotFundCode)
		}
		uid, _ := strconv.Atoi(data.UserId)
		user.Id = uid
		user.GetUserInfo()
		this.Data["data"] = data
		this.Data["uinfo"] = user
		this.Data["listTen"], _ = item.GetNewItem(10)
		this.SetTpl(HomeBaseLayout, HomeTplPath, "/item/detail.html")
	}
}

//添加
func (this *ItemController) HomeItemAdd() {

	this.checkHomeLogin()
	switch this.requestMethod {
	case "GET":
		if itemType := this.GetString("type"); itemType == "" {
			this.Abort(NotFundCode)
		} else {
			this.GetContentByType(ActAdd, itemType)
			this.SetTpl(HomeBaseLayout, HomeTplPath, "/item/add.html")
		}
	case "POST":
		var item models.Item
		if err := this.ParseForm(&item); err != nil {
			this.ReturnJson(BindErrCode, err.Error(), nil)
		}
		//校验必填参数
		item.UserId = strconv.Itoa(this.userInfo.Id)
		this.filterParams(&item)
		if err := item.InsertOrUpdate(); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		}
		this.ReturnJson(SuccessCode, SuccessMessage, nil)
	}
}

//更改状态
func (this *ItemController) HomeChangeStatus() {
	this.checkHomeLogin()
	var item models.Item
	if itemId := this.GetString("itemId"); itemId == "" {
		this.ReturnJson(FilterErrCode, ItemId+ErrNotEmpty, nil)
	} else {
		intId, _ := strconv.Atoi(itemId)
		item.Id = intId
		if err := item.UpdateStatus(1); err != nil {
			this.ReturnJson(InsertUpdateErrCode, err.Error(), nil)
		} else {
			this.ReturnJson(SuccessCode, SuccessMessage, nil)
		}
	}

}

//更改状态
func (this *ItemController) HomeDel() {
	this.checkHomeLogin()
	var item models.Item
	if itemId := this.GetString("itemId"); itemId == "" {
		this.ReturnJson(FilterErrCode, ItemId+ErrNotEmpty, nil)
	} else {
		intId, _ := strconv.Atoi(itemId)
		item.Id = intId
		if err := item.Delete(); err != nil {
			this.ReturnJson(DelErrCode, err.Error(), nil)
		}
		this.ReturnJson(SuccessCode, SuccessMessage, nil)
	}
}
