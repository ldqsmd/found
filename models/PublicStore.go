package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)


type PublicStore struct {

	StoreId         int		`form:"storeId" orm:"pk"`
	StoreRemark 	string	`form:"storeRemark"`
	StoreCreateTime string  `form:"storeCreateTime"`
	StoreUpdateTime	string  `form:"_"`
	StoreName       string	`form:"storeName"`
	Number        	string	`form:"number"`
	Brand        	int		`form:"brand"`
	Status        	int		`form:"status"`
	Manager        	string	`form:"manager"`
	ManagerPhone    string	`form:"managerPhone"`
	Address        	string	`form:"address"`
	AreaManager     string	`form:"areaManager"`
	RegionalManager string	`form:"regionalManager"`
	Ip       		string	`form:"ip"`
	Network        	int		`form:"network"`
	OpenTime       	string	`form:"openTime"`
	DecorationTime 	string 	`form:"decorationTime"`
	CloseTime       string	`form:"closeTime"`
	BookStartTime	string  `form:"bookStartTime"`
	EmployeeStartTime	string  `form:"employeeStartTime"`
	WaitTime        string	`form:"waitTime"`
	DeviceTime 		string	`form:"deviceTime"`
	BuildName 		string	`form:"buildName"`
	TempCloseTime 	string	`form:"tempCloseTime"`
	ForbiddenStatus int		`form:"forbiddenStatus"`
	SignFlag		int		`form:"signFlag"`
}


type StoreOption struct {
	StoreId         int 	`json:"storeId"`
	StoreName       string 	`json:"storeName"`
}


func (this *PublicStore) TableName() string {
	return "store"
}

//获取餐厅列表
func (this *PublicStore) GetStoreList()([]PublicStore,error){
	var storeList []PublicStore
	_, err := orm.NewOrm().Raw("SELECT * from store order by forbidden_status, sign_flag desc,open_time desc,wait_time desc").QueryRows(&storeList)
	if err == nil {
		return storeList,err
	}
	return storeList, nil
}

//获取餐厅option 信息
func (this  StoreOption)GetStoreOption()[]StoreOption  {

	var storeOption []StoreOption

	num, err :=  orm.NewOrm().Raw("SELECT store_id, store_name FROM store order by wait_time desc ").QueryRows(&storeOption)
	if err == nil {
		fmt.Println("user nums: ", num)
	}

	return  storeOption

}

func (this *PublicStore)InsertOrUpdate()error {

	if this.StoreId == 0 {
		this.StoreCreateTime  = time.Now().Format("2006-01-02 15:04:05")
		storeId,err :=  orm.NewOrm().Insert(this)
		if err != nil{
			return err
		}
		this.StoreId = int(storeId)
	}else {
		this.StoreUpdateTime  = time.Now().Format("2006-01-02 15:04:05")
		_,err :=  orm.NewOrm().Update(this)
		if err != nil{
			return err
		}
	}
	return  nil
}
//软删除
func (this *PublicStore)DeleteStore()error  {
	this.ForbiddenStatus = 1
	_,err :=  orm.NewOrm().Update(this,"forbidden_status")
	return err
}

//标记店为重要
func (this *PublicStore)SignStore()error  {

	_,err :=  orm.NewOrm().Update(this,"SignFlag")
	return err
}


//获取公共餐厅信息
func (this *PublicStore)GetStoreInfo(storeId string)PublicStore{
	//获取
	var publicStore PublicStore
	orm.NewOrm().Raw("SELECT * FROM store WHERE store_id=?",storeId).QueryRow(&publicStore)
	return publicStore
}












