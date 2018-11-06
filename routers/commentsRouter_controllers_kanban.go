package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:AttendanceController"],
		beego.ControllerComments{
			Method: "GetAttendance",
			Router: `/attendance`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:AttendanceController"],
		beego.ControllerComments{
			Method: "Leave",
			Router: `/leave`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:CourseController"],
		beego.ControllerComments{
			Method: "GetClassOneDayCourse",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"],
		beego.ControllerComments{
			Method: "GetAbnormalList",
			Router: `/archives`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"],
		beego.ControllerComments{
			Method: "GetAbnormalInfo",
			Router: `/archives/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"],
		beego.ControllerComments{
			Method: "GetDrugList",
			Router: `/drug`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:HealthyController"],
		beego.ControllerComments{
			Method: "GetDrugInfo",
			Router: `/drug/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"],
		beego.ControllerComments{
			Method: "PostSchedule",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"],
		beego.ControllerComments{
			Method: "GetScheduleList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"],
		beego.ControllerComments{
			Method: "PutSchedule",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"],
		beego.ControllerComments{
			Method: "DeleteSchedule",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:ScheduleController"],
		beego.ControllerComments{
			Method: "GetScheduleInfo",
			Router: `/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:VisitorsController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:VisitorsController"],
		beego.ControllerComments{
			Method: "GetVisitorsList",
			Router: `/visitors_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:VisitorsController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/kanban:VisitorsController"],
		beego.ControllerComments{
			Method: "GetVisitorsNum",
			Router: `/visitors_num`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
