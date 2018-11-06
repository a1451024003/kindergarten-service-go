package models

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"

	"github.com/hprose/hprose-golang/rpc"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id                    int       `json:"student_id" orm:"column(student_id);auto"`
	Name                  string    `json:"name" orm:"column(name);size(20)" description:"姓名"`
	Age                   int       `json:"age" orm:"column(age)" description:"年龄"`
	Sex                   int       `json:"sex" orm:"column(sex)" description:"性别 0男 1女"`
	NativePlace           string    `json:"native_place" orm:"column(native_place);size(20)" description:"籍贯"`
	NationOrReligion      string    `json:"nation_or_religion" orm:"column(nation_or_religion);size(20)" description:"民族或宗教"`
	Number                string    `json:"number" orm:"column(number);size(11)" description:"学号"`
	ClassInfo             string    `json:"class_info" orm:"column(class_info);size(20)" description:"所在班级"`
	Address               string    `json:"address" orm:"column(address);size(50)" description:"住址"`
	Avatar                string    `json:"avatar" orm:"column(avatar);size(150)" description:"头像"`
	Status                int8      `json:"status" orm:"column(status)" description:"状态 0未分班 1已分班 2离园"`
	BabyId                int       `json:"baby_id" orm:"column(baby_id)" description:"宝宝ID"`
	Birthday              string    `json:"birthday" orm:"column(birthday)" description:"生日"`
	KindergartenId        int       `json:"kindergarten_id" orm:"column(kindergarten_id)" description:"幼儿园ID"`
	Phone                 string    `json:"phone" orm:"column(phone);size(11)" description:"手机号"`
	HealthStatus          string    `json:"health_status" orm:"column(health_status);size(150)" description:"健康状况，多个以逗号隔开"`
	Hobby                 string    `json:"hobby" orm:"column(hobby);size(150)" description:"兴趣爱好，多个以逗号隔开"`
	IsMuslim              int       `json:"is_muslim" orm:"column(is_muslim)"`
	BabyName              string    `json:"baby_name" orm:"column(baby_name)"`
	IdNumber              string    `json:"id_number" orm:"column(id_number)"`
	Relation              string    `json:"relation" orm:"column(relation)"`
	EmergencyContact      string    `json:"emergency_contact" orm:"column(emergency_contact)"`
	EmergencyContactPhone string    `json:"emergency_contact_phone" orm:"column(emergency_contact_phone)"`
	CreatedAt             time.Time `json:"Created_at" orm:"auto_now_add"`
	UpdatedAt             time.Time `json:"updated_at" orm:"auto_now"`
	DeletedAt             time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime);null"`
}

type inviteStudent struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Avatar         string `json:"avatar"`
	BabyId         int    `json:"baby_id"`
	Birthday       string `json:"birthday"`
	KindergartenId int    `json:"kindergarten_id"`
}

func (t *Student) TableName() string {
	return "student"
}

func init() {
	orm.RegisterModel(new(Student))
}

/*
学生列表
*/
func GetStudent(id int, status int, search string, page int, prepage int, date string) (ml map[string]interface{}, err error) {
	var condition []interface{}
	where := "1=1 "
	if id == 0 {
		where += " AND s.kindergarten_id = ?"
	} else {
		where += " AND s.kindergarten_id = ?"
		condition = append(condition, id)
	}
	if status != -1 {
		where += " AND s.status = ?"
		condition = append(condition, status)
	}
	if search != "" {
		where += " AND s.name like ?"
		condition = append(condition, "%"+search+"%")
	}
	if date != "" {
		where += " AND s.updated_at like ?"
		condition = append(condition, "%"+date+"%")
	}
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	sql := qb.Select("count(*)").From("student as s").Where(where).And("isnull(deleted_at)").String()
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
		sql := qb.Select("s.name", "s.avatar", "s.student_id", "s.number", "s.phone").
			From("student as s").Where(where).And("isnull(s.deleted_at) ").Limit(prepage).Offset(limit).String()
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
学生班级列表
*/
func GetStudentClass(id int, class_type int, page int, prepage int) (ml map[string]interface{}, err error) {
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
	sql := qb.Select("count(*)").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("s.status = 1 and om.is_principal = 0 and om.type = 1").And("isnull(s.deleted_at)").String()
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
		sql := qb.Select("s.name", "s.avatar", "s.student_id", "s.number", "s.phone", "o.name as class").From("student as s").LeftJoin("organizational_member as om").
			On("s.student_id = om.member_id").LeftJoin("organizational as o").
			On("om.organizational_id = o.id").Where(where).And("s.status = 1 and om.is_principal = 0 and om.type = 1").And("isnull(s.deleted_at)").Limit(prepage).Offset(limit).String()
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
删除学生
*/
func DeleteStudent(id int, status int, ty int, class_type int) error {
	o := orm.NewOrm()
	var stu []orm.Params
	v := Student{Id: id}
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.baby_id").
		From("student as s").Where("s.student_id = ? and isnull(s.deleted_at)").String()
	_, err := o.Raw(sql, id).Values(&stu)
	if err = o.Read(&v); err == nil {
		if status == 0 {
			v.Status = 2
			o.QueryTable("baby_kindergarten").Filter("baby_id", stu[0]["baby_id"]).Update(orm.Params{
				"status": 1,
			})
		} else if status == 2 {
			v.DeletedAt, _ = time.ParseInLocation(timeLayout, timenow, loc)
		}

		if class_type == 3 || class_type == 2 || class_type == 1 {
			v.Status = 0
		}
		if _, err = o.Update(&v); err == nil {
			_, err = o.QueryTable("organizational_member").Filter("member_id", id).Filter("type", 1).Delete()
			if err == nil {
				_, err = o.QueryTable("exceptional_child").Filter("student_id", id).Delete()
				if err != nil {
					return err
				} else {
					return nil
				}
			}
		}
	}
	return nil
}

/*
学生详情
*/
func GetStudentInfo(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var kinships []Kinship
	//学生信息
	student := Student{Id: id}
	err = o.Read(&student)
	//亲属信息
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("kinship").Where("student_id = ?").String()
	_, err = o.Raw(sql, id).QueryRows(&kinships)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["student"] = student
		paginatorMap["kinship"] = kinships
		return paginatorMap, nil
	}
	err = errors.New("获取失败")
	return nil, err
}

/*
学生详情
*/
func GetBabytInfo(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var kinships []Kinship
	//学生信息
	student := Student{Id: id}
	err = o.Read(&student)
	//亲属信息
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("kinship").Where("baby_id = ?").String()
	num, err := o.Raw(sql, id).QueryRows(&kinships)
	if err == nil && num > 0 {
		paginatorMap := make(map[string]interface{})
		paginatorMap["student"] = student
		paginatorMap["kinship"] = kinships
		return paginatorMap, nil
	}
	return nil, err
}

/*
app学生详情
*/
func GetStudentOne(id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var kinships []orm.Params
	//学生信息
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("name", "sex", "age", "id_number", "phone", "address", "relation", "emergency_contact", "emergency_contact_phone", "kindergarten_id", "baby_id", "student_id").From("student").Where("baby_id = ? and isnull(deleted_at) and status != ?").String()
	num, err := o.Raw(sql, id, 2).Values(&student)
	if student != nil {
		//亲属信息
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = qb.Select("kinship_id", "student_id", "baby_id", "type", "relation", "name", "address", "unit_name", "contact_information").From("kinship").Where("student_id = ?").String()
		_, err = o.Raw(sql, student[0]["student_id"]).Values(&kinships)
	}
	if err == nil && num > 0 {
		paginatorMap := make(map[string]interface{})
		paginatorMap["student"] = student
		paginatorMap["kinship"] = kinships
		return paginatorMap, nil
	}
	return nil, err
}

/*
学生编辑
*/
func UpdateStudent(id int, student string, kinship string) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	o.Begin()
	v := Student{Id: id}
	//编辑学生信息
	var s *Student
	json.Unmarshal([]byte(student), &s)
	if err := o.Read(&v); err == nil {
		s.Id = v.Id
		if _, err = o.Update(s); err == nil {
			//写入亲属表
			_, err = o.QueryTable("kinship").Filter("student_id", id).Delete()
			if err == nil {
				if kinship != "[]" {
					var k []Kinship
					json.Unmarshal([]byte(kinship), &k)
					o.InsertMulti(100, &k)
				}
				o.Commit()
				return nil, err
			}
		} else {
			o.Rollback()
			return nil, err
		}
	}
	return nil, err
}

/*
前台编辑学生
*/
func SaveStudent(id int, student string, kinship string, ty int, member_ids string, is_principal int, organizational_id int, class int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	v := Student{Id: id}
	//编辑学生信息
	var s *Student
	json.Unmarshal([]byte(student), &s)
	if err := o.Read(&v); err == nil {
		s.Id = v.Id
		if _, err = o.Update(s); err == nil {
			//写入亲属表
			_, err = o.QueryTable("kinship").Filter("student_id", id).Delete()
			if err == nil {
				var k []Kinship
				json.Unmarshal([]byte(kinship), &k)
				_, err = o.InsertMulti(100, &k)
				if err == nil {
					if class == 1 {
						_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", s.BabyId).Update(orm.Params{
							"invite_status": 2, "actived": 0,
						})
						_, err = o.QueryTable("student").Filter("baby_id", s.BabyId).Update(orm.Params{
							"status": 0,
						})
						if err == nil {
							o.Commit()
							return nil
						}
					} else {
						var v []orm.Params
						sl := strings.Split(member_ids, ",")
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
							for _, value := range sl {
								if value == "" {
									break
								}
								sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
								_, err = o.Raw(sql, organizational_id, ty, value, is_principal).Exec()
								if err == nil {
									if v[0]["type"] == "2" && v[0]["level"] == "3" {
										//0教师 1学生
										if ty == 1 {
											if v[0]["class_type"] == "3" {
												class_info := "大班" + v[0]["name"].(string) + ""
												_, err := o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
													"status": 1, "class_info": class_info,
												})
												if err != nil {
													o.Rollback()
													return err
												}
											} else if v[0]["class_type"] == "2" {
												class_info := "中班" + v[0]["name"].(string) + ""
												_, err := o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
													"status": 1, "class_info": class_info,
												})
												if err != nil {
													o.Rollback()
													return err
												}
											} else {
												class_info := "小班" + v[0]["name"].(string) + ""
												_, err := o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
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
							}
							if err == nil {
								_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", s.BabyId).Update(orm.Params{
									"actived": 0,
								})
								_, err = o.QueryTable("student").Filter("baby_id", s.BabyId).Update(orm.Params{
									"status": 1,
								})
								if err == nil {
									o.Commit()
									return nil
								}
							}
						}
					}
				}
			}
		} else {
			o.Rollback()
			return err
		}
	}
	return nil
}

/*
后台认证学生
*/
func UpStudent(id int, student string, kinship string, ty int, member_ids string, is_principal int, organizational_id int, class int) (err error) {
	o := orm.NewOrm()
	v := Student{Id: id}
	//编辑学生信息
	var s *Student
	json.Unmarshal([]byte(student), &s)
	if err := o.Read(&v); err == nil {
		s.Id = v.Id
		if _, err = o.Update(s); err == nil {
			//写入亲属表
			_, err = o.QueryTable("kinship").Filter("student_id", id).Delete()

			if err == nil {
				if kinship != "[]" {
					var k []Kinship
					json.Unmarshal([]byte(kinship), &k)
					_, err = o.InsertMulti(100, &k)
				}
				if err == nil {
					if class == 1 { //未分班
						_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", s.BabyId).Update(orm.Params{
							"invite_status": 1, "actived": 0,
						})
						_, err = o.QueryTable("student").Filter("baby_id", s.BabyId).Update(orm.Params{
							"status": 0,
						})
					} else { //分班
						var v []orm.Params
						sl := strings.Split(member_ids, ",")
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
							for _, value := range sl {
								if value == "" {
									break
								}
								sql := "insert into organizational_member set organizational_id = ?,type = ?,member_id = ?,is_principal = ?"
								_, err = o.Raw(sql, organizational_id, ty, value, is_principal).Exec()
								if err == nil {
									if v[0]["type"] == "2" && v[0]["level"] == "3" {
										//0教师 1学生
										if ty == 1 {
											if v[0]["class_type"] == "3" {
												class_info := "大班" + v[0]["name"].(string) + ""
												_, err = o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
													"status": 1, "class_info": class_info,
												})
											} else if v[0]["class_type"] == "2" {
												class_info := "中班" + v[0]["name"].(string) + ""
												_, err = o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
													"status": 1, "class_info": class_info,
												})
											} else {
												class_info := "小班" + v[0]["name"].(string) + ""
												_, err = o.QueryTable("student").Filter("student_id", value).Update(orm.Params{
													"status": 1, "class_info": class_info,
												})
											}
										}
									}
								}
							}
							if err == nil {
								_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", s.BabyId).Update(orm.Params{
									"actived": 0, "invite_status": 1,
								})
								_, err = o.QueryTable("student").Filter("baby_id", s.BabyId).Update(orm.Params{
									"status": 1,
								})
							}
						}
					}
				}
			}
		} else {
			return err
		}
	}
	return nil
}

/*
学生-录入信息
*/
func AddStudent(student string, kinship string) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	o.Begin()
	paginatorMap = make(map[string]interface{})
	//写入学生表
	var s Student
	json.Unmarshal([]byte(student), &s)
	id, err := o.Insert(&s)
	if err != nil {
		o.Rollback()
		return nil, err
	}
	if kinship != "" {
		ids := strconv.FormatInt(id, 10)
		kid, _ := strconv.Atoi(ids)
		//写入亲属表
		var k []Kinship
		json.Unmarshal([]byte(kinship), &k)
		for key, _ := range k {
			k[key].StudentId = kid
		}
		id, err = o.InsertMulti(100, &k)
		if err != nil {
			o.Rollback()
			return nil, err
		}
	}
	_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", s.BabyId).Update(orm.Params{
		"actived": 0,
	})
	if err != nil {
		o.Rollback()
		return nil, err
	} else {
		o.Commit()
		return nil, nil
	}
}

/*
app学生-录入信息
*/
func AddAppStudent(sex int, age int, name string, phone string, address string, relation string, emergency_contact string, emergency_contact_phone string, baby_id int, kindergarten_id int, ty int, id_number string, notice_id int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	var Notice *NoticeService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	err = Notice.NoticeRead(notice_id)
	if err != nil {
		return err
	}
	if ty == 1 {
		//写入学生表
		sql := "insert into student set sex = ?,age = ?,name = ?,phone = ?,address = ?,baby_id = ?,kindergarten_id = ?,status = ?,emergency_contact = ?,emergency_contact_phone = ?,relation = ?,id_number = ?"
		res, err := o.Raw(sql, sex, age, name, phone, address, baby_id, kindergarten_id, 3, emergency_contact, emergency_contact_phone, relation, id_number).Exec()
		sId, _ := res.LastInsertId()
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", baby_id).Update(orm.Params{
			"actived": 1, "invite_status": 2,
		})

		_, err = o.QueryTable("student").Filter("student_id", sId).Update(orm.Params{
			"status": 3,
		})
		if err != nil {
			o.Rollback()
			return err
		} else {
			o.Commit()
			return nil
		}
	} else {
		_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", baby_id).Update(orm.Params{
			"actived": 1, "invite_status": 3,
		})
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
移除学生
*/
func RemoveStudent(class_id int, student_id int) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.QueryTable("organizational_member").Filter("organizational_id", class_id).Filter("member_id", student_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("exceptional_child").Filter("student_id", student_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.QueryTable("student").Filter("student_id", student_id).Update(orm.Params{
		"status": 0,
	})
	if err != nil {
		o.Rollback()
		return err
	} else {
		o.Commit()
		return nil
	}
	return err
}

/*
app-邀请学生
*/
func Invite(student string, kindergarten_id int) error {
	o := orm.NewOrm()
	var Notice *NoticeService
	var User *UserService
	var someError error
	var baby []orm.Params
	var s []inviteStudent
	json.Unmarshal([]byte(student), &s)
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation(timeLayout, timenow, loc)
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	client = rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	for _, v := range s {
		valid := validation.Validation{}
		valid.Required(student, "student").Message("学生信息不能为空")
		valid.Required(v.Name, "name").Message("幼儿姓名不能为空")
		valid.Required(v.Phone, "phone").Message("手机号不能为空")
		valid.Required(v.BabyId, "baby_id").Message("宝宝id不能为空")
		valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
		if valid.HasErrors() {
			err := errors.New(valid.Errors[0].Message)
			return err
		}
		t, _ := time.Parse("2006-01-02 15:04:05", v.Birthday)
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("*").From("baby_kindergarten").Where("baby_id = ?").And("status = 0").String()
		_, err := o.Raw(sql, v.BabyId).Values(&baby)
		if err == nil {
			if baby != nil {
				someError = errors.New("" + string(v.Name) + "已被邀请过")
				continue
			} else {
				sql = "insert into baby_kindergarten set kindergarten_id = ?,baby_id = ?,baby_name = ?,created_at = ?,birthday = ?,phone = ?,invite_status = ?,avatar = ?"
				_, err := o.Raw(sql, kindergarten_id, v.BabyId, v.Name, createTime, t, v.Phone, 0, beego.AppConfig.String("AVATAR")).Exec()
				if err == nil {
					var teacher []orm.Params
					qb, _ = orm.NewQueryBuilder("mysql")
					sql = qb.Select("t.name", "k.name as kinder_name", "k.kindergarten_id").From("teacher as t").LeftJoin("organizational_member as om").
						On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
						On("om.organizational_id = o.id").LeftJoin("kindergarten as k").
						On("o.kindergarten_id = k.kindergarten_id").Where("o.kindergarten_id = ? and t.status = 1 and om.is_principal = 1 and o.level = 2 and o.type = 1 and o.is_fixed = 1 and om.type = 0 and isnull(t.deleted_at)").String()
					_, err = o.Raw(sql, kindergarten_id).Values(&teacher)
					if teacher == nil {
						err = errors.New("幼儿园不存在或未设置园长")
						return err
					}
					if err == nil {
						userId, _ := User.GetOne(v.Phone)
						data := make(map[string]interface{})
						data["user_id"] = userId
						data["baby_id"] = v.BabyId
						data["title"] = v.Name + "小朋友"
						data["type"] = 2
						data["kindergarten_id"] = teacher[0]["kindergarten_id"]
						data["content"] = teacher[0]["name"].(string) + "园长邀请你加入" + teacher[0]["kinder_name"].(string)
						data["notice_type"] = 6
						data["choice_type"] = 3
						result, _ := json.Marshal(data)
						Notice.InviteSystem(string(result))
					}
				}
			}
		}
	}
	return someError
}

//邀请学生-admin
func Invites(student string, kindergarten_id int) error {
	o := orm.NewOrm()
	var User *UserService
	var Notice *NoticeService
	var someError error
	var baby []orm.Params
	var s []inviteStudent
	json.Unmarshal([]byte(student), &s)
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	loc, _ := time.LoadLocation("")
	timenow := time.Now().Format("2006-01-02 15:04:05")
	createTime, _ := time.ParseInLocation(timeLayout, timenow, loc)
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	client = rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	for _, v := range s {
		valid := validation.Validation{}
		valid.Required(student, "student").Message("学生信息不能为空")
		valid.Required(v.Name, "name").Message("幼儿姓名不能为空")
		valid.Required(v.Phone, "phone").Message("手机号不能为空")
		valid.Required(v.BabyId, "baby_id").Message("宝宝id不能为空")
		valid.Required(kindergarten_id, "kindergarten_id").Message("幼儿园id不能为空")
		if valid.HasErrors() {
			err := errors.New(valid.Errors[0].Message)
			return err
		}
		t, _ := time.Parse("2006-01-02 15:04:05", v.Birthday)
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("*").From("baby_kindergarten").Where("baby_id = ?").And("status = 0").String()
		_, err := o.Raw(sql, v.BabyId).Values(&baby)
		if err == nil {
			if baby != nil {
				someError = errors.New("" + string(v.Name) + "已被邀请过")
				continue
			} else {
				sql = "insert into baby_kindergarten set kindergarten_id = ?,baby_id = ?,baby_name = ?,created_at = ?,birthday = ?,phone = ?,invite_status = ?,avatar = ?"
				_, err := o.Raw(sql, kindergarten_id, v.BabyId, v.Name, createTime, t, v.Phone, 0, beego.AppConfig.String("AVATAR")).Exec()
				if err == nil {
					var teacher []orm.Params
					qb, _ = orm.NewQueryBuilder("mysql")
					sql = qb.Select("t.name", "k.name as kinder_name", "k.kindergarten_id").From("teacher as t").LeftJoin("organizational_member as om").
						On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
						On("om.organizational_id = o.id").LeftJoin("kindergarten as k").
						On("o.kindergarten_id = k.kindergarten_id").Where("o.kindergarten_id = ? and t.status = 1 and om.is_principal = 1 and o.level = 2 and o.type = 1 and o.is_fixed = 1 and om.type = 0 and isnull(t.deleted_at)").String()
					_, err = o.Raw(sql, kindergarten_id).Values(&teacher)
					if teacher == nil {
						err = errors.New("幼儿园不存在或未设置园长")
						return err
					}
					if err == nil {
						userId, _ := User.GetOne(v.Phone)
						data := make(map[string]interface{})
						data["user_id"] = userId
						data["baby_id"] = v.BabyId
						data["title"] = v.Name + "小朋友"
						data["type"] = 2
						data["kindergarten_id"] = teacher[0]["kindergarten_id"]
						data["content"] = teacher[0]["name"].(string) + "园长邀请你加入" + teacher[0]["kinder_name"].(string)
						data["notice_type"] = 6
						data["choice_type"] = 3
						result, _ := json.Marshal(data)
						Notice.InviteSystem(string(result))
					}
				}
			}
		}
	}
	return someError
}

/*
未激活baby
*/
func GetBabyInfo(kindergarten_id int, invite int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var baby []orm.Params
	var condition []interface{}
	where := "1=1 "
	if invite > -1 {
		where += " AND invite_status = ?"
		condition = append(condition, invite)
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("baby_kindergarten").Where(where).And("status = 0 and kindergarten_id = ? and actived = ?").String()
	_, err = o.Raw(sql, condition, kindergarten_id, 1).Values(&baby)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = baby
		return paginatorMap, nil
	}
	err = errors.New("获取失败")
	return nil, err
}

/*
学生名字获取班级
*/
func GetNameClass(name string, kindergarten_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	var condition []interface{}
	where := "1=1 "
	if name != "" {
		where += " AND s.name like ?"
		condition = append(condition, "%"+name+"%")
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "class_type", "s.name").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("om.type = 1").And("s.kindergarten_id = ?").String()
	_, err = o.Raw(sql, condition, kindergarten_id).Values(&class)
	if class == nil || class[0]["class_id"] == nil {
		err = errors.New("该同学未分班")
		return nil, err
	}
	for key, val := range class {
		if val["class_type"].(string) == "3" {
			class[key]["class"] = "大班" + val["class_name"].(string) + ""
		} else if val["class_type"].(string) == "2" {
			class[key]["class"] = "中班" + val["class_name"].(string) + ""
		} else {
			class[key]["class"] = "小班" + val["class_name"].(string) + ""
		}
	}
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = class
		return paginatorMap, nil
	}
	err = errors.New("获取失败")
	return nil, err
}

/*
班级下的学生
*/
func ClassStudent(kindergarten_id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	data := make(map[string][]interface{})
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "o.class_type", "s.name", "s.avatar", "s.phone as student_phone").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where("o.kindergarten_id = ?").And("om.type = 1").String()
	_, err = o.Raw(sql, kindergarten_id).Values(&class)
	if err != nil {
		return nil, err
	}
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
baby是否在幼儿园
*/
func BabyKinder(baby_id int) (paginatorMap interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("baby_kindergarten").Where("baby_id = ?").And("actived = 0").String()
	_, err = o.Raw(sql, baby_id).Values(&class)
	return class, err
}

/*
baby是否在幼儿园
*/
func BabyActived(baby_id int) (paginatorMap interface{}, err error) {
	o := orm.NewOrm()
	var class []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("baby_kindergarten").Where("baby_id = ?").And("actived = 1").String()
	_, err = o.Raw(sql, baby_id).Values(&class)
	return class, err
}

/*
app baby状态列表
*/
func BabyKinderList(page int, prepage int, kindergarten_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("count(*)").From("baby_kindergarten as bk").Where("bk.actived = 1 and bk.status = 0 and bk.kindergarten_id = ?").String()
	var total int64
	err = o.Raw(sql, kindergarten_id).QueryRow(&total)
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
		sql := qb.Select("bk.*").From("baby_kindergarten as bk").Where("bk.actived = 1 and bk.status = 0 and bk.kindergarten_id = ?").Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, kindergarten_id).Values(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total //总条数
			paginatorMap["data"] = v
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
取消邀请
*/
func ResetInvite(ty int, baby_id int, student string, kindergarten_id int) (err error) {
	o := orm.NewOrm()
	var User *UserService
	var Notice *NoticeService
	var someError error
	var s []inviteStudent
	json.Unmarshal([]byte(student), &s)
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	client = rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&Notice)
	if ty == 1 {
		_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", baby_id).Update(orm.Params{
			"invite_status": 4,
		})
	} else {
		for _, v := range s {
			_, err = o.QueryTable("baby_kindergarten").Filter("baby_id", baby_id).Update(orm.Params{
				"invite_status": 0,
			})
			if err == nil {
				var teacher []orm.Params
				qb, _ := orm.NewQueryBuilder("mysql")
				sql := qb.Select("t.name", "k.name as kinder_name", "k.kindergarten_id").From("teacher as t").LeftJoin("organizational_member as om").
					On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
					On("om.organizational_id = o.id").LeftJoin("kindergarten as k").
					On("o.kindergarten_id = k.kindergarten_id").Where("o.kindergarten_id = ? and t.status = 1 and om.is_principal = 1 and o.level = 2 and o.type = 1 and o.is_fixed = 1 and om.type = 0 and isnull(t.deleted_at)").String()
				_, err = o.Raw(sql, kindergarten_id).Values(&teacher)
				if teacher == nil {
					err = errors.New("幼儿园不存在或未设置园长")
					return err
				}
				if err == nil {
					data := make(map[string]interface{})
					userId, _ := User.GetOne(v.Phone)
					data["user_id"] = userId
					data["baby_id"] = baby_id
					data["title"] = v.Name + "小朋友"
					data["type"] = 2
					data["kindergarten_id"] = teacher[0]["kindergarten_id"]
					data["content"] = teacher[0]["name"].(string) + "园长邀请你加入" + teacher[0]["kinder_name"].(string)
					data["notice_type"] = 6
					data["choice_type"] = 3
					result, _ := json.Marshal(data)
					Notice.InviteSystem(string(result))
				}
			}
		}
	}
	return someError
}

/*
班级成员
*/
func OragnizationalStudent(class_id int, page int, prepage int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var condition []interface{}
	where := "1=1 "
	if class_id > 0 {
		where += " AND o.id = ?"
		condition = append(condition, class_id)
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.*", "o.name as class_name", "o.class_type").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("om.type = 1").And("s.status = 1 and om.is_principal = 0").And("isnull(s.deleted_at)").String()
	_, err = o.Raw(sql, condition).Values(&student)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = student
		return paginatorMap, nil
	}
	return nil, err
}

/*
删除亲属信息
*/
func DeleteKinship(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("kinship").Filter("kinship_id", id).Delete()
	return err
}
