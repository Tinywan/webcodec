package controllers

import (
	"github.com/astaxie/beego"
)

type InputController struct {
	beego.Controller
}

func (this *InputController) Prepare(){

}

func (this *InputController) Get(){
	this.Data["content"] = "value"
	this.Layout = "layout/layout.html"
	this.TplName = "input/index.html"
}

func (this *InputController) InputGet(){
	name := this.GetString("name")
	 // 不使用模版，直接用 this.Ctx.WriteString 输出字符串
	this.Ctx.WriteString(name)
}