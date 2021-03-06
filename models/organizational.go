package models

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Organizational struct {
	Id             int       `json:"id" orm:"column(id);auto"`
	KindergartenId int       `json:"kindergarten_id" orm:"column(kindergarten_id)" description:"幼儿园id"`
	ParentId       int       `json:"parent_id" orm:"column(parent_id)" description:"父级id"`
	Name           string    `json:"name" orm:"column(name);size(20)" description:"组织架构名字"`
	IsFixed        int8      `json:"is_fixed" orm:"column(is_fixed)" description:"是否固定的：0不是，1是"`
	Level          int8      `json:"level" orm:"column(level)" description:"等级"`
	ParentIds      string    `json:"parent_ids" orm:"column(parent_ids);size(50)" description:"父级所有id"`
	Type           int8      `json:"type" orm:"column(type)" description:"类型：0普通，1管理层，2年级组"`
	ClassType      int8      `json:"class_type" orm:"column(class_type)" description:"班级类型：1小班，2中班，3大班"`
	CreatedAt      time.Time `json:"created_at" orm:"column(created_at);type(datetime)" description:"添加时间"`
	UpdatedAt      time.Time `json:"updated_at" orm:"column(updated_at);type(datetime)" description:"修改时间"`
}

type OrganizationalTree struct {
	Id             int                  `json:"id" orm:"column(id);auto"`
	KindergartenId int                  `json:"kindergarten_id" orm:"column(kindergarten_id)" description:"幼儿园id"`
	ParentId       int                  `json:"parent_id" orm:"column(parent_id)" description:"父级id"`
	Name           string               `json:"name" orm:"column(name);size(20)" description:"组织架构名字"`
	IsFixed        int8                 `json:"is_fixed" orm:"column(is_fixed)" description:"是否固定的：0不是，1是"`
	Level          int8                 `json:"level" orm:"column(level)" description:"等级"`
	ParentIds      string               `json:"parent_ids" orm:"column(parent_ids);size(50)" description:"父级所有id"`
	Type           int8                 `json:"type" orm:"column(type)" description:"类型：0普通，1管理层，2年级组"`
	ClassType      int8                 `json:"class_type" orm:"column(class_type)" description:"班级类型：1小班，2中班，3大班"`
	CreatedAt      time.Time            `json:"created_at" orm:"column(created_at);type(datetime)" description:"添加时间"`
	UpdatedAt      time.Time            `json:"updated_at" orm:"column(updated_at);type(datetime)" description:"修改时间"`
	Children       []OrganizationalTree `json:"children" orm:"null"`
}

type ClassStr struct {
	Amount    int `json:"amount"`
	ClassType int `json:"class_type"`
}

func (t *Organizational) TableName() string {
	return "organizational"
}

func init() {
	orm.RegisterModel(new(Organizational))
}

/*
班级搜索
*/
func GetClassAll(kindergarten_id int, class_type int, page int, prepage int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	var condition []interface{}
	where := "1=1 "
	if kindergarten_id == 0 {
		where += " AND kindergarten_id = ?"
	} else {
		where += " AND kindergarten_id = ?"
		condition = append(condition, kindergarten_id)
	}
	if class_type == 4 {
		where += " AND is_fixed = ?"
		condition = append(condition, 0)
	} else if class_type == 1 || class_type == 2 || class_type == 3 {
		where += " AND class_type = ?"
		condition = append(condition, class_type)
	}
	// 构建查询对象
	sql := qb.Select("count(*)").From("organizational").Where(where).And("type = 2").And("level = 3").String()
	var total int64
	err = o.Raw(sql, condition).QueryRow(&total)
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
		sql := qb.Select("*").From("organizational").Where(where).And("type = 2").And("level = 3").Limit(prepage).Offset(limit).String()
		_, err := o.Raw(sql, condition).Values(&v)
		var re []map[string]interface{}
		var res map[string]interface{}
		data := Organizational{Id: 0, Name: "全部", ClassType: 4}
		organ, _ := json.Marshal(data)
		json.Unmarshal(organ, &res)
		re = append(re, res)
		for _, val := range v {
			re = append(re, val)
		}
		if err == nil {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = total         //总条数
			paginatorMap["data"] = re             //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err
}

/*
后台班级搜索
*/
func GetAdminClassAll(kindergarten_id int, class_type int, page int, prepage int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	var condition []interface{}
	where := "1=1 "
	if kindergarten_id == 0 {
		where += " AND kindergarten_id = ?"
	} else {
		where += " AND kindergarten_id = ?"
		condition = append(condition, kindergarten_id)
	}
	if class_type != 0 {
		where += " AND class_type = ?"
		condition = append(condition, class_type)
	}
	// 构建查询对象
	sql := qb.Select("count(*)").From("organizational").Where(where).And("type = 2").And("level = 3").String()
	var total int64
	err = o.Raw(sql, condition).QueryRow(&total)
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
		sql := qb.Select("*").From("organizational").Where(where).And("type = 2").And("level = 3").OrderBy("class_type").Desc().Limit(prepage).Offset(limit).String()
		num, err := o.Raw(sql, condition).Values(&v)
		if num != 0 {
			for _, val := range v {
				if val["class_type"].(string) == "3" {
					val["class_name"] = "大班" + val["name"].(string)
				} else if val["class_type"].(string) == "2" {
					val["class_name"] = "中班" + val["name"].(string)
				} else if val["class_type"].(string) == "1" {
					val["class_name"] = "小班" + val["name"].(string)
				}
			}
		}
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
班级成员
*/
func ClassMember(kindergarten_id int, class_type int, class_id int, page int, prepage int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var student []orm.Params
	var teacher []orm.Params
	var condition []interface{}
	where := "1=1 "
	if kindergarten_id == 0 {
		where += " AND o.kindergarten_id = ?"
	} else {
		where += " AND o.kindergarten_id = ?"
		condition = append(condition, kindergarten_id)
	}
	if class_type > 0 {
		where += " AND o.class_type = ?"
		condition = append(condition, class_type)
	}
	if class_id > 0 {
		where += " AND o.id = ?"
		condition = append(condition, class_id)
	}
	where += " AND o.type = ?"
	condition = append(condition, 2)
	where += " AND o.level = ?"
	condition = append(condition, 3)

	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.student_id", "s.name", "s.avatar", "s.number", "s.phone",
		"o.name as class_name", "o.class_type", "om.id").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("om.type = 1").And("s.status = 1 and is_principal = 0").And("isnull(s.deleted_at)").String()
	_, err = o.Raw(sql, condition).Values(&student)

	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("t.teacher_id", "t.name", "t.avatar", "t.number", "t.phone",
		"o.name as class_name", "om.id").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where(where).And("t.status = 1 and is_principal = 0 ").And("om.type = 0").And("isnull(t.deleted_at)").String()
	_, err = o.Raw(sql, condition).Values(&teacher)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["student"] = student
		paginatorMap["teacher"] = teacher
		return paginatorMap, nil
	}
	err = errors.New("获取失败")
	return nil, err
}

/*
删除班级
*/
func Destroy(class_id int) error {
	o := orm.NewOrm()
	o.Begin()
	var t []orm.Params
	var s []orm.Params
	var v Organizational
	o.QueryTable("organizational").Filter("id", class_id).All(&v)
	if v.IsFixed == 1 {
		err := errors.New("不能删除")
		return err
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("member_id").From("organizational_member").Where("organizational_id = ?").And("type = 0").String()
	_, err := o.Raw(sql, class_id).Values(&t)
	if err == nil {
		//修改teacher
		for key, _ := range t {
			_, err = o.QueryTable("teacher").Filter("teacher_id", t[key]["member_id"]).Update(orm.Params{
				"status": 0, "class_info": "",
			})
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}

	//修改学生
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("member_id").From("organizational_member").Where("organizational_id = ?").And("type = 1").String()
	_, err = o.Raw(sql, class_id).Values(&s)
	if err == nil {
		for key, _ := range s {
			_, err = o.QueryTable("student").Filter("student_id", s[key]["member_id"]).Update(orm.Params{
				"status": 0, "class_info": "",
			})
			if err == nil {
				_, err = o.QueryTable("exceptional_child").Filter("student_id", s[key]["member_id"]).Delete()
				if err != nil {
					o.Rollback()
					return err
				}
			}
		}
	}

	//删除班级
	_, err = o.QueryTable("organizational").Filter("id", class_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	//删除班级成员
	_, err = o.QueryTable("organizational_member").Filter("organizational_id", class_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	//删除班级圈
	_, err = o.QueryTable("group_view").Filter("class_id", class_id).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	if err == nil {
		o.Commit()
		return nil
	}
	err = errors.New("删除失败")
	return err
}

/*
创建班级
*/
func StoreClass(class_type int, kindergarten_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	var or []orm.Params
	var max_name []orm.Params
	paginatorMap = make(map[string]interface{})
	var condition []interface{}
	where := "1=1 "
	if class_type > 0 {
		where += " AND class_type = ?"
		condition = append(condition, class_type)
	}
	if kindergarten_id > 0 {
		where += " AND kindergarten_id = ?"
		condition = append(condition, kindergarten_id)
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.*").From("organizational as o").Where(where).And("type = 2").And("level = 2").String()
	num, err := o.Raw(sql, condition).Values(&or)
	//查出最大班级
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("max(CONVERT(o.name,SIGNED)) as m").From("organizational as o").Where(where).And("type = 2").And("level = 3").String()
	_, err = o.Raw(sql, condition).Values(&max_name)
	if max_name[0]["m"] == nil {
		max_name[0]["m"] = "0"
	}
	m := max_name[0]["m"].(string)
	//班级号
	ml := strings.Replace(m, "班", "", -1)
	new_name, _ := strconv.Atoi(ml)
	name_number := new_name + 1
	name := strconv.Itoa(name_number)
	if num == 0 {
		err = errors.New("班级不存在")
		return nil, err
	}
	//interface 转 int
	le := or[0]["level"].(string)
	level, _ := strconv.Atoi(le)
	lev := level + 1
	//创建 name+1
	parent_ids := or[0]["parent_ids"].(string)
	ids := or[0]["id"].(string)
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = "insert into organizational set kindergarten_id = ?,name = ?,level = ?,parent_ids = ?,class_type = ?,type = ?,parent_id = ?,is_fixed =?"
	res, err := o.Raw(sql, kindergarten_id, ""+name+"班", lev, ""+parent_ids+""+ids+",", class_type, 2, or[0]["id"], 0).Exec()
	//	or[0]["is_fixed"]
	id, _ := res.LastInsertId()
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	if err == nil {
		paginatorMap["class_id"] = id
		paginatorMap["name"] = "" + name + "班"
		paginatorMap["class_type"] = class_type
		return paginatorMap, nil
	}
	err = errors.New("创建失败")
	return nil, err
}

/*
创建班级
*/
func Store_Class(kindergarten_id int, class string) (err error) {
	o := orm.NewOrm()
	o.Begin()
	var or []orm.Params
	var max_name []orm.Params
	var cl []ClassStr
	json.Unmarshal([]byte(class), &cl)
	for _, v := range cl {
		var condition []interface{}
		where := "1=1 "
		if v.ClassType > 0 {
			where += " AND class_type = ?"
			condition = append(condition, v.ClassType)
		}
		if kindergarten_id > 0 {
			where += " AND kindergarten_id = ?"
			condition = append(condition, kindergarten_id)
		}
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("o.*").From("organizational as o").Where(where).And("type = 2").And("level = 2").String()
		num, err := o.Raw(sql, condition).Values(&or)

		qb, _ = orm.NewQueryBuilder("mysql")
		sql = qb.Select("max(CONVERT(o.name,SIGNED)) as m").From("organizational as o").Where(where).And("type = 2").And("level = 3").String()
		_, err = o.Raw(sql, condition).Values(&max_name)

		if max_name[0]["m"] == nil {
			max_name[0]["m"] = "0"
		}
		m := max_name[0]["m"].(string)
		ml := strings.Replace(m, "班­", "", -1)
		new_name, _ := strconv.Atoi(ml)
		if num == 0 {
			err = errors.New("班级不存在")
			return err
		}
		//interface 转 int
		le := or[0]["level"].(string)
		level, _ := strconv.Atoi(le)
		lev := level + 1
		parent_ids := or[0]["parent_ids"].(string)
		ids := or[0]["id"].(string)
		//	fmt.Println(new_name)
		//	var id int64
		//	var name string
		for a := 1; a <= v.Amount; a++ {
			name_number := a + new_name
			name := strconv.Itoa(name_number)
			qb, _ = orm.NewQueryBuilder("mysql")
			sql = "insert into organizational set kindergarten_id = ?,name = ?,level = ?,parent_ids = ?,class_type = ?,type = ?,parent_id = ?,is_fixed =?"
			_, err := o.Raw(sql, kindergarten_id, ""+name+"班", lev, ""+parent_ids+""+ids+",", v.ClassType, 2, or[0]["id"], 0).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}
	o.Commit()
	return nil
}

/*
教师组织架构
*/
func GetWebOrganization(kindergarten_id int, page int, prepage int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var organ []orm.Params
	qs := o.QueryTable(new(Organizational))
	var posts []Organizational
	var Organizational []OrganizationalTree
	if _, err := qs.Filter("kindergarten_id", kindergarten_id).All(&posts); err == nil {
		ml := make(map[string]interface{})
		for _, val := range posts {
			if val.ParentId == 0 {
				if val.Name != "管理层" {
					next := getNext(posts, val.Id)
					var tree OrganizationalTree
					tree.Id = val.Id
					tree.KindergartenId = val.KindergartenId
					tree.ClassType = val.ClassType
					tree.CreatedAt = val.CreatedAt
					tree.IsFixed = val.IsFixed
					tree.Level = val.Level
					tree.Type = val.Type
					tree.Name = val.Name
					tree.ParentId = val.ParentId
					tree.ParentIds = val.ParentIds
					tree.UpdatedAt = val.UpdatedAt
					tree.Children = next
					Organizational = append(Organizational, tree)
				}
			}
		}
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("organizational.*").From("organizational").Where("kindergarten_id = ? and is_fixed = ? and level = ? and type = ?").String()
		_, err = o.Raw(sql, kindergarten_id, 1, 1, 1).Values(&organ)
		if err == nil {
			ml["class"] = Organizational
			ml["manage"] = organ
			return ml, nil
		}
	}
	return nil, err
}

/*
组织架构列表
*/
func GetTeacherOrganization(kindergarten_id int, page int, prepage int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Organizational))
	var posts []Organizational
	var Organizational []OrganizationalTree
	if _, err := qs.Filter("kindergarten_id", kindergarten_id).All(&posts); err == nil {
		ml := make(map[string]interface{})
		for _, val := range posts {
			if val.ParentId == 0 {
				next := getNext(posts, val.Id)
				var tree OrganizationalTree
				tree.Id = val.Id
				tree.KindergartenId = val.KindergartenId
				tree.ClassType = val.ClassType
				tree.CreatedAt = val.CreatedAt
				tree.IsFixed = val.IsFixed
				tree.Level = val.Level
				tree.Type = val.Type
				tree.Name = val.Name
				tree.ParentId = val.ParentId
				tree.ParentIds = val.ParentIds
				tree.UpdatedAt = val.UpdatedAt
				tree.Children = next
				Organizational = append(Organizational, tree)
			}
		}
		if err == nil {
			ml["data"] = Organizational
			return ml, nil
		}
	}
	return nil, err
}

/*
组织架构列表
*/
func GetOrganization(kindergarten_id int, page int, prepage int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Organizational))
	var posts []Organizational
	var Organizational []OrganizationalTree
	if _, err := qs.Filter("kindergarten_id", kindergarten_id).All(&posts); err == nil {
		ml := make(map[string]interface{})
		for _, val := range posts {
			if val.ParentId == 0 {
				next := getNext(posts, val.Id)
				var tree OrganizationalTree
				tree.Id = val.Id
				tree.KindergartenId = val.KindergartenId
				tree.ClassType = val.ClassType
				tree.CreatedAt = val.CreatedAt
				tree.IsFixed = val.IsFixed
				tree.Level = val.Level
				tree.Type = val.Type
				tree.Name = val.Name
				tree.ParentId = val.ParentId
				tree.ParentIds = val.ParentIds
				tree.UpdatedAt = val.UpdatedAt
				tree.Children = next
				Organizational = append(Organizational, tree)
			}
		}
		if err == nil {
			ml["data"] = Organizational
			return ml, nil
		}
	}
	return nil, err
}

func getNext(posts []Organizational, id int) (Organizational []OrganizationalTree) {
	for _, val := range posts {
		if val.ParentId == id {
			next := getNext(posts, val.Id)
			var tree OrganizationalTree
			tree.Id = val.Id
			tree.KindergartenId = val.KindergartenId
			tree.ClassType = val.ClassType
			tree.CreatedAt = val.CreatedAt
			tree.IsFixed = val.IsFixed
			tree.Level = val.Level
			tree.Name = val.Name
			tree.ParentId = val.ParentId
			tree.ParentIds = val.ParentIds
			tree.UpdatedAt = val.UpdatedAt
			tree.Type = val.Type
			tree.Children = next
			Organizational = append(Organizational, tree)
		}
	}
	return Organizational
}

/*
添加组织架构
*/
func AddOrganization(name string, ty int, parent_id int, kindergarten_id int, class_type int) error {
	o := orm.NewOrm()
	var v []orm.Params
	var kinder []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("kindergarten").Where("kindergarten_id = ?").String()
	num, err := o.Raw(sql, kindergarten_id).Values(&kinder)
	if num == 0 {
		err = errors.New("幼儿园不存在")
		return err
	}
	if parent_id != 0 {
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = qb.Select("*").From("organizational").Where("id = ?").String()
		num, err = o.Raw(sql, parent_id).Values(&v)
		if num == 0 {
			err = errors.New("上一级架构不存在")
			return err
		}
		//interface 转 int
		parent_ids := v[0]["parent_ids"].(string)
		id := v[0]["id"].(string)

		t := v[0]["type"].(string)
		typ, _ := strconv.Atoi(t)

		lev := v[0]["level"].(string)
		leve, _ := strconv.Atoi(lev)
		le := leve + 1
		if typ == 1 {
			if leve >= 2 {
				err = errors.New("管理层不能超过2级")
				return err
			}
		} else {
			if leve >= 3 {
				err = errors.New("管理层不能超过3级")
				return err
			}
		}
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = "insert into organizational set parent_id = ?,name = ?,level = ?,parent_ids = ?,type = ?,kindergarten_id =?,class_type = ?"
		_, err = o.Raw(sql, parent_id, name, le, ""+parent_ids+""+id+",", ty, kindergarten_id, class_type).Exec()
	} else {
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = "insert into organizational set name = ?,type = ?,kindergarten_id =?"
		_, err = o.Raw(sql, name, ty, kindergarten_id).Exec()
	}
	if err == nil {
		return nil
	}
	return err
}

/*
删除组织架构
*/
func DelOrganization(organization_id int) error {
	o := orm.NewOrm()
	var v []orm.Params
	var val []orm.Params
	var organ []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("organizational").Where("id = ?").String()
	_, err := o.Raw(sql, organization_id).Values(&v)
	is_fixe := v[0]["is_fixed"].(string)
	if is_fixe == "1" {
		err = errors.New("不能删除")
		return err
	}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("count(*) as num").From("organizational").Where("parent_ids = ?").String()
	_, err = o.Raw(sql, organization_id).Values(&val)
	num := val[0]["num"].(string)
	nums, _ := strconv.Atoi(num)
	if nums > 0 {
		qb, _ = orm.NewQueryBuilder("mysql")
		sql = qb.Select("organizational.*").From("organizational").Where("parent_ids = ?").String()
		_, err = o.Raw(sql, organization_id).Values(&organ)
		for key, _ := range organ {
			_, err = o.QueryTable("organizational").Filter("id", organ[key]["id"]).Delete()
			_, err = o.QueryTable("organizational_member").Filter("organizational_id", organ[key]["id"]).Delete()
		}
	} else {
		_, err = o.QueryTable("organizational").Filter("id", organization_id).Delete()
		_, err = o.QueryTable("teacher").Filter("teacher_id", organization_id).Delete()
	}
	if err == nil {
		return nil
	}
	return err
}

/*
编辑组织架构
*/
func UpOrganization(organization_id int, name string) error {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("organizational").Where("id = ?").String()
	_, err := o.Raw(sql, organization_id).Values(&v)
	is_fixe := v[0]["is_fixed"].(string)
	if is_fixe == "1" {
		err = errors.New("不能编辑")
		return err
	}
	_, err = o.QueryTable("organizational").Filter("id", organization_id).Update(orm.Params{
		"name": name,
	})
	if err == nil {
		return nil
	}
	err = errors.New("编辑组织架构失败")
	return err
}

/*
班级成员
*/
func Principal(class_id int, page int, prepage int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var teacher []orm.Params
	var condition []interface{}
	where := "1=1 "
	if class_id > 0 {
		where += " AND om.organizational_id = ?"
		condition = append(condition, class_id)
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("t.*", "om.id").From("teacher as t").LeftJoin("organizational_member as om").
		On("t.teacher_id = om.member_id").Where(where).And("om.type = 0").String()
	num, err := o.Raw(sql, condition).Values(&teacher)
	if err == nil && num > 0 {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = teacher
		return paginatorMap, nil
	}
	return nil, err
}

/*
幼儿园所有班级
*/
func GetkinderClass(kindergarten_id int, user_id int, class_type int) (paginatorMapmap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	var teacher []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	paginatorMap := make(map[string]interface{})
	var condition []interface{}
	where := "1=1 "
	if class_type > 0 {
		where += " AND o.class_type = ? and o.level = 3 and o.type = 2"
		condition = append(condition, class_type)
	}
	sql := qb.Select("o.class_type", "o.id as class_id", "o.name as class_name").From("organizational as o").Where(where).And("o.kindergarten_id = ? and o.type = 2").And("o.level = 3").String()
	_, err = o.Raw(sql, condition, kindergarten_id).Values(&v)
	if v == nil {
		err = errors.New("未创建班级")
		return nil, err
	}
	for key, val := range v {
		if val["class_type"].(string) == "3" {
			v[key]["class"] = "大班" + val["class_name"].(string) + ""
		} else if val["class_type"].(string) == "2" {
			v[key]["class"] = "中班" + val["class_name"].(string) + ""
		} else {
			v[key]["class"] = "小班" + val["class_name"].(string) + ""
		}
	}
	qb, _ = orm.NewQueryBuilder("mysql")
	sql = qb.Select("teacher_id").From("teacher").Where("user_id = ? and kindergarten_id = ?").String()
	_, err = o.Raw(sql, user_id, kindergarten_id).Values(&teacher)
	if teacher == nil && err == nil {
		paginatorMap["data"] = v
		paginatorMap["teacher_id"] = nil
		return paginatorMap, nil
	} else {
		paginatorMap["data"] = v
		paginatorMap["teacher_id"] = teacher[0]["teacher_id"]
		return paginatorMap, nil
	}
	return nil, err
}

/*
幼儿园所有班级学生
*/
func GetClassStudent(class_id int) (paginatorMapmap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	paginatorMap := make(map[string]interface{})
	qb, _ = orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "o.class_type", "s.name", "s.avatar", "s.phone").From("student as s").LeftJoin("organizational_member as om").
		On("s.student_id = om.member_id").LeftJoin("organizational as o").
		On("om.organizational_id = o.id").Where("om.organizational_id = ?").And("om.type = 1").String()
	num, err := o.Raw(sql, class_id).Values(&v)
	if err == nil && num > 0 {
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}

/*
筛选学生
*/
func FilterStudent(class_id int, class_type int, kindergarten_id int, student_type int) (paginatorMapmap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	data := make(map[string][]interface{})
	paginatorMap := make(map[string]interface{})
	if class_id != 0 {
		qb, _ = orm.NewQueryBuilder("mysql")
		sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "o.class_type", "s.name", "s.avatar").From("student as s").LeftJoin("organizational_member as om").
			On("s.student_id = om.member_id").LeftJoin("organizational as o").
			On("om.organizational_id = o.id").Where("om.organizational_id = ?").And("om.type = 1").String()
		num, err := o.Raw(sql, class_id).Values(&v)
		if err != nil {
			return nil, err
		}
		for _, val := range v {
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
		if err == nil && num > 0 {
			paginatorMap["data"] = data
			return paginatorMap, nil
		}
		return nil, err
	}
	if class_id == 0 && class_type != 0 && kindergarten_id != 0 {
		if student_type == 4 && class_type == 4 {
			qb, _ = orm.NewQueryBuilder("mysql")
			sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "o.class_type", "s.name", "s.avatar").From("student as s").LeftJoin("organizational_member as om").
				On("s.student_id = om.member_id").LeftJoin("organizational as o").
				On("om.organizational_id = o.id").Where("o.kindergarten_id = ?").And("om.type = 1").String()
			num, err := o.Raw(sql, kindergarten_id).Values(&v)
			if err == nil && num > 0 {
				paginatorMap["data"] = v
				return paginatorMap, nil
			}
		} else {
			qb, _ = orm.NewQueryBuilder("mysql")
			sql := qb.Select("s.student_id", "o.id as class_id", "o.name as class_name", "o.class_type", "s.name", "s.avatar").From("student as s").LeftJoin("organizational_member as om").
				On("s.student_id = om.member_id").LeftJoin("organizational as o").
				On("om.organizational_id = o.id").Where("o.class_type = ? and o.kindergarten_id = ?").And("om.type = 1").String()
			num, err := o.Raw(sql, class_type, kindergarten_id).Values(&v)
			if err == nil && num > 0 {
				paginatorMap["data"] = v
				return paginatorMap, nil
			}
		}
	}
	return nil, err
}

/*
宝宝所在班级
*/
func GetBabyClass(babyIds string) (paginatorMapmap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	baby_id := strings.Split(babyIds, ",")
	paginatorMapmap = make(map[string]interface{})
	var class []map[string]interface{}
	strc := make(map[int]bool)
	for _, s := range baby_id {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select("s.name", "s.student_id", "s.baby_id", "o.name as class_name", "o.id as class_id", "class_type", "s.kindergarten_id", "k.name as kinder_name").From("student as s").LeftJoin("organizational_member as om").
			On("s.student_id = om.member_id").LeftJoin("organizational as o").
			On("om.organizational_id = o.id").LeftJoin("kindergarten as k").
			On("s.kindergarten_id = k.kindergarten_id").Where("s.baby_id = ?").And("s.status = 1").And("om.type = 1").And("om.is_principal = 0").And("isnull(s.deleted_at)").String()
		_, err = o.Raw(sql, s).Values(&v)
		for key, val := range v {
			class_id, _ := strconv.Atoi(val["class_id"].(string))
			if val["class_type"] == "3" {
				if val["class_name"] != nil {
					v[key]["class"] = "大班" + val["class_name"].(string) + ""
				}
			} else if val["class_type"] == "2" {
				if val["class_name"] != nil {
					v[key]["class"] = "中班" + val["class_name"].(string) + ""
				}
			} else if val["class_type"] == "1" {
				if val["class_name"] != nil {
					v[key]["class"] = "小班" + val["class_name"].(string) + ""
				}
			}
			if _, ok := strc[class_id]; !ok {
				class = append(class, val)
				strc[class_id] = true
			}
		}
	}
	paginatorMapmap["data"] = class
	return paginatorMapmap, err
}

/*
班级列表
*/
func AuthClass(kindergarten_id int) interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Organizational))
	var posts []Organizational
	var Organizational []OrganizationalTree
	if _, err := qs.Filter("kindergarten_id", kindergarten_id).All(&posts); err == nil {
		for _, val := range posts {
			if val.ParentId == 0 {
				next := getAuth(posts, val.Id)
				var tree OrganizationalTree
				tree.Id = val.Id
				tree.KindergartenId = val.KindergartenId
				tree.ClassType = val.ClassType
				tree.CreatedAt = val.CreatedAt
				tree.IsFixed = val.IsFixed
				tree.Level = val.Level
				tree.Name = val.Name
				tree.ParentId = val.ParentId
				tree.ParentIds = val.ParentIds
				tree.UpdatedAt = val.UpdatedAt
				tree.Children = next
				Organizational = append(Organizational, tree)
			}
		}
		if err == nil {
			return Organizational[1].Children
		}
	}
	return nil
}

func getAuth(posts []Organizational, id int) (Organizational []OrganizationalTree) {
	for _, val := range posts {
		if val.ParentId == id {
			next := getAuth(posts, val.Id)
			var tree OrganizationalTree
			tree.Id = val.Id
			tree.KindergartenId = val.KindergartenId
			tree.ClassType = val.ClassType
			tree.CreatedAt = val.CreatedAt
			tree.IsFixed = val.IsFixed
			tree.Level = val.Level
			tree.Name = val.Name
			tree.ParentId = val.ParentId
			tree.ParentIds = val.ParentIds
			tree.UpdatedAt = val.UpdatedAt
			tree.Type = val.Type
			tree.Children = next
			Organizational = append(Organizational, tree)
		}
	}
	return Organizational
}

/*
班级幼师，幼儿人数
*/
func GetkinderAll(kindergarten_id int, class_type int) (paginatorMapmap map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	var s []orm.Params
	var t []orm.Params
	var condition []interface{}
	where := "1=1 "
	if class_type > 0 {
		where += " AND o.class_type = ? and o.level = 3 and o.type = 2"
		condition = append(condition, class_type)
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("o.class_type", "o.id as class_id", "o.name as class_name").From("organizational as o").Where(where).And("o.kindergarten_id = ? and o.type = 2").And("o.level = 3").String()
	_, err = o.Raw(sql, condition, kindergarten_id).Values(&v)
	for _, val := range v {
		qb, _ := orm.NewQueryBuilder("mysql")
		sql1 := qb.Select("s.avatar").From("student as s").LeftJoin("organizational_member as om").
			On("s.student_id = om.member_id").Where("om.organizational_id = ?").And("om.type = 1 and om.is_principal = 0").String()
		_, err = o.Raw(sql1, val["class_id"]).Values(&s)

		qb, _ = orm.NewQueryBuilder("mysql")
		sql2 := qb.Select("t.avatar").From("teacher as t").LeftJoin("organizational_member as om").
			On("t.teacher_id = om.member_id").Where("om.organizational_id = ?").And("om.type = 0 and om.is_principal = 0").String()
		_, err = o.Raw(sql2, val["class_id"]).Values(&t)
		val["student"] = s
		val["teacher"] = t
	}
	if v == nil {
		err = errors.New("未创建班级")
		return nil, err
	}
	for key, val := range v {
		if val["class_type"].(string) == "3" {
			v[key]["class"] = "大班" + val["class_name"].(string) + ""
		} else if val["class_type"].(string) == "2" {
			v[key]["class"] = "中班" + val["class_name"].(string) + ""
		} else {
			v[key]["class"] = "小班" + val["class_name"].(string) + ""
		}
	}
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}
