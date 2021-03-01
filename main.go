package main

import (
	"github.com/astaxie/beego"
	_ "microsvcdb02/conf"
	_ "microsvcdb02/filter"
	"microsvcdb02/routers"
)

func main() {
	routers.Router()
	beego.Run()

}
