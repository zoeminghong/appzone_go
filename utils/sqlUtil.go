package utils

import (
	"github.com/astaxie/beego/orm"
)
/**
	获取Model数据
 */
func GetBean(sql string, m interface{}, v ...interface{}) error {
	o := orm.NewOrm()
	err := o.Raw(sql, v...).QueryRow(m)
	return err
}
