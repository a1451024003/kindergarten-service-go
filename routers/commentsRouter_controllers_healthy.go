package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:BodyController"],
		beego.ControllerComments{
			Method: "Push",
			Router: `/push/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:ClassController"],
		beego.ControllerComments{
			Method: "Post_info",
			Router: `/add_info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:DrugController"],
		beego.ControllerComments{
			Method: "DrugInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Inspect",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "AbnormalList",
			Router: `/abnormallist/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Abnormal",
			Router: `/archives/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Content",
			Router: `/content/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Country",
			Router: `/country/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Counts",
			Router: `/counts/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "DeleteStudent",
			Router: `/deleteStudent/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Height",
			Router: `/height/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Project",
			Router: `/project/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "ProjectNew",
			Router: `/projectNew/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:InspectController"],
		beego.ControllerComments{
			Method: "Weight",
			Router: `/weight/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:SituationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:StatisticsController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/healthy:StatisticsController"],
		beego.ControllerComments{
			Method: "Statistics",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
