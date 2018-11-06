package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:AttendanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetTimelist",
			Router: `/class_course`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetTimeOne",
			Router: `/class_day`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetClassCourseware",
			Router: `/class_daycourse`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetTimeDay",
			Router: `/classday`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetPlan",
			Router: `/plan`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetPlanInfo",
			Router: `/plan_info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseClassController"],
		beego.ControllerComments{
			Method: "GetPlanInfoNew",
			Router: `/planinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "DelCourse",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "PostTime",
			Router: `/add_time`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Add_Alltime",
			Router: `/addalltime`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Add_time",
			Router: `/addtime`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Add_use",
			Router: `/adduse`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetCourse",
			Router: `/courseinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "DelTimeInfo",
			Router: `/del_time`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetTimeInfo",
			Router: `/time_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"],
		beego.ControllerComments{
			Method: "Add_info",
			Router: `/addinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:CourseInfoController"],
		beego.ControllerComments{
			Method: "Edit_info",
			Router: `/editinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "GetSearch",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "AllergenPreparation",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "DelAllergen",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"],
		beego.ControllerComments{
			Method: "GetNoticeList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:FacilitiesDisplayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "GetKindergartenLifeList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "GetKindergartenLifeInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "ActivityPicture",
			Router: `/activity`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenVisitorsController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:KindergartenVisitorsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "GetNoticeList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "GetNoticeInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetOrganization",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "UpOrganization",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "AddOrganization",
			Router: `/addorgan`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetBabyClass",
			Router: `/class_baby`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetKC",
			Router: `/class_kinder`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetCS",
			Router: `/class_student`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "DelOrganization",
			Router: `/delorgan`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "FilterStudent",
			Router: `/filter_student`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "Member",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetTeacherOrganization",
			Router: `/organ_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalController"],
		beego.ControllerComments{
			Method: "Principal",
			Router: `/principal`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "OrganizationList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "WebOrganizationList",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "MyKinderTeacher",
			Router: `/my_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "OrganizationalTeachers",
			Router: `/organizational_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "MyKindergarten",
			Router: `/teacher_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"],
		beego.ControllerComments{
			Method: "Ping",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"],
		beego.ControllerComments{
			Method: "PingOnemoreRpc",
			Router: `/rpc/onemore`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:PingController"],
		beego.ControllerComments{
			Method: "PingUserRpc",
			Router: `/rpc/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"],
		beego.ControllerComments{
			Method: "GetSideShowList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"],
		beego.ControllerComments{
			Method: "GetSideShow",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"],
		beego.ControllerComments{
			Method: "DeleteSideShow",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:SideShowController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Student",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "UpdateStudent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "DeleteStudent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetStudentOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetBaby",
			Router: `/baby`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "BabyKinderList",
			Router: `/baby_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetStudentClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetNameClass",
			Router: `/get_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Invite",
			Router: `/invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:StudentController"],
		beego.ControllerComments{
			Method: "RemoveStudent",
			Router: `/remove`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "GetTeacher",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "GetTeacherInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "GetClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "Teacher",
			Router: `/filter_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "Invite",
			Router: `/invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "OrganizationalTeacher",
			Router: `/organizational_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "RemoveTeacher",
			Router: `/remove`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeacherController"],
		beego.ControllerComments{
			Method: "GetTeacherOne",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"],
		beego.ControllerComments{
			Method: "TeachersShowAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"],
		beego.ControllerComments{
			Method: "TeachersShowOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:TeachersShowController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetGroupPermission",
			Router: `/group/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "GroupAll",
			Router: `/group_all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetUserPermission",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
