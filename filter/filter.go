package filter

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var logFilter = func(ctx *context.Context){
	url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	formParams, _ := json.Marshal(ctx.Request.Form)
	params := ctx.Input.Params()
	delete(params,":splat")
	paramsStr, _ := json.Marshal(params)
	outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
	divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	topDivider := "┌" + divider
	middleDivider := "├" + divider
	bottomDivider := "└" + divider
	outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(formParams) + "\n│ 路径参数:"+string(paramsStr)+"\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
	logs.Info(outputStr)
}

func init() {
   beego.InsertFilter("/*",beego.FinishRouter,logFilter,false)
}
