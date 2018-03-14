package routers

import (
	"github.com/Tinywan/webcodec/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//// GET 请求
	//beego.Router("/get", &controllers.MainController{}, "get:Get")
	//
	//// 登陆
	//beego.Router("/", &controllers.MainController{}, "get:Login")
	//
	//// POST 请求 http://127.0.0.1:8080/api/list
	//beego.Router("/post",&controllers.MainController{},"post:Post")

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
}
