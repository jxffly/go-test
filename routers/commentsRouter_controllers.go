package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego-test/controllers:AnnotationController"] = append(beego.GlobalControllerRouter["beego-test/controllers:AnnotationController"],
		beego.ControllerComments{
			Method: "Find",
			Router: `/anno/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
