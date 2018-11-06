package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type CourseInfo struct {
	Id           int       `json:"id" orm:"column(id);auto;"`
	Name         string    `json:"name" orm:"column(name);size(30)"; description:"标题"`
	CourseId     int       `json:"course_id" orm:"column(course_id)"`
	TearcherId   int       `json:"tearcher_id" orm:"column(tearcher_id)"`
	TearcherName string    `json:"tearcher_name" orm:"column(tearcher_name);size(30)"`
	Domain       string    `json:"domain" orm:"column(domain);size(30)"`
	Intro        string    `json:"intro" orm:"column(intro);size(30)"`
	Url          string    `json:"url" orm:"column(url)"`
	CoursewareId string    `json:"courseware_id" orm:"column(courseware_id)`
	Aim          string    `json:"aim" orm:"column(aim)`
	Plan         string    `json:"plan" orm:"column(plan)`
	Activity     string    `json:"activity" orm:"column(activity)`
	Job          string    `json:"job" orm:"column(job)`
	Etc          string    `json:"etc" orm:"column(etc)`
	List         string    `json:"list" orm:"column(list)`
	Type         int       `json:"type" orm:"column(type)"`
	Times        string    `json:"times" orm:"column(times)`
	CreatedAt    time.Time `json:"created_at" orm:"auto_now_add"`
}

func (t *CourseInfo) TableName() string {
	return "course_info"
}

func init() {
	orm.RegisterModel(new(CourseInfo))
}

/*
添加
*/
func AddCourseInfo(m *CourseInfo) (map[string]interface{}, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = id //返回数据
		return paginatorMap, err
	}
	return nil, err
}

/*
列表
*/
func GetCourseInfoList(class_type int, kindergarten_id int, date string) (map[string]interface{}, error) {
	var v []orm.Params
	o := orm.NewOrm()
	sql := "select a.id,a.name,a.type,a.course_id,b.id as zid,b.name as zname,b.begin_date,b.end_date from course_info a  left join course c on a.course_id = c.id left join course b on b.id= c.parent_id where b.begin_date <='" + date + "' and b.end_date >= '" + date + "' and b.class_type=" + strconv.Itoa(class_type) + " and b.kindergarten_id =" + strconv.Itoa(kindergarten_id)
	_, err := o.Raw(sql).Values(&v)
	if err == nil && len(v) > 0 {
		var ml []map[string]interface{}
		var ids string
		var ids_arr []string
		for key,val := range v{
			if key == 0 {
				ids = val["zid"].(string)
			} else {
				ids_arr = strings.Split(ids, val["zid"].(string))
				if len(ids_arr) == 1 {
					ids += ","+val["zid"].(string)
				}
			}
		}
		ids_arr = strings.Split(ids, ",")
		if len(ids_arr) > 0 {
			for _,val := range ids_arr {
				list := make(map[string]interface{})
				var list_one []interface{}
				for _,va := range v{
					if val == va["zid"] {
						list["name"] = va["zname"]
						list["begin_date"] = va["begin_date"]
						list["end_date"] = va["end_date"]
						list_one = append(list_one,va)
					}
				}
				list["data"] = list_one
				ml = append(ml,list)
			}
		}

		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = ml
		return paginatorMap, nil
	}
	return nil, err

}

/*
列表
*/
func GetCourseInfoListf(class_type int, kindergarten_id int, date string) (map[string]interface{}, error) {
	var v []CourseInfo
	o := orm.NewOrm()
	sql := "select a.* from course_info a left join course b on a.course_id = b.id where b.begin_date <='" + date + "' and b.end_date >= '" + date + "' and b.class_type=" + strconv.Itoa(class_type) + " and b.kindergarten_id =" + strconv.Itoa(kindergarten_id)

	_, err := o.Raw(sql).QueryRows(&v)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v //分页数据
		return paginatorMap, nil

	}
	return nil, err

}


/*
Web -详情
*/
func GetCourseInfoInfo(id int) map[string]interface{} {
	var v []CourseInfo
	o := orm.NewOrm()
	err := o.QueryTable("course_info").Filter("Id", id).One(&v)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap
	}
	return nil
}

/*
Web -详情
*/
func GetCourseInfoEdit(id int, m CourseInfo) (err error) {
	v := CourseInfo{Id: id}
	o := orm.NewOrm()
	if err := o.Read(&v); err == nil {
		if m.CourseId > 0 {
			v.CourseId = m.CourseId
		}
		if m.TearcherId > 0 {
			v.TearcherId = m.TearcherId
		}
		if m.TearcherName != "" {
			v.TearcherName = m.TearcherName
		}
		if m.Name != "" {
			v.Name = m.Name
		}
		if m.Aim != "" {
			v.Aim = m.Aim
		}
		if m.Type > 0 {
			v.Type = m.Type
		}
		if m.Domain != "" {
			v.Domain = m.Domain
		}
		if m.Intro != "" {
			v.Intro = m.Intro
		}
		if m.Url != "" {
			v.Url = m.Url
		}
		if m.CoursewareId != "" {
			v.CoursewareId = m.CoursewareId
		}
		if m.Plan != "" {
			v.Plan = m.Plan
		}
		if m.Activity != "" {
			v.Activity = m.Activity
		}
		if m.Etc != "" {
			v.Etc = m.Etc
		}
		if m.Times != "" {
			v.Times = m.Times
		}
		if m.List != "" {
			v.List = m.List
		}
		if m.Job != "" {
			v.Job = m.Job
		}
		if _, err = o.Update(&v); err == nil {
			return nil
		}
	}
	return err
}

/*
删除
*/
func DeleteCourseInfo(id int) (err error) {
	o := orm.NewOrm()
	v := CourseInfo{Id: id}
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&CourseInfo{Id: id}); err == nil {
			return nil
		}
	}
	return err
}
