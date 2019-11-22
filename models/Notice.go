package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Notice struct {
	Id         int    `form:"id" orm:"pk"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	CreateTime string `json:"createTime" `
	UpdateTime string `json:"updateTime" `
}

const (
	NoticeTable = "notice"
)

//获取列表
func (this *Notice) GetList() ([]Notice, error) {
	var list []Notice
	_, err := orm.NewOrm().QueryTable(NoticeTable).OrderBy("-create_time").All(&list)
	return list, err
}

//详情
func (this *Notice) GetDetail(id string) (Notice, error) {
	var detail Notice
	err := orm.NewOrm().QueryTable(NoticeTable).Filter("id", id).One(&detail)
	return detail, err
}

//新增或刪除
func (this *Notice) InsertOrUpdate() error {

	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	this.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if this.Id == 0 {
		id, err := orm.NewOrm().Insert(this)
		if err != nil {
			return err
		}
		this.Id = int(id)
	} else {
		if _, err := orm.NewOrm().Update(this); err != nil {
			return err
		}
	}
	return nil
}

//刪除
func (this *Notice) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>home>>>>>>>>>>>>>>>>>>>>>>>

//获取列表
func (this *Notice) HomeGetList(limit, offset int) (int64, []Notice, error) {
	var list []Notice

	query := orm.NewOrm().QueryTable(NoticeTable)
	count, _ := query.Count()
	_, err := query.OrderBy("-create_time").Limit(limit, offset).All(&list)
	return count, list, err
}

//获取最新的num 条
func (this *Notice) GetTheLatest(num int) ([]Notice, error) {
	var list []Notice
	_, err := orm.NewOrm().QueryTable(NoticeTable).OrderBy("-create_time").Limit(num).All(&list)
	return list, err
}
