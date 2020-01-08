package main

import (
	"html/template"
	"net/http"

	"qnsoft/web_api/controllers"
	jobs "qnsoft/web_api/controllers/Jobs"
	_ "qnsoft/web_api/routers"

	"github.com/astaxie/beego"
)

func main() {
	//初始化信用卡智能代还任务计划
	jobs.InitJobs()
	//初始化淘宝客任务计划
	controllers.Taobao_joblist()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("401", page_note_permission)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles("views/404.html")
	data := make(map[string]interface{})
	//data["content"] = "page not found"
	t.Execute(rw, data)
}

func page_note_permission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.html").ParseFiles("views/401.html")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
