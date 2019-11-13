package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Item struct {
	Id         int    `form:"id" orm:"pk"`
	UserId     string `json:"userId" form:"userId"`
	Type       int    `json:"type" form:"type"`
	Title      string `json:"title" form:"title"`
	Name       string `json:"name" form:"name"`
	Time       string `json:"time" form:"time"`
	Place      string `json:"place" form:"place"`
	Image      string `json:"image" form:"image"`
	Trait      string `json:"trait" form:"trait"`
	Status     int    `json:"status" form:"status"`
	Remark     string `json:"remark" form:"remark"`
	CreateTime string `json:"createTime" `
	UpdateTime string `json:"updateTime" `
}

//获取列表
func (this *Item) GetList(itemType int) ([]Item, error) {
	var list []Item
	this.Type = itemType
	query := orm.NewOrm().QueryTable("item")
	_, err := query.OrderBy("-create_time").Filter("type", this.Type).All(&list)

	return list, err
}

//獲取詳情
func (this *Item) GetNewItem(num int) ([]Item, error) {
	var list []Item
	query := orm.NewOrm().QueryTable("item")
	_, err := query.OrderBy("-create_time").Limit(num).All(&list)
	return list, err
}

//獲取詳情
func (this *Item) GetDetail(id string) (Item, error) {
	var item Item
	err := orm.NewOrm().Raw("select * from item where id=?", id).QueryRow(&item)
	return item, err
}

//新增或刪除
func (this *Item) InsertOrUpdate() error {

	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	this.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if this.Id == 0 {
		id, err := orm.NewOrm().Insert(this)
		if err != nil {
			return err
		}
		this.Id = int(id)
	} else {
		if _, err := orm.NewOrm().Update(this, "Title", "Name", "Time", "Place", "Trait", "Remark", "Image", "UpdateTime"); err != nil {
			return err
		}
	}
	return nil
}

//标记完成
func (this *Item) UpdateStatus(status int) error {
	this.Status = status
	if _, err := orm.NewOrm().Update(this, "Status", "UpdateTime"); err != nil {
		return err
	}
	return nil
}

//刪除
func (this *Item) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}
