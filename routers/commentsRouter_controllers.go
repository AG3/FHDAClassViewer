package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"] = append(beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"] = append(beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetCourse",
			Router: `/course`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"] = append(beego.GlobalControllerRouter["BetterClassViewer/controllers:MainController"],
		beego.ControllerComments{
			Method: "UpdateDatas",
			Router: `/admin/update`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
