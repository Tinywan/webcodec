package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"time"
	"strings"
	"strconv"
	"github.com/Tinywan/webcodec/models"
	"github.com/Tinywan/webcodec/libs"
)

type MainController struct {
	BaseController
}

// 首页
func (this *MainController) Index() {
	this.Data["pageTitle"] = "系统概况"  // this.Data 是一个用来存储输出数据的 map，可以赋值任意类型的值
	this.display()
}

func (this *MainController) Get() {
	name := this.GetString("name")
	this.Ctx.WriteString(name)
}

// 个人信息
func (this *MainController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user, _ := models.UserGetById(this.userId)

	if this.isPost() {
		flash := beego.NewFlash()
		user.Email = this.GetString("email")
		user.Update()
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")
		if password1 != "" {
			if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else if password2 != password1 {
				flash.Error("两次输入的密码不一致")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else {
				user.Salt = string(utils.RandomCreateBytes(10))
				user.Password = libs.Md5([]byte(password1 + user.Salt))
				user.Update()
			}
		}
		flash.Success("修改成功！")
		flash.Store(&this.Controller)
		this.redirect(beego.URLFor(".Profile"))
	}

	this.Data["pageTitle"] = "个人信息"
	this.Data["user"] = user
	this.display()
}

// 用户登陆
func (this *MainController)  Login() {
	if this.userId > 0 {
		this.redirect("/")
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		flash := beego.NewFlash()

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if username != "" && password != "" {
			user, err := models.UserGetByName(username)
			errorMsg := ""
			if err != nil || user.Password != libs.Md5([]byte(password + user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				models.UserUpdate(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + authkey, 7 * 86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + authkey)
				}

				this.redirect(beego.URLFor("TaskController.List"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("MainController.Login"))
		}
	}

	this.TplName = "main/login.html" // this.TplName 就是需要渲染的模板，当然也可以不使用模版，直接用 this.Ctx.WriteString 输出字符串
}

// 退出登录
func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("MainController.Login"))
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}



