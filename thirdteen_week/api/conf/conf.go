package conf

import "github.com/astaxie/beego"

var (
	Runmode string
)

func init() {
	Runmode = beego.AppConfig.String("runmode")
}
