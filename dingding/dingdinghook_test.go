package dingding

import (
	"fmt"
	"testing"
	"time"
)

func _TestSend(t *testing.T) {
	ding := Ding{AccessToken: "b123b96dbc43937b653ec3e005612d50b504f28773035f0194469d28cf244ef7"}
	msg := Message{Content: "测试"}
	result := ding.Send(msg)
	fmt.Println(result)

	link := Link{Title: "link测试", Content: "测试", ContentURL: "https://www.baidu.com"}
	result = ding.Send(link)
	fmt.Println(result)

	markdown := Markdown{Title: "markdown测试", Content: "#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"}
	result = ding.Send(markdown)
	fmt.Println(result)
}

func _TestDingQueue(t *testing.T) {
	ding := &DingQueue{Title: "queue测试", Interval: 3, AccessToken: "b123b96dbc43937b653ec3e005612d50b504f28773035f0194469d28cf244ef7"}
	ding.Init()

	go ding.Start()

	ding.Push("#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	time.Sleep(time.Second * 5)
	ding.Push("#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	ding.Push("#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")

	time.Sleep(time.Second * 10)

	ding.Push("#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")
	ding.Push("#### 杭州天气\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](http://image.jpg)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n")

	time.Sleep(time.Second * 10)
}
