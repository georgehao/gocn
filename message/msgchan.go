package message

import "errors"

// TextUrl 每日新闻内容
type TextUrl struct {
	Text string
	Url  string
}

type Message struct {
	TextUrls   []TextUrl
	DailyTitle string
	Author     string // 编辑
	PostUrl    string // 原文地址
}

var messageChain chan Message

func Push(m Message) {
	messageChain <- m
}

func Pop() (Message, error) {
	select {
	case m := <-messageChain:
		return m, nil
	default:
		return Message{}, errors.New("nil")
	}
}

func init() {
	messageChain = make(chan Message, 10000)
}
