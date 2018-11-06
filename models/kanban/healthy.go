package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"strings"
)

//喂药申请列表
func GetDrugList(kindergartenId int, classId int, role int, date string, page int, perPage int) (ml interface{}, code int, err error) {
	where := "LEFT(healthy_drug.created_at,10) = '" + date + "'"
	if kindergartenId != 0 {
		where += " AND healthy_drug.kindergarten_id = " + strconv.Itoa(kindergartenId)
	}
	if role != 1 {
		if classId != 0 {
			where += " AND healthy_drug.class_id = " + strconv.Itoa(classId)
		}
	}
	o := orm.NewOrm()
	offset := (page - 1) * perPage //偏移量
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("healthy_drug.*,student.name").From("healthy_drug").
		LeftJoin("student").On("healthy_drug.student_id = student.student_id").
		Where(where).
		OrderBy("id").Desc().
		Limit(perPage).Offset(offset).String()
	var drug []orm.Params
	_, err = o.Raw(sql).Values(&drug)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	} else {
		for _, val := range drug {
			if val["url"] != nil {
				val["url"] = strings.Split(val["url"].(string), ",")
			}
		}
		var total orm.ParamsList
		totalSql := "SELECT COUNT(*) FROM healthy_drug WHERE " + where
		o.Raw(totalSql).ValuesFlat(&total)
		totalInt, _ := strconv.Atoi(total[0].(string))             //总条数
		pageNum := math.Ceil(float64(totalInt) / float64(perPage)) //总页数
		ml = Page{Data: drug, Total: totalInt, PerPage: perPage, PageNum: int(pageNum)}
		return ml, code, err
	}
}

//喂药申请详情
func GetDrugInfo(drugId int) (ml interface{}, code int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("healthy_drug").Where("id = ?").String()
	var drug []orm.Params
	_, err = o.Raw(sql, drugId).Values(&drug)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	} else {
		for _, val := range drug {
			if val["url"] != nil {
				val["url"] = strings.Split(val["url"].(string), ",")
			}
		}
		ml = drug
	}
	return ml, code, err
}

//异常列表
func GetAbnormalList(kindergartenId int, classId int, classType int, date string, pType int, page int, perPage int) (ml interface{}, code int, err error) {
	where := "hi.student_id != 0 AND hi.abnormal != ''"
	if kindergartenId != 0 {
		where += " AND hi.kindergarten_id = " + strconv.Itoa(kindergartenId)
	}
	if classType != 0 {
		where += " AND hi.types = " + strconv.Itoa(classType)
	}
	where += " AND LEFT(hi.date,10) = '" + date + "'"
	if pType == 2 {
		if classId != 0 {
			where += " AND hi.class_id = " + strconv.Itoa(classId)
		}
	}
	o := orm.NewOrm()
	offset := (page - 1) * perPage //偏移量
	qb, _ := orm.NewQueryBuilder("mysql")
	column := "hi.*,hc.abnormal1,hc.abnormal2,hc.abnormal3,hc.abnormal4,hc.abnormal5,hc.abnormal6,o.name as className,s.name AS student_name,s.avatar,s.class_info AS class_name,t.name AS teacher_name"
	sql := qb.Select(column).From("healthy_inspect AS hi").
		LeftJoin("student AS s").On("hi.student_id = s.student_id").
		LeftJoin("teacher AS t").On("hi.teacher_id = t.teacher_id").
		LeftJoin("organizational AS o").On("hi.class_id = o.id").
		LeftJoin("healthy_column AS hc").On("hi.id = hc.inspect_id").
		Where(where).
		OrderBy("hi.id").Desc().
		Limit(perPage).Offset(offset).String()
	var abnormal []orm.Params
	_, err = o.Raw(sql).Values(&abnormal)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	} else {
		for _, val := range abnormal {
			if val["url"] != nil {
				val["url"] = strings.Split(val["url"].(string), ",")
			}
		}
		var total orm.ParamsList
		qb, _ := orm.NewQueryBuilder("mysql")
		totalSql := qb.Select("COUNT(*)").From("healthy_inspect AS hi").
			LeftJoin("student AS s").On("hi.student_id = s.student_id").
			LeftJoin("teacher AS t").On("hi.teacher_id = t.teacher_id").
			LeftJoin("healthy_column AS hc").On("hi.id = hc.inspect_id").
			Where(where).String()
		o.Raw(totalSql).ValuesFlat(&total)
		totalInt, _ := strconv.Atoi(total[0].(string))             //总条数
		pageNum := math.Ceil(float64(totalInt) / float64(perPage)) //总页数
		ml = Page{Data: abnormal, Total: totalInt, PerPage: perPage, PageNum: int(pageNum)}
		return ml, code, err
	}
}

//异常详情
func GetAbnormalInfo(abnormalId int) (ml interface{}, code int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	column := "hi.*,hc.abnormal1,hc.abnormal2,hc.abnormal3,hc.abnormal4,hc.abnormal5,hc.abnormal6,o.name as className,s.name AS student_name,s.avatar,s.class_info AS class_name,t.name AS teacher_name"
	sql := qb.Select(column).From("healthy_inspect AS hi").
		LeftJoin("student AS s").On("hi.student_id = s.student_id").
		LeftJoin("teacher AS t").On("hi.teacher_id = t.teacher_id").
		LeftJoin("organizational AS o").On("hi.class_id = o.id").
		LeftJoin("healthy_column AS hc").On("hi.id = hc.inspect_id").
		Where("hi.id = ?").String()
	var abnormal []orm.Params
	_, err = o.Raw(sql, abnormalId).Values(&abnormal)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	}
	for _, val := range abnormal {
		if val["url"] != nil {
			val["url"] = strings.Split(val["url"].(string), ",")
		}
	}
	ml = abnormal
	return ml, code, err
}
