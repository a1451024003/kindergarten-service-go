package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"

	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id                         int       `json:"teacher_id" orm:"column(teacher_id);auto" description:"自增id"`
	Name                       string    `json:"name" orm:"column(name);size(20)" description:"姓名"`
	Age                        int       `json:"age" orm:"column(age)" description:"年龄"`
	Sex                        int       `json:"sex" orm:"column(sex)" description:"性别 0男  1女"`
	Avatar                     string    `json:"avatar" orm:"column(avatar);size(150)" description:"头像"`
	Number                     string    `json:"number" orm:"column(number);size(20)" description:"教职工编号"`
	NationOrReligion           string    `json:"nation_or_religion" orm:"column(nation_or_religion);size(10)" description:"民族或宗教"`
	NativePlace                string    `json:"native_place" orm:"column(native_place);size(20)" description:"籍贯"`
	UserId                     int       `json:"user_id" orm:"column(user_id)" description:"用户id"`
	ClassInfo                  string    `json:"class_info" orm:"column(class_info);size(10)" description:"班级信息"`
	Phone                      string    `json:"phone" orm:"column(phone);size(11)" description:"联系电话"`
	EnterGardenTime            string    `json:"enter_garden_time" orm:"column(enter_garden_time);type(date)" description:"进入本园时间"`
	EnterJobTime               string    `json:"enter_job_time" orm:"column(enter_job_time);type(date)" description:"参加工作时间"`
	KindergartenId             int       `json:"kindergarten_id" orm:"column(kindergarten_id)" description:"幼儿园id"`
	Address                    string    `json:"address" orm:"column(address);size(191)" description:"住址"`
	IdNumber                   string    `json:"id_number" orm:"column(id_number);size(18)" description:"身份证号"`
	EmergencyContact           string    `json:"emergency_contact" orm:"column(emergency_contact);size(20)" description:"紧急联系人"`
	EmergencyContactPhone      string    `json:"emergency_contact_phone" orm:"column(emergency_contact_phone);size(11)" description:"紧急联系人电话"`
	Post                       string    `json:"post" orm:"column(post);size(10)" description:"职务"`
	Source                     string    `json:"source" orm:"column(source);size(191)" description:"来源"`
	TeacherCertificationNumber string    `json:"teacher_certification_number" orm:"column(teacher_certification_number);size(20)" description:"教师资格认证编号"`
	TeacherCertificationStatus int8      `json:"teacher_certification_status" orm:"column(teacher_certification_status)" description:"教师资格证书状态，是否认证"`
	Status                     int8      `json:"status" orm:"column(status)" description:"状态：0未分班，1已分班，2离职"`
	Birthday                   string    `json:"birthday" orm:"column(birthday)" description:"出生年月日"`
	CreatedAt                  time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt                  time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt                  time.Time `json:"deleted_at" orm:""`
}

type UserService struct {
	GetOne                   func(string) (int, error)
	UpdateUK                 func(userId int, InviteStatus int, Actived int) error
	DeleteUkByUserId         func(userId int) error
	UpdateByUkId             func(ukId int, userId int, kindergartenId int, role int) error
	GetUKByUser              func(userId int, actived int) (interface{}, error)
	GetUKByUserId            func(userId int) (interface{}, error)
	GetOneByUserId           func(userId int) (interface{}, error)
	GetUsers                 func(types int) ([]map[string]interface{}, error)
	CreateUK                 func(userId int, kindergartenId int, role int) (int64, error)
	ChangeUK                 func(UserKindergartenId int, InviteStatus int) error
	GetUK                    func(string) error
	UpdateUser               func(userId int, name string) error
	DeleteUKByKindergartenId func(KindergartenId int) int64
	UpdateUkByUserId         func(UserId int, Role int) int64
}

type PublishService struct {
	Pub func(notice_range int, id int, text string) error
}

type NoticeService struct {
	InviteSystem       func(value string) error
	PostSystem         func(value string) error
	AttestKindergarten func(value string) error
	NoticeRead         func(notice_id int) error
}

type inviteTeacher struct {
	Name               string `json:"name"`
	Phone              string `json:"phone"`
	Avatar             string `json:"avatar"`
	Role               int    `json:"role"`
	UserKindergartenId int    `json:"user_kindergarten_id"`
	KindergartenId     int    `json:"kindergarten_id"`
}

func (t *Teacher) TableName() string {
	return "teacher"
}

func init() {
	orm.RegisterModel(new(Teacher))
}

/*
教师列表
*/
func GetTeacher(id int, status int, search string, page int, prepage int, date string) (ml map[string]interface{}, err error) {
	var condition []interface{}
	where := "1=1 "
	if id != 0 && status == -1 {
		where += " AND t.kindergarten_id = ? and t.class_info != ? and t.status != ?"
		condition = append(condition, id, "园长", 2)
	}
	if id == 0 {
		where += " AND t.kindergarten_id = ?"
	} else {
		where += " AND t.kindergarten_id = ?"
		condition = append(condition, id)
	}
	if status != -1 {
		where += " AND t.status = ?"
		condition = append(condition, status)
	}
	if search != "" {
		where += " AND t.name like ?"
		condition = append(condition, "%"+search+"%")
	}
	if date != "" {
		where += " AND t.updated_at like ?"
		condition = append(condition, "%"+date+"%")
	}
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	sql := qb.Select("count(*)").From("teacher as t").Where(where).And("isnull(deleted_at)").String()
	var total int64
	err = o.Raw(sql, condition).QueryRow(&total)
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
		sql := qb.Select("t.*").
			From("teacher as t").Where(where).And("isnull(t.deleted_at)").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, condition).Values(&v)
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
班级列表
*/
func GetClass(id int, class_type int, page int, prepage int) (ml map[string]interface{}, err error) {
	var condition []interface{}
	where := "1=1 "
	if id == 0 {
		where += " AND o.kindergarten_id = ?"
	} else {
		where += " AND o.kindergarten_id = ?"
		condition = append(condition, id)
	}
	if class_type != 0 {
		where += " AND o.class_type = ?"
		condition = append(condition, class_type)
	}
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	sql := qb.Select("count(*)").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("t.status = 1").And("o.type = 2").And("o.level = 3").And("om.is_principal = 0").And("isnull(t.deleted_at)").String()
	var total int64
	err = o.Raw(sql, condition).QueryRow(&total)
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
		sql := qb.Select("t.name", "t.avatar", "t.teacher_id", "t.number", "t.phone", "o.name as class").From("teacher as t").LeftJoin("organizational_member as om").
			On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
			On("om.organizational_id = o.id").Where(where).And("isnull(t.deleted_at)").And("om.is_principal = 0 and om.type = 0").And("o.type = 2").And("o.level = 3").And("status = 1").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, condition).Values(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			data := make(map[string][]interface{})
			paginatorMap["total"] = total //总条数
			for _, val := range v {
				if strc, ok := val["class"].(string); ok {
					data[strc] = append(data[strc], val)
				}
			}
			//分页数据
			paginatorMap["data"] = data
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
删除教师
*/
func DeleteTeacher(id int, status int, class_type int) error {
	o := orm.NewOrm()
	v := Teacher{Id: id}
	var tea []orm.Params
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.user_id").
		From("teacher as t").Where("t.teacher_id = ? and isnull(t.deleted_at)").String()
	_, err := o.Raw(sql, id).Values(&tea)
	if err = o.Read(&v); err == nil {
		if status == 0 {
			v.Status = 2
			var User *UserService
			client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
			client.UseService(&User)
			userId, _ := strconv.Atoi(tea[0]["user_id"].(string))
			err := User.DeleteUkByUserId(userId)
			fmt.Println(err)
		} else if status == 2 {
			v.DeletedAt, _ = time.ParseInLocation(timeLayout, timenow, loc)
		}
		if class_type == 3 || class_type == 2 || class_type == 1 {
			v.Status = 0
		}
		if _, err = o.Update(&v); err == nil {
			_, err = o.QueryTable("teachers_show").Filter("teacher_id", id).Delete()
			_, err = o.QueryTable("organizational_member").Filter("member_id", id).Filter("type", 0).Delete()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
教师详情
*/
func GetTeacherInfo(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	v := &Teacher{Id: id}
	if err := o.Read(v); err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}

/*
app教师详情
*/
func GetTeacherOne(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.name", "t.sex", "t.age", "t.phone", "t.address", "t.id_number", "t.teacher_certification_number", "t.kindergarten_id", "t.user_id", "t.teacher_id").From("teacher as t").Where("t.user_id = ? and isnull(t.deleted_at) and t.status != ?").String()
	_, err = o.Raw(sql, id, 2).Values(&v)
	if err == nil {
		paginatorMap = make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}

/*
教师编辑
*/
func UpdateTeacher(m *Teacher) error {
	o := orm.NewOrm()
	v := Teacher{Id: m.Id}
	if m.Post == "" {
		m.Post = "普通教师"
	}
	if err := o.Read(&v); err == nil {
		if _, err = o.Update(m); err != nil {
			return err
		}
	}
	return nil
}

/*
前台教师编辑
*/
func SaveTeacher(teacher string, ty int, member_ids string, is_principal int, organizational_id int, id int, class int) error {
	o := orm.NewOrm()
	var t *Teacher
	json.Unmarshal([]byte(teacher), &t)
	v := Teacher{Id: id}
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	if t.Post == "" {
		t.Post = "普通教师"
	}
	t.Avatar = beego.AppConfig.String("AVATAR")
	if err := o.Read(&v); err == nil {
		t.Id = v.Id
		_, err = o.Update(t)
		if err != nil {
			return err
		} else {
			if class == 1 {
				_, err := o.QueryTable("teacher").Filter("teacher_id", id).Update(orm.Params{
					"status": 0,
				})
				err = User.UpdateUK(t.UserId, 1, 0)
				if err != nil {
					return err
				}
			} else {
				var v []orm.Params
				s := strings.Split(member_ids, ",")
				qb, _ := orm.NewQueryBuilder("mysql")
				sql := qb.Select("o.*").From("organizational as o").Where("id = ?").String()
				_, err := o.Raw(sql, organizational_id).Values(&v)
				if v == nil {
					err = errors.New("没有该班级")
					return err
				}
				//组织架构为园长不能添加
				if v[0]["type"] == "1" && v[0]["is_fixed"] == "1" {
					err = errors.New("不能添加")
					return err
				} else {
					for _, value := range s {
						if value == "" {
							break
						}
						sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
						_, err = o.Raw(sql, organizational_id, ty, value, is_principal).Exec()
						if err == nil {
							if v[0]["type"] == "2" && v[0]["level"] == "3" {
								//0教师
								if ty == 0 {
									if v[0]["class_type"] == "3" {
										class_info := "大班" + v[0]["name"].(string) + ""
										_, err = o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
											"status": 1, "class_info": class_info,
										})
									} else if v[0]["class_type"] == "2" {
										class_info := "中班" + v[0]["name"].(string) + ""
										_, err = o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
											"status": 1, "class_info": class_info,
										})
									} else {
										class_info := "小班" + v[0]["name"].(string) + ""
										_, err = o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
											"status": 1, "class_info": class_info,
										})
									}
								}
							}
							if err == nil {
								err = User.UpdateUK(t.UserId, 1, 0)
								if err != nil {
									return err
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

/*
教师-录入信息
*/
func AddTeacher(m *Teacher) error {
	var User *UserService
	o := orm.NewOrm()
	if m.Post == "" {
		m.Post = "普通教师"
	}
	IdCard := beego.Substr(m.IdNumber, 6, 4)
	idCard, _ := strconv.Atoi(IdCard)
	timeNow := beego.Substr(time.Now().Format("2006-01-02 15:04:05"), 0, 4)
	TimeNow, _ := strconv.Atoi(timeNow)
	m.Age = (TimeNow) - (idCard)
	_, err := o.Insert(m)
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	err = User.UpdateUK(m.UserId, 1, 0)
	if err != nil {
		return err
	}
	return nil
}

/*
教师-app录入信息
*/
func AddAppTeacher(sex int, age int, name string, phone string, address string, id_number string, user_id int, kindergarten_id int, teacher_certification_number string, ty int, notice_id int) (err error) {
	var User *UserService
	var Notice *NoticeService
	o := orm.NewOrm()
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	client = rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	err = Notice.NoticeRead(notice_id)
	if err != nil {
		return err
	}
	if ty == 1 {
		sql := "insert into teacher set sex = ?,age = ?,name = ?,phone = ?,address = ?,id_number = ?,user_id = ?,kindergarten_id =?,teacher_certification_number = ?,status = ?"
		_, err := o.Raw(sql, sex, age, name, phone, address, id_number, user_id, kindergarten_id, teacher_certification_number, 3).Exec()
		if err != nil {
			return err
		}
		err = User.UpdateUK(user_id, 2, 1)
	} else {
		err = User.UpdateUK(user_id, 3, 1)
	}

	return err
}

/*
组织框架移除教师
*/
func RemoveTeacher(teacher_id int, class_id int) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.QueryTable("organizational_member").Filter("organizational_id", class_id).Filter("member_id", teacher_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("teacher").Filter("teacher_id", teacher_id).Update(orm.Params{
		"status": 0,
	})
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("teachers_show").Filter("teacher_id", teacher_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
		return nil
	}
}

/*
教师列表
*/
func OrganizationalTeacher(id int, ty int, person int, class_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	var teacher []orm.Params
	var condition []interface{}
	where := "1=1 "
	paginatorMap = make(map[string]interface{})
	if ty == 1 {
		if person == 1 {
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("t.name", "t.avatar", "t.teacher_id", "t.number", "t.phone").
				From("teacher as t").Where("kindergarten_id = ?").And("isnull(deleted_at)").And("status = 0").String()
			_, err := o.Raw(sql, id).Values(&v)
			if class_id > 0 {
				where += " AND om.organizational_id = ?"
				condition = append(condition, class_id)
			}
			qb, _ = orm.NewQueryBuilder("mysql")
			sql = qb.Select("t.name", "t.avatar", "t.teacher_id", "t.number", "t.phone", "om.id").From("teacher as t").LeftJoin("organizational_member as om").
				On("t.teacher_id = om.member_id").Where(where).And("t.kindergarten_id = ?").And("om.type = 0").And("is_principal = 0").String()
			_, err = o.Raw(sql, condition, id).Values(&teacher)
			for _, val := range teacher {
				v = append(v, val)
			}
			if err == nil {
				paginatorMap["data"] = v
				return paginatorMap, nil
			}
		} else {
			qb, _ := orm.NewQueryBuilder("mysql")
			sql := qb.Select("t.name", "t.avatar", "t.teacher_id", "t.number", "t.phone").
				From("teacher as t").Where("kindergarten_id = ?").And("isnull(deleted_at)").And("status = 0").String()
			num, err := o.Raw(sql, id).Values(&v)
			if err == nil && num > 0 {
				paginatorMap["data"] = v
				return paginatorMap, nil
			}
		}
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("t.name", "t.avatar", "t.teacher_id", "t.number", "t.phone").
			From("teacher as t").Where("kindergarten_id = ?").And("isnull(deleted_at)").String()
		num, err := o.Raw(sql, id).Values(&v)
		if err == nil && num > 0 {
			paginatorMap["data"] = v
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
筛选教师
*/
func FilterTeacher(class_type int, kindergarten_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.teacher_id", "o.id as class_id", "o.name as class_name", "o.class_type", "t.name", "t.avatar").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where("o.class_type = ? and o.level = 3 and o.type =2 and om.is_principal = 0").And("o.kindergarten_id = ?").And("om.type = 0").And("isnull(t.deleted_at)").String()
	_, err = o.Raw(sql, class_type, kindergarten_id).Values(&class)
	if err != nil {
		return nil, err
	}
	data := make(map[string][]interface{})
	for _, val := range class {
		if val["class_type"].(string) == "3" {
			if strc, ok := val["class_name"].(string); ok {
				data["大班"+strc] = append(data["大班"+strc], val)
			}
		} else if val["class_type"].(string) == "2" {
			if strc, ok := val["class_name"].(string); ok {
				data["中班"+strc] = append(data["中班"+strc], val)
			}
		} else if val["class_type"].(string) == "1" {
			if strc, ok := val["class_name"].(string); ok {
				data["小班"+strc] = append(data["小班"+strc], val)
			}
		}
	}
	ml = make(map[string]interface{})
	ml["data"] = data
	return ml, nil
}

/*
用户id获取教师id
*/
func GetUt(user_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var teacher []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.teacher_id").From("teacher as t").Where("t.user_id = ?").And("isnull(t.deleted_at)").String()
	_, err = o.Raw(sql, user_id).Values(&teacher)
	if err == nil {
		ml = make(map[string]interface{})
		ml["data"] = teacher
		return ml, nil
	}
	return nil, err
}

/*
筛选教师
*/
func FilterGetTeacher(class_type int, kindergarten_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.user_id", "t.name").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where("o.class_type = ? and o.level = 3 and o.type =2").And("o.kindergarten_id = ?").And("om.type = 0").And("isnull(t.deleted_at)").String()
	_, err = o.Raw(sql, class_type, kindergarten_id).Values(&class)
	if err != nil {
		return nil, err
	}

	ml = make(map[string]interface{})
	ml["data"] = class
	return ml, nil
}

/*
删除用户关联表
*/
func ResetUserId(user_id int) error {
	o := orm.NewOrm()
	o.Begin()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("teacher as t").Where("user_id = ?").And("isnull(t.deleted_at)").String()
	_, err := o.Raw(sql, user_id).Values(&v)
	if v == nil {
		return nil
	} else {
		_, err = o.QueryTable("organizational_member").Filter("member_id", v[0]["teacher_id"]).Filter("type", 0).Delete()
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.QueryTable("teacher").Filter("teacher_id", v[0]["teacher_id"]).Delete()
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.QueryTable("teachers_show").Filter("teacher_id", v[0]["teacher_id"]).Delete()
		if err != nil {
			o.Rollback()
			return err
		} else {
			o.Commit()
			return nil
		}
	}
}

/*
教师编辑
*/
func UpTeacher(teacher string, ty int, member_ids string, is_principal int, organizational_id int, id int, class int) error {
	o := orm.NewOrm()
	var t *Teacher
	json.Unmarshal([]byte(teacher), &t)
	v := Teacher{Id: id}
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	if t.Post == "" {
		t.Post = "普通教师"
	}
	if err := o.Read(&v); err == nil {
		t.Id = v.Id
		_, err = o.Update(t)
		if err != nil {
			return err
		} else {
			_, err = o.QueryTable("organizational_member").Filter("member_id", id).Filter("is_principal", 0).Delete()
			if err != nil {
				return err
			}
			if class == 1 {
				_, err := o.QueryTable("teacher").Filter("teacher_id", id).Update(orm.Params{
					"status": 0,
				})
				err = User.UpdateUK(t.UserId, 1, 0)
				if err != nil {
					return err
				}
			} else {
				var v []orm.Params
				s := strings.Split(member_ids, ",")
				if ty == 0 {
					qb, _ := orm.NewQueryBuilder("mysql")
					sql := qb.Select("o.*").From("organizational as o").Where("o.id = ?").String()
					_, err := o.Raw(sql, organizational_id).Values(&v)
					if v == nil {
						err = errors.New("没有该班级")
						return err
					}
					//组织架构为园长不能添加
					if v[0]["type"] == "1" && v[0]["is_fixed"] == "1" {
						err = errors.New("不能添加")
						return err
					} else {
						for _, value := range s {
							if value == "" {
								break
							}
							sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
							_, err = o.Raw(sql, organizational_id, ty, value, is_principal).Exec()
							if err == nil {
								if v[0]["type"] == "2" && v[0]["level"] == "3" {
									//0教师
									if ty == 0 {
										if v[0]["class_type"] == "3" {
											class_info := "大班" + v[0]["name"].(string) + ""
											_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
												"status": 1, "class_info": class_info,
											})
											if err != nil {
												o.Rollback()
												return err
											}
										} else if v[0]["class_type"] == "2" {
											class_info := "中班" + v[0]["name"].(string) + ""
											_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
												"status": 1, "class_info": class_info,
											})
											if err != nil {
												o.Rollback()
												return err
											}
										} else {
											class_info := "小班" + v[0]["name"].(string) + ""
											_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
												"status": 1, "class_info": class_info,
											})
											if err != nil {
												o.Rollback()
												return err
											}
										}
									}
								}
								if err == nil {
									User.UpdateUK(t.UserId, 1, 0)
									User.UpdateUkByUserId(t.UserId, 5)
									o.Commit()
									return err
								}
							}
						}
					}
				} else {
					qb, _ := orm.NewQueryBuilder("mysql")
					sql := qb.Select("o.*").From("organizational as o").Where("o.level = 2 and o.type = 3 and o.kindergarten_id = ?").String()
					_, err := o.Raw(sql, t.KindergartenId).Values(&v)
					for _, value := range s {
						if value == "" {
							break
						}
						sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
						_, err = o.Raw(sql, v[0]["id"], ty, value, is_principal).Exec()
						if err == nil {
							_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
								"status": 0, "class_info": "保健医",
							})
							if err != nil {
								o.Rollback()
								return err
							}
							User.UpdateUK(t.UserId, 1, 0)
							User.UpdateUkByUserId(t.UserId, 7)
							o.Commit()
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

/*
教师在次分班
*/
func Duties(ty int, member_ids string, is_principal int, organizational_id int, user_id int, kindergarten_id int, teacher_id int, teacher string) error {
	o := orm.NewOrm()
	var v []orm.Params
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	s := strings.Split(member_ids, ",")
	_, err := o.QueryTable("organizational_member").Filter("member_id", teacher_id).Filter("is_principal", 0).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	if ty == 0 {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("o.*").From("organizational as o").Where("o.id = ?").String()
		_, err := o.Raw(sql, organizational_id).Values(&v)
		if v == nil {
			err = errors.New("没有该班级")
			return err
		}
		//组织架构为园长不能添加
		if v[0]["type"] == "1" && v[0]["is_fixed"] == "1" {
			err = errors.New("不能添加")
			return err
		} else {
			for _, value := range s {
				if value == "" {
					break
				}
				sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
				_, err = o.Raw(sql, organizational_id, ty, value, is_principal).Exec()
				if err == nil {
					if v[0]["type"] == "2" && v[0]["level"] == "3" {
						//0教师
						if ty == 0 {
							if v[0]["class_type"] == "3" {
								class_info := "大班" + v[0]["name"].(string) + ""
								_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
									"status": 1, "class_info": class_info,
								})
								if err != nil {
									o.Rollback()
									return err
								}
							} else if v[0]["class_type"] == "2" {
								class_info := "中班" + v[0]["name"].(string) + ""
								_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
									"status": 1, "class_info": class_info,
								})
								if err != nil {
									o.Rollback()
									return err
								}
							} else {
								class_info := "小班" + v[0]["name"].(string) + ""
								_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
									"status": 1, "class_info": class_info,
								})
								if err != nil {
									o.Rollback()
									return err
								}
							}
						}
					}
				}
				if err == nil {
					User.UpdateUkByUserId(user_id, 5)
					o.Commit()
					return err
				}
			}
		}
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("o.*").From("organizational as o").Where("o.level = 2 and o.type = 3 and o.kindergarten_id = ?").String()
		_, err := o.Raw(sql, kindergarten_id).Values(&v)
		for _, value := range s {
			if value == "" {
				break
			}
			sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
			_, err = o.Raw(sql, v[0]["id"], ty, value, is_principal).Exec()
			if err == nil {
				_, err := o.QueryTable("teacher").Filter("teacher_id", value).Update(orm.Params{
					"status": 0, "class_info": "保健医",
				})
				if err != nil {
					o.Rollback()
					return err
				}
				User.UpdateUkByUserId(user_id, 7)
				o.Commit()
				return err
			}
		}
	}
	return nil
}

/*
删除用户关联表
*/
func InviteReset(user_id int) (err error) {
	o := orm.NewOrm()
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	err = User.DeleteUkByUserId(user_id)
	_, err = o.QueryTable("teacher").Filter("user_id", user_id).Delete()
	return err
}

/*
取消邀请
*/
func TeacherResetInvite(ty int, user_id int, teacher string, kindergarten_id int) (err error) {
	var User *UserService
	var Notice *NoticeService
	var t []inviteTeacher
	json.Unmarshal([]byte(teacher), &t)
	//rpc服务
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	client = rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	if ty == 1 {
		User.UpdateUK(user_id, 4, 1)
	} else {
		//获取用户关联表
		for _, value := range t {
			if err == nil {
				err = User.UpdateUser(user_id, value.Name)
				o := orm.NewOrm()
				var teacher []orm.Params
				qb, _ := orm.NewQueryBuilder("mysql")
				sql := qb.Select("t.name", "k.name as kinder_name").From("teacher as t").LeftJoin("organizational_member as om").
					On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
					On("om.organizational_id = o.id").LeftJoin("kindergarten as k").
					On("o.kindergarten_id = k.kindergarten_id").Where("o.kindergarten_id = ? and t.status = 1 and om.is_principal = 1 and o.level = 2 and o.type = 1 and o.is_fixed = 1 and om.type = 0 and isnull(t.deleted_at)").String()
				_, err = o.Raw(sql, value.KindergartenId).Values(&teacher)
				if teacher == nil {
					err = errors.New("幼儿园不存在或未设置园长")
					return err
				}
				if err == nil {
					data := make(map[string]interface{})
					data["user_id"] = user_id
					data["type"] = 1
					data["baby_id"] = 0
					data["title"] = value.Name + "老师"
					data["kindergarten_id"] = teacher[0]["kindergarten_id"]
					data["content"] = teacher[0]["name"].(string) + "园长邀请你加入" + teacher[0]["kinder_name"].(string)
					data["notice_type"] = 6
					data["choice_type"] = 3
					result, _ := json.Marshal(data)
					err = Notice.InviteSystem(string(result))
					if err == nil {
						err := User.UpdateUK(user_id, 0, 1)
						return err
					}
				}
			}
		}
	}
	return err
}
