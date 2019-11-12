package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"strconv"
)

const (
	SuccessMessage = "操作成功"
	ErrNotEmpty    = "不能为空"

	SuccessCode         = 0
	BindErrCode         = -1
	FilterErrCode       = -2
	InsertUpdateErrCode = -3
	DelErrCode          = -4

	//上传文件类型
	UploadExcel   = 1
	UploadPicture = 2

	BaseLayoutPage = "base/layout_page.html"
	HomeBaseLayout = "home/base/layout.html"
	HomeTplPath    = "home"
)

// 通过map主键唯一的特性过滤重复元素
func FilterStrSlice(slc []string) []string {

	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//发送邮件
func SendEmail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"from": beego.AppConfig.String("email::from"),
		"user": beego.AppConfig.String("email::user"),
		"pass": beego.AppConfig.String("email::pass"),
		"host": beego.AppConfig.String("email::host"),
		"port": beego.AppConfig.String("email::port"),
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["from"]+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                   //发送给多个用户
	m.SetHeader("Subject", subject)                                //设置邮件主题
	m.SetBody("text/html", body)                                   //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)

	return err

}
