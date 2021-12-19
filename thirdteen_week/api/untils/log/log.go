package log

import (
	"api/conf"

	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func init() {
	if conf.Runmode == "dev" {
		consoleLog()
	}

}

//输出
func consoleLog() {
	Logger = logs.NewLogger()
}

//写入文件
func fileLog() {
	Logger = logs.NewLogger(10000)
	jsonConfig := `{
        "filename" : "test.log", // 文件名
        "maxlines" : 1000,       // 最大行
        "maxsize"  : 10240       // 最大Size
    }`
	Logger.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	Logger.SetLevel(logs.LevelDebug)     // 设置日志写入缓冲区的等级
	Logger.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）
}
