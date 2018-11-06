package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Attendance struct {
	Id             int    `json:"id" orm:"column(id);auto"`
	SignStartTime  string `json:"sign_start_time" orm:"column(sign_start_time); description:"签到开始时间""`
	SignEndTime    string `json:"sign_end_time" orm:"column(sign_end_time)"; description:"签到结束时间"`
	BackStartTime  string `json:"back_start_time" orm:"column(back_start_time)" ;description:"离园开始时间"`
	BackEndTime    string `json:"back_end_time" orm:"column(back_end_time)"; description:"离园结束时间"`
	UseTime        string `json:"use_time" orm:"column(use_time)";description:"重复使用时间"`
	KindergartenId string `json:"kindergarten_id" orm:"column(kindergarten_id)";description:"幼儿园ID"`
}

type AttendanceDetailed struct {
	Id             int    `json:"id" orm:"column(id);auto"`
	AttenId        int    `json:"atten_id" orm:"column(atten_id); description:"使用时间/规则ID""`
	StudentId      int    `json:"student_id" orm:"column(student_id); description:"学生ID""`
	SignTime       string `json:"sign_time" orm:"column(sign_time)"; description:"入园时间"`
	BackTime       string `json:"back_time" orm:"column(back_time)" ;description:"离园时间"`
	AttenType      int    `json:"atten_type" orm:"column(atten_type)"; description:"考勤状态：1正常；2事假；3病假；4异常"`
	LeaveStartTime string `json:"leave_start_time" orm:"column(leave_start_time)";description:"请假开始时间"`
	LeaveEndtime   string `json:"leave_end_time" orm:"column(leave_end_time)";description:"请假结束时间"`
	LeaveReason    string `json:"leave_reason" orm:"column(leave_reason)";description:"请假原因"`
	AtteTime       string `json:"atte_time" orm:"column(atte_time)";description:"考勤时间"`
}
type JSONStruct struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func (t *Attendance) TableName() string {
	return "attendance"
}
func init() {
	orm.RegisterModel(new(Attendance))
	orm.RegisterModel(new(AttendanceDetailed))

}

// 新增考勤规则
func AddAttendance(m *Attendance) (id int64, err error) {
	o := orm.NewOrm()
	//查找该幼儿园是否存在考勤规则
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	where := "kindergarten_id = " + m.KindergartenId
	sql := qb.Select("id").From("attendance").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		if len(maps) > 0 {
			num, err := o.QueryTable("attendance").Filter("kindergarten_id", m.KindergartenId).Update(orm.Params{
				"sign_start_time": m.SignStartTime,
				"sign_end_time":   m.SignEndTime,
				"back_start_time": m.BackStartTime,
				"back_end_time":   m.BackEndTime,
				"use_time":        m.UseTime,
			})
			return num, err
		} else {
			id, err = o.Insert(m)
		}
	}

	return
}

//新增考勤
func AddDetailed(m *AttendanceDetailed) (id int64, err error) {
	o := orm.NewOrm()
	//查询学生当天考勤是否存在，存在则修改，不存在则新增
	StudentId := m.StudentId
	AtteTime := m.AtteTime
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	where := "student_id = " + strconv.Itoa(StudentId)
	where += " AND atte_time = \"" + AtteTime + "\""
	sql := qb.Select("id").From("attendance_detailed").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		if len(maps) > 0 {
			for _, v := range maps {
				i, _ := strconv.Atoi(v["id"].(string))
				m.Id = i
				if _, err = o.Update(m); err == nil {
					id = int64(m.Id)
				}
			}
		} else {
			id, err = o.Insert(m)
		}
	}
	return
}

//一键入园
func TopspeedDetailed(AttenId int, StudentId string, SignTime string, BackTime string, AtteTime string) (id int64, err error) {
	o := orm.NewOrm()
	str := strings.Split(StudentId, ",")
	t := time.Now()
	date := t.Format("2006-01-02")
	for _, v := range str {
		i, _ := strconv.Atoi(v)
		o.QueryTable("attendance_detailed").Filter("student_id", i).Filter("atte_time", date).Update(orm.Params{
			"atten_id":   AttenId,
			"sign_time":  SignTime,
			"back_time":  BackTime,
			"atte_time":  AtteTime,
			"atten_type": 1,
		})
	}

	return 0, nil
}

func GetSignList(ClassInfo string, KindergartenId int) (s map[string]interface{}, err error) {
	//查询当前时间是否上课
	t := time.Now()
	NotSignToDb(KindergartenId)
	var Week int
	WeekDay := t.Weekday().String()
	if WeekDay == "Sunday" {
		Week = 7
	} else if WeekDay == "Monday" {
		Week = 1
	} else if WeekDay == "Tuesday" {
		Week = 2
	} else if WeekDay == "Wednesday" {
		Week = 3
	} else if WeekDay == "Thursday" {
		Week = 4
	} else if WeekDay == "Friday" {
		Week = 5
	} else if WeekDay == "Saturday" {
		Week = 6
	}
	//查询当前幼儿园规则
	o := orm.NewOrm()
	var maps []orm.Params
	var smaps []orm.Params
	var xmaps []orm.Params
	var ml []interface{}
	var mld []interface{}
	qb, _ := orm.NewQueryBuilder("mysql")
	where := "kindergarten_id = " + strconv.Itoa(KindergartenId)
	sql := qb.Select("id as atten_id,use_time,sign_end_time").From("attendance").Where(where).String()
	wherk_type := 0
	var AttenId int
	var sign_end_time string
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		for _, v := range maps {
			n := v["use_time"]
			atten_id, _ := strconv.Atoi(v["atten_id"].(string))
			sign_end_time = v["sign_end_time"].(string)
			AttenId = atten_id
			wherk := strings.Split(n.(string), ",")
			for _, val := range wherk {
				i, _ := strconv.Atoi(val)
				if Week == i {
					wherk_type = 1
					break
				}
			}
		}
	}

	normal := 0      //正常
	abnormal := 0    //异常
	Absenteeism := 0 //旷课1
	MatterLeave := 0 //事假2
	SickLeave := 0   //病假3
	m := make(map[string]interface{})
	if wherk_type == 1 {
		m["holiday"] = 0
		//判断当前时间是否在签到结束时间之前
		time := time.Now().Format("15:04")
		date := t.Format("2006-01-02")

		if time < sign_end_time {
			//查询该班级学生
			sqb, _ := orm.NewQueryBuilder("mysql")
			swhere := "o.kindergarten_id = " + strconv.Itoa(KindergartenId)
			swhere += " AND s.class_info = " + "\"" + ClassInfo + "\""
			swhere += " AND ad.atte_time = " + "\"" + date + "\""
			swhere += " AND s.status = 1"
			swhere += " AND om.type = 1"
			swhere += " AND s.deleted_at IS NULL "
			ssql := sqb.Select("ad.sign_time,ad.back_time,s.name,s.student_id,s.avatar,ad.atten_type,ad.leave_reason").From("student AS s").LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(swhere).OrderBy("ad.atten_type").String()

			if _, err := o.Raw(ssql).Values(&smaps); err == nil {
				for _, v := range smaps {
					if v["atten_type"] == "0" || v["atten_type"] == "1" || v["atten_type"] == "5" {
						normal += 1
						ml = append(ml, v)
					} else {
						if v["atten_type"] == "4" {
							Absenteeism += 1
						}
						if v["atten_type"] == "3" {
							SickLeave += 1
						}
						if v["atten_type"] == "2" {
							MatterLeave += 1
						}
						abnormal += 1
						mld = append(mld, v)
					}
				}
			}
		} else {
			//先将没签到所有的学生放入attendance_detailed表中
			//从attendance_detailed表中查出签到列表
			sqb, _ := orm.NewQueryBuilder("mysql")
			swhere := "o.kindergarten_id = " + strconv.Itoa(KindergartenId)
			swhere += " AND s.class_info = " + "\"" + ClassInfo + "\""
			swhere += " AND ad.atte_time = " + "\"" + date + "\""
			swhere += " AND s.status = 1"
			swhere += " AND om.type = 1"
			swhere += " AND s.deleted_at IS NULL "
			ssql := sqb.Select("ad.sign_time,ad.back_time,ad.id,s.name,s.student_id,s.avatar,ad.atten_type,ad.leave_reason").From("student AS s").LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(swhere).OrderBy("ad.atten_type").String()

			if _, err := o.Raw(ssql).Values(&smaps); err == nil {
				for _, v := range smaps {
					//sign_time
					//sign_end_time
					//迟到：入园时间大于入园结束时间且以签到，属于异常
					if v["sign_time"].(string) != "" {
						if v["sign_time"].(string) > sign_end_time {
							v["atten_type"] = 6
						}
					}
					if v["atten_type"] == "1" || v["atten_type"] == "5" {
						normal += 1
						ml = append(ml, v)
					} else {
						if v["atten_type"] == "4" || v["atten_type"] == "6" {
							Absenteeism += 1
						}
						if v["atten_type"] == "3" {
							SickLeave += 1
						}
						if v["atten_type"] == "2" {
							MatterLeave += 1
						}
						abnormal += 1
						mld = append(mld, v)
					}
				}
			}
		}
		m["normal"] = normal
		m["abnormal"] = abnormal
		m["atten_id"] = AttenId
		m["absenteeism"] = Absenteeism
		m["sickleave"] = SickLeave
		m["matterleave"] = MatterLeave
		m["normal_attendance_list"] = ml
		m["abnormal_attendance_list"] = mld
		return m, nil
	} else {
		//查询该班级学生
		xsqb, _ := orm.NewQueryBuilder("mysql")
		xwhere := "o.kindergarten_id = " + strconv.Itoa(KindergartenId)
		xwhere += " AND s.class_info = " + "\"" + ClassInfo + "\""
		xwhere += " AND s.status = 1"
		xwhere += " AND om.type = 1"
		xwhere += " AND s.deleted_at IS NULL "
		xsql := xsqb.Select("s.name,s.student_id,s.avatar").From("student AS s").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(xwhere).OrderBy("s.student_id").Desc().String()
		if _, err := o.Raw(xsql).Values(&xmaps); err == nil {
			for _, v := range xmaps {
				normal += 1
				ml = append(ml, v)
			}
		}
		m["holiday"] = 1
		m["normal"] = normal
		m["abnormal"] = 0
		m["atten_id"] = AttenId
		m["absenteeism"] = Absenteeism
		m["sickleave"] = SickLeave
		m["matterleave"] = MatterLeave
		m["normal_attendance_list"] = ml
		m["abnormal_attendance_list"] = mld
		return m, nil
	}

	return nil, err
}

//将没签到的学生写入attendance_detailed
func NotSignToDb(KindergartenId int) {
	o := orm.NewOrm()
	var maps []orm.Params
	var smaps []orm.Params
	var Detailed AttendanceDetailed
	var DetailedRe []AttendanceDetailed
	qb, _ := orm.NewQueryBuilder("mysql")
	where := " s.status = 1"
	where += " AND o.kindergarten_id = " + strconv.Itoa(KindergartenId)
	where += " AND om.type = 1"
	where += " AND s.deleted_at IS NULL "
	sql := qb.Select("s.kindergarten_id,s.student_id,s.avatar").From("student AS s").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(where).String()
	date := time.Now().Format("2006-01-02")
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		for _, v := range maps {
			//查询该学生是否在考勤表中
			sqb, _ := orm.NewQueryBuilder("mysql")
			swhere := "student_id = " + v["student_id"].(string)
			swhere += " AND atte_time = \"" + date + "\""
			ssql := sqb.Select("id").From("attendance_detailed").Where(swhere).String()
			if _, serr := o.Raw(ssql).Values(&smaps); serr == nil {
				if len(smaps) == 0 {
					v["atten_type"] = nil
				} else {
					v["atten_type"] = 1
				}
			}
			if v["atten_type"] == nil {
				//通过幼儿园ID拿到规则ID
				var kmaps []orm.Params
				kqb, _ := orm.NewQueryBuilder("mysql")
				where := "kindergarten_id = " + v["kindergarten_id"].(string)
				ksql := kqb.Select("id AS attendance_id").From("attendance").Where(where).String()
				if k, err := o.Raw(ksql).Values(&kmaps); err == nil {
					if k > 0 {
						for _, va := range kmaps {
							AttenId, _ := strconv.Atoi(va["attendance_id"].(string))
							StudentId, _ := strconv.Atoi(v["student_id"].(string))
							Detailed.AttenId = AttenId
							Detailed.StudentId = StudentId
							//Detailed.AttenType = 4
							Detailed.AtteTime = date
						}
						DetailedRe = append(DetailedRe, Detailed)
					}
				}

			}
		}
		o.InsertMulti(1, DetailedRe)
	}
}

//请假
func LeaveDetailed(m *AttendanceDetailed, KindergartenId int) (id int64, err error) {
	//通过KindergartenId获取规则ID  atten_id
	o := orm.NewOrm()
	time := time.Now().Format("2006-01-02")
	//查询当前时间该学生有无考勤操作
	var smaps []orm.Params
	num, _ := o.QueryTable("attendance_detailed").Filter("student_id", m.StudentId).Filter("atte_time", time).Values(&smaps)
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	where := "kindergarten_id = " + strconv.Itoa(KindergartenId)
	sql := qb.Select("id AS atten_id").From("attendance").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		for _, v := range maps {
			i, _ := strconv.Atoi(v["atten_id"].(string))
			m.AttenId = i
			m.AtteTime = time
		}
		for _, res := range smaps {
			Id := res["Id"].(int64)
			m.Id = int(Id)
		}

		if num == 0 {
			id, err = o.Insert(m)
		} else {
			id, err = o.Update(m)
		}

	}

	return
}

//园长端考勤列表
func GetAllAttendance(ClassType int, AtteTime string, KindergartenId int) (m map[string]interface{}, err error) {
	NotSignToDb(KindergartenId)
	o := orm.NewOrm()
	//通过班级名称获得下面的所有班
	var maps []orm.Params
	var cmaps []orm.Params
	//获取考勤规则
	SignEndTime, _ := GetRUle(KindergartenId)
	var sign_end_time string
	for _, v := range SignEndTime {
		sign_end_time = v["sign_end_time"].(string)
	}

	qb, _ := orm.NewQueryBuilder("mysql")
	cqb, _ := orm.NewQueryBuilder("mysql")

	cwhere := " class_type = " + strconv.Itoa(ClassType)
	cwhere += " AND kindergarten_id = " + strconv.Itoa(KindergartenId)
	cwhere += " AND level=2"
	var ClassName string
	csql := cqb.Select("name AS class_name").From("organizational").Where(cwhere).String()
	if _, err := o.Raw(csql).Values(&cmaps); err == nil {
		for _, val := range cmaps {
			ClassName = val["class_name"].(string)
		}
	}
	where := " class_type = " + strconv.Itoa(ClassType)
	where += " AND kindergarten_id = " + strconv.Itoa(KindergartenId)
	where += " AND level=3"
	sql := qb.Select("name").From("organizational").Where(where).String()
	mle := make(map[string]interface{})
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		for _, v := range maps {
			name := ClassName + v["name"].(string)
			m, _ := GetNum(name, AtteTime)
			Normal := 0     //正常1
			ThingThing := 0 //事假2
			SickLeave := 0  //病假3
			Abnormal := 0   //异常4
			ml := make(map[string]interface{})
			for _, val := range m {
				atten_type, _ := strconv.Atoi(val["atten_type"].(string))
				//sign_time
				//sign_end_time
				//迟到：入园时间大于入园结束时间且以签到，属于异常
				if val["sign_time"].(string) > sign_end_time {
					atten_type = 6
				}
				if atten_type == 1 || atten_type == 5 {
					Normal += 1
				} else if atten_type == 2 {
					ThingThing += 1
				} else if atten_type == 3 {
					SickLeave += 1
				} else if atten_type == 4 || atten_type == 6 {
					Abnormal += 1
				}
				ml["normal"] = Normal
				ml["ThingThing"] = ThingThing
				ml["SickLeave"] = SickLeave
				ml["abnormal"] = Abnormal
			}
			mle[v["name"].(string)] = ml
		}
	}
	return mle, err
}

//通过班级名称获取考勤数据（大班1班）
func GetNum(name string, AtteTime string) (m []orm.Params, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	//分别用名字取考勤数据
	where := " s.class_info=\"" + name + "\""
	where += " AND ad.atte_time=\"" + AtteTime + "\""
	where += " AND s.status = 1"
	where += " AND s.deleted_at IS NULL "
	sql := qb.Select("ad.atten_type,ad.sign_time ").From("attendance_detailed AS ad").LeftJoin("attendance as a").On("ad.atten_id=a.id").LeftJoin("student AS s ").On("ad.student_id=s.student_id").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		return maps, nil
	}
	return nil, err
}

//离园
func LeaveGarden(StudentId int) (num int64, err error) {
	//查询该学生是否入园
	o := orm.NewOrm()
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	//分别用名字取考勤数据
	date := time.Now().Format("2006-01-02")
	time := time.Now().Format("15:04")
	where := " student_id=" + strconv.Itoa(StudentId)
	where += " AND atte_time = \"" + date + "\""
	sql := qb.Select("student_id").From("attendance_detailed").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		if len(maps) > 0 {
			num, err := o.QueryTable("attendance_detailed").Filter("student_id", StudentId).Filter("atte_time", date).Update(orm.Params{
				"back_time":  time,
				"atten_type": 5,
			})
			if err == nil {
				return num, err
			} else {
				return 0, err
			}
		}
	}
	return
}

//一键离园
func TopspeedLeaveGarden(ClassInfo string) (err error) {
	//查询该学生是否入园
	o := orm.NewOrm()
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	//查询该班级下所有已入园学生
	date := time.Now().Format("2006-01-02")
	time := time.Now().Format("15:04")
	where := " s.class_info=\"" + ClassInfo + "\""
	where += " AND ad.atte_time = \"" + date + "\""
	where += " AND ad.back_time=''"
	where += " AND ad.atten_type=1"
	where += " AND s.status = 1"
	where += " AND s.deleted_at IS NULL "
	sql := qb.Select("s.student_id").From("student AS s").LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		if len(maps) > 0 {
			o.Begin()
			for _, v := range maps {
				_, err := o.QueryTable("attendance_detailed").Filter("student_id", v["student_id"]).Filter("atte_time", date).Update(orm.Params{
					"back_time":  time,
					"atten_type": 5,
				})
				if err != nil {
					o.Rollback()
					return err
				} else {
					o.Commit()
				}
			}
		}
	}
	return err
}

//园长端考勤表图（整个幼儿园）
func AttendanceGarden(KindergartenId int) (m map[string]interface{}, err error) {
	//根据KindergartenId查询学生总数
	//获取考勤规则
	SignEndTime, _ := GetRUle(KindergartenId)
	var sign_end_time string
	for _, v := range SignEndTime {
		sign_end_time = v["sign_end_time"].(string)
	}
	o := orm.NewOrm()
	var maps []orm.Params
	var smaps []orm.Params
	ml := make(map[string]interface{})
	qb, _ := orm.NewQueryBuilder("mysql")
	sqb, _ := orm.NewQueryBuilder("mysql")
	where := "o.kindergarten_id=" + strconv.Itoa(KindergartenId)
	where += " AND s.status = 1"
	where += " AND s.deleted_at IS NULL "
	where += " AND om.type = 1"
	sql := qb.Select("count(*) AS studentnum").From("student AS s").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(where).String()
	allstudent := 0    //全部学生
	allattendance := 0 //出勤学生
	matterleave := 0   //事假
	sickleave := 0     //病假
	abnormal := 0      //异常
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		for _, v := range maps {
			i, _ := strconv.Atoi(v["studentnum"].(string))
			allstudent = i
		}
	}
	//查询考勤学生
	swhere := "o.kindergarten_id=" + strconv.Itoa(KindergartenId)
	//atte_time
	swhere += " AND ad.atte_time=\"" + time.Now().Format("2006-01-02") + "\""
	swhere += " AND s.status = 1"
	swhere += " AND om.type = 1"
	swhere += " AND s.deleted_at IS NULL "
	ssql := sqb.Select("ad.atten_type,ad.sign_time").From("student AS s").LeftJoin("attendance_detailed AS ad").On("s.student_id=ad.student_id").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").Where(swhere).String()
	if _, err := o.Raw(ssql).Values(&smaps); err == nil {
		for _, val := range smaps {
			//sign_time
			//sign_end_time
			//迟到：入园时间大于入园结束时间且以签到，属于异常
			if val["sign_time"].(string) != "" {
				if val["sign_time"].(string) > sign_end_time {
					val["atten_type"] = "4"
				}
			}
			if val["atten_type"] == "1" || val["atten_type"] == "5" {
				allattendance += 1
			} else if val["atten_type"] == "2" {
				matterleave += 1
			} else if val["atten_type"] == "3" {
				sickleave += 1
			} else if val["atten_type"] == "4" || val["atten_type"] == "0" {
				abnormal += 1
			}
		}
	}
	ml["allstudent"] = allstudent
	ml["allattendance"] = allattendance
	ml["matterleave"] = matterleave
	ml["sickleave"] = sickleave
	ml["abnormal"] = abnormal

	return ml, err
}

//园长获取某班级考勤详情
func GetDetailed(KindergartenId int, Class string, ClassName string, AtteTime string) (m map[string]interface{}, err error) {
	o := orm.NewOrm()
	//获取考勤规则
	SignEndTime, _ := GetRUle(KindergartenId)
	var sign_end_time string
	for _, v := range SignEndTime {
		sign_end_time = v["sign_end_time"].(string)
	}
	var maps []orm.Params
	normal := 0       //正常
	abnormal := 0     //异常
	matterleave := 0  //事假
	sickleave := 0    //病假
	abnormal_num := 0 //异常
	var ml []interface{}
	var mld []interface{}
	qb, _ := orm.NewQueryBuilder("mysql")
	name := Class + ClassName
	where := " s.class_info=\"" + name + "\""
	where += " AND o.kindergarten_id=" + strconv.Itoa(KindergartenId)
	where += " AND ad.atte_time=\"" + AtteTime + "\""
	where += " AND s.status = 1"
	where += " AND om.type = 1"
	where += " AND s.deleted_at IS NULL "
	sql := qb.Select("s.name,s.student_id,s.avatar,ad.atten_type,ad.leave_reason,ad.sign_time,ad.back_time").From("attendance_detailed AS ad").LeftJoin("student AS s ").LeftJoin("organizational_member as om").On("s.student_id = om.member_id").LeftJoin("organizational as o").On("om.organizational_id = o.id").On("ad.student_id=s.student_id").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		m := make(map[string]interface{})
		for _, v := range maps {
			time_now := time.Now().Format("15:04")
			//迟到：入园时间大于入园结束时间且以签到，属于异常
			if time_now > sign_end_time {
				if v["atten_type"] != "2" && v["atten_type"] != "3" {
					if v["sign_time"].(string) > sign_end_time || v["sign_time"].(string) == "" {
						v["atten_type"] = 6
					}
				}
			}
			if v["atten_type"] == "1" || v["atten_type"] == "5" || v["atten_type"] == "0" {
				normal += 1
				ml = append(ml, v)
			} else {
				abnormal += 1
				mld = append(mld, v)
			}
			if v["atten_type"] == "2" {
				matterleave += 1
			} else if v["atten_type"] == "3" {
				sickleave += 1
			} else if v["atten_type"] == "4" || v["atten_type"] == 6 {
				abnormal_num += 1
			}
		}
		m["normal"] = normal
		m["abnormal"] = abnormal
		m["normal_attendance_list"] = ml
		m["abnormal_attendance_list"] = mld
		m["matterleave"] = matterleave
		m["sickleave"] = sickleave
		m["abnormal_num"] = abnormal_num
		return m, err
	}
	return nil, err
}

//园长获取考勤规则
func GetRUle(KindergartenId int) (m []orm.Params, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	where := "kindergarten_id=" + strconv.Itoa(KindergartenId)
	sql := qb.Select("*").From("attendance").Where(where).String()
	if _, err := o.Raw(sql).Values(&maps); err == nil {
		return maps, err
	}
	return nil, err
}

func UpdateAttensById(id int, AttenType string, LeaveReason string) (number int64, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("attendance_detailed").Filter("id", id).Update(orm.Params{
		"atten_type":   AttenType,
		"leave_reason": LeaveReason,
	})
	return num, err
}
