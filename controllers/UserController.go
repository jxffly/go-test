package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
	}

func (c *UserController )Find()  {
	id:=c.Ctx.Input.Params()
	logs.Info("find method called,params:{}",id)
	c.Ctx.WriteString("auto mapping")
}

