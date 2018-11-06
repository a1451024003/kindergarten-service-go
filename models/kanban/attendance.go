package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris/core/errors"
)

//学生总数/出勤数
func GetAttendance(KindergartenId int, date string) (m map[string]interface{}, code int, err error) {
	ml := make(map[string]interface{})
	//根据kindergartenId查询学生总数
	//获取考勤规则
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("COUNT(Distinct s.student_id) AS student_num").From("student AS s").
		LeftJoin("organizational_member as om").On("s.student_id = om.member_id").
		LeftJoin("organizational as o").On("om.organizational_id = o.id").
		Where("o.kindergarten_id = ?").And("s.status = 1").And("s.deleted_at IS NULL").And("om.type = 1").
		String()
	var student orm.ParamsList
	_, err = o.Raw(sql, KindergartenId).ValuesFlat(&student)
	if err != nil {
		code = 1005
		err = errors.New("学生总数获取失败！")
		return ml, code, err
	} else {
		ml["student_num"] = student[0]
	}
	//查询考勤学生
	attendanceQb, _ := orm.NewQueryBuilder("mysql")
	attendanceSql := attendanceQb.Select("COUNT(Distinct ad.id) AS attendance_num").From("student AS s").
		LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").
		LeftJoin("organizational_member AS om").On("s.student_id = om.member_id").
		LeftJoin("organizational AS o").On("om.organizational_id = o.id").
		Where("o.kindergarten_id = ?").And("ad.atte_time = ?").And("s.status = 1").
		And("s.deleted_at IS NULL").And("om.type = 1").And("ad.atten_type").In("1,5").
		String()
	var attendance orm.ParamsList
	_, err = o.Raw(attendanceSql, KindergartenId, date).ValuesFlat(&attendance)
	if err != nil {
		code = 1005
		err = errors.New("学生出勤数获取失败！")
		return ml, code, err
	} else {
		ml["attendance_num"] = attendance[0]
	}
	return ml, code, err
}

//今日请假
func Leave(kindergartenId int, ClassInfo string, date string) (ml map[string]interface{}, code int, err error) {
	ml = make(map[string]interface{})
	var mll []interface{}
	//查询该班级学生
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("s.name,s.student_id,s.avatar,ad.atten_type,ad.leave_reason").From("student AS s").
		LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").
		LeftJoin("organizational_member AS om").On("s.student_id = om.member_id").
		LeftJoin("organizational AS o").On("om.organizational_id = o.id").
		Where("o.kindergarten_id = ?").And("s.class_info = ?").And("ad.atte_time = ?").
		And("s.status = 1").And("s.deleted_at IS NULL").And("om.type = 1").
		OrderBy("ad.atten_type").Asc().String()
	var leave []orm.Params
	_, err = o.Raw(sql, kindergartenId, ClassInfo, date).Values(&leave)
	if err != nil {
		code = 1005
		err = errors.New("学生总数获取失败！")
		return ml, code, err
	}
	MatterLeave := 0 //事假
	SickLeave := 0   //病假
	for _, v := range leave {
		if v["atten_type"] == "2" {
			MatterLeave += 1
			v["sign"] = "事假"
			mll = append(mll, v)
		}
		if v["atten_type"] == "3" {
			SickLeave += 1
			v["sign"] = "病假"
			mll = append(mll, v)
		}
	}
	ml["sick_leave"] = SickLeave
	ml["matter_leave"] = MatterLeave
	ml["data"] = mll
	return ml, code, err
}
