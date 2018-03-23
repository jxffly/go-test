package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AnnotationController struct {
	beego.Controller
	}
// @router /anno/:key [get]
func (c *AnnotationController )Find()  {
	id:=c.Ctx.Input.Params()
	logs.Info("find method called,params:{}",id)
	c.Ctx.WriteString("annotation mapping")
}
