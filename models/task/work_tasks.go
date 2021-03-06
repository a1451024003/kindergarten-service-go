package task

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/astaxie/beego"
	"strconv"
)

type WorkTasks struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Describe       string    `json:"describe"`
	Deadline       time.Time `json:"deadline"`
	SaveFolderId   int       `json:"save_folder_id"`
	SaveFolderName string    `json:"save_folder_name"`
	Publisher      int       `json:"publisher"`
	PublisherName  string    `json:"publisher_name"`
	TaskNum        int       `json:"task_num"`
	FinishNum      int       `json:"finish_num"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt      time.Time `json:"updated_at" orm:"auto_now"`
}

type WorkTasksOperator struct {
	Id             int       `json:"id"`
	Operator       int       `json:"operator"`
	OperatorName   string    `json:"operator_name"`
	OperatorAvatar string    `json:"operator_avatar"`
	CoursewareId   string    `json:"courseware_id"`
	CoursewareName string    `json:"courseware_name"`
	CoursewareThumb string    `json:"courseware_thumb"`
	UploadTime     time.Time `json:"upload_time" orm:"auto_now_add"`
	WorkTasksId    int       `json:"work_tasks_id"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt      time.Time `json:"updated_at" orm:"auto_now"`
}

type WorkTasksCc struct {
	Id          int       `json:"id"`
	Cc          int       `json:"cc"`
	CcName      string    `json:"cc_name"`
	WorkTasksId int       `json:"work_tasks_id"`
	CreatedAt   time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt   time.Time `json:"updated_at" orm:"auto_now"`
}

type NoticeServer struct {
	AddWork func(title map[string]interface{}, name string, urlId int, workNotice map[string]interface{}, noticeType string) error
}

func (wt *WorkTasks) TableName() string {
	return "work_tasks"
}

func (wt *WorkTasksOperator) TableName() string {
	return "work_tasks_operator"
}

func (wt *WorkTasksCc) TableName() string {
	return "work_tasks_cc"
}

func init() {
	orm.RegisterModel(new(WorkTasks), new(WorkTasksOperator), new(WorkTasksCc))
}

func (wt *WorkTasks) Save(operator, cc []map[string]interface{}) error {
	o := orm.NewOrm()
	o.Begin()

	if _, err := o.Insert(wt); err != nil {
		o.Rollback()

		return err
	}

	noticeTitle := map[string]interface{}{
		"operator":fmt.Sprintf(`你有一个任务"%s"需要执行`, wt.Title),
		"cc":fmt.Sprintf(`你有一个任务"%s"`, wt.Title),
	}

	workNotice := map[string]interface{}{}

	var wtos []WorkTasksOperator
	var oIds []int
	for _, value := range operator {
		recipientId := int(value["id"].(float64))
		wtr := WorkTasksOperator{
			Operator:       recipientId,
			OperatorName:   value["name"].(string),
			OperatorAvatar: value["avatar"].(string),
			WorkTasksId:    wt.Id,
		}
		wtos = append(wtos, wtr)
		oIds = append(oIds, recipientId)
	}
	if _, err := o.InsertMulti(len(wtos), wtos); err != nil {
		o.Rollback()

		return err
	}
	workNotice["operator"] = oIds

	var cIds []int
	if len(cc) != 0 {
		var wtcs []WorkTasksCc
		for _, value := range cc {
			ccId := int(value["id"].(float64))
			wtc := WorkTasksCc{Cc: ccId, CcName: value["name"].(string), WorkTasksId: wt.Id}
			wtcs = append(wtcs, wtc)
			cIds = append(cIds, ccId)
		}

		if _, err := o.InsertMulti(len(wtcs), wtcs); err != nil {
			o.Rollback()

			return err
		}
	}
	workNotice["cc"] = cIds
	o.Commit()
	var noticeServer *NoticeServer
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&noticeServer)
	err := noticeServer.AddWork(noticeTitle, wt.PublisherName, wt.Id, workNotice, "4")
	fmt.Println(err)

	return nil
}

func (wt *WorkTasks) Get(uId int) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	var tasks []*WorkTasks

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("wt.*").
		From("work_tasks as wt").
		LeftJoin("work_tasks_cc as wtc").On("wt.id = wtc.work_tasks_id").
		LeftJoin("work_tasks_operator as wto").On("wt.id = wto.work_tasks_id").
		Where("(wt.publisher = ? or wto.operator = ? or wtc.cc = ?) and wt.status <> 2").
		GroupBy("wt.id").OrderBy("wt.id").Desc().String()

	if num, err := o.Raw(sql, uId, uId, uId).QueryRows(&tasks); err != nil {
		return res, err
	} else if num <= 0 {
		return res, err
	}

	var taskIds []int
	for _, value := range tasks {
		taskIds = append(taskIds, value.Id)
	}

	var operators []WorkTasksOperator
	o.QueryTable(new(WorkTasksOperator)).Filter("work_tasks_id__in", taskIds).All(&operators)

	var ccs []WorkTasksCc
	o.QueryTable(new(WorkTasksCc)).Filter("work_tasks_id__in", taskIds).All(&ccs)

	for _, value := range tasks {
		var op []WorkTasksOperator
		for _, ov := range operators {
			if ov.WorkTasksId == value.Id {
				op = append(op, ov)
			}
		}
		var oc []WorkTasksCc
		for _, cv := range ccs {
			if cv.WorkTasksId == value.Id {
				oc = append(oc, cv)
			}
		}
		var maps map[string]interface{}
		jsons, _ := json.Marshal(value)
		json.Unmarshal(jsons, &maps)
		maps["deadline"] = value.Deadline.Format("2006-01-02 15:04:05")
		maps["operator"] = op
		maps["cc"] = oc
		res = append(res, maps)
	}

	return res, nil
}

func (wt *WorkTasks) GetInfoById() (map[string]interface{}, error) {
	var res map[string]interface{}
	o := orm.NewOrm()

	if err := o.QueryTable(wt).Filter("id", wt.Id).Exclude("status", 2).One(wt); err != nil {
		return res, err
	}
	jsons, _ := json.Marshal(wt)
	json.Unmarshal(jsons, &res)
	res["deadline"] = wt.Deadline.Format("2006-01-02 15:04:05")

	var wto []WorkTasksOperator
	if _, err := o.QueryTable(new(WorkTasksOperator)).Filter("work_tasks_id", wt.Id).All(&wto); err != nil {
		return res, err
	}
	res["operator"] = wto

	var wtc []WorkTasksCc
	if _, err := o.QueryTable(new(WorkTasksCc)).Filter("work_tasks_id", wt.Id).All(&wtc); err != nil {
		return res, err
	}
	res["cc"] = wtc

	return res, nil
}

type CoursewareServer struct {
	CopyCourseware_kg func (int, int, int, string, string) (string, error)
} 

func (wt *WorkTasks) Complete(operator int, coursewareId, coursewareName, coursewareThumb string) error {
	o := orm.NewOrm()
	tempCoursewareId := fmt.Sprintf("%s%s", strconv.Itoa(operator), strconv.FormatInt(time.Now().Unix(), 10))

	wto := WorkTasksOperator{Operator: operator, WorkTasksId: wt.Id}
	if err := o.Read(&wto, "operator", "work_tasks_id"); err != nil {
		return err
	}
	if wto.Status != 0 {
		return errors.New("completed")
	}
	if coursewareId != "" {
		wto.CoursewareId = tempCoursewareId
		wto.CoursewareName = coursewareName
		wto.CoursewareThumb = coursewareThumb
		wto.UploadTime = time.Now()
	}
	wto.Status = 1
	o.Begin()

	if _, err := o.Update(&wto); err != nil {
		o.Rollback()

		return err
	}

	if _, err := o.QueryTable(wt.TableName()).Filter("id", wt.Id).Update(orm.Params{
		"finish_num": orm.ColValue(orm.ColAdd, 1),
		"updated_at": time.Now(),
	}); err != nil {
		o.Rollback()

		return err
	}

	if err := o.Read(wt); err != nil {
		o.Rollback()

		return err
	}
	if wt.Status != 0 {
		o.Rollback()

		return errors.New("exceed")
	}
	if wt.TaskNum == wt.FinishNum {
		wt.Status = 1
		if _, err := o.Update(wt, "status"); err != nil {
			o.Rollback()

			return err
		}
	} else if wt.FinishNum > wt.TaskNum {
		o.Rollback()

		return errors.New("exceed")
	}
	if coursewareId != "" {
		var coursewareServer *CoursewareServer
		clientC := rpc.NewHTTPClient(beego.AppConfig.String("COURSEWARE_SERVER"))
		clientC.UseService(&coursewareServer)
		if _, err := coursewareServer.CopyCourseware_kg(0, wt.SaveFolderId, operator, coursewareId, tempCoursewareId); err != nil {
			o.Rollback()

			return err
		}
	}
	o.Commit()

	var wtc []WorkTasksCc
	cIds := []int{wt.Publisher}
	if num, err := o.QueryTable(new(WorkTasksCc)).Filter("work_tasks_id", wt.Id).All(&wtc); err == nil && num > 0 {
		for _, val := range wtc {
			cIds = append(cIds, val.Cc)
		}
	}
	noticeTitle := map[string]interface{}{
		"operator":"",
		"cc":fmt.Sprintf(`%s已完成"%s"任务`, wto.OperatorName, wt.Title),
	}
	workNotice := map[string]interface{}{"operator":[]int{}, "cc":cIds}
	var noticeServer *NoticeServer
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_NOTICE_SERVER"))
	client.UseService(&noticeServer)
	err := noticeServer.AddWork(noticeTitle, wt.PublisherName, wt.Id, workNotice, "4")
	fmt.Println(err)

	return nil
}

func (wto *WorkTasksOperator) Schedule() ([]WorkTasksOperator, error) {
	var wtos []WorkTasksOperator
	_, err := orm.NewOrm().QueryTable(wto).Filter("work_tasks_id", wto.WorkTasksId).All(&wtos)

	return wtos, err
}

func (wt *WorkTasks) Delete() error {
	o := orm.NewOrm()

	if err := o.Read(wt); err != nil {
		return err
	}

	wt.Status = 2
	if _, err := o.Update(wt, "status"); err != nil {
		return err
	}

	return nil
}

func (wt *WorkTasks) Finish() error {
	o := orm.NewOrm()

	if err := o.Read(wt); err != nil {
		return err
	}

	wt.Status = 3
	if _, err := o.Update(wt, "status"); err != nil {
		return err
	}

	return nil
}
