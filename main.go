package main

import (
	"github.com/robfig/cron/v3"
	"monitor/core/handler"
	"monitor/core/task"
	"net/http"
)

func main() {

	// 定时任务
	c := cron.New()
	_, _ = c.AddFunc("*/1 * * * *", task.UpdateUserOperation)
	_, _ = c.AddFunc("*/15 * * * *", task.UpdatePVUV)
	c.Start()
	defer c.Stop()

	// 网页路由
	http.HandleFunc("/", handler.HelloHandler)

	err := http.ListenAndServe(":8004", nil)
	if err != nil {
		panic("Failed to start server , err: " + err.Error())
	}
}
