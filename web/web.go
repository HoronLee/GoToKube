package web

import (
	"VDController/docker"
	"VDController/logger"
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	// Web端日志记录器
	wLogger *logger.Logger
	// 全局 Web 对象
	server *http.Server
)

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
	wLogger = logger.NewLogger(logger.INFO)
	wLogger.Log(logger.INFO, "启动Web程序")
	http.HandleFunc("/", vdIndex) // 设置访问的路由
	server = &http.Server{Addr: ":8080"}
	// http 携程
	go func() {
		if err := server.ListenAndServe(); err != nil {
			wLogger.Log(logger.ERROR, "创建监听端口失败")
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

// 未使用
func StopWeb() {
	if server != nil {
		wLogger.Log(logger.INFO, "关闭Web程序")
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Fatal("Server shutdown error: ", err)
		}
	}
}
