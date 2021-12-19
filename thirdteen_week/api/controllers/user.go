package controllers

import (
	p "api/parameters"
	"api/untils/log"

	"github.com/astaxie/beego"
)

//UserController 账号服务
type UserController struct {
	beego.Controller
}

//Create 创建账号
func (c *UserController) Create() {
	log.Logger.Debug("debug:user create")
	c.Data["json"] = &p.UserCreateParam{
		Name:     "test",
		Password: "123456",
	}
	c.ServeJSON()
}
