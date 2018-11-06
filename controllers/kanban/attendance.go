package kanban

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/kanban"
	"time"
)

//考勤
type AttendanceController struct {
	beego.Controller
}

func (a *AttendanceController) URLMapping() {
	a.Mapping("GetAttendance", a.GetAttendance)
	a.Mapping("Leave", a.Leave)
}

// GetAttendance ...
// @Title GetAttendance
// @Description 学生总数/出勤数
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param date            query string false "日期"
// @Success 0    获取成功！
// @Failure 1005 学生总数获取失败！/ 学生出勤数获取失败！
// @router /attendance [get]
func (a *AttendanceController) GetAttendance() {
	kindergartenId, _ := a.GetInt("kindergarten_id")
	date := a.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	if valid.HasErrors() {
		a.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetAttendance(kindergartenId, date)
		if err != nil {
			a.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			a.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	a.ServeJSON()
}

// Leave ...
// @Title Leave
// @Description 请假
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param class_info      query string true  "班级名称"
// @Param date            query string false "日期"
// @Success 0    获取成功！
// @Failure 1003 获取失败！
// @Failure 1006 用户没有权限！
// @router /leave [get]
func (a *AttendanceController) Leave() {
	kindergartenId, _ := a.GetInt("kindergarten_id")
	classInfo := a.GetString("class_info")
	date := a.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	valid.Required(classInfo, "class_info").Message("班级名称 必须填写！")
	if valid.HasErrors() {
		a.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := models.Leave(kindergartenId, classInfo, date)
		if err != nil {
			a.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			a.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	a.ServeJSON()
}
