package models

import "github.com/astaxie/beego/orm"

type Channel_Params struct{
	Id string `orm:"pk"`
	ChannelId string `orm:"column(channelId)"`
	ChannelCode string `orm:"column(channelCode)"`
	ParamKey string `orm:"column(paramKey)"`
	ParamValue string `orm:"column(paramValue)"`
	Description string
	Isdeleted int16
}

func init() {
	orm.RegisterModel(new(Channel_Params))
}