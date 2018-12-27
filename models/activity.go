package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/validation"
	"errors"
)
//数据库字段映射表
type PinActivity struct {
	Id int `orm:"column(id);pk"`
	Title string `orm:"column(title)" valid:"Required"`
	JoinCount int `orm:"column(join_count)" valid:"Required"`
	OwnerPrice int `orm:"column(owner_price)" valid:"Required"`
	MemberPrice int `orm:"column(member_price)" valid:"Required"`
	PriceType int `orm:"column(price_type)"`
	StartTime int64 `orm:"column(start_time)" valid:"Required"`
	EndTime int64 `orm:"column(end_time)" valid:"Required"`
	Status int `orm:"column(status)"`
	Img string `orm:"column(img)" valid:"Required"`
	CreatedAt time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func (pa *PinActivity) TableName() string {
	return TableName("pin_activity")
}
//入库
func (pa *PinActivity) PinActivityAdd() error{
	if err := CheckPinActivity(pa); err != nil {
		return  err
	}
	if _, err := orm.NewOrm().Insert(pa); err != nil {
		return err
	}
	return nil
}
/**
检测数据
*/
func CheckPinActivity(pa *PinActivity) (err error) {
	valid := validation.Validation{}
	b,_ := valid.Valid(pa)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Field + " " +err.Message)
		}
	}
	return nil
}
/**
根据id获取详情
*/
func (pa *PinActivity) GetActivityInfo(activityId int)(*PinActivity,error) {
	err := orm.NewOrm().QueryTable(TableName("pin_activity")).Filter("id",activityId).One(pa)
	if err !=nil {
		return nil,err
	}
	return pa ,nil
}
/**
根据id修改内容
*/
func (pa *PinActivity) UpdateActivityInfo(fields ...string) error {
	if err := CheckPinActivity(pa); err != nil {
		return  err
	}
	_,err := orm.NewOrm().Update(pa,fields...)
	if err != nil {
		return err
	}
	return nil
}
/**
获取列表
*/
func (pa *PinActivity) GetActivityList(page,offset, pageSize int, filters ...interface{})([]*PinActivity,int64) {
	list := make([]*PinActivity,0)
	query := orm.NewOrm().QueryTable(TableName("pin_activity"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	//fmt.Println(list)
	return list, total
}

