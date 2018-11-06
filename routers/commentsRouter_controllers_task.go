package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkPlanController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Complete",
			Router: `/complete`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Finished",
			Router: `/finish/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/task:WorkTaskController"],
		beego.ControllerComments{
			Method: "Schedule",
			Router: `/schedule`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
