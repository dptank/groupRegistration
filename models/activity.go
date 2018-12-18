package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/validation"
	"errors"
	"fmt"
)

type PinActivity struct {
	Id int `orm:"column(id);pk"`
	Title string `orm:"column(title)" valid:"Required"`
	JoinCount int `orm:"column(join_count)" valid:"Required"`
	OwnerPrice int `orm:"column(owner_price)" `
	MemberPrice int `orm:"column(member_price)"`
	PriceType int `orm:"column(price_type)"`
	StartTime int64 `orm:"column(start_time)"`
	EndTime int64 `orm:"column(end_time)"`
	Status int `orm:"column(status)"`
	Img string `orm:"column(img)"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" `
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func (pa *PinActivity) TableName() string {
	return TableName("pin_activity")
}
//入库
func (pa *PinActivity) PinActivityAdd() error{
	if err := checkPinActivity(pa); err != nil {
		return  err
	}
	if _, err := orm.NewOrm().Insert(pa); err != nil {
		return err
	}
	return nil
}
func checkPinActivity(pa *PinActivity) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(pa)
	fmt.Println(b)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

