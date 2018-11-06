package kanban

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/kanban"
	"time"
)

//日程
type ScheduleController struct {
	beego.Controller
}

func (s *ScheduleController) URLMapping() {
	s.Mapping("PostSchedule", s.PostSchedule)
	s.Mapping("GetScheduleList", s.GetScheduleList)
	s.Mapping("GetScheduleInfo", s.GetScheduleInfo)
	s.Mapping("PutSchedule", s.PutSchedule)
	s.Mapping("DeleteSchedule", s.DeleteSchedule)
}

type JSONStruct struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

// PostSchedule ...
// @Title PostSchedule
// @Description 日程添加
// @Param content         query string true "日程内容"
// @Param time            query string true "日程时间"
// @Param kindergarten_id query int    true "幼儿园ID"
// @Param type            query int    true "类型：1日程，2备忘录"
// @Param user_id         query int    true "用户ID"
// @Param role            query int    true "用户角色"
// @Success 0    添加成功！
// @Failure 1003 添加失败！
// @Failure 1006 用户没有权限！
// @router / [post]
func (s *ScheduleController) PostSchedule() {
	content := s.GetString("content")
	times := s.GetString("time")
	kindergartenId, _ := s.GetInt("kindergarten_id")
	ty, _ := s.GetInt("type")
	userId, _ := s.GetInt("user_id")
	role, _ := s.GetInt("role")

	valid := validation.Validation{}
	valid.Required(content, "content").Message("日程内容 必须填写！")
	valid.Required(times, "time").Message("日程时间 必须填写！")
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	valid.Range(ty, 1, 2, "type").Message("事件类型 必须填写！")
	valid.Required(userId, "user_id").Message("用户ID 必须填写！")
	valid.Required(role, "role").Message("用户角色 必须填写！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := models.PostSchedule(content, times, kindergartenId, ty, userId, role)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, []string{}, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, data, "添加成功！"}
		}
	}
	s.ServeJSON()
}

// GetScheduleList ...
// @Title GetScheduleList
// @Description 日程列表
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param user_id         query int    true  "用户ID"
// @Param date            query string false "日期"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router / [get]
func (s *ScheduleController) GetScheduleList() {
	kindergartenId, _ := s.GetInt("kindergarten_id")
	userId, _ := s.GetInt("user_id")
	date := s.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	valid.Required(userId, "user_id").Message("用户ID 必须填写！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetScheduleList(kindergartenId, userId, date)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	s.ServeJSON()
}

// GetScheduleInfo ...
// @Title GetScheduleInfo
// @Description 日程详情
// @Param schedule_id query int true "日程ID"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /details [get]
func (s *ScheduleController) GetScheduleInfo() {
	scheduleId, _ := s.GetInt("schedule_id")

	valid := validation.Validation{}
	valid.Required(scheduleId, "schedule_id").Message("日程ID 必须填写！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetScheduleInfo(scheduleId)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, []string{}, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	s.ServeJSON()
}

// PutSchedule ...
// @Title PutSchedule
// @Description 日程修改
// @Param schedule_id     query int    true "日程ID"
// @Param content         query string true "日程内容"
// @Param time            query string true "日程时间"
// @Param kindergarten_id query int    true "幼儿园ID"
// @Param user_id         query int    true "用户ID"
// @Param role            query int    true "用户角色"
// @Success 0    修改成功！
// @Failure 1003 修改失败！
// @router / [put]
func (s *ScheduleController) PutSchedule() {
	scheduleId, _ := s.GetInt("schedule_id")
	content := s.GetString("content")
	times := s.GetString("time")
	kindergartenId, _ := s.GetInt("kindergarten_id")
	userId, _ := s.GetInt("user_id")
	role, _ := s.GetInt("role")

	valid := validation.Validation{}
	valid.Required(scheduleId, "schedule_id").Message("日程ID 必须填写！")
	valid.Required(content, "content").Message("日程内容 必须填写！")
	valid.Required(times, "time").Message("日程时间 必须填写！")
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	valid.Required(userId, "user_id").Message("用户ID 必须填写！")
	valid.Range(role, 1, 1, "role").Message("用户没有权限！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		code, err := models.PutSchedule(scheduleId, content, times, kindergartenId, userId)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, []string{}, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, []string{}, "修改成功！"}
		}
	}
	s.ServeJSON()
}

// DeleteSchedule ...
// @Title DeleteSchedule
// @Description 日程删除
// @Param schedule_id query int true "日程ID"
// @Success 0    删除成功！
// @Failure 1003 删除失败！
// @router / [delete]
func (s *ScheduleController) DeleteSchedule() {
	scheduleId, _ := s.GetInt("schedule_id")

	valid := validation.Validation{}
	valid.Required(scheduleId, "schedule_id").Message("日程ID 必须填写！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		code, err := models.DeleteSchedule(scheduleId)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, []string{}, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, []string{}, "删除成功！"}
		}
	}
	s.ServeJSON()
}
