package routers

import (
	"api/controllers"

	"github.com/astaxie/beego"
)

var v1 = beego.NewNamespace("/v1",
	// user controller
	beego.NSNamespace("/user",
		beego.NSRouter("/create", &controllers.UserController{}, "*:Create"),
	),
)
