package routers

import (
	"beego-test/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/main", &controllers.MainController{})
	beego.Router("/rest/?:id", &controllers.RestController{})
	beego.Router("/index/:id", &controllers.RestController{},"*:Index")

	beego.Router("/unmap/", &controllers.RestController{})
	beego.AutoRouter(&controllers.UserController{})
	//初始化 namespace
	sentry := func(c *context.Context) {
		c.Input.Context.WriteString("aaa")
	}
	ns :=
		beego.NewNamespace("/v1",
			beego.NSRouter("/version", &controllers.AdminController{}, "get:ShowAPIVersion"),
			beego.NSRouter("/changepassword", &controllers.UserController{}),
			beego.NSNamespace("/shop",
				beego.NSBefore(sentry),
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("notAllowed"))
				}),
			),
			beego.NSNamespace("/cms",
				beego.NSInclude(
					&controllers.AnnotationController{},
				),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
