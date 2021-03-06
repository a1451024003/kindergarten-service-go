package healthy

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/healthy"
	"strconv"
)

//餐检
type InspectController struct {
	beego.Controller
}

// URLMapping ...
func (c *InspectController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 添加检查
// @Description 添加喂药申请
// @Param   class_id     			formData    int  	true        "班级ID"
// @Param   student_id     			formData    int  	true        "学生ID"
// @Param   types     				formData    int	    true        "检查类型"
// @Param   abnormal     			formData    string  true        "异常情况"
// @Param   handel     				formData    string  true        "处理方式"
// @Param   url     				formData    string  true        "照片留档"
// @Param   infect     				formData    string  true        "是否传染（1，否2，是）"
// @Success 0 {int} healthy.Inspect
// @Failure 1001 补全信息
// @Failure 1003 保存失败
// @router / [post]
func (c *InspectController) Post() {
	class_name := c.GetString("class_name")
	class_id, _ := c.GetInt("class_id")
	student_id, _ := c.GetInt("student_id")
	types, _ := c.GetInt("types")
	abnormal := c.GetString("abnormal")
	handel := c.GetString("handel")
	url := c.GetString("url")
	infect, _ := c.GetInt("infect")
	drug_id, _ := c.GetInt("drug_id")
	teacher_id, _ := c.GetInt("teacher_id")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	date := c.GetString("date")
	content := c.GetString("content")

	valid := validation.Validation{}
	valid.Required(class_name, "class_name").Message("班级名称不能为空")
	valid.Required(student_id, "student_id").Message("学生ID不能为空")
	valid.Required(class_id, "class_id").Message("班级ID不能为空")
	valid.Required(teacher_id, "teacher_id").Message("教师ID不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")

	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, "", valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}

	w := healthy.Inspect{
		StudentId:      student_id,
		ClassId:        class_id,
		Types:          types,
		Abnormal:       abnormal,
		Handel:         handel,
		Url:            url,
		Infect:         infect,
		DrugId:         drug_id,
		TeacherId:      teacher_id,
		KindergartenId: kindergarten_id,
		ClassName:      class_name,
		Date:           date,
		Content:        content,
	}
	if err := w.Save(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, "", "记录成功"}
	} else {
		fmt.Println(err)
		c.Data["json"] = JSONStruct{"error", 1003, "", "记录失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 餐检列表
// @Param	page			query	int		false		"第几页"
// @Param	kindergarten_id	query	int		true		"幼儿园ID"
// @Param	per_page		query	int		true		"页数"
// @Param	class_id		query	int		false		"班级ID"
// @Param	role			query	int		true		"身份类型"
// @Param	date			query	string	true		"餐检时间"
// @Success 0 {object} 		healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router / [get]
func (c *InspectController) GetAll() {
	var f *healthy.Inspect
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_id, _ := c.GetInt("class_id")
	types, _ := c.GetInt("types")
	perPage, _ := c.GetInt("per_page")
	role, _ := c.GetInt("role")
	date := c.GetString("date")
	bady_id, _ := c.GetInt("baby_id")
	search := c.GetString("search")

	//验证参数是否为空
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	valid.Required(kindergarten_id, "role").Message("用户身份不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	fmt.Println(date)
	if works, err := f.GetAll(page, perPage, kindergarten_id, class_id, types, role, bady_id, date, search); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description 删除
// @Param	id		path 	string	true		"自增ID"
// @Success 0 {string} delete success!
// @Failure 1003 id is empty
// @router /:id [delete]
func (c *InspectController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := healthy.DeleteInspect(id)
	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1003, nil, "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 统计
// @Param	kindergarten_id	query	int	true		"幼儿园ID"
// @Success 0 {object} 	healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1003 		获取失败
// @router /counts/ [get]
func (c *InspectController) Counts() {
	kindergarten_id, _ := c.GetInt("kindergarten_id")

	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}
	v := healthy.Countss(kindergarten_id)

	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1003, nil, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 详情
// @Param	kindergarten_id	query	int	true		"幼儿园ID"
// @Success 0 {object} 	healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1003 		获取失败
// @router /:id
func (c *InspectController) Inspect() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := healthy.InspectInfo(id)
	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// @Title 编辑检查
// @Description 添加喂药申请
// @Param   class_id     			formData    int  	true        "班级ID"
// @Param   student_id     			formData    int  	true        "学生ID"
// @Param   types     				formData    int	    true        "检查类型"
// @Param   abnormal     			formData    string  true        "异常情况"
// @Param   handel     				formData    string  true        "处理方式"
// @Param   url     				formData    string  true        "照片留档"
// @Param   infect     				formData    string  true        "是否传染（1，否2，是）"
// @Success 0 {int} healthy.Inspect
// @Failure 1001 补全信息
// @Failure 1003 保存失败
// @router /:id [put]
func (c *InspectController) Put() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")

	class_name := c.GetString("class_name")
	class_id, _ := c.GetInt("class_id")
	student_id, _ := c.GetInt("student_id")
	types, _ := c.GetInt("types")
	abnormal := c.GetString("abnormal")
	handel := c.GetString("handel")
	url := c.GetString("url")
	infect, _ := c.GetInt("infect")
	drug_id, _ := c.GetInt("drug_id")
	teacher_id, _ := c.GetInt("teacher_id")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	content := c.GetString("content")
	date := c.GetString("date")

	valid := validation.Validation{}
	valid.Required(student_id, "student_id").Message("学生ID不能为空")
	valid.Required(class_id, "class_id").Message("班级ID不能为空")
	valid.Required(types, "types").Message("检查类型不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")

	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, "", valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}

	w := healthy.Inspect{
		Id:             id,
		StudentId:      student_id,
		ClassId:        class_id,
		Types:          types,
		Abnormal:       abnormal,
		Handel:         handel,
		Url:            url,
		Infect:         infect,
		DrugId:         drug_id,
		TeacherId:      teacher_id,
		KindergartenId: kindergarten_id,
		ClassName:      class_name,
		Content:        content,
		Date:           date,
	}
	if err := w.SaveInspect(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, "", "申请成功"}
	} else {
		fmt.Println(err)
		c.Data["json"] = JSONStruct{"error", 1003, "", "申请失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 异常档案列表
// @Param	page			query	int		false		"第几页"
// @Param	kindergarten_id	query	int		true		"幼儿园ID"
// @Param	per_page		query	int		true		"页数"
// @Param	class_id		query	int		false		"班级ID"
// @Param	role			query	int		true		"身份类型"
// @Param	date			query	string	true		"餐检时间"
// @Success 0 {object} 		healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /archives/ [get]
func (c *InspectController) Abnormal() {
	var f *healthy.Inspect
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_id, _ := c.GetInt("class_id")
	perPage, _ := c.GetInt("per_page")
	date := c.GetString("date")
	search := c.GetString("search")
	types, _ := c.GetInt("types")
	role, _:= c.GetInt("role")
	pType := 2

	//验证参数是否为空
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	if works, err := f.Abnormals(pType, types, role, page, perPage, kindergarten_id, class_id, date, search); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 项目详情
// @Param	page			query	int		false		"第几页"
// @Param	per_page		query	int		true		"页数"
// @Param	class_id		query	int		false		"班级ID"
// @Param	role			query	int		true		"身份类型"
// @Param	date			query	string	true		"餐检时间"
// @Success 0 {object} 		healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /project/ [get]
func (c *InspectController) Project() {
	var f *healthy.Inspect
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_id, _ := c.GetInt("class_id")
	perPage, _ := c.GetInt("per_page")
	body_id, _ := c.GetInt("body_id")
	baby_id, _ := c.GetInt("baby_id")
	search := c.GetString("search")

	if works, err := f.Projects(page, perPage, kindergarten_id, class_id, body_id, baby_id, search); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 项目详情
// @Param	page			query	int		false		"第几页"
// @Param	per_page		query	int		true		"页数"
// @Param	class_id		query	int		false		"班级ID"
// @Param	role			query	int		true		"身份类型"
// @Param	date			query	string	true		"餐检时间"
// @Success 0 {object} 		healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /projectNew/ [get]
func (c *InspectController) ProjectNew() {
	var f *healthy.Inspect
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_id, _ := c.GetInt("class_id")
	perPage, _ := c.GetInt("per_page")
	body_id, _ := c.GetInt("body_id")
	baby_id, _ := c.GetInt("baby_id")
	column := c.GetString("column")

	if works, err := f.ProjectNew(page, perPage, kindergarten_id, class_id, body_id, baby_id, column); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 体重健康统计
// @Param	kindergarten_id		query	int		true		"幼儿园ID"
// @Param	types				query	string	true		"类型"
// @Success 0 {object} 			healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /weight/ [get]
func (c *InspectController) Weight() {
	var f *healthy.Inspect

	kindergarten_id, _ := c.GetInt("kindergarten_id")
	date := c.GetString("date")

	//验证参数是否为空
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	if works, err := f.Weights(kindergarten_id, date); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description 身高健康统计
// @Param	kindergarten_id		query	int		true		"幼儿园ID"
// @Param	types				query	string	true		"类型"
// @Success 0 {object} 			healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /height/ [get]
func (c *InspectController) Height() {
	var f *healthy.Inspect

	kindergarten_id, _ := c.GetInt("kindergarten_id")
	date := c.GetString("date")

	//验证参数是否为空
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	if works, err := f.Heights(kindergarten_id, date); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description 全园统计
// @Param	kindergarten_id		query	int		true		"幼儿园ID"
// @Param	types				query	string	true		"类型"
// @Success 0 {object} 			healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /country/ [get]
func (c *InspectController) Country() {
	var f *healthy.Inspect
	kindergarten_id, _ := c.GetInt("kindergarten_id")

	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	if works, err := f.Country(kindergarten_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, works, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()

}

// Delete ...
// @Title Delete
// @Description 删除
// @Param	id		path 	int	true		"自增ID"
// @Success 0 {string} delete success!
// @Failure 1003 id is empty
// @router /deleteStudent/:id [delete]
func (c *InspectController) DeleteStudent() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := healthy.DeleteStudent(id)
	if v == nil {
		c.Data["json"] = JSONStruct{"error", 1003, nil, "删除失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	}
	c.ServeJSON()
}

// PUT ...
// @Title PUT
// @Description 添加备注
// @Param	id		path 	int	true		"体检ID"
// @Success 200 {string} put success!
// @Failure 403 id is empty
// @router /content/:id [put]
func (c *InspectController) Content() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	f := healthy.Inspect{Id: id}
	content := c.GetString("content")
	if err := f.Contents(content); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, f, "操作成功"}
	} else if err == orm.ErrNoRows {
		c.Data["json"] = JSONStruct{"error", 1002, err, "用户不存在"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1003, err, "操作失败"}
	}

	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description 幼儿园异常儿童
// @Param	kindergarten_id		query	int		true		"幼儿园ID"
// @Success 0 {object} 			healthy.Inspect
// @Failure 1001 		参数不能为空
// @Failure 1005 		获取失败
// @router /abnormallist/ [get]
func (c *InspectController) AbnormalList() {
	var f *healthy.Inspect
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	class_id, _:= c.GetInt("class_id")

	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, struct{}{}, valid.Errors[0].Message}
		c.ServeJSON()
		c.StopRun()
	}
	if body,err := f.AbnormalList(kindergarten_id, class_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, body, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}

	c.ServeJSON()
}
