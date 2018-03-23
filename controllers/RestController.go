package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RestController struct {
	beego.Controller
}

func (c *RestController) Get() {
	id:=c.Ctx.Input.Param(":id")
	logs.Info("the id from paths is ",id)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["author"]="jinxiaofei"
	c.Ctx.WriteString("rest get!	")
}
func (c *RestController) Post() {
	c.Ctx.WriteString("hello world!!	")
}

func (c *RestController) Index() {
	id:=c.Ctx.Input.Param(":id")
	logs.Info("the id from paths is ",id)
	c.Ctx.WriteString("my personal method!!	")
}

func (c *RestController) UnMapping() {
	c.Ctx.WriteString("my personal method un mapping!!	")
}