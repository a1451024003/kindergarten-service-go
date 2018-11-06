package models

import (
	"math"
	"time"

	"github.com/astaxie/beego/orm"
)

type Notice struct {
	Id             int       `json:"id" orm:"column(id);auto;"`
	Title          string    `json:"title" orm:"column(title);size(50)"; description:"标题"`
	Content        string    `json:"content" orm:"column(content);size(255)" description:"公告内容"`
	KindergartenId int       `json:"kindergarten_id" orm:"column(kindergarten_id)";description:"幼儿园ID"`
	CreatedAt      time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt      time.Time `json:"updated_at" orm:"auto_now"`
}

func (t *Notice) TableName() string {
	return "notice"
}

func init() {
	orm.RegisterModel(new(Notice))
}

/*
添加公告
*/
func AddNotice(title string, content string, kindergarten_id int) (err error) {
	o := orm.NewOrm()
	m := Notice{Title: title, Content: content, KindergartenId: kindergarten_id}
	_, err = o.Insert(&m)
	return err
}

/*
公告列表
*/
func GetNoticeList(page int, prepage int, kindergarten_id int) (ml map[string]interface{}, err error) {
	var v []Notice
	o := orm.NewOrm()
	nums, err := o.QueryTable("notice").Filter("kindergarten_id", kindergarten_id).All(&v)
	if err == nil && nums > 0 {
		//根据nums总数，和prepage每页数量 生成分页总数
		totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
		if page > totalpages {
			page = totalpages
		}
		if page <= 0 {
			page = 1
		}
		limit := (page - 1) * prepage
		num, err := o.QueryTable("notice").Filter("kindergarten_id", kindergarten_id).Limit(prepage, limit).All(&v)
		if err == nil && num > 0 {
			paginatorMap := make(map[string]interface{})
			paginatorMap["total"] = nums          //总条数
			paginatorMap["data"] = v              //分页数据
			paginatorMap["page_num"] = totalpages //总页数
			return paginatorMap, nil
		}
	}
	return nil, err

}

/*
Web -公告详情
*/
func GetNoticeInfo(id int) (ml map[string]interface{}, err error) {
	var v []Notice
	o := orm.NewOrm()
	err = o.QueryTable("notice").Filter("Id", id).One(&v)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}

/*
删除公告
*/
func DeleteNotice(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("notice").Filter("id", id).Delete()
	return err
}

/*
编辑公告
*/
func UpdateNotice(id int, title string, content string, kindergarten_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("notice").Filter("id", id).Update(orm.Params{
		"title": title, "content": content, "kindergarten_id": kindergarten_id,
	})
	return err
}
