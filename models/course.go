package models

import (
	"time"

	"encoding/json"
	"math"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id             int       `json:"id" orm:"column(id);auto;"`
	Name           string    `json:"name" orm:"column(name);size(30)"; description:"标题"`
	KindergartenId int       `json:"kindergarten_id" orm:"column(kindergarten_id)";description:"幼儿园ID"`
	Status         int       `json:"status" orm:"column(status)"`
	BeginDate      string    `json:"begin_date" orm:"column(begin_date);size(30)"`
	EndDate        string    `json:"end_date" orm:"column(end_date);size(30)"`
	Aim            string    `json:"aim" orm:"column(aim);size(30)"`
	Url            string    `json:"url" orm:"column(url)"`
	Leval          int       `json:"leval" orm:"column(leval)`
	ParentId       int       `json:"parent_id" orm:"column(parent_id)"`
	ClassType      int       `json:"class_type" orm:"column(class_type)"`
	CreatedAt      time.Time `json:"created_at" orm:"auto_now_add"`
}

func (t *Course) TableName() string {
	return "course"
}

func init() {
	orm.RegisterModel(new(Course))
}

/*
添加
*/
func AddCourse(m *Course) (map[string]interface{}, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	if err == nil {
		if m.Status == 1 {
			err = UseCourse(int(id), m.KindergartenId)
		}
		return nil, err
	}
	return nil, err
}

//修改
func UpdataCourse(id int, begin_date string, end_date string, name string, url string, status int, date int, aim string) error {
	o := orm.NewOrm()
	m := Course{Id: id}
	err := o.Read(&m)
	if err == nil {
		if date == 1 {
			if begin_date == "0000-00-00" {
				m.BeginDate = ""
				m.EndDate = ""
			} else {
				if begin_date != "" {
					m.BeginDate = begin_date
				}
				if end_date != "" {
					m.EndDate = end_date
				}
			}
		}
		if aim != "" {
			m.Aim = aim
		}
		if name != "" {
			m.Name = name
		}
		if url != "" {
			m.Url = url
		}
		if status == -1 {
			m.Status = 0
		}
		if status == 1 {
			m.Status = 1
			err = UseCourse(m.Id, m.KindergartenId)
		}
		o.Update(&m)
		return nil
	}
	return err
}

//批量设置
func UpdataAllCourse(kindergarten_id int, class_type int) error {
	o := orm.NewOrm()
	p, _ := o.Raw("UPDATE course SET end_date = '',begin_date='' WHERE class_type=" + strconv.Itoa(class_type) + " and kindergarten_id = " + strconv.Itoa(kindergarten_id)).Prepare()
	p.Exec()
	p.Close()

	return nil
}

//采用
func UseCourse(id int, kindergarten_id int) error {
	o := orm.NewOrm()
	m := Course{Id: id}
	err := o.Read(&m)
	if err == nil {
		p, _ := o.Raw("UPDATE course SET status = 0 WHERE parent_id=0 and kindergarten_id = " + strconv.Itoa(kindergarten_id)).Prepare()
		p.Exec()
		p.Close()
		m.Status = 1
		o.Update(&m)

		return nil
	}
	return err
}

//列表
func GetCourseList(parent_id int, kindergarten_id int, status int, page, per_page int, class_type int) (map[string]interface{}, error) {
	var v []Course
	o := orm.NewOrm()
	qs := o.QueryTable("course")
	qsl := o.QueryTable("course")
	if class_type > 0 {
		qs = qs.Filter("class_type", class_type)
		qsl = qsl.Filter("class_type", class_type)
	}
	nums, err := qs.Filter("kindergarten_id", kindergarten_id).Filter("parent_id", parent_id).All(&v)
	if err == nil && nums > 0 {
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(nums) / float64(per_page))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * per_page
		num, err := qsl.Limit(per_page, limit).Filter("kindergarten_id", kindergarten_id).Filter("parent_id", parent_id).All(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = nums          //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err

}

//删除
func DeleteCourse(id int) (err error) {
	o := orm.NewOrm()
	v := Course{Id: id}
	if err := o.Read(&v); err == nil {
		if _, err = o.Delete(&Course{Id: id}); err == nil {

			return nil
		}
	}
	return err
}

// 专题详情
func InfoCourse(id int, sass int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var list []Course
	if _, err = o.Raw("select * from course where parent_id=" + strconv.Itoa(id)).QueryRows(&list); err == nil && len(list) > 0 {
		var ids string
		for key, val := range list {
			if key == 0 {
				ids = strconv.Itoa(val.Id)
			} else {
				ids += "," + strconv.Itoa(val.Id)
			}
		}
		var data []map[string]interface{}
		list_json, _ := json.Marshal(list)
		json.Unmarshal(list_json, &data)
		var c_info []CourseInfo
		o.Raw("select * from course_info where course_id in(" + ids + ")").QueryRows(&c_info)
		if len(c_info) > 0 {
			for key, val := range data {
				var c_one []CourseInfo
				for _, v := range c_info {
					if v.CourseId == int(val["id"].(float64)) {
						c_one = append(c_one, v)
					}
				}
				data[key]["list"] = c_one
			}
		}
		ml = make(map[string]interface{})
		ml["data"] = data
		if sass == 1 {
			var cv Course
			cv.Id = id
			o.Read(&cv)
			ml["info"] = cv
		}

		return ml, err
	}

	return nil, err
}
