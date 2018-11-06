package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//班级某一天课程
func GetClassOneDayCourse(classId int, date string) (ml map[string]interface{}, code int, err error) {
	o := orm.NewOrm()
	sql := "select * from course_class where class_id=" + strconv.Itoa(classId) + " and begin_date <= '" + date + "' and end_date >= '" + date + "'"
	var courseClass []orm.Params
	_, err = o.Raw(sql).Values(&courseClass)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	}
	if courseClass == nil {
		code = 1002
		err = errors.New("暂无数据！")
		return ml, code, err
	}
	var courseClassMap interface{}
	for _, val := range courseClass {
		json.Unmarshal([]byte(val["content"].(string)), &courseClassMap)
	}
	var data []interface{}
	for _, val := range courseClassMap.([]interface{}) {
		course := val.(map[string]interface{})["course"]
		for _, v := range course.([]interface{}) {
			if v.(map[string]interface{})["date"] == date {
				v.(map[string]interface{})["time"] = val.(map[string]interface{})["time"]
				data = append(data, v)
			}
		}
	}
	ml = make(map[string]interface{})
	ml["data"] = data
	ml["total"] = len(data)
	return ml, code, err
}
