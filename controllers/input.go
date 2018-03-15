package controllers

type InputController struct {
	BaseController
}

func (this *InputController) InputGet(){
	name := this.GetString("name")
	 // 不使用模版，直接用 this.Ctx.WriteString 输出字符串
	this.Ctx.WriteString(name)
}