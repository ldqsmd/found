package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLInit()  {

	host 		:= beego.AppConfig.String("mysql::host")
	port 		:= beego.AppConfig.String("mysql::port")
	userName 	:= beego.AppConfig.String("mysql::userName")
	password 	:= beego.AppConfig.String("mysql::password")
	database 	:= beego.AppConfig.String("mysql::databse")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", userName+":"+password+"@tcp("+host+":"+ port+")/"+database+"?charset=utf8", 30)

	if  err != nil {
		fmt.Println(err.Error())
	}
}


