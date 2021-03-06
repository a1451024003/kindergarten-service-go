package admin

import (
	"kindergarten-service-go/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type ExceptionalChildController struct {
	BaseController
}

func (c *ExceptionalChildController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.Get)
}

// GetAll ...
// @Title 特殊儿童列表/搜索特殊儿童
// @Description 特殊儿童列表/搜索特殊儿童
// @Param 	type				query	int		true	"获取数据类型1过敏儿童数据，2体质类型儿童数据"
// @Param	child_name			query	string	false	"特殊儿童姓名"
// @Param	page				query	int		false	"当前页，默认为1"
// @Param	per_page			query	string	false	"每页显示条数，默认为10"
// @Param	keyword				query	string	false	"关键字(特殊儿童姓名/特殊儿童过敏源)"
// @Param	kindergarten_id		query	int		false	"幼儿园ID"
// @Success 0 			{object} 	models.ExceptionalChild
// @Failure 1005 获取失败
// @router / [get]
func (c *ExceptionalChildController) GetAll() {
	// 获取类型
	get_type, _ := c.GetInt("type")
	child_name := c.GetString("child_name")
	// 关键字
	keyword := c.GetString("keyword")
	// page_num
	page, _ := c.GetInt64("page")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")

	// limit
	limit, _ := c.GetInt64("per_page")
	if info, err := models.GetAllExceptionalChild(child_name, get_type, page, limit, keyword, kindergarten_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, info, "获取成功"}

	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}
	c.ServeJSON()
}

// Post ...
// @Title 						新增特殊儿童
// @Description 				新增特殊儿童
// @Param	child_name			formData  string	true		"特殊儿童姓名"
// @Param	class				formData  int 		true		"特殊儿童班级"
// @Param	somatotype			formData  int		true		"体质类型"
// @Param	allergen			formData  string	false		"过敏源"
// @Param	source				formData  int		true		"信息来源"
// @Param	kindergarten_id		formData  int		true		"幼儿园ID"
// @Param	creator				formData  int		true		"创建人"
// @Param	student_id			formData  int		true		"学生ID"
// @Success 0					{json} JSONSturct
// @Failure 1003 				新增失败
// @router / [post]
func (c *ExceptionalChildController) Post() {
	// 特殊儿童姓名
	child_name := c.GetString("child_name")
	// 特殊儿童班级
	class, _ := c.GetInt("class")
	// 体质类型
	somatotype, _ := c.GetInt8("somatotype")
	// 过敏源
	allergen := c.GetString("allergen")
	// 信息来源
	source, _ := c.GetInt8("source")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	// 创建人
	creator, _ := c.GetInt("creator")
	// 学生ID
	student_id, _ := c.GetInt("student_id")

	valid := validation.Validation{}

	valid.Required(child_name, "child_name").Message("儿童姓名不能为空")
	valid.Required(class, "class").Message("所在班级不能为空")
	valid.Required(somatotype, "somatotype").Message("体质类型不能为空")
	valid.Required(source, "source").Message("信息来源不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	valid.Required(creator, "creator").Message("创建人不能为空")
	valid.Required(student_id, "student_id").Message("学生ID不能为空")

	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		if id, err := models.AddExceptionalChild(child_name, class, somatotype, allergen, source, kindergarten_id, creator, student_id); err == nil {
			if id == 0 {
				c.Data["json"] = JSONStruct{"error", 1007, err, "已有此数据"}
			} else {
				c.Data["json"] = JSONStruct{"success", 0, nil, "新增成功"}
			}
		} else {
			c.Data["json"] = JSONStruct{"error", 1003, err, "新增失败"}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title 			按ID查询特殊儿童
// @Description 	按ID查询特殊儿童
// @Param	id		path 	string	true		"特殊儿童ID"
// @Success 0		{object}  models.ExceptionalChild
// @Failure 1005 	获取失败
// @router /:id [get]
func (c *ExceptionalChildController) GetOne() {
	// 主键ID
	idStr := c.Ctx.Input.Param(":id")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	v, err := models.GetExceptionalChildById(idStr, kindergarten_id)
	if err == nil {
		if v != nil {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		} else {
			c.Data["json"] = JSONStruct{"error", 1002, err, "没有相关数据"}
		}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}
	c.ServeJSON()
}

// Put ...
// @Title 					更新特殊儿童
// @Description 				更新特殊儿童
// @Param	id				path 	string	false		"特殊儿童主键自增ID"
// @Param	child_name		body 	string	false		"特殊儿童姓名"
// @Param	class			body 	int		false		"特殊儿童班级"
// @Param	somatotype		body 	int		false		"体质类型"
// @Param	allergen		body 	string	false		"过敏源"
// @Param	student_id		body 	int		false		"学生ID"
// @Success 0 				{string} 	success
// @Failure 1003			更新失败
// @router /:id [put]
func (c *ExceptionalChildController) Put() {
	// 主键ID
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	// 姓名
	child_name := c.GetString("child_name")
	// 班级ID
	class, _ := c.GetInt("class")
	// 体质类型
	somatotype, _ := c.GetInt8("somatotype")
	// 过敏源
	allergen := c.GetString("allergen")
	// 学生ID
	student_id, _ := c.GetInt("student_id")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")

	if num, err := models.UpdateExceptionalChildById(id, child_name, class, somatotype, allergen, student_id, kindergarten_id); err == nil {
		if num == 0 {
			c.Data["json"] = JSONStruct{"error", 1007, err, "已有此数据"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "更新成功"}
		}
	} else {
		c.Data["json"] = JSONStruct{"error", 1003, err, "更新失败"}
	}
	c.ServeJSON()
}

// Put ...
// @Title 					更新体检特殊儿童
// @Description 				更新体检特殊儿童
// @Param	id				path 	string	false		"体检主键自增ID"
// @Param	child_name		body 	string	false		"儿童姓名"
// @Param	class_id		body 	int		false		"儿童班级"
// @Param	somatotype		body 	int		false		"体质类型"
// @Success 0 				{string} 	success
// @Failure 1003			更新失败
// @router /inspect/:id [put]
func (c *ExceptionalChildController) PutInspect() {
	// 主键ID
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	// 学生ID
	student_id, _ := c.GetInt("student_id")
	// 学生姓名
	child_name := c.GetString("child_name")
	// 班级ID
	class_id, _ := c.GetInt("class_id")
	// 体质类型
	somatotype := c.GetString("somatotype")
	if err := models.UpInspect(id, student_id, child_name, somatotype, class_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, nil, "更新成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1003, err, "更新失败"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title 			删除特殊儿童
// @Description 	删除特殊儿童
// @Param	id		path 	string	true		"特殊儿童ID"
// @Param	type	query 	int		true		"删除类型1:过敏儿童数据，2:体质类型儿童数据"
// @Success 0 		{string} 	success
// @Failure 1004	删除失败
// @router /:id [delete]
func (c *ExceptionalChildController) Delete() {
	// 主键ID
	idStr := c.Ctx.Input.Param(":id")
	// 删除类型
	del_type, _ := c.GetInt("type")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteExceptionalChild(id, del_type); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1004, err, "删除失败"}
	}
	c.ServeJSON()
}
