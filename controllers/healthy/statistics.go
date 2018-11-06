package healthy

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"kindergarten-service-go/models/healthy"
	"time"
)

//健康统计
type StatisticsController struct {
	beego.Controller
}

// URLMapping ...
func (s *StatisticsController) URLMapping() {
	s.Mapping("Statistics", s.Statistics)
}

// Statistics ...
// @Title Statistics
// @Description 健康统计
// @Param kindergarten_id query int    true  "幼儿园ID"
// @Param date            query string false "日期"
// @Success 0    获取成功！
// @Failure 1005 获取失败！
// @router / [get]
func (s *StatisticsController) Statistics() {
	kindergartenId, _ := s.GetInt("kindergarten_id")
	date := s.GetString("date", time.Now().Format("2006-01-02 15:04:05"))
	date = string([]byte(date)[:10])

	valid := validation.Validation{}
	valid.Required(kindergartenId, "kindergarten_id").Message("幼儿园ID 必须填写！")
	if valid.HasErrors() {
		s.Data["json"] = JSONStruct{"error", 1001, []string{}, valid.Errors[0].Message}
	} else {
		data, code, err := healthy.Statistics(kindergartenId, date)
		if err != nil {
			s.Data["json"] = JSONStruct{"error", code, nil, err.Error()}
		} else {
			s.Data["json"] = JSONStruct{"success", 0, data, "获取成功！"}
		}
	}
	s.ServeJSON()
}
