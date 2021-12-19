package routers

import (
	"api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/user/create", &controllers.UserController{}, "*:Create")
	// beego.AddNamespace(v1)
}
