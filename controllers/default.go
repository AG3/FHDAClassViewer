package controllers

import (
	"BetterClassViewer/models"
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"strings"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Get() {
	//rawip := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	//ip := rawip[0]
	//beego.Notice("From:", ip, ", Visit")
	c.TplName = "index.html"
}

// @router /course [get]
func (c *MainController) GetCourse() {
	/*rawip := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	ip := rawip[0]

	beego.Notice("From:", ip, ", Request", sub)*/
	sub := c.GetString("subject")
	tmp := models.GetCourse(sub)
	res := strings.Join(tmp, "|")
	c.Data["json"] = res
	c.ServeJSON()
}

// @router /subjects [get]
func (c *MainController) GetSubject() {
	res := strings.Join(models.Sub_name, "|")
	c.Data["json"] = res
	c.ServeJSON()
}

// @router /class [get]
func (c *MainController) GetClass() {
	/*rawip := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	ip := rawip[0]

	beego.Notice("From:", ip, ", Request", sub, crse)*/
	sub := c.GetString("subject")
	crse := c.GetString("course")
	inds := models.GetClass(sub, crse)
	var tmp_s []string
	for i := range inds {
		tmp, _ := json.Marshal(models.Class[inds[i]])
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
