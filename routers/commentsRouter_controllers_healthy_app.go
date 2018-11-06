package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:BodyController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:ClassController"],
		beego.ControllerComments{
			Method: "Post_info",
			Router: `/add_info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:DrugController"],
		beego.ControllerComments{
			Method: "DrugInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Inspect",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Abnormal",
			Router: `/archives/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Boby",
			Router: `/boby`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Body",
			Router: `/body/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Counts",
			Router: `/counts/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Personal",
			Router: `/personal/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "PersonalInfo",
			Router: `/personalInfo/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "PersonalList",
			Router: `/personalList/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Project",
			Router: `/project/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "ProjectNew",
			Router: `/projectNew/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "Situation",
			Router: `/situation`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:InspectController"],
		beego.ControllerComments{
			Method: "SituationUrl",
			Router: `/situationUrl`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:SituationController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy/app:SituationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
