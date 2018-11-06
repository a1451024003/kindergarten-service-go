package admin

import (
	"kindergarten-service-go/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

//幼儿园
type KindergartenController struct {
	BaseController
}

// GetIntroduceInfo ...
// @Title 幼儿园介绍详情
// @Description 幼儿园介绍详情
// @Param	id		path 	string	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /:id [get]
func (c *KindergartenController) GetIntroduceInfo() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	valid := validation.Validation{}
	valid.Required(id, "id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1002, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.GetKindergartenById(id)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1005, v, "获取失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
	}
	c.ServeJSON()
}

// SetPrincipal ...
// @Title 添加园长/教师到幼儿园 未激活状态
// @Description 设置园长
// @Param	user_id		        path 	int	true		"用户ID"
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Param	role		        path 	int	true		"身份（1 园长 5 ）"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router / [post]
func (c *KindergartenController) SetPrincipal() {
	user_id, _ := c.GetInt("user_id")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	role, _ := c.GetInt("role")
	valid := validation.Validation{}
	valid.Required(user_id, "user_id").Message("用户ID不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		err := models.AddPrincipal(user_id, kindergarten_id, role)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, err, "保存失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
		}
		c.ServeJSON()
	}
}

// GetAll ...
// @Title 幼儿园列表
// @Description 幼儿园列表
// @Param	page                  query	int	     false		"页数"
// @Param	per_page              query	int	     false		"每页显示条数"
// @Success 200 {object} models.Kindergarten
// @Failure 403
// @router / [get]
func (c *KindergartenController) GetAll() {
	search := c.GetString("search")
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	v, err := models.GetAll(page, prepage, search, kindergarten_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// StudentClass ...
// @Title 学生姓名搜索班级
// @Description 学生姓名搜索班级
// @Param	name                  query	 string	     true		"姓名"
// @Param	page                  query	   int	     false		"页数"
// @Param	per_page              query	   int	     false		"每页显示条数"
// @Success 200 {object} models.Kindergarten
// @Failure 403
// @router /student_class [get]
func (c *KindergartenController) StudentClass() {
	name := c.GetString("name")
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("学生姓名不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.StudentClass(page, prepage, name)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// SetKindergarten ...
// @Title 添加幼儿园
// @Description 添加幼儿园
// @Param	name		                path 	int	true		"幼儿园名称"
// @Param	license_no   		        path 	int	true		"执照号"
// @Param	kinder_grade		        path 	int	true		"幼儿园级别"
// @Param	kinder_child_no		        path 	int	true		"分校数"
// @Param	address      		        path 	int	true		"地址"
// @Param	tenant_id    		        path 	int	true		"租户，企业编号"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /set_kindergarten [post]
func (c *KindergartenController) SetKindergarten() {
	name := c.GetString("name")
	license_no, _ := c.GetInt("license_no")
	kinder_grade := c.GetString("kinder_grade")
	kinder_child_no, _ := c.GetInt("kinder_child_no")
	address := c.GetString("address")
	tenant_id, _ := c.GetInt("tenant_id")
	phone := c.GetString("phone")
	region := c.GetString("region")
	telephone := c.GetString("telephone")
	user_id, _ := c.GetInt("user_id")
	principal := c.GetString("principal")
	id_number := c.GetString("id_number")
	school_license := c.GetString("school_license")
	tax_registration := c.GetString("tax_registration")
	catering_services := c.GetString("catering_services")
	private_non_enterprise := c.GetString("private_non_enterprise")
	err := models.AddKindergarten(name, license_no, kinder_grade, kinder_child_no, tenant_id, phone, region, address, telephone, principal, id_number, school_license, tax_registration, catering_services, private_non_enterprise, user_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err, "保存失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
	}
	c.ServeJSON()
}

// delete ...
// @Title 删除幼儿园
// @Description 删除幼儿园
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router / [delete]
func (c *KindergartenController) Delete() {
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		err := models.DeleteKinder(kindergarten_id)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "删除失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
		}
		c.ServeJSON()
	}
}

// updata ...
// @Title 编辑幼儿园
// @Description 编辑幼儿园
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /:id [put]
func (c *KindergartenController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	name := c.GetString("name")
	license_no, _ := c.GetInt("license_no")
	kinder_grade := c.GetString("kinder_grade")
	kinder_child_no, _ := c.GetInt("kinder_child_no")
	address := c.GetString("address")
	tenant_id, _ := c.GetInt("tenant_id")
	phone := c.GetString("phone")
	region := c.GetString("region")
	telephone := c.GetString("telephone")
	user_id, _ := c.GetInt("user_id")
	principal := c.GetString("principal")
	id_number := c.GetString("id_number")
	school_license := c.GetString("school_license")
	tax_registration := c.GetString("tax_registration")
	catering_services := c.GetString("catering_services")
	private_non_enterprise := c.GetString("private_non_enterprise")
	err := models.UpdataKinder(id, name, license_no, kinder_grade, kinder_child_no, tenant_id, phone, region, address, telephone, principal, id_number, school_license, tax_registration, catering_services, private_non_enterprise, user_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "编辑失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
	}
	c.ServeJSON()
}

// GetKg ...
// @Title 登陆幼儿园信息
// @Description 登陆幼儿园信息
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Param	user_id		        path 	int	true		"用户ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /getkg [get]
func (c *KindergartenController) GetKg() {
	user_id, _ := c.GetInt("user_id")
	v, err := models.GetKg(user_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// GetKinderMbmber ...
// @Title oms-幼儿园所有成员
// @Description oms-幼儿园所有成员
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /get_member [get]
func (c *KindergartenController) GetKinderMbmber() {
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.GetKinderMbmber(kindergarten_id, page, prepage)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "获取失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// FoodClass ...
// @Title 饮食班级
// @Description 饮食班级
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /food_class [get]
func (c *KindergartenController) FoodClass() {
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.FoodClass(kindergarten_id)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "获取失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// FoodScale ...
// @Title 饮食比例
// @Description 饮食比例
// @Param	class_type		    path 	int	true		"班级类型"
// @Param	is_muslim		    path 	int	true		"是否清真"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /food_scale [get]
func (c *KindergartenController) FoodScale() {
	is_muslim, _ := c.GetInt("is_muslim")
	class_type := c.GetString("class_type")
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		c.ServeJSON()
	} else {
		v, err := models.FoodScale(is_muslim, kindergarten_id, class_type)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "获取失败"}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		c.ServeJSON()
	}
}

// Introduce ...
// @Title 编辑幼儿园介绍
// @Description 编辑幼儿园介绍
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /introduce/:id [put]
func (c *KindergartenController) Introduce() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	introduce := c.GetString("introduce")
	introduce_picture := c.GetString("introduce_picture")
	err := models.UpdataKinderInduce(id, introduce, introduce_picture)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "编辑失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
	}
	c.ServeJSON()
}

// reset ...
// @Title 取消关联幼儿园
// @Description 取消关联幼儿园
// @Param	baby_id		path 	int	true		"宝宝ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /baby [post]
func (c *KindergartenController) Reset() {
	baby_id, _ := c.GetInt("baby_id")
	err := models.Reset(baby_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "取消失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "取消成功"}
	}
	c.ServeJSON()
}

// Count ...
// @Title 教师，学生，班级数量
// @Description 教师，学生，班级数量
// @Param	kindergarten_id           query	int	         true		"幼儿园ID"
// @Success 200 {object} models.Organizational
// @Failure 403
// @router /count [get]
func (o *KindergartenController) Count() {
	kindergarten_id, _ := o.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
	if valid.HasErrors() {
		o.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
		o.ServeJSON()
	} else {
		v, err := models.Count(kindergarten_id)
		if err != nil {
			o.Data["json"] = JSONStruct{"error", 1005, nil, "获取失败"}
		} else {
			o.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
		}
		o.ServeJSON()
	}
}

// AttestKindergarten ...
// @Title 申请认证幼儿园
// @Description 申请认证幼儿园
// @Param	name		                path 	int	true		"幼儿园名称"
// @Param	principal   		        path 	int	true		"执照号"
// @Param	id_number		            path 	int	true		"幼儿园级别"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /attest_kindergarten [post]
func (c *KindergartenController) AttestKindergarten() {
	name := c.GetString("name")
	phone := c.GetString("phone")
	user_id, _ := c.GetInt("user_id")
	principal := c.GetString("principal")
	err := models.AttestKindergarten(name, principal, phone, user_id)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err, "保存失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
	}
	c.ServeJSON()
}

// AttestKindergartenAll ...
// @Title 申请认证幼儿园列表
// @Description 申请认证幼儿园列表
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /attest_all [get]
func (c *KindergartenController) AttestKindergartenAll() {
	prepage, _ := c.GetInt("per_page", 20)
	page, _ := c.GetInt("page")
	search := c.GetString("search")
	v, err := models.AttestKindergartenAll(page, prepage, search)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err, "获取失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, v, "获取成功"}
	}
	c.ServeJSON()
}

// Attest ...
// @Title 认证通过/未通过幼儿园
// @Description 认证通过/未通过幼儿园
// @Param	kindergarten_id		path 	int	true		"幼儿园ID"
// @Success 200 {object} models.Kindergarten
// @Failure 403 :id is empty
// @router /attest/:id [put]
func (c *KindergartenController) Attest() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ty, _ := c.GetInt("type")
	name := c.GetString("name")
	title := c.GetString("title")
	phone := c.GetString("phone")
	region := c.GetString("region")
	user_id, _ := c.GetInt("user_id")
	address := c.GetString("address")
	content := c.GetString("content")
	telephone := c.GetString("telephone")
	principal := c.GetString("principal")
	id_number := c.GetString("id_number")
	school_license := c.GetString("school_license")
	tax_registration := c.GetString("tax_registration")
	catering_services := c.GetString("catering_services")
	private_non_enterprise := c.GetString("private_non_enterprise")
	err := models.Attest(id, name, phone, region, address, telephone, principal, id_number, school_license, tax_registration, catering_services, private_non_enterprise, ty, user_id, title, content)
	if err != nil {
		c.Data["json"] = JSONStruct{"error", 1003, err.Error(), "编辑失败"}
	} else {
		c.Data["json"] = JSONStruct{"success", 0, nil, "编辑成功"}
	}
	c.ServeJSON()
}
