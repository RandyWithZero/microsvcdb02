package conf

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.Info("配置文件配置")
	//服务器默认在请求的时候输出 server名称
	beego.BConfig.ServerName = "SVCDB02"
	//是否将错误信息进行渲染，默认值为 true，即出错会提示友好的出错页面，对于 API 类型的应用可能需要将该选项设置为 false
	beego.BConfig.EnableErrorsRender = false
	//是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。
	beego.BConfig.WebConfig.AutoRender = false
	//是否开启文档内置功能，默认是 false
	beego.BConfig.WebConfig.EnableDocs = true
	//Flash 数据设置时 Cookie 的名称
	beego.BConfig.WebConfig.FlashName = "SVCDB02_FLASH"
	//是否开启热升级，默认是 false，关闭热升级
	beego.BConfig.Listen.Graceful = true
	//设置 HTTP 的超时时间，默认是 0，不超时。
	beego.BConfig.Listen.ServerTimeOut = 30
	//session 是否开启，默认是 false。
	beego.BConfig.WebConfig.Session.SessionOn = true
	//session 的引擎，默认是 memory
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	//存在客户端的 cookie 名称
	beego.BConfig.WebConfig.Session.SessionName = "SVCDB02-sessionID"
	//session 默认存在客户端的 cookie 的时间，默认值是 3600 秒
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
}
