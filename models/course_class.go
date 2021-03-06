package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CourseClass struct {
	Id        int    `json:"id" orm:"column(id);auto;"`
	Content   string `json:"content" orm:"column(content);`
	ClassId   int    `json:"class_id" orm:"column(class_id)"`
	BeginDate string `json:"begin_date" orm:"column(begin_date)"`
	EndDate   string `json:"end_date" orm:"column(end_date)"`
}

func (t *CourseClass) TableName() string {
	return "course_class"
}

func init() {
	orm.RegisterModel(new(CourseClass))
}

/*
添加
*/
func AddCourseClass(m *CourseClass, content string) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	v := m
	if crd, _, err := o.ReadOrCreate(m, "ClassId", "BeginDate"); err == nil {
		if !crd {
			v.Id = m.Id
			v.Content = content
			o.Update(v)
			return nil, err
		}
	}
	return nil, err
}

/*
列表
*/
func GetCourseClassList(class_id int, begin_date string) (map[string]interface{}, error) {
	var v []CourseClass
	o := orm.NewOrm()
	sql := "select * from course_class where class_type=" + strconv.Itoa(class_id)

	_, err := o.Raw(sql).QueryRows(&v)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v //分页数据
		return paginatorMap, nil

	}
	return nil, err

}

//课程表，一天的课程
func GetCourseClassInfo(class_id int, date string, types int, kindergarten_id int) map[string]interface{} {
	var v CourseClass
	o := orm.NewOrm()
	sql := "select * from course_class where class_id=" + strconv.Itoa(class_id) + " and begin_date <= '" + date + "' and end_date >= '" + date + "'"
	err := o.Raw(sql).QueryRow(&v)
	fmt.Println(err)
	if err == nil {
		type Course_myinfolist struct {
			Date         string `json:"date"`
			Name         string `json:"name"`
			CourseInfoId int    `json:"course_info_id"`
			CourseId     int    `json:"course_id"`
			TypeName     string `json:"typeName"`
			Type         int    `json:"type"`
		}
		type Course_myinfo struct {
			KindergartenTimeId int                 `json:"kindergarten_time_id"`
			Time               string              `json:"time"`
			Course             []Course_myinfolist `json:"course"`
		}
		var content []Course_myinfo
		json.Unmarshal([]byte(v.Content), &content)
		if types == 1 {
			var kt []KindergartenTime

			sql := "select * from kindergarten_time where (type=0 and class_id = 0 and kindergarten_id=" + strconv.Itoa(kindergarten_id) + ") or (type=1 and class_id = " + strconv.Itoa(class_id) + " and kindergarten_id=" + strconv.Itoa(kindergarten_id) + ") order by begin_time asc"
			if _, err := o.Raw(sql).QueryRows(&kt); err == nil {
				for key, val := range kt {
					for _, va := range content {
						for _, v := range va.Course {
							if val.Id == va.KindergartenTimeId && date == v.Date {
								kt[key].Name = v.Name
							}
						}
					}
				}
			}
			paginatorMap := make(map[string]interface{})
			paginatorMap["data"] = kt
			return paginatorMap
		}
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = content
		return paginatorMap
	}
	return nil
}

//课程表，一天的课程
func GetCourseDayClassInfo(kindergarten_id int) map[string]interface{} {
	o := orm.NewOrm()
	var kt []KindergartenTime

	sql := "select * from kindergarten_time where  kindergarten_id=" + strconv.Itoa(kindergarten_id) + " and class_type = 0 order by begin_time asc"
	if _, err := o.Raw(sql).QueryRows(&kt); err == nil {

	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["data"] = kt
	return paginatorMap

	return nil
}

/*
删除
*/
func DeleteCourseClass(id int) map[string]interface{} {
	o := orm.NewOrm()
	v := CourseClass{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CourseClass{Id: id}); err == nil {
			paginatorMap := make(map[string]interface{})
			paginatorMap["data"] = num //返回数据
			return paginatorMap
		}
	}
	return nil
}

//教学计划
func PlanCourseClass(class_id int, date_time string) map[string]interface{} {
	o := orm.NewOrm()
	var v []CourseClass
	sql := "select * from course_class where class_id=" + strconv.Itoa(class_id) + " and left(begin_date,7) = '" + date_time + "' order by begin_date asc"
	o.Raw(sql).QueryRows(&v)
	ml := make(map[string]interface{})
	var list []map[string]interface{}
	vjson, _ := json.Marshal(v)
	json.Unmarshal(vjson, &list)
	for key, v := range list {
		type Course_myinfolist struct {
			Date         string `json:"date"`
			Name         string `json:"name"`
			CourseInfoId int    `json:"course_info_id"`
			CourseId     int    `json:"course_id"`
			TypeName     string `json:"typeName"`
			Type         int    `json:"type"`
		}
		type Course_myinfo struct {
			KindergartenTimeId int                 `json:"kindergarten_time_id"`
			Time               string              `json:"time"`
			Course             []Course_myinfolist `json:"course"`
		}
		var content []Course_myinfo
		json.Unmarshal([]byte(v["content"].(string)), &content)
		info_id := "0"
		course_id := "0"
		var course []int
		for _, va := range content {
			for _, v := range va.Course {
				if v.CourseId == 0 {
					course = append(course,v.CourseId)
				}
				info_id += "," + strconv.Itoa(v.CourseInfoId)
				course_id += "," + strconv.Itoa(v.CourseId)
			}
		}
		sql_info := "select a.name,a.aim,a.id from course a left join course b on a.id = b.parent_id where b.id in (" + course_id + ") group by a.id"
		var zhuanti []orm.Params
		o.Raw(sql_info).Values(&zhuanti)
		maps := make(map[string]interface{})
		if len(course) > 0 {
			maps["name"] = "无专题"
			maps["aim"] = "无目标"
			maps["id"] = "0"
			zhuanti = append(zhuanti, maps)
		}
		list[key]["data"] = zhuanti

		list[key]["date"] = beego.Substr(list[key]["begin_date"].(string), 8, 2) + "~" + beego.Substr(list[key]["end_date"].(string), 8, 2)
		delete(list[key], "content")
	}
	ml["data"] = list
	return ml
}

//计划详情
func PlanInfoCourseClass(id int) map[string]interface{} {
	var v CourseClass
	o := orm.NewOrm()
	sql := "select * from course_class where id=" + strconv.Itoa(id)
	err := o.Raw(sql).QueryRow(&v)
	if err == nil {
		type Course_myinfolist struct {
			Date         string `json:"date"`
			Name         string `json:"name"`
			CourseInfoId int    `json:"course_info_id"`
			CourseId     int    `json:"course_id"`
			TypeName     string `json:"typeName"`
			Type         int    `json:"type"`
		}
		type Course_myinfo struct {
			KindergartenTimeId int                 `json:"kindergarten_time_id"`
			Time               string              `json:"time"`
			Course             []Course_myinfolist `json:"course"`
		}
		var content []Course_myinfo
		json.Unmarshal([]byte(v.Content), &content)
		info_id := "0"
		course_id := "0"
		for _, va := range content {
			for _, v := range va.Course {
				info_id += "," + strconv.Itoa(v.CourseInfoId)
				course_id += "," + strconv.Itoa(v.CourseId)
			}
		}
		sql_info := "select b.name,b.aim,b.id,b.parent_id as fid from course b  where b.id in (" + course_id + ")"
		var mubiao []orm.Params
		_, err := o.Raw(sql_info).Values(&mubiao)
		fmt.Println(err)
		if err == nil {
			sql_info := "select a.name,a.aim,a.id from course a left join course b on a.id = b.parent_id where b.id in (" + course_id + ") group by a.id"
			var zhuanti []orm.Params
			o.Raw(sql_info).Values(&zhuanti)
			sql_course := "select b.name as bname,b.aim as baim,a.url,a.course_id,a.name,a.id,a.created_at,a.type from course_info a left join course b on a.course_id= b.id where a.id in (" + info_id + ")"
			var mapsc []orm.Params
			o.Raw(sql_course).Values(&mapsc)
			list_arr := strings.Split(course_id, "0")
			if len(list_arr) > 2 {
				mubiao_one := make(map[string]interface{})
				mubiao_one["name"] = "无目标"
				mubiao_one["aim"] = "无目标"
				mubiao_one["id"] = "0"
				mubiao_one["fid"] = "0"
				mubiao = append(mubiao, mubiao_one)
				zhuanti_one := make(map[string]interface{})
				zhuanti_one["name"] = "无专题"
				zhuanti_one["aim"] = "无目标"
				zhuanti_one["id"] = "0"
				zhuanti = append(zhuanti, zhuanti_one)
			}
			for key, val := range zhuanti {
				var list2 []map[string]interface{}
				for ke, va := range mubiao {
					if val["id"] == va["fid"] {
						var list3 []map[string]interface{}
						for _, v := range mapsc {
							if v["course_id"] == va["id"] {
								list3 = append(list3, v)
							}
						}
						mubiao[ke]["data"] = list3
						list2 = append(list2, mubiao[ke])
					}
				}
				zhuanti[key]["data"] = list2
			}
			paginatorMap := make(map[string]interface{})
			paginatorMap["data"] = zhuanti
			return paginatorMap
		}

	}
	return nil
}
func PlanInfonewCourseClass(id int, c_id int) map[string]interface{} {
	list := PlanInfoCourseClass(id)
	ljson, _ := json.Marshal(list["data"])
	fmt.Println(string(ljson))
	var l []map[string]interface{}
	json.Unmarshal(ljson, &l)
	ml := make(map[string]interface{})
	for _, val := range l {
		if val["id"].(string) == strconv.Itoa(c_id) {
			ml["data"] = val
		}
	}
	return ml
}

func GetClassCourseware(class_id int, time string) (ml []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v CourseClass
	sql := "select * from course_class where class_id=" + strconv.Itoa(class_id) + " and begin_date <= '" + time + "' and end_date >= '" + time + "'"
	err = o.Raw(sql).QueryRow(&v)
	type Course_myinfolist struct {
		Date         string `json:"date"`
		Name         string `json:"name"`
		CourseInfoId int    `json:"course_info_id"`
		CourseId     int    `json:"course_id"`
		TypeName     string `json:"typeName"`
		Type         int    `json:"type"`
	}
	type Course_myinfo struct {
		KindergartenTimeId int                 `json:"kindergarten_time_id"`
		Time               string              `json:"time"`
		Course             []Course_myinfolist `json:"course"`
	}
	var content []Course_myinfo
	ids := "0"
	if err == nil {
		json.Unmarshal([]byte(v.Content), &content)
		for _, val := range content {
			one_info := make(map[string]interface{})
			one_info["time"] = val.Time
			one_info["date"] = time
			ml = append(ml, one_info)
			for _, va := range val.Course {
				if time == va.Date {
					one_info["course_info_id"] = va.CourseInfoId
					ids += "," + strconv.Itoa(va.CourseInfoId)
				}
			}
		}
		fmt.Println(ml)
		var list_c []CourseInfo
		sql_c := "select * from course_info where id in (" + ids + ")"
		_, err = o.Raw(sql_c).QueryRows(&list_c)
		fmt.Println(list_c)
		if err == nil {
			for _, val := range ml {
				for _, va := range list_c {
					if val["course_info_id"] == va.Id {
						val["course"] = va
					}
				}
			}
			return ml, err
		}
	}
	return nil, err
}
