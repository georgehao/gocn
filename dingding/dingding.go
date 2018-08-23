package dingding

import (
	"fmt"
	"github.com/georgehao/gocn/db"
	"github.com/georgehao/gocn/message"
	"strings"
	"time"
)

func buildMessage(msg message.Message) string {
	str := fmt.Sprintf("### %s\n", msg.DailyTitle)
	for _, v := range msg.TextUrls {
		if len(v.Url) == 0 || len(v.Text) == 0 {
			continue
		}

		if strings.Contains(v.Text, "招聘") ||
			strings.Contains(v.Text, "Gopher meetup") ||
			strings.Contains(v.Text, "Gopher Meetup") ||
			strings.Contains(v.Text, "GoCN归档") ||
			strings.Contains(v.Text, "订阅新闻") {
			continue
		}
		str += fmt.Sprintf("- [%s](%s)\n", v.Text, v.Url)
	}

	str += fmt.Sprintf("\n归档：%s\n", "")
	return str
}

func Send() {
	token := ""
	ding := Ding{AccessToken: token}

	for {
		msg, err := message.Pop()
		if err != nil {
			continue
			time.Sleep(time.Second * 10)
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
