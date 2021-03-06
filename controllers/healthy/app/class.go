package app

import (
	"kindergarten-service-go/models/healthy"
	"strconv"

	"github.com/astaxie/beego"
	"encoding/json"
)

//体检主题班级
type ClassController struct {
	beego.Controller
}

// URLMapping ...
func (c *ClassController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title
// @Description 添加主题
// @Param	body_name			query	string	false	"主题名"
// @Param	kindergarten_id		query	int	false	"幼儿园id"
// @Param	test_time			query	string	false	"年-月-日"
// @Param	types		query	int	false	"类型"
// @Param	class_id		query	int	false	"班级id"
// @Param	class_total		query	int	false	"总人数"
// @Param	class_actual	query	int	false	""
// @Param	test_time	query	string	false	"测评日期"
// @Param	class_rate	query	int	false	""
// @Success 201 {int} healthy.Class
// @Failure 403 body is empty
// @router / [post]
func (c *ClassController) Post() {
	body_name := c.GetString("body_name")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	test_time := c.GetString("test_time")
	types, _ := c.GetInt("types")
	class_id, _ := c.GetInt("class_id")

	body_id, _ := healthy.CrBody(body_name, kindergarten_id, test_time, types,class_id)
	class_total, _ := c.GetInt("class_total")
	class_actual, _ := c.GetInt("class_actual")
	class_rate, _ := c.GetInt("class_rate")
	var b healthy.Class
	b.BodyId = int(body_id)
	b.ClassRate = class_rate
	b.ClassActual = class_actual
	b.ClassTotal = class_total
	b.ClassId = class_id
	if err := healthy.AddClass(&b, int(body_id), class_id, types); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, nil, "添加成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1001, err.Error(), "添加失败"}
	}
	c.ServeJSON()
}

// Post ...
// @Title
// @Description 添加主题
// @Param	data			query	string	false	"json 串"
// @Success 201 {int} healthy.Class
// @Failure 403 body is empty
// @router /add_info [post]
func (c *ClassController) Post_info() {
	data := c.GetString("data")
	var i []healthy.Inspect_add
	json.Unmarshal([]byte(data), &i)
	if i[0].StudentId == 0 {
		c.Data["json"] = JSONStruct{"success", 1002, nil, "学生ID不能为空"}
	}else {
		if err := healthy.AddlistInspect(data); err == nil {
			c.Data["json"] = JSONStruct{"success", 0, nil, "添加成功"}
		} else {
			c.Data["json"] = JSONStruct{"error", 1001, err, "添加失败"}
		}
	}

	c.ServeJSON()
}

// Put ...
// @Title 修改体检主题
// @Description 修改体检主题
// @Param	theme		query	string	false	"主题名"
// @Param	total		query	int	false	"总人数"
// @Param	actual		query	int	false	"实际人数"
// @Param	rate		query	string	false	"合格率"
// @Param	test_time	query	string	false	"测评日期"
// @Param	mechanism	query	int	false	"体检机构id"
// @Param	kindergarten_id	query	int	false	"幼儿园id"
// @Param	types	query	int	false	"类型 1，体质测评2，体检"
// @Param	project	query	string	false	"体检项目"
// @Success 200 {object} healthy.Class
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ClassController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var b healthy.Class
	body_id, _ := c.GetInt("body_id")
	class_id, _ := c.GetInt("class_id")
	class_total, _ := c.GetInt("class_total")
	//class_time := c.GetString("class_time")
	class_actual, _ := c.GetInt("class_actual")
	class_rate, _ := c.GetInt("class_rate")
	b.Id = id
	b.BodyId = body_id
	b.ClassRate = class_rate
	b.ClassActual = class_actual
	b.ClassTotal = class_total
	b.ClassId = class_id
	if err := healthy.UpdataByIdClass(&b); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, nil, "添加成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1001, err.Error(), "添加失败"}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title 体检主题列表
// @Description 体检主题列表
// @Param	page	query	int	false	"页"
// @Param	per_page	query	int	false	"每页条数"
// @Param	types	query	int	false	"类型"
// @Param	theme	query	string	false	"名字搜索"
// @Success 200 {object} healthy.Class
// @Failure 403
// @router / [get]
func (c *ClassController) GetAll() {
	page := 1
	per_page := 20
	if v, err := c.GetInt("page"); err == nil {
		page = v
	}
	if v, err := c.GetInt("per_page"); err == nil {
		per_page = v
	}
	class_id, _ := c.GetInt("class_id")
	body_id, _ := c.GetInt("body_id")
	if l, err := healthy.GetAllClass(page, per_page, class_id, body_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, l, "添加成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1001, err.Error(), "添加失败"}
	}
	c.ServeJSON()
}
