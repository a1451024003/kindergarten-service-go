package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris/core/errors"
	"strconv"
	"time"
)

type Schedule struct {
	Id             int    `json:"id" orm:"column(id);auto" description:"编号"`
	Content        string `json:"content" orm:"column(content)" description:"内容"`
	Time           string `json:"time" orm:"column(time)" description:"时间"`
	KindergartenId int    `json:"kindergarten_id" orm:"column(kindergarten_id)" description:"幼儿园ID"`
	UserId         int    `json:"user_id" orm:"column(user_id)" description:"用户ID"`
	Type           int    `json:"type" orm:"column(type)" description:"类型：1日程，2备忘录"`
	CreatedAt      string `json:"created_at" orm:"column(created_at);type(datetime)" description:"创建时间"`
	UpdatedAt      string `json:"updated_at" orm:"column(updated_at);type(datetime)" description:"修改时间"`
}

func (s *Schedule) TableName() string {
	return "schedule"
}

func init() {
	orm.RegisterModel(new(Schedule))
}

//日程添加
func PostSchedule(content string, times string, kindergartenId int, ty int, userId int, role int) (ml interface{}, code int, err error) {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	schedule := Schedule{Content: content, Time: times, KindergartenId: kindergartenId, UserId: userId, Type: ty, CreatedAt: timeNow, UpdatedAt: timeNow}
	if ty == 1 { //日程
		if role != 1 {
			code = 1006
			err = errors.New("用户没有权限！")
			return ml, code, err
		} else {
			_, err = o.Insert(&schedule)
			if err != nil {
				code = 1003
				err = errors.New("添加失败！")
				return ml, code, err
			} else {
				ml = schedule
			}
		}
	}
	if ty == 2 { //备忘录
		_, err = o.Insert(&schedule)
		if err != nil {
			code = 1003
			err = errors.New("添加失败！")
			return ml, code, err
		} else {
			ml = schedule
		}
	}
	return ml, code, err
}

//日程列表
func GetScheduleList(kindergartenId int, userId int, date string) (ml map[string]interface{}, code int, err error) {
	o := orm.NewOrm()
	var schedule []orm.Params
	sql := "SELECT * FROM `schedule` WHERE CASE `type` WHEN 1 THEN kindergarten_id = " + strconv.Itoa(kindergartenId) + " ELSE kindergarten_id = " + strconv.Itoa(kindergartenId) + " AND user_id = " + strconv.Itoa(userId) + " END AND LEFT (created_at, 10) = \"" + date + "\" ORDER BY time ASC"
	_, err = o.Raw(sql).Values(&schedule)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	} else {
		ml = make(map[string]interface{})
		if schedule == nil {
			ml["data"] = nil
		} else {
			ml["data"] = schedule
		}
	}
	return ml, code, err
}

//日程详情
func GetScheduleInfo(scheduleId int) (ml interface{}, code int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("schedule").Where("id = ?")
	sql := qb.String()
	var schedule []Schedule
	_, err = o.Raw(sql, scheduleId).QueryRows(&schedule)
	if err != nil {
		code = 1005
		err = errors.New("获取失败！")
		return ml, code, err
	} else {
		if schedule == nil {
			ml = nil
		} else {
			ml = schedule[0]
		}
	}
	return ml, code, err
}

//日程修改
func PutSchedule(scheduleId int, content string, times string, kindergartenId int, userId int) (code int, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("schedule").
		Filter("id", scheduleId).Filter("kindergarten_id", kindergartenId).Filter("user_id", userId).
		Update(orm.Params{"content": content, "time": times, "updated_at": time.Now()})
	if err != nil {
		code = 1003
		err = errors.New("修改失败！")
		return code, err
	}
	return code, err
}

//日程删除
func DeleteSchedule(scheduleId int) (code int, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("schedule").Filter("id", scheduleId).Delete()
	if err != nil {
		code = 1003
		err = errors.New("删除失败！")
		return code, err
	}
	return code, err
}
