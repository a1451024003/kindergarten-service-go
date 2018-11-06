package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:ExceptionalChildController"],
		beego.ControllerComments{
			Method: "PutInspect",
			Router: `/inspect/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "SetPrincipal",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "GetIntroduceInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Attest",
			Router: `/attest/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "AttestKindergartenAll",
			Router: `/attest_all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "AttestKindergarten",
			Router: `/attest_kindergarten`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Reset",
			Router: `/baby`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Count",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "FoodClass",
			Router: `/food_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "FoodScale",
			Router: `/food_scale`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "GetKinderMbmber",
			Router: `/get_member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "GetKg",
			Router: `/getkg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "Introduce",
			Router: `/introduce/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "SetKindergarten",
			Router: `/set_kindergarten`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenController"],
		beego.ControllerComments{
			Method: "StudentClass",
			Router: `/student_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "GetKindergartenLifeList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "GetKindergartenLifeInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:KindergartenLifeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"],
		beego.ControllerComments{
			Method: "GetNoticeList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"],
		beego.ControllerComments{
			Method: "GetNoticeInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:NoticeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetOrganization",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "UpOrganization",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "AddOrganization",
			Router: `/addorgan`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "AuthClass",
			Router: `/auth_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "GetKCAll",
			Router: `/class_kinder`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "DelOrganization",
			Router: `/delorgan`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "AddManyClass",
			Router: `/many_class`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "Member",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalController"],
		beego.ControllerComments{
			Method: "Principal",
			Router: `/principal`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "OrganizationList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "WebOrganizationList",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "MyKindergarten",
			Router: `/my_kinder`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:OrganizationalMemberController"],
		beego.ControllerComments{
			Method: "MyKinderTeacher",
			Router: `/my_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PermissionController"],
		beego.ControllerComments{
			Method: "Option",
			Router: `/option`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"],
		beego.ControllerComments{
			Method: "Ping",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"],
		beego.ControllerComments{
			Method: "PingOnemoreRpc",
			Router: `/rpc/onemore`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:PingController"],
		beego.ControllerComments{
			Method: "PingUserRpc",
			Router: `/rpc/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:RouteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "GetStudent",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "DeleteStudent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "UpdateStudent",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "UpStudent",
			Router: `/auth/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "GetBaby",
			Router: `/baby`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "BabyActived",
			Router: `/baby_actived`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "BabyKindergarten",
			Router: `/baby_kinder`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "GetStudentClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "GetNameClass",
			Router: `/get_class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "Invite",
			Router: `/invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "DeleteKinship",
			Router: `/kinship/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "OragnizationalStudent",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "RemoveStudent",
			Router: `/remove`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:StudentController"],
		beego.ControllerComments{
			Method: "ResetInvite",
			Router: `/reset_invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "GetTeacher",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "GetTeacherInfo",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "UpTeacher",
			Router: `/auth/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "GetClass",
			Router: `/class`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Duties",
			Router: `/duties`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "GetUT",
			Router: `/getut`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Invite",
			Router: `/invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "InviteReset",
			Router: `/invite_reset`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "OrganizationalTeacher",
			Router: `/organizational_teacher`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "RemoveTeacher",
			Router: `/remove`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "Reset",
			Router: `/reset`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:TeacherController"],
		beego.ControllerComments{
			Method: "ResetInvite",
			Router: `/reset_invite`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetGroupPermission",
			Router: `/group/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"] = append(beego.GlobalControllerRouter["kindergarten-service-go/controllers/admin:UserPermissionController"],
		beego.ControllerComments{
			Method: "GetUserPermission",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
