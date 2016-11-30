package controllers

import (
	"github.com/astaxie/beego/logs"
	"appzone/models"
	"appzone/utils"
	"github.com/astaxie/beego"
)
const sql = "SELECT * FROM channel_params WHERE id = ?"
type IndexController struct {
	beego.Controller
}
func (c *IndexController) Get() {
	logs.Info("开始了")
	var user models.Channel_Params
	err := utils.GetBean(sql, &user, "fengdaitest001")
	if err != nil {
		logs.Error("数据获取失败")

	}
	c.Data["json"]=utils.ToJson(user)
	c.ServeJSON()
}
