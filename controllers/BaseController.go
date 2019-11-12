package controllers

import (
	"errors"
	"fmt"
	"found/models"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
	requestMethod  string //当前接口请求方式
	userInfo       models.User
	adminInfo      models.Admin
}

func (this *BaseController) Prepare() {
	//附值
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.requestMethod = this.Ctx.Request.Method //当前接口请求方式
	//从Session里获取数据 设置用户信息
	this.adapterAdminInfo()
	this.adapterUserInfo()
	//this.checAdminLogin()
	//this.checkHomeLogin()
}

//获取session admin信息
//适配到BaseController
func (this *BaseController) adapterAdminInfo() {
	adminInfo := this.GetSession("adminInfo")
	if adminInfo != nil {
		this.adminInfo = adminInfo.(models.Admin)
		this.Data["adminInfo"] = adminInfo
	} else {
		this.Data["adminInfo"] = new(models.Admin)
	}
}

//获取session user信息
//适配到BaseController
func (this *BaseController) adapterUserInfo() {
	userInfo := this.GetSession("userInfo")
	if userInfo != nil {
		this.userInfo = userInfo.(models.User)
		this.Data["userInfo"] = userInfo
	} else {
		this.Data["userInfo"] = new(models.User)
	}
}

//检查后台用户是否登录
//没有登录返回登录页面
func (this *BaseController) checAdminLogin() {

	if this.adminInfo.Id == 0 {
		//登录页面地址
		loginUrl := this.URLFor("LoginController.Login") + "?returnURL="
		//登录成功后返回的址为当前
		returnURL := this.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if this.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := this.Ctx.Input.Refer()
			this.ReturnJson(302, "请登录", loginUrl+returnURL)
		}
		this.Redirect(loginUrl+returnURL, 302)
		this.StopRun()
	}

}

//检查前台是否登录
//没有登录返回登录页面
func (this *BaseController) checkHomeLogin() {

	if this.userInfo.Id == 0 {
		//登录页面地址
		loginUrl := this.URLFor("LoginController.HomeLogin")
		//登录成功后返回的址为当前
		//如果ajax请求则返回相应的错码和跳转的地址
		this.Redirect(loginUrl, 302)
		this.StopRun()
	}

}

// 设置模板
// 第一个参数为layout，
// 第二个参数为模板
func (this *BaseController) SetTpl(template ...string) {

	var tplName string
	layout := "base/layout_base.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		layout = template[0]
		tplName = template[1]
		//三个参数的时候
	// 第一个参数为layout，
	// 第二个参数为模板路径
	// 第三个参数为模板名称
	case len(template) == 3:
		layout = template[0]
		tplName = template[1] + template[2]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(this.controllerName[0 : len(this.controllerName)-10])
		actionName := strings.ToLower(this.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}

	this.Layout = layout
	this.TplName = tplName
	fmt.Println(tplName)
}

//JSON返回
func (this *BaseController) ReturnJson(code int, message string, data interface{}) {

	var returnJson models.ReturnJson
	returnJson.Code = code
	returnJson.Message = message
	returnJson.Data = data
	this.Data["json"] = &returnJson
	this.ServeJSON()
	this.StopRun()
}

//1 上传excel 2 pic
func (this *BaseController) UpFileTable(formFile string, fileType int) (string, error) {

	var filePath string
	f, h, err := this.GetFile(formFile) //获取上传的文件
	if err != nil {
		return filePath, err
	}
	ext := path.Ext(h.Filename)

	switch fileType {
	case 1:
		//验证后缀名是否符合要求
		var AllowExtMap map[string]bool = map[string]bool{
			".xls":  true,
			".csv":  true,
			".xlsx": true,
		}
		if _, ok := AllowExtMap[ext]; !ok {
			return filePath, errors.New("文件格式不正确,允许文件格式(.xls/.csv/.xlsx)")
		}

	case 2:
		//验证后缀名是否符合要求
		var AllowExtMap map[string]bool = map[string]bool{
			".jpg": true,
			".png": true,
			".img": true,
		}
		if _, ok := AllowExtMap[ext]; !ok {
			return filePath, errors.New("文件格式不正确,允许文件格式(.jpg/.png/.img)")
		}
	}

	//创建目录
	uploadDir := "static/upload/" + formFile + "/"
	err = os.MkdirAll(uploadDir, 777)
	if err != nil {
		return filePath, err
	}
	//构造文件名称
	fileName := time.Now().Format("20060102150405_") + fmt.Sprintf("%d", rand.Intn(9999)+1000) + ext
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	err = this.SaveToFile(formFile, uploadDir+fileName)
	if err != nil {
		return filePath, err
	}
	filePath = "/" + uploadDir + fileName
	return filePath, nil
}

//
func (this *BaseController) Error404() {
	page404Url := this.URLFor("HomeController.Page404") + "?returnUrl=" + this.Ctx.Request.URL.Path
	this.Redirect(page404Url, 302)
	this.StopRun()
}

func (c *BaseController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.tpl"
}

func (c *BaseController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.tpl"
}
