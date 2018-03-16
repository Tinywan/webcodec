package controllers

import (
	"github.com/astaxie/beego"
	"github.com/Tinywan/webcodec/models"
	"strconv"
	"strings"
	"github.com/Tinywan/webcodec/libs"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type NestPreparer interface {
	NestPreparer()
}

// 通过子 struct 的方法重写，就可以实现自己的逻辑
type BaseController struct {
	beego.Controller     // 这个控制器里面内嵌了 beego.Controller，这就是 Go 的嵌入方式，也就是 MainController 自动拥有了所有 beego.Controller 的方法。
	controllerName string
	actionName     string
	user           *models.User
	userId         int
	userName       string
	pageSize       int
	isLogin		   bool
}

// 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类
// 实现接口方法 使用指针访问
func (this *BaseController) Prepare() {
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	//this.auth()

	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

//登录状态验证
func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.UserGetById(userId)
			if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id
				this.userName = user.UserName
				this.user = user
			}
		}
	}

	if this.userId == 0 && (this.controllerName != "main" ||
		(this.controllerName == "main" && this.actionName != "logout" && this.actionName != "login")) {
		this.redirect(beego.URLFor("MainController.Login"))
	}
}

//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.Layout = "layout/layout.html"
	this.TplName = tplname
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 显示错误信息
func (this *BaseController) showMsg(args ...string) {
	this.Data["message"] = args[0]
	redirect := this.Ctx.Request.Referer()
	if len(args) > 1 {
		redirect = args[1]
	}

	this.Data["redirect"] = redirect
	this.Data["pageTitle"] = "系统提示"
	this.display("error/message")
	this.Render()
	this.StopRun()
}

// 输出json
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg

	this.jsonResult(out)
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *BaseController) Get() {
	this.Data["Website"] = "beego.tinywan.com"
	this.Data["Email"] = "756684177@qq.com"
	this.TplName = "main/index.html"
}

func (this *BaseController) Post() {
	this.Ctx.WriteString("This is Post request this") // echo
}

