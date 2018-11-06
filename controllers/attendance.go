package controllers

import (
	"kindergarten-service-go/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// URLMapping ...
func (c *AttendanceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//考勤
type AttendanceController struct {
	beego.Controller
}

// Post 学生签到
// @Title 学生签到
// @Description 学生签到
// @Param	atten_id			formData 	int		true		"规则ID"
// @Param	student_id		formData 	int		true		"学生ID"
// @Param	sign_time		formData 	string	true		"入园时间"
// @Param	back_time		formData 	string	false	"离园时间"
// @Param	atte_time		formData 	string	true		"考勤时间"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
// @router / [post]
func (c *AttendanceController) Post() {
	AttenId, _ := c.GetInt("atten_id")
	StudentId, _ := c.GetInt("student_id")
	SignTime := c.GetString("sign_time")
	BackTime := c.GetString("back_time")
	AtteTime := c.GetString("atte_time")
	valid := validation.Validation{}
	valid.Required(AttenId, "atten_id").Message("规则ID不能为空")
	valid.Required(StudentId, "student_id").Message("学生ID不能为空")
	valid.Required(SignTime, "sign_time").Message("入园时间不能为空")
	valid.Required(AtteTime, "atte_time").Message("考勤时间")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		var v models.AttendanceDetailed
		v.AttenId = AttenId
		v.StudentId = StudentId
		v.SignTime = SignTime
		v.BackTime = BackTime
		v.AtteTime = AtteTime
		v.AttenType = 1
		if _, err := models.AddDetailed(&v); err == nil {
			c.Data["json"] = models.JSONStruct{"success", 0, v, "操作成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"error", 1003, v, "操作失败！"}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Food by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Food
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AttendanceController) GetOne() {

}

// GetAll 考勤列表（园长端）
// @Title 考勤列表（园长端）
// @Description 考勤列表（园长端）
// @Param	class_name			path	string	false	"班级名称"
// @Param	atte_time			path	string	false	"考勤时间"
// @Param	kindergarten_id		path	int	false	"幼儿园ID"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403
// @router / [get]
func (c *AttendanceController) GetAll() {
	ClassType, _ := c.GetInt("class_type")
	AtteTime := c.GetString("atte_time")
	KindergartenId, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(ClassType, "class_type").Message("班级名称不能为空")
	valid.Required(AtteTime, "atte_time").Message("考勤时间不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		v, err := models.GetAllAttendance(ClassType, AtteTime, KindergartenId)
		if err != nil {
			c.Data["json"] = models.JSONStruct{"error", 1005, v, "获取失败"}
		} else {
			if len(v) == 0 {
				c.Data["json"] = models.JSONStruct{"success", 1002, v, "没有相关数据"}
			} else {
				c.Data["json"] = models.JSONStruct{"success", 0, v, "获取成功"}
			}
		}
	}
	c.ServeJSON()
}

// Put 修改考勤状态
// @Title 修改考勤状态
// @Description 修改考勤状态
// @Param	id		path 	string	true		"考勤ID"
// @Param	atten_type		path 	string	true		"考勤状态"
// @Success 200 {object} models.Food
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AttendanceController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	AttenType := c.GetString("atten_type")
	LeaveReason := c.GetString("leave_reason")
	id, _ := strconv.Atoi(idStr)
	if _, err := models.UpdateAttensById(id, AttenType, LeaveReason); err == nil {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "修改成功！"}
	} else {
		c.Data["json"] = models.JSONStruct{"error", 1003, nil, "修改失败！"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Food
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AttendanceController) Delete() {

}

// AddRule 发布规则（校长端）
// @Title AddRule  发布规则（校长端）
// @Description 发布规则（校长端）
// @Param	sign_start_time		formData 	string	true		"签到开始时间"
// @Param	sign_end_time		formData 	string	true		"签到结束时间"
// @Param	back_start_time		formData 	string	true		"离园开始时间"
// @Param	back_end_time		formData 	string	true		"离园结束时间"
// @Param	use_time				formData 	string	true		"使用时间"
// @Param	kindergarten_id		formData 	int		true		"幼儿园ID"
// @Success 0 {object} models.Food
func (c *AttendanceController) AddRule() {
	SignStartTime := c.GetString("sign_start_time")
	SignEndTime := c.GetString("sign_end_time")
	BackStartTime := c.GetString("back_start_time")
	BackEndTime := c.GetString("back_end_time")
	UseTime := c.GetString("use_time")
	KindergartenId, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(SignStartTime, "sign_start_time").Message("签到开始时间不能为空")
	valid.Required(SignEndTime, "sign_end_time").Message("签到结束时间不能为空")
	valid.Required(BackStartTime, "back_start_time").Message("离园开始时间不能为空")
	valid.Required(BackEndTime, "back_end_time").Message("离园结束时间不能为空")
	valid.Required(UseTime, "use_time").Message("重复时间")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		var v models.Attendance
		v.SignStartTime = SignStartTime
		v.SignEndTime = SignEndTime
		v.BackStartTime = BackStartTime
		v.BackEndTime = BackEndTime
		v.UseTime = UseTime
		v.KindergartenId = strconv.Itoa(KindergartenId)
		if _, err := models.AddAttendance(&v); err == nil {
			c.Data["json"] = models.JSONStruct{"success", 0, v, "新增成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"error", 1003, v, "新增失败！"}
		}
	}
	c.ServeJSON()
}

// SignList 签到列表（教师端）
// @Title 签到列表（教师端）
// @Description 签到列表（教师端）
// @Param	class_info			path 	string	true		"班级名称"
// @Param	kindergarten_id		path 	int		true		"幼儿园ID"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) SignList() {
	ClassInfo := c.GetString("class_info")
	KindergartenId, _ := c.GetInt("kindergarten_id")
	//查询当前时间有没有考勤
	//按班级查询学生列表
	l, err := models.GetSignList(ClassInfo, KindergartenId)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1005, l, "获取失败"}
	} else {
		if l != nil {
			c.Data["json"] = models.JSONStruct{"success", 0, l, "获取成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"success", 1002, l, "没有相关数据"}
		}
	}
	c.ServeJSON()

}

// TopspeedSign 一键签到
// @Title 一键签到
// @Description 一键签到
// @Param	atten_id			formData 	int		true		"规则ID"
// @Param	student_id		formData 	string	true		"学生ID"
// @Param	sign_time		formData 	string	true		"入园时间"
// @Param	back_time		formData 	string	false	"离园时间"
// @Param	atte_time		formData 	string	true		"考勤时间"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) TopspeedSign() {
	AttenId, _ := c.GetInt("atten_id")
	StudentId := c.GetString("student_id")
	SignTime := c.GetString("sign_time")
	BackTime := c.GetString("back_time")
	AtteTime := c.GetString("atte_time")
	valid := validation.Validation{}
	valid.Required(AttenId, "atten_id").Message("规则ID不能为空")
	valid.Required(StudentId, "student_id").Message("学生ID不能为空")
	valid.Required(SignTime, "sign_time").Message("入园时间不能为空")
	valid.Required(AtteTime, "atte_time").Message("考勤时间")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		if num, err := models.TopspeedDetailed(AttenId, StudentId, SignTime, BackTime, AtteTime); err == nil {
			c.Data["json"] = models.JSONStruct{"success", 0, num, "一键入园成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"error", 1003, num, "一键入园失败！"}
		}
	}
	c.ServeJSON()
}

// Leave 请假
// @Title 请假
// @Description 请假
// @Param	atten_id			formData 	int		true		"规则ID"
// @Param	student_id		formData 	string	true		"学生ID"
// @Param	sign_time		formData 	string	true		"入园时间"
// @Param	back_time		formData 	string	false	"离园时间"
// @Param	atte_time		formData 	string	true		"考勤时间"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) Leave() {
	StudentId, _ := c.GetInt("student_id")
	AttenType, _ := c.GetInt("atten_type")
	LeaveStartTime := c.GetString("leave_start_time")
	LeaveEndTime := c.GetString("leave_end_time")
	LeaveReason := c.GetString("leave_reason")
	KindergartenId, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(AttenType, "atten_type").Message("考勤状态不能为空")
	valid.Required(StudentId, "student_id").Message("学生ID不能为空")
	valid.Required(LeaveStartTime, "leave_start_time").Message("请假开始时间不能为空")
	valid.Required(LeaveEndTime, "leave_end_time").Message("请假结束时间不能为空")
	valid.Required(LeaveReason, "leave_reason").Message("请假原因不能为空")
	valid.Required(KindergartenId, "kindergarten_id").Message("幼儿园ID")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		var v models.AttendanceDetailed
		v.StudentId = StudentId
		v.AttenType = AttenType
		v.LeaveStartTime = LeaveStartTime
		v.LeaveEndtime = LeaveEndTime
		v.LeaveReason = LeaveReason
		if _, err := models.LeaveDetailed(&v, KindergartenId); err == nil {
			c.Data["json"] = models.JSONStruct{"success", 0, nil, "请假成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"error", 1003, nil, "请假失败！"}
		}
	}
	c.ServeJSON()
}

// LeaveGarden 离园
// @Title 离园
// @Description 离园
// @Param	student_id		path 	int	true		"学生ID"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) LeaveGarden() {
	StudentId, _ := c.GetInt("student_id")
	_, err := models.LeaveGarden(StudentId)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1003, nil, "操作失败"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "操作成功！"}
	}
	c.ServeJSON()

}

// TopspeedLeaveGarden 一键离园
// @Title 一键离园
// @Description 一键离园
// @Param	class_info		path 	string	true		"班级名称"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) TopspeedLeaveGarden() {
	ClassInfo := c.GetString("class_info")
	err := models.TopspeedLeaveGarden(ClassInfo)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1003, nil, "操作失败"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, nil, "操作成功！"}
	}
	c.ServeJSON()

}

// AttendanceGarden 园长端考勤表图（整个幼儿园）
// @Title 园长端考勤表图（整个幼儿园）
// @Description 园长端考勤表图（整个幼儿园）
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) AttendanceGarden() {
	KindergartenId, _ := c.GetInt("kindergarten_id")
	v, err := models.AttendanceGarden(KindergartenId)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1003, v, "获取失败"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, v, "获取成功！"}
	}
	c.ServeJSON()

}

// GetDetailed 园长端考勤详细（幼儿园某个班级）
// @Title 园长端考勤详细（幼儿园某个班级）
// @Description 园长端考勤详细（幼儿园某个班级）
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Param	class		path 	string	true		"班"
// @Param	class_name		path 	string	true		"班级"
// @Param	atte_time		path 	string	true		"考勤时间"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) GetDetailed() {
	KindergartenId, _ := c.GetInt("kindergarten_id")
	Class := c.GetString("class")
	ClassName := c.GetString("class_name")
	AtteTime := c.GetString("atte_time")
	v, err := models.GetDetailed(KindergartenId, Class, ClassName, AtteTime)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1003, v, "获取失败"}
	} else {
		c.Data["json"] = models.JSONStruct{"success", 0, v, "获取成功！"}
	}
	c.ServeJSON()

}

// GetDetailed 园长端获取考勤规则
// @Title 园长端获取考勤规则
// @Description 园长端获取考勤规则）
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 0 {object} models.AttendanceDetailed
// @Failure 403 body is empty
func (c *AttendanceController) GetRUle() {
	KindergartenId, _ := c.GetInt("kindergarten_id")
	v, err := models.GetRUle(KindergartenId)
	if err != nil {
		c.Data["json"] = models.JSONStruct{"error", 1003, v, "获取失败"}
	} else {
		if v != nil {
			c.Data["json"] = models.JSONStruct{"success", 0, v, "获取成功！"}
		} else {
			c.Data["json"] = models.JSONStruct{"success", 1002, v, "没有相关数据"}
		}
	}
	c.ServeJSON()

}
