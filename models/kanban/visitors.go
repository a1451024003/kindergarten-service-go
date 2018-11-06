package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris/core/errors"
	"math"
	"strconv"
)

type Page struct {
	Data    interface{} `json:"data" description:"数据"`
	Total   int         `json:"total" description:"总条数"`
	PerPage int         `json:"per_page" description:"每页数"`
	PageNum int         `json:"page_num" description:"总页数"`
}

//访客数量
func GetVisitorsNum(date string) (ml map[string]interface{}, code int, err error) {
	sql := "SELECT COUNT(Distinct visitor_id) FROM visitors WHERE LEFT(visitor_time, 10) >= \"" + date + "\""
	o := orm.NewOrm()
	var total orm.ParamsList
	_, err = o.Raw(sql).ValuesFlat(&total)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	}
	ml = make(map[string]interface{})
	if date != "" {
		ml["date"] = date
	}
	ml["visitors_num"] = total[0]
	return ml, code, err
}

//访客列表
func GetVisitorsList(date string, page int, perPage int) (ml interface{}, code int, err error) {
	sql := "SELECT * FROM visitors WHERE LEFT(visitor_time, 10) >= \"" + date + "\" LIMIT ? OFFSET ?"
	totalSql := "SELECT COUNT(Distinct visitor_id) FROM visitors WHERE LEFT(visitor_time, 10) >= \"" + date + "\""
	offset := (page - 1) * perPage //偏移量
	o := orm.NewOrm()
	var visitors []orm.Params
	_, err = o.Raw(sql, perPage, offset).Values(&visitors)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	}
	var total orm.ParamsList
	_, err = o.Raw(totalSql).ValuesFlat(&total)
	totalInt, _ := strconv.Atoi(total[0].(string))             //总条数
	pageNum := math.Ceil(float64(totalInt) / float64(perPage)) //总页数
	ml = Page{Data: visitors, Total: totalInt, PerPage: perPage, PageNum: int(pageNum)}
	return ml, code, err
}
