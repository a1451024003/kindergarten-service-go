package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type KindergartenCertificate struct {
	Id                   int       `json:"kindergarten_id" orm:"column(kindergarten_id);auto" description:"编号"`
	TaxRegistration      string    `json:"tax_registration" orm:"column(tax_registration);size(255)" description:"税务登记证"`
	CateringServices     string    `json:"catering_services" orm:"column(catering_services);size(255)" description:"餐饮服务许可证"`
	PrivateNonEnterprise string    `json:"private_non_enterprise" orm:"column(private_non_enterprise);size(255)" description:"民办非企业单位登记证书"`
	CreatedAt            time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt            time.Time `json:"updated_at" orm:"auto_now"`
}

func (t *KindergartenCertificate) TableName() string {
	return "kindergarten_certificate"
}

func init() {
	orm.RegisterModel(new(KindergartenCertificate))
}
