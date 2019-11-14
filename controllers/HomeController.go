package controllers

import (
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
