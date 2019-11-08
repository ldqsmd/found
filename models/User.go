package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type User struct {
	Id 				int		`form:"id" orm:"pk"`
	Account 		string	`json:"account" form:"account"`
	Password 		string	`json:"password" form:"password"`
	Name 		string	`json:"name" form:"name"`
	Email 			string	`json:"email" form:"email"`
	Phone 			int		`json:"phone" form:"phone"`
	Forbidden  		int 	`json:"forbidden" form:"forbidden"`
	CreateTime   	string 	`json:"createTime" `
	UpdateTime   	string 	`json:"updateTime" `
}

//获取列表
func (this *User) GetList()([]User,error){
	var list []User
	query := orm.NewOrm().QueryTable("user")
	_,err := query.OrderBy("-create_time").All(&list)
	return list,err
}

//獲取詳情
func (this *User) GetDetail(id string)(User,error){
	var User  User
	err := orm.NewOrm().Raw("select * from User where id=?",id).QueryRow(&User)
	return User, err
}

//新增或刪除
func (this *User)InsertOrUpdate()error  {

	this.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	this.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	if this.Id == 0 {
		id,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.Id = int(id)
	}else {
		if  _,err :=  orm.NewOrm().Update(this,"Name","Email","phone","UpdateTime"); err != nil{
			return err
		}
	}
	return nil
}

//禁用用户
func (this *User)ForbidUser(userId,forbidden int)error  {
	this.Id = userId
	this.Forbidden = forbidden
	if  _,err :=  orm.NewOrm().Update(this,"Forbidden","UpdateTime"); err != nil{
		return err
	}
	return nil
}












