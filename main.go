package main

import (
	"github.com/georgehao/gocn/db"
	"github.com/georgehao/gocn/dingding"
	"github.com/georgehao/gocn/splider"
	"time"
)

func main() {
	go db.Run()
	go dingding.Send()

	splider.Run()

	time.Sleep(time.Minute * 1)
}
