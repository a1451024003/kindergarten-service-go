package kanban

import (
	"github.com/astaxie/beego"
	"kindergarten-service-go/models/kanban"
	"time"
)

//访客
type VisitorsController struct {
	beego.Controller
}

func (v *VisitorsController) URLMapping() {
	v.Mapping("GetVisitorsNum", v.GetVisitorsNum)
	v.Mapping("GetVisitorsList", v.GetVisitorsList)
}

// GetVisitorsNum ...
// @Title GetVisitorsNum
// @Description 访客数量
// @Param date query string true "日期"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /visitors_num [get]
func (v *VisitorsController) GetVisitorsNum() {
	date := v.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	data, code, err := models.GetVisitorsNum(date)
	if err != nil {
		v.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
	} else {
		v.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
	}
	v.ServeJSON()
}

// GetVisitorsList ...
// @Title GetVisitorsList
// @Description 访客列表
// @Param date     query string true  "日期"
// @Param page     query int    false "当前页"
// @Param per_page query int    false "每页数"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router /visitors_list [get]
func (v *VisitorsController) GetVisitorsList() {
	date := v.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])
	page, _ := v.GetInt("page", 1)
	perPage, _ := v.GetInt("per_page", 10)

	data, code, err := models.GetVisitorsList(date, page, perPage)
	if err != nil {
		v.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
	} else {
		v.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
	}
	v.ServeJSON()
}
