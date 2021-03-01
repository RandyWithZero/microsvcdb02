package controllers

import (
	"github.com/astaxie/beego"
	"os"
)

type MicroSVCDB02Controller struct {
	beego.Controller
}
func (c *MicroSVCDB02Controller) Get() {
	getEnv := os.Getenv("HOSTNAME")
	msg :="--->调用svcdb02"+"["+getEnv+"]"
	c.Ctx.WriteString(msg)
}
