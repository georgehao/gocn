package main

import (
	"gocn/db"
	"gocn/dingding"
	"gocn/splider"
	"time"
)

func main() {
	go db.Run()
	go dingding.Send()

	splider.Run()

	// 当使用定时任务启动时，使用这里，等待输入写入文件完成
	time.Sleep(time.Minute * 1)

	// 当成常驻进程启动是，把这里打开
	// select {}
}
