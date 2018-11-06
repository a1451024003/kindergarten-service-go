package services

import (
	"encoding/json"
	"fmt"
	"kindergarten-service-go/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hprose/hprose-golang/rpc"
)

var User *UserService

type KindergartenServer struct {
}

func (c *KindergartenServer) Init() {
	server := rpc.NewHTTPService()
	server.AddAllMethods(&KindergartenServer{})
	beego.Handler("/rpc/kindergarten", server)
}

type UserService struct {
	GetBabyInfo func(id int) (map[string]interface{}, error)
}

//班级信息
func (c *KindergartenServer) GetKg(user_id int, kindergarten_id int) (value map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	var ident []orm.Params
	var kinder []orm.Params
	var permission []orm.Params
	//幼儿园信息
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("k.name").From("kindergarten as k").Where("kindergarten_id = ?").String()
	_, err = o.Raw(sql, kindergarten_id).Values(&kinder)
	//权限信息
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("p.identification").From("user_permission as up").LeftJoin("permission as p").
		On("up.permission_id = p.id").Where("up.user_id = ?").String()
	_, err = o.Raw(sql, user_id).Values(&permission)
	//班级信息
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("o.id as class_id", "o.name as class_name", "o.class_type", "t.teacher_id").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where("t.user_id = ?").And("o.type = 2").And("o.level = 3").String()
	_, err = o.Raw(sql, user_id).Values(&v)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("p.identification").From("user_permission as up").LeftJoin("permission as p").
		On("up.permission_id = p.id").Where("up.user_id = ? and p.type = 0").String()
	_, err = o.Raw(sql, user_id).Values(&ident)

	if err == nil {
		if v == nil {
			value := make(map[string]interface{})
			value["kindergarten_name"] = kinder[0]["name"]
			jsons, _ := json.Marshal(permission)
			value["permission"] = jsons
			if ident != nil {
				value["ident"] = 1
			} else {
				value["ident"] = 0
			}
			return value, nil
		} else {
			value := v[0]
			value["kindergarten_name"] = kinder[0]["name"]
			jsons, _ := json.Marshal(permission)
			value["permission"] = jsons
			if ident != nil {
				value["ident"] = 1
			} else {
				value["ident"] = 0
			}
			return value, nil
		}
	}
	return nil, err
}

//班级成员
func (c *KindergartenServer) GetMember(organizational_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var teacher []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.*").From("organizational_member as om").LeftJoin("teacher as t").
		On("om.member_id = t.teacher_id").Where("om.organizational_id = ?").And("om.type = 0").String()
	_, err = o.Raw(sql, organizational_id).Values(&teacher)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("s.*").From("organizational_member as om").LeftJoin("student as s").
		On("om.member_id = s.student_id").Where("om.organizational_id = ?").And("om.type = 1").String()
	_, err = o.Raw(sql, organizational_id).Values(&student)
	for _, v := range teacher {
		student = append(student, v)
	}
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = student
		return paginatorMap, nil
	}
	return nil, err
}

func (c *KindergartenServer) GetClass(kindergarten_id int) (ml map[string]interface{}) {
	v := models.GetGroupPermission(kindergarten_id)
	ml = make(map[string]interface{})
	ml["data"] = v
	return ml
}

func (c *KindergartenServer) GetAllergenChild(allergen string, kindergarten_id int) (ml interface{}) {

	if allergenChild, err := models.GetAllergenChild(allergen, kindergarten_id); err == nil {

		jsonData, _ := json.Marshal(allergenChild)
		return string(jsonData)
	} else {
		return nil
	}
}

//班级成员
func (c *KindergartenServer) GetClassName(organizational_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.name as class_name").From("organizational as o").Where("o.id = ?").String()
	_, err = o.Raw(sql, organizational_id).Values(&class)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = class
		return paginatorMap, nil
	}
	return nil, err
}

//宝宝是否在幼儿园
func (c *KindergartenServer) GetBaby(baby_id int) interface{} {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("b.*").From("baby_kindergarten as b").Where("b.baby_id = ?").And("b.status = 0").String()
	_, err := o.Raw(sql, baby_id).Values(&v)
	if err == nil {
		return v
	}
	return nil
}

//教师通知
func (c *KindergartenServer) TeacherNotice(class_type int, kindergarten_id int) (interface{}, error) {
	v, err := models.ClassTeacher(class_type, kindergarten_id)
	if err == nil {
		return v, nil
	}
	return nil, err
}

//学生通知
func (c *KindergartenServer) StudentNotice(class_type int, kindergarten_id int) (interface{}, error) {
	v, err := models.Classtudent(class_type, kindergarten_id)
	if err == nil {
		return v, nil
	}
	return nil, err
}

//获取教师列表
func (c *KindergartenServer) GetTeacher(class_type int, kindergarten_id int) (ml map[string]interface{}, err error) {
	list, err := models.FilterGetTeacher(class_type, kindergarten_id)
	return list, err
}

//班级名称
func (c *KindergartenServer) ClassName(kindergarten_id int, class_type int) (ml interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.name as class_name").From("organizational as o").Where("o.kindergarten_id = ?").And("o.class_type = ? and o.level = 2").String()
	_, err = o.Raw(sql, kindergarten_id, class_type).Values(&class)
	return class, err
}

//获取帖子接收用户
func (c *KindergartenServer) UserPost(user_id int) (ml string, err error) {
	o := orm.NewOrm()
	var teacher []orm.Params
	var student []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("om.organizational_id").From("organizational_member as om").LeftJoin("teacher as t").
		On("om.member_id = t.teacher_id").Where("t.user_id = ?").And("om.type = 0").String()
	_, err = o.Raw(sql, user_id).Values(&teacher)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("s.baby_id").From("organizational_member as om").LeftJoin("student as s").
		On("om.member_id = s.student_id").Where("om.organizational_id = ?").And("om.type = 1").String()
	_, err = o.Raw(sql, teacher[0]["organizational_id"]).Values(&student)

	var creatd string
	var creator []int
	for key, v := range student {
		var crea string
		baby := v["baby_id"].(string)
		UserId, _ := strconv.Atoi(baby)
		client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
		client.UseService(&User)
		uid, err := User.GetBabyInfo(UserId)
		if uid["creator"] != nil {
			uids := uid["creator"].(int)
			creator = append(creator, uids)
			crea = strings.Replace(strings.Trim(fmt.Sprint(creator), "[]"), " ", ",", -1)
		}
		if creatd == "" {
			creatd = crea
		} else {
			creatd += "," + crea
		}
		if key == len(student)-2 {
			createdd := strings.TrimLeft(creatd, ",")
			return createdd, err
		}
	}
	return creatd, err
}

//幼儿园所有人员
func (c *KindergartenServer) TeacherAll(kindergarten_id int) (ml interface{}, err error) {
	o := orm.NewOrm()
	var count []orm.Params
	var teacher []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*) as num").From("teacher").Where("kindergarten_id = ?").And("isnull(deleted_at) and status != ?").String()
	_, err = o.Raw(sql, kindergarten_id, 2).Values(&count)
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("*").From("teacher").Where("kindergarten_id = ?").And("isnull(deleted_at) and status != ?").String()
	_, err = o.Raw(sql, kindergarten_id, 2).Values(&teacher)
	if teacher != nil {
		for _, v := range teacher {
			count = append(count, v)
		}
	}
	data, _ := json.Marshal(count)
	return data, err
}
