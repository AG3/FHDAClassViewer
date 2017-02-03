package controllers

import (
	"BetterClassViewer/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Get() {
	c.TplName = "htmltest.html"
}

// @router /course [get]
func (c *MainController) GetCourse() {
	sub := c.GetString("subject")
	tmp := models.GetCourse(sub)
	res := strings.Join(tmp, "|")
	c.Data["json"] = res
	c.ServeJSON()
}

// @route /subject [get]
func (c *MainController) GetSubject() {
	res := strings.Join(models.Sub_name, "|")
	c.Data["json"] = res
	c.ServeJSON()
}

// @route /class [get]
func (c *MainController) GetClass() {
	sub := c.GetString("subject")
	crse := c.GetString("course")
	inds := models.GetClass(sub, crse)
	var tmp_s []string
	for i := range inds {
		tmp, _ := json.Marshal(models.Class[i])
		tmp_s = append(tmp_s, string(tmp))
	}
	res := strings.Join(tmp_s, "|")
	c.Data["json"] = res
	c.ServeJSON()
}

// @router /admin/update [get]
func (c *MainController) UpdateDatas() {
	usrname := c.GetString("whosyourdaddy")
	if usrname == "Pikabbit" {
		models.UpdateData()
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = "Failed"
	}
	c.ServeJSON()
}
