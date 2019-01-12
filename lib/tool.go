package lib

import (
	"github.com/astaxie/beego/validation"
	"errors"
)
/**
数据校验
*/
func CheckData(pa interface{}) (err error) {
	valid := validation.Validation{}
	b,_ := valid.Valid(pa)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Field + " " +err.Message)
		}
	}
	return nil
}