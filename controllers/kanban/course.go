package kanban

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/kanban"
	"time"
)

//课程
type CourseController struct {
	beego.Controller
}

func (c *CourseController) URLMapping() {
	c.Mapping("GetClassOneDayCourse", c.GetClassOneDayCourse)
}

// GetClassOneDayCourse ...
// @Title GetClassOneDayCourse
// @Description 班级某天课程
// @Param class_id query int    true  "班级ID"
// @Param date     query string false "日期"
// @Success 0    获取成功！
// @Failure 1002 暂无数据！
// @Failure 1005 获取失败！
// @router / [get]
func (c *CourseController) GetClassOneDayCourse() {
	classId, _ := c.GetInt("class_id")
	date := c.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	valid := validation.Validation{}
	valid.Required(classId, "class_id").Message("班级ID 必须填写！")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := models.GetClassOneDayCourse(classId, date)
		if err != nil {
			c.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			c.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	c.ServeJSON()
}
