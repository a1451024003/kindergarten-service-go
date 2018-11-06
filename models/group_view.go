package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type GroupView struct {
	Id        int       `json:"id";orm:"column(id);auto"`
	UserId    int       `json:"user_id";orm:"column(user_id)"`
	ClassType int8      `json:"class_type";orm:"column(class_type)"`
	ClassId   int       `json:"class_id";orm:"column(class_id)"`
	CreatedAt time.Time `json:"created_at";orm:"auto_now"`
}

func (t *GroupView) TableName() string {
	return "group_view"
}

func init() {
	orm.RegisterModel(new(GroupView))
}
