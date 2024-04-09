package web

import (
	"VDController/docker"
	"VDController/logger"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Web端日志记录器
var wLogger *logger.Logger

func vdIndex(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	tmpl, err := template.ParseFiles("./web/template/index.tmpl")
	if err != nil {
		wLogger.Log(logger.ERROR, "创建网页模版失败")
		fmt.Println("create template failed, err:", err)
		return
	}
	envInfo := docker.GetEnvInfo()
	tmpl.Execute(w, envInfo)
}

func StartWeb() {
	wLogger := logger.NewLogger(logger.INFO)
	wLogger.Log(logger.INFO, "启动Web程序")
	http.HandleFunc("/", vdIndex)            // 设置访问的路由
	err := http.ListenAndServe(":8080", nil) // 设置监听的端口
	if err != nil {
		wLogger.Log(logger.ERROR, "创建监听端口失败")
		log.Fatal("ListenAndServe: ", err)
	}
}
