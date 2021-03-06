// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"kindergarten-service-go/controllers"
	"kindergarten-service-go/controllers/admin"

	"kindergarten-service-go/controllers/healthy"
	"kindergarten-service-go/controllers/healthy/app"
	"kindergarten-service-go/controllers/task"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"kindergarten-service-go/controllers/kanban"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/api/v2/kg",

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/course_class",
			beego.NSInclude(
				&controllers.CourseClassController{},
			),
		),

		beego.NSNamespace("/course_info",
			beego.NSInclude(
				&controllers.CourseInfoController{},
			),
		),

		beego.NSNamespace("/healthy/drug",
			beego.NSInclude(
				&healthy.DrugController{},
			),
		),

		beego.NSNamespace("/app/healthy/drug",
			beego.NSInclude(
				&app.DrugController{},
			),
		),

		beego.NSNamespace("/healthy/inspect",
			beego.NSInclude(
				&healthy.InspectController{},
			),
		),

		beego.NSNamespace("/app/healthy/inspect",
			beego.NSInclude(
				&app.InspectController{},
			),
		),

		beego.NSNamespace("/healthy/body",
			beego.NSInclude(
				&healthy.BodyController{},
			),
		),

		beego.NSNamespace("/app/healthy/body",
			beego.NSInclude(
				&app.BodyController{},
			),
		),

		beego.NSNamespace("/healthy/class",
			beego.NSInclude(
				&healthy.ClassController{},
			),
		),

		beego.NSNamespace("/app/healthy/class",
			beego.NSInclude(
				&app.ClassController{},
			),
		),

		beego.NSNamespace("/healthy/situation",
			beego.NSInclude(
				&healthy.SituationController{},
			),
		),

		beego.NSNamespace("/app/healthy/situation",
			beego.NSInclude(
				&app.SituationController{},
			),
		),

		beego.NSNamespace("/healthy/column",
			beego.NSInclude(
				&healthy.ColumnController{},
			),
		),

		beego.NSNamespace("/app/healthy/column",
			beego.NSInclude(
				&app.ColumnController{},
			),
		),

		beego.NSNamespace("/admin/kindergarten",
			beego.NSInclude(
				&admin.KindergartenController{},
			),
		),

		beego.NSNamespace("/admin/kindergarten_life",
			beego.NSInclude(
				&admin.KindergartenLifeController{},
			),
		),

		beego.NSNamespace("/kindergarten_life",
			beego.NSInclude(
				&controllers.KindergartenLifeController{},
			),
		),

		beego.NSNamespace("/facilities_display",
			beego.NSInclude(
				&controllers.FacilitiesDisplayController{},
			),
		),

		beego.NSNamespace("/slide_show",
			beego.NSInclude(
				&controllers.SideShowController{},
			),
		),

		beego.NSNamespace("/admin/notice",
			beego.NSInclude(
				&admin.NoticeController{},
			),
		),

		beego.NSNamespace("/notice",
			beego.NSInclude(
				&controllers.NoticeController{},
			),
		),

		beego.NSNamespace("/organizational",
			beego.NSInclude(
				&controllers.OrganizationalController{},
			),
		),

		beego.NSNamespace("/organizational_member",
			beego.NSInclude(
				&controllers.OrganizationalMemberController{},
			),
		),

		beego.NSNamespace("/admin/organizational",
			beego.NSInclude(
				&admin.OrganizationalController{},
			),
		),

		beego.NSNamespace("/admin/organizational_member",
			beego.NSInclude(
				&admin.OrganizationalMemberController{},
			),
		),

		beego.NSNamespace("/admin/permission",
			beego.NSInclude(
				&admin.PermissionController{},
			),
		),

		beego.NSNamespace("/admin/role",
			beego.NSInclude(
				&admin.RoleController{},
			),
		),

		beego.NSNamespace("/admin/route",
			beego.NSInclude(
				&admin.RouteController{},
			),
		),

		beego.NSNamespace("/admin/student",
			beego.NSInclude(
				&admin.StudentController{},
			),
		),

		beego.NSNamespace("/admin/teacher",
			beego.NSInclude(
				&admin.TeacherController{},
			),
		),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),

		beego.NSNamespace("/teachers_show",
			beego.NSInclude(
				&controllers.TeachersShowController{},
			),
		),

		beego.NSNamespace("/teacher",
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),

		beego.NSNamespace("/user_permission",
			beego.NSInclude(
				&controllers.UserPermissionController{},
			),
		),

		beego.NSNamespace("/admin/user_permission",
			beego.NSInclude(
				&admin.UserPermissionController{},
			),
		),

		beego.NSNamespace("/ping",
			beego.NSInclude(
				&admin.PingController{},
			),
		),

		beego.NSNamespace("/admin/ping",
			beego.NSInclude(
				&admin.PingController{},
			),
		),
		beego.NSNamespace("/app/visitors",
			beego.NSInclude(
				&controllers.KindergartenVisitorsController{},
			),
		),
		beego.NSNamespace("/app/special_child",
			beego.NSInclude(
				&controllers.ExceptionalChildController{},
			),
		),

		beego.NSNamespace("/admin/special_child",
			beego.NSInclude(
				&admin.ExceptionalChildController{},
			),
		),

		beego.NSNamespace("/task/work_tasks",
			beego.NSInclude(
				&task.WorkTaskController{},
			),
		),
		beego.NSNamespace("/app/attendance",
			beego.NSInclude(
				&controllers.AttendanceController{},
			),
		),

		beego.NSNamespace("/task/work_plan",
			beego.NSInclude(
				&task.WorkPlanController{},
			),
		),

		beego.NSNamespace("/kanban/schedule",
			beego.NSInclude(
				&kanban.ScheduleController{},
			),
		),
		beego.NSNamespace("/kanban/attendance",
			beego.NSInclude(
				&kanban.AttendanceController{},
			),
		),
		beego.NSNamespace("/kanban/visitors",
			beego.NSInclude(
				&kanban.VisitorsController{},
			),
		),
		beego.NSNamespace("/kanban/course",
			beego.NSInclude(
				&kanban.CourseController{},
			),
		),
		beego.NSNamespace("/kanban/healthy",
			beego.NSInclude(
				&kanban.HealthyController{},
			),
		),

		beego.NSNamespace("/healthy/statistics",
			beego.NSInclude(
				&healthy.StatisticsController{},
			),
		),

		//编辑考勤
		beego.NSRouter("/app/addrule", &controllers.AttendanceController{}, "post:AddRule"),
		//签到列表（教师端）
		beego.NSRouter("/app/signlist", &controllers.AttendanceController{}, "get:SignList"),
		//一键入园（教师端）
		beego.NSRouter("/app/topspeedsign", &controllers.AttendanceController{}, "post:TopspeedSign"),
		//请假（家长端）
		beego.NSRouter("/app/leave", &controllers.AttendanceController{}, "post:Leave"),
		//离园
		beego.NSRouter("/app/leavegarden", &controllers.AttendanceController{}, "get:LeaveGarden"),
		//一键离园
		beego.NSRouter("/app/topspeedleave", &controllers.AttendanceController{}, "get:TopspeedLeaveGarden"),
		//首页图表
		beego.NSRouter("/app/attendancegarden", &controllers.AttendanceController{}, "get:AttendanceGarden"),
		//园长获取某班级考勤详情
		beego.NSRouter("/app/getdetailed", &controllers.AttendanceController{}, "get:GetDetailed"),
		//园长获取考勤规则
		beego.NSRouter("/app/getrle", &controllers.AttendanceController{}, "get:GetRUle"),

		// 根据过敏源获取过敏儿童
		beego.NSRouter("/app/allergen_child", &controllers.ExceptionalChildController{}, "get:GetAllergenChild"),
		// 根据宝宝ID获取过敏源
		beego.NSRouter("/app/get_allergen", &controllers.ExceptionalChildController{}, "get:GetAllergen"),
	)
	beego.AddNamespace(ns)
}
