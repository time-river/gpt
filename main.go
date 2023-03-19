package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"scancenter/cronTask"
	"scancenter/route"
	"github.com/robfig/cron"
)


//main函数作为主入口，后续可以通过选择编译来判断进行不同模式的启动
func main(){
	startWebApp()

}

//执行app方式
func startWebApp(){
	router := mux.NewRouter()
	route.SetRoute(router)
	http.ListenAndServe(":8080",router)
}

//执行进程模式
func startDaemon(){

}

//执行cron模式
func startCron(){
	c := cron.New()
	c.AddFunc("*/5 * * * * ?",func(){
		cronTask.Test()})
	c.Start()
	select {}
}




