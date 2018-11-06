package models

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"github.com/hprose/hprose-golang/rpc"
)

type Kindergarten struct {
	Id               int       `json:"kindergarten_id" orm:"column(kindergarten_id);auto" description:"编号"`
	Name             string    `json:"name" orm:"column(name);size(50)" description:"幼儿园名称"`
	LicenseNo        int       `json:"license_no" orm:"column(license_no)" description:"执照号"`
	KinderGrade      string    `json:"kinder_grade" orm:"column(kinder_grade);size(45)" description:"幼儿园级别"`
	KinderChildNo    int       `json:"kinder_child_no" orm:"column(kinder_child_no)" description:"分校数"`
	Address          string    `json:"address" orm:"column(address);size(50)" description:"地址"`
	TenantId         int       `json:"tenant_id" orm:"column(tenant_id)" description:"租户，企业编号"`
	Status           int8      `json:"status" orm:"column(status)" description:"状态：0:正常，1:删除"`
	Introduce        string    `json:"introduce" orm:"column(introduce);size(255)" description:"幼儿园介绍"`
	Phone            string    `json:"phone" orm:"column(phone);size(255)" description:"手机号"`
	Region           string    `json:"region" orm:"column(region);size(255)" description:"地区"`
	Telephone        string    `json:"telephone" orm:"column(telephone);size(255)" description:"幼儿园介绍"`
	Principal        string    `json:"principal" orm:"column(principal);size(255)" description:"负责人"`
	IdNumber         string    `json:"id_number" orm:"column(id_number);size(255)" description:"身份证号"`
	SchoolLicense    string    `json:"school_license" orm:"column(school_license);size(255)" description:"办学许可证"`
	UserId           int       `json:"user_id" orm:"column(user_id);size(255)" description:"幼儿园介绍"`
	IntroducePicture string    `json:"introduce_picture" orm:"column(introduce_picture);size(255)" description:"幼儿园介绍图"`
	CreatedAt        time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt        time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt        time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime);null"`
}

func (t *Kindergarten) TableName() string {
	return "kindergarten"
}

func init() {
	orm.RegisterModel(new(Kindergarten))
}

type OnemoreService struct {
	Tenant func(tenant_id int, kindergarten_id int) error
}

type CourseService struct {
	CoursewareNum func(k_id int) interface{}
}

/*
web-幼儿园介绍详情
*/
func GetKindergartenById(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("k.*", "kc.tax_registration", "kc.catering_services", "kc.private_non_enterprise").From("kindergarten as k").LeftJoin("kindergarten_certificate as kc").
		On("k.kindergarten_id = kc.kindergarten_id").Where("k.kindergarten_id = ?").String()
	_, err = o.Raw(sql, id).Values(&v)
	if err == nil {
		paginatorMap = make(map[string]interface{})
		paginatorMap["data"] = v //返回数据
	}
	return paginatorMap, err
}

/*
oms-设置园长/教师
*/
func AddPrincipal(user_id int, kindergarten_id int, role int) error {
	o := orm.NewOrm()
	o.Begin()
	var or []orm.Params
	var t []orm.Params
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("organizational").Where("kindergarten_id = ?").And("type = 1").And("is_fixed = 1").String()
	_, err := o.Raw(sql, kindergarten_id).Values(&or)
	if err == nil {
		if or == nil {
			err = errors.New("幼儿园没有组织")
			return err
		}
	}
	uk, _ := User.GetUKByUser(user_id, 1)
	if uk == nil {
		User.CreateUK(user_id, kindergarten_id, role)
	} else {
		uks, _ := json.Marshal(uk)
		var ukss map[string]interface{}
		json.Unmarshal(uks, &ukss)
		uk_id := ukss["id"].(float64)
		ukId := int(uk_id)
		err := User.UpdateByUkId(ukId, user_id, kindergarten_id, role)
		if err != nil {
			err := errors.New("用户不存在")
			return err
		}
	}
	if role == 1 || role == 5 {
		userInfo, err := User.GetOneByUserId(user_id)
		if err != nil {
			return err
		}
		userinfo, _ := json.Marshal(userInfo)
		var user map[string]interface{}
		json.Unmarshal(userinfo, &user)
		if err == nil {
			o.QueryTable("teacher").Filter("user_id", user_id).Filter("status", 0).Update(orm.Params{
				"status": 1,
			})
			if err != nil {
				o.Rollback()
				return err
			}
		}
		if role == 1 {
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("t.teacher_id").From("teacher as t").Where("user_id = ?").And("status = 0").And("kindergarten_id = ?").String()
			_, err := o.Raw(sql, user_id, kindergarten_id).Values(&t)
			if err == nil {
				if t == nil {
					sql := "insert into teacher set user_id = ?,name = ?,phone = ?,sex = ?,age = ?,address = ?,kindergarten_id = ?,avatar = ?,status = ?,class_info = ?"
					id, err := o.Raw(sql, user_id, user["name"].(string), user["phone"].(string), user["sex"], user["age"], user["address"].(string), kindergarten_id, user["avatar"].(string), 1, "园长").Exec()
					if err != nil {
						o.Rollback()
						return err
					}
					if err == nil {
						teacherId, _ := id.LastInsertId()
						for _, value := range or {
							sql := "insert into organizational_member set organizational_id = ?,member_id = ?,is_principal = ?,type = ?"
							o.Raw(sql, value["id"], teacherId, 1, 0).Exec()
						}
					} else {
						o.Rollback()
						return err
					}
				}
			}
		} else {
			sql := "insert into teacher set user_id = ?,name = ?,phone = ?,sex = ?,age = ?,address = ?,kindergarten_id = ?,avatar = ?"
			_, err := o.Raw(sql, user_id, user["name"].(string), user["phone"].(string), user["sex"], user["age"], user["address"].(string), kindergarten_id, user["avatar"].(string)).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}
	if err == nil {
		o.Commit()
		return nil
	}
	return err
}

/*
oms幼儿园列表
*/
func GetAll(page int, prepage int, search string, kindergarten_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	var condition []interface{}
	where := "1=1 "
	if search != "" {
		where += " AND k.name like ?"
		condition = append(condition, "%"+search+"%")
	}
	if kindergarten_id != 0 {
		where += " AND k.kindergarten_id = ?"
		condition = append(condition, kindergarten_id)
	}
	// 构建查询对象
	sql := qb.Select("count(*)").From("kindergarten as k").LeftJoin("kindergarten_certificate as kc").
		On("k.kindergarten_id = kc.kindergarten_id").Where(where).And("status = ?").String()
	var total int64
	err = o.Raw(sql, condition, 0).QueryRow(&total)
	if err == nil {
		var v []orm.Params
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(total) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("k.*", "kc.tax_registration", "kc.catering_services", "kc.private_non_enterprise").From("kindergarten as k").LeftJoin("kindergarten_certificate as kc").
			On("k.kindergarten_id = kc.kindergarten_id").Where(where).And("k.status = ?").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, condition, 0).Values(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total         //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
学生姓名搜索班级
*/
func StudentClass(page int, prepage int, name string) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	sql := qb.Select("count(*)").From("organizational_member as om").LeftJoin("student as t").
		On("om.member_id = t.student_id").LeftJoin("organizational as o").
		On("o.id = om.organizational_id").Where("t.name = ?").And("om.type = 1").String()
	var total int64
	err = o.Raw(sql, name).QueryRow(&total)
	if err == nil {
		var v []orm.Params
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(total) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("o.id as class_id", "o.name as class_name").From("organizational_member as om").LeftJoin("student as t").
			On("om.member_id = t.student_id").LeftJoin("organizational as o").
			On("o.id = om.organizational_id").Where("t.name = ?").And("om.type = 1").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, name).Values(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total         //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
添加幼儿园
*/
func AddKindergarten(name string, license_no int, kinder_grade string, kinder_child_no int, tenant_id int, phone string, region string, address string, telephone string, principal string, id_number string, school_license string, tax_registration string, catering_services string, private_non_enterprise string, user_id int) error {
	o := orm.NewOrm()
	o.Begin()
	var Onemore *OnemoreService
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation(timeLayout, timenow, loc)
	sql := "insert into kindergarten set name = ?,license_no = ?,kinder_grade = ?,kinder_child_no = ?,address = ?,tenant_id = ?,created_at = ?,phone = ?,region = ?,telephone = ?,principal = ?,id_number = ?,school_license = ?,user_id = ?"
	id, err := o.Raw(sql, name, license_no, kinder_grade, kinder_child_no, address, tenant_id, createTime, phone, region, telephone, principal, id_number, school_license, user_id).Exec()
	kindergartern_id, _ := id.LastInsertId()
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_SMS_SERVER"))
	client.UseService(&Onemore)
	kid := strconv.FormatInt(kindergartern_id, 10)
	kId, _ := strconv.Atoi(kid)
	err = Onemore.Tenant(tenant_id, kId)
	if err != nil {
		o.Rollback()
		return err
	}
	sql1 := "insert into kindergarten_certificate set tax_registration = ?,catering_services = ?,private_non_enterprise = ?,kindergarten_id =?"
	_, err = o.Raw(sql1, tax_registration, catering_services, private_non_enterprise, kId).Exec()

	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err = o.Raw(sql, kindergartern_id, 0, "管理层", 1, 1, 0, 1).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizational_id, err := id.LastInsertId()
	oid := strconv.FormatInt(organizational_id, 10)
	oid = oid + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
		id, err = o.Raw(sql, kindergartern_id, organizational_id, "园长", 1, 2, oid, 1).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//年级组
	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err = o.Raw(sql, kindergartern_id, 0, "年级组", 1, 1, 0, 2).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizationalId, err := id.LastInsertId()
	o_id := strconv.FormatInt(organizationalId, 10)
	o_id = o_id + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
		_, err = o.Raw(sql, kindergartern_id, organizationalId, "大班", 1, 2, o_id, 2, 3).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		if err == nil {
			sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
			_, err = o.Raw(sql, kindergartern_id, organizationalId, "中班", 1, 2, o_id, 2, 2).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
			if err == nil {
				sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
				_, err = o.Raw(sql, kindergartern_id, organizationalId, "小班", 1, 2, o_id, 2, 1).Exec()
				if err != nil {
					o.Rollback()
					return err
				}
			}
		}
	}

	//其他（保健医）
	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err = o.Raw(sql, kindergartern_id, 0, "其他", 1, 1, 0, 3).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizationalIds, err := id.LastInsertId()
	org_id := strconv.FormatInt(organizationalIds, 10)
	or_id := org_id + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
		_, err = o.Raw(sql, kindergartern_id, organizationalIds, "保健医", 1, 2, or_id, 3, 0).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}
	if err == nil {
		o.Commit()
		return nil
	}
	return err
}

/*
删除幼儿园
*/
func DeleteKinder(kindergarten_id int) (err error) {
	o := orm.NewOrm()
	var v []orm.Params
	o.Begin()
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	_, err = o.QueryTable("kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"status": 1,
	})
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("kindergarten_certificate").Filter("kindergarten_id", kindergarten_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("id").From("organizational").Where("kindergarten_id = ?").String()
	_, err = o.Raw(sql, kindergarten_id).Values(&v)
	for _, val := range v {
		_, err = o.QueryTable("organizational_member").Filter("organizational_id", val["id"]).Delete()
	}
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.QueryTable("organizational").Filter("kindergarten_id", kindergarten_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.QueryTable("student").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"status": 1,
	})
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("teacher").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"status": 1,
	})
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.QueryTable("baby_kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"status": 1,
	})
	if err != nil {
		o.Rollback()
		return err
	}
	User.DeleteUKByKindergartenId(kindergarten_id)
	o.Commit()
	return err
}

/*
编辑幼儿园
*/
func UpdataKinder(kindergarten_id int, name string, license_no int, kinder_grade string, kinder_child_no int, tenant_id int, phone string, region string, address string, telephone string, principal string, id_number string, school_license string, tax_registration string, catering_services string, private_non_enterprise string, user_id int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	_, err = o.QueryTable("kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"name": name, "license_no": license_no, "kinder_grade": kinder_grade, "kinder_child_no": kinder_child_no,
		"address": address, "tenant_id": tenant_id, "phone": phone, "region": region, "telephone": telephone, "principal": principal, "id_number": id_number,
		"school_license": school_license, "user_id": user_id,
	})
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("kindergarten_certificate").Filter("kindergarten_id", kindergarten_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}

	sql := "insert into kindergarten_certificate set tax_registration = ?,catering_services = ?,private_non_enterprise = ?,kindergarten_id = ?"
	_, err = o.Raw(sql, tax_registration, catering_services, private_non_enterprise, kindergarten_id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

/*
登陆幼儿园信息
*/
func GetKg(user_id int) (value map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	var ident []orm.Params
	var kinder []orm.Params
	var permission []orm.Params
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	role, _ := User.GetUKByUserId(user_id)
	rol := make(map[string]interface{})
	ro, _ := json.Marshal(role)
	json.Unmarshal([]byte(ro), &rol)
	//幼儿园信息
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("k.name", "k.kindergarten_id").From("teacher as t").LeftJoin("kindergarten as k").
		On("t.kindergarten_id = k.kindergarten_id").Where("t.user_id = ?").String()
	_, err = o.Raw(sql, user_id).Values(&kinder)

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
			if kinder == nil {
				value["kindergarten_name"] = nil
				value["kindergarten_id"] = nil
			} else {
				value["kindergarten_name"] = kinder[0]["name"]
				value["kindergarten_id"] = kinder[0]["kindergarten_id"]
			}
			value["permission"] = permission
			value["role"] = rol["role"]
			if ident != nil {
				value["ident"] = 1
			} else {
				value["ident"] = 0
			}
			return value, nil
		} else {
			value := v[0]
			if kinder == nil {
				value["kindergarten_name"] = nil
				value["kindergarten_id"] = nil
			} else {
				value["kindergarten_name"] = kinder[0]["name"]
				value["kindergarten_id"] = kinder[0]["kindergarten_id"]
			}
			value["permission"] = permission
			value["role"] = rol["role"]
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

/*
oms-幼儿园所有成员
*/
func GetKinderMbmber(kindergarten_id int, page int, prepage int) (value map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var teacher []orm.Params
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	sql := qb.Select("count(*)").From("teacher as t").Where("t.kindergarten_id = ?").And("isnull(t.deleted_at)").String()
	var total int64
	err = o.Raw(sql, kindergarten_id).QueryRow(&total)
	var Stotal int64
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*)").From("student as s").LeftJoin("baby_kindergarten as bk").
		On("s.baby_id = bk.baby_id").Where("s.kindergarten_id = ?").And("isnull(s.deleted_at)").String()
	err = o.Raw(sql, kindergarten_id).QueryRow(&Stotal)
	total = total + Stotal
	if err == nil {
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(total) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("s.*", "bk.actived").From("student as s").LeftJoin("baby_kindergarten as bk").
			On("s.baby_id = bk.baby_id").Where("s.kindergarten_id = ?").And("isnull(s.deleted_at)").Limit(prepage).Offset(limit).String()
		_, err = o.Raw(sql, kindergarten_id).Values(&student)
		if student == nil {
			err = errors.New("该幼儿园没有学生")
			return nil, err
		}
		for _, v := range student {
			role := 6
			v["role"] = role
		}
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = qb.Select("t.*").From("teacher as t").Where("t.kindergarten_id = ?").And("isnull(t.deleted_at)").Limit(prepage).Offset(limit).String()
		_, err = o.Raw(sql, kindergarten_id).Values(&teacher)
		if teacher == nil {
			err = errors.New("该幼儿园没有教师")
			return nil, err
		}
		rol := make(map[string]interface{})
		for _, v := range teacher {
			user_id, _ := strconv.Atoi(v["user_id"].(string))
			role, _ := User.GetUKByUserId(user_id)
			ro, _ := json.Marshal(role)
			json.Unmarshal([]byte(ro), &rol)
			v["role"] = rol["role"]
			v["actived"] = rol["actived"]
			student = append(student, v)
		}
		value = make(map[string]interface{})
		value["total"] = total //总条数
		value["data"] = student
		value["page_num"] = totalpages //总页数
		return value, err
	}
	return nil, err
}

/*
饮食班级
*/
func FoodClass(kindergarten_id int) (value map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.*").From("organizational as o").Where("o.kindergarten_id = ?").And("o.is_fixed = 1 and o.level = 2 and o.type = 2").String()
	_, err = o.Raw(sql, kindergarten_id).Values(&student)
	value = make(map[string]interface{})
	value["data"] = student
	return value, err
}

/*
饮食班级
*/
func FoodScale(is_muslim int, kindergarten_id int, class_type string) (value map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var girl []orm.Params
	var boy []orm.Params
	where := "1 "
	where += " and o.class_type in (" + class_type + ") "
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*) as num", "sum(s.age) as age").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("s.is_muslim = ?").And("om.type = 1").And("s.status = 1 and o.level = 3 and o.kindergarten_id = ?").String()
	_, err = o.Raw(sql, is_muslim, kindergarten_id).Values(&student)
	if err != nil {
		return nil, err
	}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*) as num").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("om.type = 1").And("s.is_muslim = ?").And("s.status = 1 and o.level = 3 and o.kindergarten_id = ? and s.sex = ?").String()
	_, err = o.Raw(sql, is_muslim, kindergarten_id, 0).Values(&boy)
	if err != nil {
		return nil, err
	}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*) as num").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("om.type = 1").And("s.is_muslim = ?").And("s.status = 1 and o.level = 3 and o.kindergarten_id = ? and s.sex = ?").String()
	_, err = o.Raw(sql, is_muslim, kindergarten_id, 1).Values(&girl)
	if err != nil {
		return nil, err
	}
	var num string
	var age string
	if student[0]["age"] == nil {
		age = "0"
	} else {
		age = student[0]["age"].(string)
	}
	if student[0]["num"] == nil {
		num = "0"
	} else {
		num = student[0]["num"].(string)
	}
	ages, _ := strconv.Atoi(age)
	nums, _ := strconv.Atoi(num)
	NewAge := int(math.Ceil(float64(ages) / float64(nums)))
	if NewAge < 1 {
		NewAge = 0
	}
	value = make(map[string]interface{})
	value["num"] = nums
	value["age"] = NewAge
	value["scale"] = "" + boy[0]["num"].(string) + ":" + girl[0]["num"].(string) + ""
	return value, err
}

/*
班级
*/
func Class(kindergarten_id int, class_type int) (value interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.name as class_name").From("organizational as o").Where("o.class_type = ?").And("o.kindergarten_id = ? and o.level = 2").String()
	_, err = o.Raw(sql, class_type, kindergarten_id).Values(&class)
	if err != nil {
		return nil, err
	}
	return class, err
}

/*
编辑幼儿园
*/
func UpdataKinderInduce(kindergarten_id int, introduce string, introduce_picture string) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
		"introduce": introduce, "introduce_picture": introduce_picture,
	})
	if err != nil {
		return err
	}
	return err
}

/*
取消关联幼儿园
*/
func Reset(baby_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", baby_id).Delete()
	if err == nil {
		var student []orm.Params
		timeLayout := "2006-01-02 15:04:05" //转化所需模板
		loc, _ := time.LoadLocation("")
		timenow := time.Now().Format("2006-01-02 15:04:05")
		TimeDel, _ := time.ParseInLocation(timeLayout, timenow, loc)
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("s.student_id").From("student as s").Where("s.baby_id = ?").String()
		_, err = o.Raw(sql, baby_id).Values(&student)
		if err == nil && student != nil {
			_, err = o.QueryTable("organizational_member").Filter("member_id", student[0]["student_id"]).Delete()
			if err == nil {
				_, err = o.QueryTable("student").Filter("baby_id", baby_id).Update(orm.Params{
					"deleted_at": TimeDel,
				})
				return err
			}
		}
	}
	return err
}

//学生，教师，班级数量
func Count(kindergarten_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	var teacher []orm.Params
	var student []orm.Params
	paginatorMap = make(map[string]interface{})
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*) as num").From("teacher").Where("kindergarten_id = ? and class_info != ? and status != ? and isnull(deleted_at)").String()
	_, err = o.Raw(sql, kindergarten_id, "园长", 2).Values(&teacher)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*) as num").From("student").Where("kindergarten_id = ? and status != ? and isnull(deleted_at)").String()
	_, err = o.Raw(sql, kindergarten_id, 2).Values(&student)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*) as num").From("organizational").Where("kindergarten_id = ? and level = 3 and type = 2").String()
	_, err = o.Raw(sql, kindergarten_id).Values(&class)
	if err == nil {
		paginatorMap["class"] = class[0]["num"]
		paginatorMap["teacher"] = teacher[0]["num"]
		paginatorMap["student"] = student[0]["num"]
		return paginatorMap, nil
	}
	return nil, err
}

/*
认证幼儿园
*/
func AttestKindergarten(name string, principal string, phone string, user_id int) error {
	o := orm.NewOrm()
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation(timeLayout, timenow, loc)
	sql := "insert into kindergarten set name = ?,principal = ?,phone = ?,created_at = ?,status = ?,user_id = ?"
	res, err := o.Raw(sql, name, principal, phone, createTime, 2, user_id).Exec()
	kgId, _ := res.LastInsertId()
	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err := o.Raw(sql, kgId, 0, "管理层", 1, 1, 0, 1).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizational_id, err := id.LastInsertId()
	oid := strconv.FormatInt(organizational_id, 10)
	oid = oid + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
		id, err = o.Raw(sql, kgId, organizational_id, "园长", 1, 2, oid, 1).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//年级组
	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err = o.Raw(sql, kgId, 0, "年级组", 1, 1, 0, 2).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizationalId, err := id.LastInsertId()
	o_id := strconv.FormatInt(organizationalId, 10)
	o_id = o_id + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
		_, err = o.Raw(sql, kgId, organizationalId, "大班", 1, 2, o_id, 2, 3).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		if err == nil {
			sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
			_, err = o.Raw(sql, kgId, organizationalId, "中班", 1, 2, o_id, 2, 2).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
			if err == nil {
				sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
				_, err = o.Raw(sql, kgId, organizationalId, "小班", 1, 2, o_id, 2, 1).Exec()
				if err != nil {
					o.Rollback()
					return err
				}
			}
		}
	}

	//其他（保健医）
	sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?"
	id, err = o.Raw(sql, kgId, 0, "其他", 1, 1, 0, 3).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	organizationalIds, err := id.LastInsertId()
	org_id := strconv.FormatInt(organizationalIds, 10)
	or_id := org_id + ","
	if err == nil {
		sql = "insert into organizational set kindergarten_id = ?,parent_id = ?,name = ?,is_fixed = ?,level = ?,parent_ids = ?,type = ?,class_type = ?"
		_, err = o.Raw(sql, kgId, organizationalIds, "保健医", 1, 2, or_id, 3, 0).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}
	if err == nil {
		o.Commit()
		return nil
	}
	return err
}

/*
未认证幼儿园列表
*/
func AttestKindergartenAll(page int, prepage int, search string) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	var con []interface{}
	where := "1 "
	if search != "" {
		where += "AND k.name like ? "
		con = append(con, "%"+search+"%")
	}
	// 构建查询对象
	sql := qb.Select("count(*)").From("kindergarten as k").LeftJoin("kindergarten_certificate as kc").
		On("k.kindergarten_id = kc.kindergarten_id").Where(where).And("k.status = ? and isnull(k.deleted_at)").String()
	var total int64
	err = o.Raw(sql, con, 2).QueryRow(&total)
	if err == nil {
		var v []orm.Params
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(total) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("k.*", "kc.tax_registration", "kc.catering_services", "kc.private_non_enterprise").From("kindergarten as k").LeftJoin("kindergarten_certificate as kc").
			On("k.kindergarten_id = kc.kindergarten_id").Where(where).And("k.status = ? and isnull(k.deleted_at)").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, con, 2).Values(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total         //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
认证通过/未通过幼儿园
*/
func Attest(kindergarten_id int, name string, phone string, region string, address string, telephone string, principal string, id_number string, school_license string, tax_registration string, catering_services string, private_non_enterprise string, ty int, user_id int, title string, content string) (err error) {
	o := orm.NewOrm()
	var role int = 1
	var Notice *NoticeService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	if ty == 1 {
		_, err = o.QueryTable("kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
			"name": name, "phone": phone, "region": region, "address": address, "telephone": telephone, "principal": principal, "id_number": id_number, "school_license": school_license, "status": 0,
		})
		if err != nil {
			return err
		}
		sql := "insert into kindergarten_certificate set tax_registration = ?,catering_services = ?,private_non_enterprise = ?,kindergarten_id = ?"
		_, err := o.Raw(sql, tax_registration, catering_services, private_non_enterprise, kindergarten_id).Exec()
		if err != nil {
			return err
		}
		if err == nil {
			var or []orm.Params
			var t []orm.Params
			var User *UserService
			client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
			client.UseService(&User)
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("*").From("organizational").Where("kindergarten_id = ?").And("type = 1").And("is_fixed = 1").String()
			_, err := o.Raw(sql, kindergarten_id).Values(&or)
			if err == nil {
				if or == nil {
					err = errors.New("幼儿园没有组织")
					return err
				}
			}
			uk, _ := User.GetUKByUser(user_id, 1)
			if uk == nil {
				_, err := User.CreateUK(user_id, kindergarten_id, role)
				if err != nil {
					return err
				}
			} else {
				uks, err := json.Marshal(uk)
				if err != nil {
					return err
				}
				var ukss map[string]interface{}
				json.Unmarshal(uks, &ukss)
				uk_id := ukss["id"].(float64)
				ukId := int(uk_id)
				err = User.UpdateByUkId(ukId, user_id, kindergarten_id, role)
				if err != nil {
					err := errors.New("用户不存在")
					return err
				}
			}
			userInfo, err := User.GetOneByUserId(user_id)
			if err != nil {
				return err
			}
			userinfo, err := json.Marshal(userInfo)
			if err != nil {
				return err
			}
			var user map[string]interface{}
			json.Unmarshal(userinfo, &user)
			if err == nil {
				_, err = o.QueryTable("teacher").Filter("user_id", user_id).Filter("status", 0).Update(orm.Params{
					"status": 1,
				})
				if err != nil {
					return err
				}
			}
			qb, _ = orm.NewQueryBuilder("mysql")
			sql = qb.Select("t.teacher_id").From("teacher as t").Where("user_id = ?").And("status = 0").And("kindergarten_id = ?").String()
			_, err = o.Raw(sql, user_id, kindergarten_id).Values(&t)
			if err == nil {
				if t == nil {
					sql := "insert into teacher set user_id = ?,name = ?,phone = ?,sex = ?,age = ?,address = ?,kindergarten_id = ?,avatar = ?,status = ?,class_info = ?"
					id, err := o.Raw(sql, user_id, user["name"].(string), user["phone"].(string), user["sex"], user["age"], user["address"].(string), kindergarten_id, user["avatar"].(string), 1, "园长").Exec()
					if err != nil {
						return err
					}
					if err == nil {
						teacherId, _ := id.LastInsertId()
						for _, value := range or {
							sql := "insert into organizational_member set organizational_id = ?,member_id = ?,is_principal = ?,type = ?"
							o.Raw(sql, value["id"], teacherId, 1, 0).Exec()
						}
					} else {
						return err
					}
				}
			}
			if err == nil {
				data := make(map[string]interface{})
				data["title"] = "幼儿园权限认证完成通知"
				data["user_id"] = user_id
				data["content"] = "您好。恭喜您在蓝天白云-智慧幼儿园中认证了幼儿园。您的账号获得园长功能权限。"
				data["notice_type"] = 6
				data["choice_type"] = 3
				result, _ := json.Marshal(data)
				User.UpdateUK(user_id, 1, 0)
				Notice.AttestKindergarten(string(result))
				return nil
			}
		}
	} else {
		_, err = o.QueryTable("kindergarten").Filter("kindergarten_id", kindergarten_id).Update(orm.Params{
			"status": 1,
		})
		if err != nil {
			return err
		}
		_, err = o.QueryTable("organizational").Filter("kindergarten_id", kindergarten_id).Delete()
		if err != nil {
			return err
		}
		if err == nil {
			data := make(map[string]interface{})
			data["title"] = title
			data["user_id"] = user_id
			data["content"] = content
			data["notice_type"] = 6
			data["choice_type"] = 3
			result, _ := json.Marshal(data)
			Notice.AttestKindergarten(string(result))
			return nil
		}
	}
	return err
}
