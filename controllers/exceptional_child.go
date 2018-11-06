package controllers

import (
	"kindergarten-service-go/models"
	"strconv"

	"github.com/astaxie/beego/validation"
)

type ExceptionalChildController struct {
	BaseController
}

func (c *ExceptionalChildController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @Title 特殊儿童列表/搜索特殊儿童
// @Description 特殊儿童列表/搜索特殊儿童
// @Param	keyword	   	 query	string	 false	"关键字(特殊儿童姓名/特殊儿童过敏源)"
// @Param	page	  	 query	int	 	 false	"当前页"
// @Success 0 			 {object}  models.ExceptionalChild
// @Failure 1005 获取失败
// @router / [get]
func (c *ExceptionalChildController) GetSearch() {
	// 关键字
	keyword := c.GetString("keyword")
	// page_num
	page, _ := c.GetInt64("page")

	// limit
	limit, _ := c.GetInt64("per_page")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")

	if info, err := models.GetExceptionalInspect(page, limit, keyword, kindergarten_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, info, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title 根据过敏源获取特殊儿童
// @Description 根据过敏源获取特殊儿童
// @Param	allergen	query	string	true	"过敏源信息，多个过敏源以','分隔"
// @Success 0 			{string} 	success
// @Failure 1005 获取失败
func (c *ExceptionalChildController) GetAllergenChild() {
	allergen := c.GetString("allergen")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	valid := validation.Validation{}
	valid.Required(allergen, "allergen").Message("过敏源信息不能为空")
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")

	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		if allergenChild, err := models.GetAllergenChild(allergen, kindergarten_id); err == nil {
			c.Data["json"] = JSONStruct{"success", 0, allergenChild, "获取成功"}
		} else {
			c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
		}
	}
	c.ServeJSON()
}

// @Title 过敏食物报备
// @Description 过敏食物报备
// @Param 		class				formData  	int    	true		"班级ID"
// @Param 		kindergarten_id		formData  	int    	true		"幼儿园ID"
// @Param 		creator				formData  	int    	true		"创建人ID"
// @Param 		student_id			formData  	int    	true		"宝宝ID"
// @Param 		source				formData  	int    	true		"来源信息"
// @Param 		somatotype			formData  	int    	true		"体质类型"
// @Param		child_name			formData	string	true		"特殊儿童姓名"
// @Param 		allergen			formData  	string 	true		"过敏源，多个过敏源以','分隔"
// @Success 0					{json} JSONSturct
// @Failure 1003 				新增失败
// @router / [post]
func (c *ExceptionalChildController) AllergenPreparation() {
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	// 创建人ID
	creator, _ := c.GetInt("creator")
	// 宝宝ID
	baby_id, _ := c.GetInt("baby_id")
	// 来源信息
	source, _ := c.GetInt8("source")
	// 体质类型
	somatotype, _ := c.GetInt8("somatotype")
	// 特殊儿童姓名
	child_name := c.GetString("child_name")
	// 过敏源
	allergen := c.GetString("allergen")

	valid := validation.Validation{}
	valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园ID不能为空")
	valid.Required(creator, "creator").Message("创建人ID不能为空")
	valid.Required(baby_id, "baby_id").Message("宝宝ID不能为空")
	valid.Required(source, "source").Message("来源信息不能为空")
	valid.Required(somatotype, "somatotype").Message("体质类型不能为空")
	valid.Required(allergen, "allergen").Message("过敏源不能为空")

	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		if num, err := models.AllergenPreparation(child_name, somatotype, allergen, source, kindergarten_id, creator, baby_id); err == nil {
			if num == 0 {
				c.Data["json"] = JSONStruct{"error", 1007, nil, "已有此数据"}
			} else {
				c.Data["json"] = JSONStruct{"success", 0, nil, "保存成功"}
			}

		} else {
			c.Data["json"] = JSONStruct{"error", 1003, err, "保存失败"}
		}
	}
	c.ServeJSON()
}

// GetAllergen ...
// @Title 			根据宝宝ID过敏源
// @Description 		根据宝宝ID过敏源
// @Param	id		query 	int		true		"宝宝ID"
// @Success 0 		{string} 	success
// @Failure 1004		获取失败
func (c *ExceptionalChildController) GetAllergen() {
	// 宝宝ID
	id, _ := c.GetInt("baby_id")
	// 幼儿园ID
	kindergarten_id, _ := c.GetInt("kindergarten_id")
	if info, err := models.GetAllergen(id, kindergarten_id); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, info, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, err, "获取失败"}
	}
	c.ServeJSON()
}

// DelAllergen ...
// @Title 			删除过敏源
// @Description 		删除过敏源
// @Param	id		path 	string	true		"特殊儿童ID"
// @Success 0 		{string} 	success
// @Failure 1004		删除失败
// @router /:id [delete]
func (c *ExceptionalChildController) DelAllergen() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteExceptionalChild(id, 1); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, nil, "删除成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1004, err, "删除失败"}
	}
	c.ServeJSON()
}
