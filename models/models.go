package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type ClassRecord struct {
	Crn            string `json:"crn"`
	Subject        string `json:"subject"`
	Course         string `json:"course"`
	Section        string `json:"section"`
	Credit         string `json:"credit"`
	Title          string `json:"title"`
	Days           string `json:"days"`
	Start_time     string `json:"start_time"`
	End_time       string `json:"end_time"`
	Instructor     string `json:"instructor"`
	Start_date     string `json:"start_date"`
	End_date       string `json:"end_date"`
	Location       string `json:"location"`
	Online_content string `json:"online_content"`
}

var Class []ClassRecord
var subject map[string]int
var Sub_name []string

type node struct {
	Name2Int map[string]int
	Children []int
}

var nodes [3000]node

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./data.db")
	subject = make(map[string]int)
	UpdateData()
	Sub_name = make([]string, 0, len(subject))
	for k := range subject {
		Sub_name = append(Sub_name, k)
	}
}

func UpdateData() {
	o := orm.NewOrm()
	o.Using("default")

	num_classes, err := o.Raw("SELECT * FROM class").QueryRows(&Class)
	checkErr(err)

	var i int64
	tot_p := 0
	for i = 0; i < num_classes; i++ {
		sub := Class[i].Subject
		ind, ok := subject[sub]

		if ok == false {
			ind = tot_p
			subject[sub] = ind
			nodes[ind].Name2Int = make(map[string]int)
			tot_p++
		}

		crse := Class[i].Course
		ind_c, ok := nodes[ind].Name2Int[crse]
		if ok == false {
			ind_c = tot_p
			nodes[ind_c].Name2Int = make(map[string]int)
			nodes[ind].Name2Int[crse] = ind_c
			tot_p++
		}
		//fmt.Printf("add rec %d to %d\n", int(i), ind_c)
		nodes[ind_c].Children = append(nodes[ind_c].Children, int(i))
	}
}

func GetClass(sub string, crse string) []int {
	ind := subject[sub]
	ind_c := nodes[ind].Name2Int[crse]
	return nodes[ind_c].Children
}

func GetCourse(sub string) []string {
	res := make([]string, 0, len(nodes[subject[sub]].Name2Int))
	for e := range nodes[subject[sub]].Name2Int {
		//fmt.Println(e)
		res = append(res, e)
	}
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
