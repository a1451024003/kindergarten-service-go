package models

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExceptionalChild struct {
	Id             int    `json:"id";orm:"column(id);auto"`
	ChildName      string `json:"child_name";orm:"column(child_name);size(20)"`
	Class          int    `json:"class";orm:"column(class)"`
	Somatotype     int8   `json:"somatotype";orm:"column(somatotype)"`
	Allergen       string `json:"allergen";orm:"column(allergen);size(20)"`
	Source         int8   `json:"source";orm:"column(source)"`
	KindergartenId int    `json:"kindergarten_id";orm:"column(kindergarten_id)"`
	Creator        int    `json:"creator";orm:"column(creator)"`
	StudentId      int    `json:"student_id";orm:"column(student_id)"`
	CreatedAt      string `json:"created_at";orm:"column(created_at);type(datetime)"`
	UpdatedAt      string `json:"updated_at";orm:"column(updated_at);type(datetime)"`
}

func (t *ExceptionalChild) TableName() string {
	return "exceptional_child"
}

func init() {
	orm.RegisterModel(new(ExceptionalChild))
}

// AddExceptionalChild insert a new ExceptionalChild into database and returns
// last inserted Id on success.
func AddExceptionalChild(child_name string, class int, somatotype int8, allergen string, source int8, kindergarten_id int, creator int, student_id int) (id int64, err error) {
	var exceptionalChild ExceptionalChild
	o := orm.NewOrm()
	var infos []orm.Params
	where := " allergen = \"" + string(allergen) + "\" AND somatotype = " + strconv.Itoa(int(somatotype)) + " AND student_id = " + strconv.Itoa(student_id) + " AND kindergarten_id = " + strconv.Itoa(kindergarten_id)
	if n, er := o.Raw("SELECT allergen FROM `exceptional_child` WHERE " + where).Values(&infos); er == nil && n > 0 {
		// 已存在相同数据
		return 0, err
	} else {
		exceptionalChild.ChildName = child_name
		exceptionalChild.Class = class
		exceptionalChild.Somatotype = somatotype
		exceptionalChild.Allergen = allergen
		exceptionalChild.Source = source
		exceptionalChild.KindergartenId = kindergarten_id
		exceptionalChild.Creator = creator
		exceptionalChild.StudentId = student_id
		exceptionalChild.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		exceptionalChild.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		if id, err := o.Insert(&exceptionalChild); err != nil && id <= 0 {
			return id, errors.New("保存失败")
		} else {
			return id, nil
		}
	}
	return
}

// GetExceptionalChildById retrieves ExceptionalChild by Id. Returns error if
// Id doesn't exist
func GetExceptionalChildById(id string, kindergarten_id int) (exceptionalChild interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	var where string
	where = " ex.id=" + id
	if kindergarten_id != 0 {
		where += " AND ex.kindergarten_id = " + strconv.Itoa(kindergarten_id)
	}
	sql := qb.Select("ex.id, ex.child_name, ex.somatotype, ex.allergen, ex.source, ex.kindergarten_id, ex.creator, ex.student_id, ex.created_at, ex.updated_at, stu.class_info as class_name").From("exceptional_child as ex").LeftJoin("student as stu").On("stu.student_id = ex.student_id").Where(where).String()
	var maps []orm.Params
	if num, err := o.Raw(sql).Values(&maps); err == nil && num > 0 {

		return maps, err
	} else {
		return nil, err
	}

}

// GetAllExceptionalChild retrieves all ExceptionalChild matches certain condition. Returns empty list if
// no records exist
func GetAllExceptionalChild(child_name string, get_type int, page int64, limit int64, keyword string, kindergarten_id int) (Page, error) {
	o := orm.NewOrm()
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	var maps []orm.Params
	var err error
	var num int64
	// 查询过敏信息表
	if get_type == 1 {
		where := "1=1 "
		if kindergarten_id != 0 {
			where += " AND ex.kindergarten_id = " + strconv.Itoa(kindergarten_id)
		}

		if child_name != "" {
			where += " AND ex.child_name like \"%" + string(child_name) + "%\""
		}

		// 特殊儿童姓名或者特殊儿童过敏源
		if keyword != "" {
			where += " AND ex.child_name like \"%" + string(keyword) + "%\" OR ex.allergen like \"%" + string(keyword) + "%\""

		}
		where += " AND stu.student_id IS NOT NULL "
		totalqb, _ := orm.NewQueryBuilder("mysql")
		tatolsql := totalqb.Select("count(*)").
			From("exceptional_child as ex").
			LeftJoin("student as stu").
			On("stu.student_id = ex.student_id").
			Where(where).
			String()

		err := o.Raw(tatolsql).QueryRow(&total)

		if err == nil {
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("ex.id, ex.class as class_id, ex.child_name, ex.somatotype, ex.allergen, ex.source, ex.kindergarten_id, ex.creator, ex.student_id, ex.created_at, ex.updated_at, stu.student_id, stu.class_info").
				From("exceptional_child as ex").
				LeftJoin("student as stu").
				On("stu.student_id = ex.student_id").
				Where(where).
				OrderBy("ex.id").
				Desc().
				Limit(int(limit)).
				Offset(int(offset)).
				String()

			num, err = o.Raw(sql).Values(&maps)

		}

	} else {
		where := "1=1 "
		if kindergarten_id != 0 {
			where += " AND he.kindergarten_id = " + strconv.Itoa(kindergarten_id)
		}

		if child_name != "" {
			where += " AND stu.name like \"%" + string(child_name) + "%\""
		}

		var nums []orm.Params

		if n, er := o.Raw("SELECT id FROM healthy_body AS he WHERE " + where + " ORDER BY he.id DESC LIMIT 1").Values(&nums); er == nil && n > 0 {
			where += " AND he.body_id = " + nums[0]["id"].(string)
		}

		where += " AND stu.student_id IS NOT NULL "
		totalqb, _ := orm.NewQueryBuilder("mysql")
		tatolsql := totalqb.Select("count(*)").
			From("healthy_inspect as he").
			LeftJoin("student as stu").
			On("stu.student_id = he.student_id").
			Where(where).
			GroupBy("he.student_id").
			String()
		var totals []orm.Params
		if n, er := o.Raw(tatolsql).Values(&totals); er == nil && n > 0 {
			// 查询体检表
			where += " AND ( he.abnormal_weight != '正常' Or he.abnormal_height = '正常' ) AND ( he.abnormal_weight != '' Or he.abnormal_height != '' )  "

			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("he.id, he.class_id, he.abnormal_weight as somatot_info, he.student_id, max(he.created_at) as updated_at, he.evaluate, stu.name as child_name, stu.student_id, stu.class_info").
				From("healthy_inspect as he").
				LeftJoin("student as stu").
				On("stu.student_id = he.student_id").
				Where(where).
				GroupBy("he.student_id").
				Limit(int(limit)).
				Offset(int(offset)).
				String()

			total = int64(len(totals))
			num, err = o.Raw(sql).Values(&maps)
		}

	}

	if err == nil && num > 0 {
		var newMap []orm.Params
		for k, v := range maps {
			// 序号
			v["index"] = (int(page)*10 - 10) + (k + 1)
			t := time.Now()
			currentTime := t.Unix() - (24 * 3600 * 7) //当前时间戳
			BeCharge := v["updated_at"]
			toBeCharge := BeCharge.(string)
			timeLayout := "2006-01-02 15:04:05"
			loc, _ := time.LoadLocation("Local")
			theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
			sr := theTime.Unix() //将更新时间转为时间戳
			if sr > currentTime {
				v["new"] = 1 //最新
			} else {
				v["new"] = 0
			}

			delete(v, "updated_at")
			newMap = append(newMap, v)
		}
		pageNum := int(math.Ceil(float64(total) / float64(limit)))

		return Page{newMap, total, pageNum}, nil

	}
	return Page{}, nil
}

// UpdateExceptionalChild updates ExceptionalChild by Id and returns error if
// the record to be updated doesn't exist
func UpdateExceptionalChildById(id int, child_name string, class int, somatotype int8, allergen string, student_id int, kindergarten_id int) (num int, err error) {
	o := orm.NewOrm()
	exceptionalChild := ExceptionalChild{Id: id}
	if err = o.Read(&exceptionalChild); err == nil {
		var infos []orm.Params
		where := " allergen = \"" + string(allergen) + "\" AND somatotype = " + strconv.Itoa(int(somatotype)) + " AND student_id = " + strconv.Itoa(student_id) + " AND kindergarten_id = " + strconv.Itoa(kindergarten_id)
		if n, er := o.Raw("SELECT allergen FROM `exceptional_child` WHERE " + where).Values(&infos); er == nil && n > 0 {
			// 已存在相同数据
			return 0, nil
		} else {
			if child_name != "" {
				exceptionalChild.ChildName = child_name
			}

			if class != 0 {
				exceptionalChild.Class = class
			}

			if somatotype != 0 {
				exceptionalChild.Somatotype = somatotype
			}

			if allergen != "" {
				exceptionalChild.Allergen = allergen
			}

			exceptionalChild.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

			if student_id != 0 {
				exceptionalChild.StudentId = student_id
			}

			if num, err := o.Update(&exceptionalChild); err == nil {
				return int(num), err
			} else {
				return 1, err
			}
		}
	}
	return 1, err
}

// 修改体检数据
func UpInspect(id int, student_id int, child_name string, somatotype string, class_id int) (err error) {

	o := orm.NewOrm()
	var sql string
	if somatotype != "" && class_id == 0 {
		sql += " abnormal_weight = \"" + somatotype + "\" "
	}

	if class_id != 0 && somatotype == "" {
		sql += " class_id = " + strconv.Itoa(class_id)
	}

	if somatotype != "" && class_id != 0 {
		sql += " abnormal_weight = \"" + somatotype + "\",class_id = " + strconv.Itoa(class_id)
	}

	if sql != "" {
		if _, err := o.Raw("UPDATE healthy_inspect SET " + sql + " WHERE id =" + strconv.Itoa(id)).Exec(); err == nil {
			if child_name != "" {
				if _, err = o.QueryTable("student").Filter("student_id", student_id).Update(orm.Params{
					"name": child_name,
				}); err == nil {
					return nil
				} else {
					return err
				}
			}
			return err
		} else {
			return err
		}
	}
	return err
}

// DeleteExceptionalChild deletes ExceptionalChild by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExceptionalChild(id int, del_type int) (err error) {
	o := orm.NewOrm()
	if del_type == 1 {
		e := &ExceptionalChild{Id: id}
		// ascertain id exists in the database
		if err = o.Read(e); err == nil {
			if _, err = o.Delete(&ExceptionalChild{Id: id}); err == nil {
				return nil
			}
		}
	} else {
		if _, err = o.QueryTable("healthy_inspect").Filter("id", id).Update(orm.Params{
			"evaluate": 1,
		}); err == nil {
			return nil
		}
	}

	return err
}

// 根据过敏源获取过敏儿童
func GetAllergenChild(allergen string, kindergarten_id int) (allergenChild []map[string]interface{}, err error) {
	param := strings.Split(allergen, ",")
	o := orm.NewOrm()
	var allergens []orm.Params
	for _, v := range param {
		if v != "" {
			maps := make(map[string]interface{})

			where := " allergen like \"%" + string(v) + "%\" AND kindergarten_id =" + strconv.Itoa(kindergarten_id)
			if _, err := o.Raw("SELECT allergen FROM `exceptional_child` WHERE " + where).Values(&allergens); err == nil {
				if childName, childNum, err := GetChildName(v); err == nil {
					maps["allergen"] = v
					maps["child_name"] = childName
					maps["child_num"] = childNum
					allergenChild = append(allergenChild, maps)
				} else {
					return allergenChild, err
				}
			} else {
				return allergenChild, err
			}
		}
	}

	return allergenChild, nil
}

// 获取过敏儿童名称
func GetChildName(val string) (childName string, childNum int64, err error) {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	childNum, errs := o.QueryTable("exceptional_child").Filter("allergen__contains", val).Count()
	if errs == nil && childNum > 0 {
		where := " allergen like \"%" + string(val) + "%\""
		if _, err := o.Raw("SELECT ex.child_name,stu.class_info FROM `exceptional_child` as ex LEFT JOIN student as stu ON stu.student_id = ex.student_id WHERE " + where).ValuesList(&lists); err == nil {
			var str string
			for _, row := range lists {
				if row[1] != nil {
					str += row[0].(string) + ":(" + row[1].(string) + ")" + ","
				} else {
					str += row[0].(string) + ","
				}
				s := []rune(str)
				childName = string(s[:len(s)-1])
			}
			return childName, childNum, err
		}
	}

	return childName, childNum, err
}

// 根据baby_id获取过敏源
func GetAllergen(id int, kindergarten_id int) (allergen []interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	idStr := strconv.Itoa(id)
	where := " stu.baby_id=" + idStr + " AND kindergarten_id=" + strconv.Itoa(kindergarten_id)
	var maps []orm.Params
	sql := qb.Select("ex.allergen, ex.id, COUNT(DISTINCT ex.allergen) as field").From("student as stu").LeftJoin("exceptional_child as ex").On("ex.student_id = stu.student_id").Where(where).GroupBy("ex.`allergen`").String()
	if num, err := o.Raw(sql).Values(&maps); err == nil && num > 0 {
		for _, row := range maps {
			if row["allergen"] != nil {
				delete(row, "field")
				allergen = append(allergen, row)
			}
		}
		return allergen, err
	}
	return nil, err
}

// 过敏食物报备
func AllergenPreparation(child_name string, somatotype int8, allergens string, source int8, kindergarten_id int, creator int, baby_id int) (id int64, err error) {
	var exceptionalChildList []ExceptionalChild
	o := orm.NewOrm()
	allergen := strings.Split(allergens, ",")
	var maps []orm.Params
	if num, err := o.Raw("SELECT stu.student_id,org.member_id FROM student as stu LEFT JOIN organizational_member as org ON stu.student_id = org.member_id WHERE baby_id = ? LIMIT 1", baby_id).Values(&maps); err == nil && num > 0 {
		if maps[0]["student_id"] != nil && maps[0]["member_id"] != nil {
			student_id, _ := strconv.Atoi(maps[0]["student_id"].(string))
			class, _ := strconv.Atoi(maps[0]["member_id"].(string))
			for _, v := range allergen {
				if v != "" {
					var infos []orm.Params
					where := " allergen like \"%" + string(v) + "%\" AND somatotype = " + strconv.Itoa(int(somatotype)) + " AND student_id = ? AND kindergarten_id = ? "
					if n, er := o.Raw("SELECT allergen FROM `exceptional_child` WHERE "+where, student_id, kindergarten_id).Values(&infos); er != nil || n == 0 {
						var exceptionalChild ExceptionalChild
						exceptionalChild.Id = 0
						exceptionalChild.ChildName = child_name
						exceptionalChild.Class = class
						exceptionalChild.Somatotype = somatotype
						exceptionalChild.Allergen = v
						exceptionalChild.Source = source
						exceptionalChild.KindergartenId = kindergarten_id
						exceptionalChild.Creator = creator
						exceptionalChild.StudentId = student_id
						exceptionalChild.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
						exceptionalChild.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
						exceptionalChildList = append(exceptionalChildList, exceptionalChild)
					} else {
						// 已存在相同数据
						return 0, nil
					}
				}
			}
			id, err = o.InsertMulti(1, exceptionalChildList)
		}
		return num, err
	}
	err = errors.New("该学生不存在")
	return id, err
}

// app端特殊儿童列表
func GetExceptionalInspect(page int64, limit int64, keyword string, kindergarten_id int) (Page, error) {
	o := orm.NewOrm()
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	var total int64
	where := "1"
	exWhere := "1"
	if kindergarten_id != 0 {
		kid := strconv.Itoa(kindergarten_id)
		where += " AND he.kindergarten_id = " + kid
		exWhere += " AND ex.kindergarten_id = " + kid
	}

	var num []orm.Params

	if n, er := o.Raw("SELECT id FROM healthy_body AS he WHERE " + where + " ORDER BY he.id DESC LIMIT 1").Values(&num); er == nil && n > 0 {
		where += " AND he.body_id = " + num[0]["id"].(string)
	}

	where += " AND ( he.abnormal_weight != '正常' Or he.abnormal_height = '正常' ) AND ( he.abnormal_weight != '' Or he.abnormal_height != '' )"

	if keyword != "" {
		where += " AND he.abnormal_weight LIKE \"%" + string(keyword) + "%\" OR stu.name LIKE \"%" + string(keyword) + "%\""
		exWhere += " AND ex.child_name like \"%" + string(keyword) + "%\" OR ex.allergen like \"%" + string(keyword) + "%\""
	}
	where += " AND stu.student_id IS NOT NULL "
	exWhere += " AND stu.student_id IS NOT NULL "
	var totalNum []orm.Params
	if _, totalErr := o.Raw("SELECT ll.* FROM (SELECT stu.name AS child_name,he.class_id,he.abnormal_weight,he.student_id,MAX(he.created_at) AS updated_at,stu.class_info as class_name FROM healthy_inspect AS he LEFT JOIN student AS stu ON stu.student_id = he.student_id LEFT JOIN organizational AS o ON o.id = he.class_id WHERE " + where + " GROUP BY student_id UNION ALL SELECT ex.child_name,ex.class AS class_id,ex.allergen AS abnormal_weight,ex.student_id,ex.updated_at,stu.class_info as class_name FROM exceptional_child AS ex LEFT JOIN student as stu ON stu.student_id = ex.student_id where " + exWhere + ") AS ll").Values(&totalNum); totalErr == nil {
		total = int64(len(totalNum))
	}
	var maps []orm.Params
	if _, err := o.Raw("SELECT ll.* FROM (SELECT stu.name AS child_name,he.class_id,he.abnormal_weight,he.student_id,MAX(he.created_at) AS updated_at,stu.class_info as class_name FROM healthy_inspect AS he LEFT JOIN student AS stu ON stu.student_id = he.student_id LEFT JOIN organizational AS o ON o.id = he.class_id WHERE " + where + " GROUP BY student_id UNION ALL SELECT ex.child_name,ex.class AS class_id,ex.allergen AS abnormal_weight,ex.student_id,ex.updated_at,stu.class_info as class_name FROM exceptional_child AS ex LEFT JOIN student as stu ON stu.student_id = ex.student_id where " + exWhere + ") AS ll ORDER BY ll.updated_at DESC LIMIT " + strconv.Itoa(int(offset)) + "," + strconv.Itoa(int(limit))).Values(&maps); err == nil {
		var newMap []orm.Params
		for _, v := range maps {

			if v["abnormal_weight"].(string) != "肥胖" && v["abnormal_weight"].(string) != "瘦小" {
				v["somatot_info"] = v["abnormal_weight"].(string) + "过敏"
			} else {
				v["somatot_info"] = v["abnormal_weight"].(string)
			}

			delete(v, "abnormal_weight")
			delete(v, "updated_at")
			newMap = append(newMap, v)
		}
		pageNum := int(math.Ceil(float64(total) / float64(limit)))
		return Page{newMap, total, pageNum}, nil
	} else {
		return Page{}, err
	}

	return Page{}, nil
}
