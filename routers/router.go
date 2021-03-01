package routers

import (

//	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"microsvcdb02/controllers"
)

func Router()  {
	namespace:= beego.NewNamespace("/api/svcdb")
	namespace.Router("/",&controllers.MicroSVCDB02Controller{},"Get:Get")
	beego.AddNamespace(namespace)
}