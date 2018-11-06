package admin

import (
	"kindergarten-service-go/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

//学生
type StudentController struct {
	BaseController
}

// GetStudent ...
// @Title 学生列表
// @Description 学生列表
// @Param	kindergarten_id       query	int	     true		"幼儿园ID"
// @Param	status                query	int	     false		"状态"
// @Param	search                query	int	     false		"搜索条件"
// @Param	page                  query	int	     false		"页数"
// @Param	per_page              query	int	     false		"每页显示条数"
// @Success 200 {object} models.Student
// @Failure 403
// @router / [get]
func (c *StudentController) GetStudent() {
	search := c.GetString("search")
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	date := c.GetString("date")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	status, _ := c.GetInt("status", -1)
	v, err := models.GetStudent(kindergarten_id, status, search, page, prepage, date)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// GetStudentClass ...
// @Title 学生班级搜索
// @Description 学生班级搜索
// @Param	kindergarten_id       query	int	     true		"幼儿园ID"
// @Param	class_type            query	int	     true		"班级类型"
// @Param	page                  query	int	     false		"页数"
// @Param	per_page              query	int	     false		"每页显示条数"
// @Success 200 {object} models.Student
// @Failure 403
// @router /class [get]
func (c *StudentController) GetStudentClass() {
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_type, _ := c.GetInt("class_type")
	v, err := models.GetStudentClass(kindergarten_id, class_type, page, prepage)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// RemoveStudent ...
// @Title RemoveStudent
// @Description 移除学生
// @Param	student_id		path 	    int	true		"学生ID"
// @Param	class_id		    path 	    int	true		"班级ID"
// @Success 200 {string} delete success!
// @Failure 403 student_id is empty
// @router /remove [delete]
func (c *StudentController) RemoveStudent() {
	student_id, _ := c.GetInt("student_id")
	class_id, _ := c.GetInt("class_id")
	err := models.RemoveStudent(class_id, student_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1004, nil, err.Error()}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "移除成功"}
	}
	c.ServeJSON()
}

// GetStudentInfo ...
// @Title Get Student Info
// @Description 学生详情
// @Param	student_id       query	int	 true		"学生编号"
// @Success 200 {object} models.Student
// @Failure 403 :学生编号为空
// @router /:id [get]
func (c *StudentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetStudentInfo(id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// UpdateStudent ...
// @Title 编辑学生
// @Description 编辑学生
// @Param	id		    path 	int	               true		    "学生编号"
// @Param	body		body 	models.Student	       true		"param(json)"
// @Success 200 {object} models.Student
// @Failure 403 :id is not int
// @router /:id [put]
func (c *StudentController) UpdateStudent() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	student := c.GetString("student")
	kinship := c.GetString("kinship")
	valid := validation.Validation{}
	valid.Required(student, "student").Message("学生信息不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		_, err := models.UpdateStudent(id, student, kinship)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
		}
		c.ServeJSON()
	}
}

// Post ...
// @Title 学生-录入信息
// @Description 学生-录入信息
// @Param	body		body 	models.Student	true		"json"
// @Success 201 {int} models.Student
// @Failure 403 body is empty
// @router / [post]
func (c *StudentController) Post() {
	student := c.GetString("student")
	kinship := c.GetString("kinship")
	valid := validation.Validation{}
	valid.Required(student, "student").Message("学生信息不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		_, err := models.AddStudent(student, kinship)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
		}
		c.ServeJSON()
	}
}

// Invite ...
// @Title 邀请学生/批量邀请
// @Description 邀请学生/批量邀请
// @Param	name		        body 	int   	true		"学生姓名(json)"
// @Param	baby_id		        body 	int   	true		"宝宝id(json)"
// @Param	kindergarten_id		body 	int   	true		"幼儿园id(json)"
// @Success 201 {int} models.Student
// @Failure 403 body is empty
// @router /invite [post]
func (c *StudentController) Invite() {
	student := c.GetString("student")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	err := models.Invites(student, kindergarten_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, nil, err.Error()}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "邀请成功"}
	}
	c.ServeJSON()
}

// DeleteStudent ...
// @Title DeleteStudent
// @Description 删除学生
// @Param	student_id		path 	int	true		"学生ID"
// @Param	status		    path 	int	true		"状态(status 0:未分班 2:离园)"
// @Param	type		        path 	int	true		"删除类型（type 0:学生离园 1:删除档案）"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *StudentController) DeleteStudent() {
	class_type, _ := c.GetInt("class_type")
	status, _ := c.GetInt("status")
	ty, _ := c.GetInt("type")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteStudent(id, status, ty, class_type)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1004, nil, "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}

// GetNameClass ...
// @Title GetNameClass
// @Description 学生名字获取班级
// @Param	name       query	 int	 true		"学生姓名"
// @Success 200 {object} models.Student
// @Failure 403 :幼儿园id不能为空
// @router /get_class [get]
func (c *StudentController) GetNameClass() {
	name := c.GetString("name")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("学生姓名不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.GetNameClass(name, kindergarten_id)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// BabyKindergarten ...
// @Title BabyKindergarten
// @Description baby是否加入幼儿园
// @Param	baby_id       query	 int	 true		"宝贝id"
// @Success 200 {object} models.Student
// @Failure 403 :幼儿园id不能为空
// @router /baby_kinder [get]
func (c *StudentController) BabyKindergarten() {
	baby_id, _ := c.GetInt("baby_id")
	v, err := models.BabyKinder(baby_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// BabyActived ...
// @Title baby是否在幼儿园
// @Description baby是否在幼儿园
// @Param	baby_id       query	 int	 true		"宝贝id"
// @Success 200 {object} models.Student
// @Failure 403 :幼儿园id不能为空
// @router /baby_actived [get]
func (c *StudentController) BabyActived() {
	baby_id, _ := c.GetInt("baby_id")
	v, err := models.BabyActived(baby_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// UpStudent ...
// @Title 编辑学生
// @Description 编辑学生
// @Param	id		    path 	int	                   true		    "学生编号"
// @Param	body		    body 	models.Student	       true		"param(json)"
// @Success 200 {object} models.Student
// @Failure 403 :id is not int
// @router /auth/:id [put]
func (c *StudentController) UpStudent() {
	ty, _ := c.GetInt("type")
	class, _ := c.GetInt("class")
	student := c.GetString("student")
	kinship := c.GetString("kinship")
	member_ids := c.GetString("member_ids")
	is_principal, _ := c.GetInt("is_principal")
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	organizational_id, _ := c.GetInt("organizational_id")
	valid := validation.Validation{}
	valid.Required(student, "student").Message("学生信息不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		err := models.UpStudent(id, student, kinship, ty, member_ids, is_principal, organizational_id, class)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
		}
		c.ServeJSON()
	}
}

// ResetInvite ...
// @Title ResetInvite
// @Description 取消邀请·在次邀请
// @Param	baby_id		path 	int	true		"宝宝id"
// @Param	type		path 	int	true		"1 取消邀请  2 再次邀请"
// @Success 200 {string} success!
// @Failure 403 id is empty
// @router /reset_invite [post]
func (c *StudentController) ResetInvite() {
	ty, _ := c.GetInt("type")
	student := c.GetString("student")
	baby_id, _ := c.GetInt("baby_id")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	err := models.ResetInvite(ty, baby_id, student, kindergarten_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1004, nil, "失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "成功"}
	}
	c.ServeJSON()
}

// GetBaby ...
// @Title GetBaby
// @Description 未激活baby
// @Param	kindergarten_id       query	 int	 true		"幼儿园id"
// @Success 200 {object} models.Student
// @Failure 403 :幼儿园id不能为空
// @router /baby [get]
func (c *StudentController) GetBaby() {
	invite, _ := c.GetInt("invite", -1)
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.GetBabyInfo(kindergarten_id, invite)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// OragnizationalStudent ...
// @Title 班级所有学生
// @Description 班级所有学生
// @Param	class_id                  query	int	     false		"班级id"
// @Param	page                      query	int	     false		"页数"
// @Param	per_page                  query	int	     false		"每页显示条数"
// @Success 200 {object} models.Organizational
// @Failure 403
// @router /member [get]
func (o *StudentController) OragnizationalStudent() {
	prepage, _ := o.GetInt("per_page", 20)
	page, _ := o.GetInt("page")
	class_id, _ := o.GetInt("class_id")
	v, err := models.OragnizationalStudent(class_id, page, prepage)
	if err != nil {
		o.Data["json"] = JSONStruct{"error", 1005, nil, err.Error()}
	} else {
		o.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	o.ServeJSON()
}

// DeleteKinship ...
// @Title DeleteKinship
// @Description 删除亲属信息
// @Param	kinship_id		path 	int	true		"学生ID"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /kinship/:id [delete]
func (c *StudentController) DeleteKinship() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteKinship(id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1004, nil, "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}
