package models

import (
	"math"
	"time"

	"github.com/astaxie/beego/orm"
)

type SlideShow struct {
	Id             int       `json:"id" orm:"column(id);auto"`
	Title          string    `json:"title" orm:"column(Title);" description:"标题"`
	KindergartenId int       `json:"kindergarten_id" orm:"column(kindergarten_id)"`
	Content        string    `json:"content" orm:"column(Content)" description:"内容"`
	Picture        string    `json:"picture" orm:"column(picture)" description:"图片"`
	CreatedAt      time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt      time.Time `json:"update_at" orm:"auto_now"`
}

func (t *SlideShow) TableName() string {
	return "slide_show"
}

func init() {
	orm.RegisterModel(new(SlideShow))
}

/*
添加轮播图
*/
func AddSlideShow(title string, content string, kindergarten_id int, picture string) (err error) {
	o := orm.NewOrm()
	m := SlideShow{Title: title, Content: content, KindergartenId: kindergarten_id, Picture: picture}
	_, err = o.Insert(&m)
	return err
}

/*
轮播图列表
*/
func GetSlideShowList(page int, prepage int, kindergarten_id int) (paginatorMap map[string]interface{}, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	sql := qb.Select("count(*)").From("slide_show").Where("kindergarten_id = ?").String()
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
		sql := qb.Select("*").From("slide_show").Where("kindergarten_id = ?").Limit(prepage).Offset(limit).String()
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
轮播图详情
*/
func GetSlideShow(id int) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	var v []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("slide_show").Where("id = ?").String()
	_, err = o.Raw(sql, id).Values(&v)
	if err == nil {
		paginatorMap := make(map[string]interface{})
		paginatorMap["data"] = v
		return paginatorMap, nil
	}
	return nil, err
}

/*
删除轮播图
*/
func DeleteSlideShow(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("slide_show").Filter("id", id).Delete()
	return err
}

/*
编辑轮播图
*/
func UpdateSlideShow(id int, title string, content string, kindergarten_id int, picture string) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("slide_show").Filter("id", id).Update(orm.Params{
		"Content": content, "Title": title, "KindergartenId": kindergarten_id, "Picture": picture,
	})
	return err
}
