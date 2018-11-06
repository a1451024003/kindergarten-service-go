package kanban

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/kanban"
	"time"
)

//健康
type HealthyController struct {
	beego.Controller
}

// URLMapping ...
func (h *HealthyController) URLMapping() {
	h.Mapping("GetDrugList", h.GetDrugList)
	h.Mapping("GetDrugInfo", h.GetDrugInfo)
	h.Mapping("GetAbnormalList", h.GetAbnormalList)
	h.Mapping("GetAbnormalInfo", h.GetAbnormalInfo)
}

// GetDrugList ...
// @Title GetDrugList
// @Description 喂药申请列表
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param class_id        query int    true  "班级ID"
// @Param role            query int    true  "用户身份"
// @Param date            query string false "日期"
// @Param page            query int    false "当前页"
// @Param per_page        query int    false "每页数"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /drug [get]
func (h *HealthyController) GetDrugList() {
	kindergartenId, _ := h.GetInt("kindergarten_id")
	classId, _ := h.GetInt("class_id")
	role, _ := h.GetInt("role")
	date := h.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])
	page, _ := h.GetInt("page", 1)
	perPage, _ := h.GetInt("per_page", 10)

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 不能为空！")
	valid.Required(role, "role").Message("用户身份 不能为空！")
	if valid.HasErrors() {
		h.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetDrugList(kindergartenId, classId, role, date, page, perPage)
		if err != nil {
			h.Data["json"] = JSONStruct{"error", code, err, err.Error()}
		} else {
			h.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	h.ServeJSON()
}

// GetDrugInfo ...
// @Title GetDrugInfo
// @Description 喂药申请详情
// @Param id query int true "ID"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /drug/details id
func (h *HealthyController) GetDrugInfo() {
	drugId, _ := h.GetInt("id")

	valid := validation.Validation{}
	valid.Required(drugId, "id").Message("ID 不能为空！")
	if valid.HasErrors() {
		h.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetDrugInfo(drugId)
		if err != nil {
			h.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			h.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	h.ServeJSON()
}

// GetAbnormalList ...
// @Title GetAbnormalList
// @Description 异常列表
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param class_id        query int    true  "班级ID"
// @Param class_type      query int    false "班级类型"
// @Param date            query string false "日期"
// @Param page            query int    false "当前页"
// @Param per_page        query int    false "每页数"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /archives [get]
func (h *HealthyController) GetAbnormalList() {
	kindergartenId, _ := h.GetInt("kindergarten_id")
	classId, _ := h.GetInt("class_id")
	classType, _ := h.GetInt("class_type")
	date := h.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])
	pType := 2
	page, _ := h.GetInt("page", 1)
	perPage, _ := h.GetInt("per_page", 10)

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 不能为空")
	if valid.HasErrors() {
		h.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetAbnormalList(kindergartenId, classId, classType, date, pType, page, perPage)
		if err != nil {
			h.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			h.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	h.ServeJSON()
}

// GetAbnormalInfo ...
// @Title GetAbnormalInfo
// @Description 异常详情
// @Param id query int true "ID"
// @Failure 1005 获取失败！
// @router /archives/details [get]
func (h *HealthyController) GetAbnormalInfo() {
	abnormalId, _ := h.GetInt("id")

	valid := validation.Validation{}
	valid.Required(abnormalId, "id").Message("ID 不能为空")
	if valid.HasErrors() {
		h.Data["json"] = JSONStruct{"error", 1001, nil, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetAbnormalInfo(abnormalId)
		if err != nil {
			h.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			h.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	h.ServeJSON()
}
