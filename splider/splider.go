package splider

import (
	"fmt"
	"gocn/config"
	"gocn/message"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var c *colly.Collector

func Run() {
	url := ""
	// On every a element which has href attribute call callback
	c.OnHTML("div[class=\"aw-mod aw-question-detail aw-item\"]", func(e *colly.HTMLElement) {
		dailyTitle := e.ChildText("h1")
		var textUrls []message.TextUrl
		var author string
		e.ForEach("div[class=\"content markitup-box\"]", func(i int, e *colly.HTMLElement) {
			e.ForEach("li", func(i int, e *colly.HTMLElement) {
				url := e.ChildText("a[href]")
				urlIndex := strings.Index(e.Text, url)
				if urlIndex <= len(e.Text) && urlIndex != -1 {
					//println(urlIndex, len(e.Text), e.Text)
					text := e.Text[0:urlIndex]
					textUrls = append(textUrls, message.TextUrl{text, url})
				}
			})

			// div
			e.ForEach("div", func(i int, element *colly.HTMLElement) {
				text := element.Text
				if text != "" && !strings.Contains(text, "GoCN每日新闻(") {
					httpIndex := strings.Index(text, "http")
					if httpIndex <= len(element.Text) && httpIndex != -1 {
						text := element.Text[0:httpIndex]
						url := element.Text[httpIndex:]
						textUrls = append(textUrls, message.TextUrl{text, url})
					}
				}
			})

			// p
			e.ForEach("p", func(i int, element *colly.HTMLElement) {
				url := element.ChildText("a[href]")
				text := element.Text
				if text != "" && !strings.Contains(text, "GoCN每日新闻(") && !strings.Contains(text, "p.p1 {margin: 0.0px 0.0px 0.0px 0.0px; line-height") {
					// 处理text带的url
					urlIndex := strings.Index(text, url)
					if urlIndex <= len(element.Text) && urlIndex != -1 {
						text := element.Text[0:urlIndex]
						textUrls = append(textUrls, message.TextUrl{text, url})
					}
				}
			})

			e.ForEach("*", func(i int, element *colly.HTMLElement) {
				authorIndex1 := strings.Index(element.Text, "编辑:")
				authorIndex2 := strings.Index(element.Text, "编辑：")

				index := -1
				authorLen := -1
				if authorIndex1 >= 0 {
					index = authorIndex1
					authorLen = len("编辑:")
				} else if authorIndex2 >= 0 {
					index = authorIndex2
					authorLen = len("编辑：")
				}

				if index >= 0 {
					author = element.Text[index+authorLen:]
				}
			})
		})

		var newTextUrls []message.TextUrl
		uniMap := make(map[string]bool)
		for _, v := range textUrls {
			if _, ok := uniMap[v.Url]; !ok {
				uniMap[v.Url] = true
				newTextUrls = append(newTextUrls, v)
			}
		}

		message.Push(message.Message{
			DailyTitle: dailyTitle,
			TextUrls:   newTextUrls,
			Author:     author,
			PostUrl:    e.Request.URL.String(),
		})
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if /*e.Text == ">" ||*/ strings.Contains(e.Text, "GoCN每日新闻") {
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
			_ = c.Visit(e.Request.AbsoluteURL(link))
			url = e.Request.AbsoluteURL(link)
		}
	})

	if config.Config.GetBool("splider.all") {
		c.OnHTML("div[class=\"page-control\"]", func(element *colly.HTMLElement) {
			isHasNext := false
			element.ForEach("a[href]", func(i int, e *colly.HTMLElement) {
				link := e.Attr("href")
				if link == "javascript:;" {
					isHasNext = true
				}

				if isHasNext == true && link != "javascript:;" {
					_ = c.Visit(e.Request.AbsoluteURL(link))
					isHasNext = false
				}
			})
		})
	}

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	_ = c.Visit("https://gocn.vip/explore/category-14")
}

func init() {
	// Instantiate default collector
	c = colly.NewCollector(
		colly.AllowedDomains("gocn.vip"),
		colly.MaxDepth(1),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	parallelism := config.Config.GetInt("splider.parallelism")
	delay := time.Duration(config.Config.GetInt("splider.delay")) * time.Second

	c.Limit(&colly.LimitRule{
		DomainGlob:  "gocn.vip",
		Parallelism: parallelism,
		Delay:       delay,
	})
}
