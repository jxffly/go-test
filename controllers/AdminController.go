package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
	}

func (c *AdminController)ShowAPIVersion()  {
	c.Ctx.WriteString("v1")
}