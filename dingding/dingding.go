package dingding

import (
	"fmt"
	"gocn/config"
	"gocn/db"
	"gocn/message"
	"strings"
	"time"
)

func buildMessage(msg message.Message) string {
	str := fmt.Sprintf("## %s", msg.DailyTitle)
	str += fmt.Sprintln()
	for _, v := range msg.TextUrls {
		if len(v.Url) == 0 || len(v.Text) == 0 {
			continue
		}

		if strings.Contains(v.Text, "GoCN归档") || strings.Contains(v.Text, "订阅新闻") {
			continue
		}

		textValue := strings.Replace(v.Text, "\n", "", -1)
		realText := strings.Replace(textValue, " ", "", -1)

		str += fmt.Sprintf("- [%s](%s)", realText, v.Url)
		str += fmt.Sprintln()
	}

	index := strings.Index(msg.Author, "订阅新闻")
	author := msg.Author
	if index > 0 {
		author = msg.Author[:index]
	}

	str += fmt.Sprintln()
	str += fmt.Sprintf("编辑：%s", author)
	str += fmt.Sprintln()
	str += fmt.Sprintln()
	str += fmt.Sprintf("原文地址: %s", msg.PostUrl)
	return str
}

func Send() {
	ding := Ding{AccessToken: config.Config.GetString("dingding.token")}

	for {
		msg, err := message.Pop()
		if err != nil {
			continue
		}

		if !db.CheckSend(msg.DailyTitle) {
			db.Push(msg.DailyTitle)
			content := buildMessage(msg)
			db.PushMarkdown(content)
			markdown := Markdown{Title: "GoCN每日新闻", Content: content}
			ding.Send(markdown)
			time.Sleep(time.Second * 1)
		}
	}
}
