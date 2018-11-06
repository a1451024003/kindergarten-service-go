package healthy

import (
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris/core/errors"
	"math"
	"strconv"
)

//健康统计
func Statistics(kindergartenId int, date string) (ml map[string]interface{}, code int, err error) {
	o := orm.NewOrm()

	//幼儿园学生总数
	totalQb, _ := orm.NewQueryBuilder("mysql")
	totalSql := totalQb.Select("COUNT(Distinct s.student_id)").From("student AS s").
		LeftJoin("organizational_member AS om").On("s.student_id = om.member_id").
		LeftJoin("organizational AS o").On("om.organizational_id = o.id").
		Where("o.kindergarten_id = ?").And("om.type = 1").String()
	var totalNum orm.ParamsList
	_, err = o.Raw(totalSql, kindergartenId).ValuesFlat(&totalNum)
	if err != nil {
		code = 1005
		err = errors.New("学生总数获取失败！")
		return ml, code, err
	}
	totalNumInt, _ := strconv.Atoi(totalNum[0].(string))
	if totalNumInt < 1 {
		code = 1002
		err = errors.New("幼儿园暂无学生！")
		return ml, code, err
	} else {
		//病假人数
		sickLeaveQb, _ := orm.NewQueryBuilder("mysql")
		sickLeaveSql := sickLeaveQb.Select("COUNT(Distinct s.student_id)").From("student AS s").
			LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").
			LeftJoin("organizational_member AS om").On("s.student_id = om.member_id").
			LeftJoin("organizational AS o").On("om.organizational_id = o.id").
			Where("o.kindergarten_id = ?").And("ad.atte_time = ?").And("ad.atten_type = 3").
			And("s.status = 1").And("s.deleted_at IS NULL").And("om.type = 1").
			String()
		var sickLeaveNum orm.ParamsList
		_, err = o.Raw(sickLeaveSql, kindergartenId, date).ValuesFlat(&sickLeaveNum)
		if err != nil {
			code = 1005
			err = errors.New("病假人数获取失败！")
			return ml, code, err
		}

		abnormalWhere := "hi.student_id != 0 AND hi.abnormal != '' AND LEFT(hi.date,10) = '" + date + "'"
		drugWhere := "LEFT(created_at,10) = '" + date + "'"
		if kindergartenId != 0 {
			abnormalWhere += " AND hi.kindergarten_id = " + strconv.Itoa(kindergartenId)
			drugWhere += " AND kindergarten_id = " + strconv.Itoa(kindergartenId)
		}

		//异常人数
		abnormalQb, _ := orm.NewQueryBuilder("mysql")
		abnormalSql := abnormalQb.Select("COUNT(Distinct hi.id)").From("healthy_inspect AS hi").
			LeftJoin("student AS s").On("hi.student_id = s.student_id").
			LeftJoin("teacher AS t").On("hi.teacher_id = t.teacher_id").
			LeftJoin("healthy_column AS hc").On("hi.id = hc.inspect_id").
			Where(abnormalWhere).String()
		var abnormalNum orm.ParamsList
		_, err = o.Raw(abnormalSql).ValuesFlat(&abnormalNum)
		if err != nil {
			code = 1005
			err = errors.New("异常人数获取失败！")
			return ml, code, err
		}

		//喂药通知数量
		var drugNum orm.ParamsList
		drugSql := "SELECT COUNT(Distinct id) FROM healthy_drug WHERE " + drugWhere
		_, err = o.Raw(drugSql).ValuesFlat(&drugNum)
		if err != nil {
			code = 1005
			err = errors.New("喂药通知获取失败！")
			return ml, code, err
		}

		//健康人数
		sickLeaveNumInt, _ := strconv.Atoi(sickLeaveNum[0].(string))
		abnormalNumInt, _ := strconv.Atoi(abnormalNum[0].(string))
		healthNum := totalNumInt - (sickLeaveNumInt + abnormalNumInt)

		//健康率
		rate := float64(healthNum) / float64(totalNumInt)
		rates := math.Trunc(rate*1e3+0.5) * 1e-3
		healthRate := float64(rates) * 100

		ml = make(map[string]interface{})
		ml["total_num"] = totalNum[0]          //学生总数
		ml["sick_leave_num"] = sickLeaveNum[0] //病假人数
		ml["abnormal_num"] = abnormalNum[0]    //异常人数
		ml["drug_num"] = drugNum[0]            //喂药通知数量
		ml["health_num"] = healthNum           //健康人数
		ml["health_rate"] = healthRate         //健康率
	}
	return ml, code, err
}
