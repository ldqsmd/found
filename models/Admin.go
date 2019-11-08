package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id          	int		`form:"adminId"`
	Account        	string	`form:"account"`
	UserName    	string	`form:"userName"`
	Password     	string	`form:"password"`
	Email     		string 	`form:"email"`
	Status     		int
	CreateTime     	string
}


func (this *Admin) GetAdminList()([]Admin,int64)  {
	var adminList []Admin
	query := orm.NewOrm().QueryTable("admin")
	total ,_ :=query.OrderBy("id").Filter("status","0").All(&adminList)
	return adminList, total
}


func (this *Admin) GetAdminDetail(id string)(Admin){
	var admin  Admin
	orm.NewOrm().Raw("select * from admin where id=?",id).QueryRow(&admin)
	return admin
}

func (this *Admin)InsertOrUpdate()error  {

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if this.Id == 0 {
		this.CreateTime = nowTime
		if id,err :=  orm.NewOrm().Insert(this); err != nil{
			return err
		}else{
			this.Id  = int(id)
		}
	}else {
		if  _,err :=  orm.NewOrm().Update(this,"UserName","Email"); err != nil{
			return err
		}
	}
	return nil
}

//获取管理员信息
func (this *Admin)ForbidAdmin(id int)(error){

	this.Id  	= id
	this.Status = 1
	if  _,err :=  orm.NewOrm().Update(this,"Status"); err != nil{
		return err
	}
	return nil
}

//获取管理员信息
func (this *Admin)GetAdminInfo()(error){
	//登录获取用户信息
	if this.Id == 0 {
		return orm.NewOrm().Raw("SELECT id,account,user_name,email,status FROM admin WHERE account=? and password=?", this.Account,this.Password).QueryRow(&this)
	}else{
		//根据adminId获取用户信息
		return orm.NewOrm().Raw("SELECT id,account,user_name,email,status FROM admin WHERE id=?", this.Id).QueryRow(&this)
	}
}












