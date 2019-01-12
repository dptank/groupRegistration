package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"groupRegistration/lib"
)

type  PinEvent struct {
	Id int `orm:"column(id);pk"`
	PinActivityId int `orm:"column(pin_activity_id)" valid:"Required"`
	JoinCount int `orm:"column(join_count)" valid:"Required"`
	CountLimit int `orm:"column(count_limit)" valid:"Required"`
	Owner int `orm:"column(owner)" valid:"Required"`
	EndTime int64 `orm:"column(end_time)" valid:"Required"`
	Status int `orm:"column(status)"`
	CreatedAt time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func (pa *PinEvent) TableName() string {
	return TableName("pin_event")
}
/**
插入拼团事件表
*/
func (pa *PinEvent) AddPinEvent() error{
	if err := lib.CheckData(pa) ;err !=nil{
		return  err
	}
	if _,err := orm.NewOrm().Insert(pa);err!=nil {
		return err
	}
	return nil
}
