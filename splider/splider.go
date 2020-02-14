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
	var url, title string
	// On every a element which has href attribute call callback
	c.OnHTML("div[class=\"card-body markdown markdown-toc\"]", func(e *colly.HTMLElement) {
		var textUrls []message.TextUrl
		var author string
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			url := e.ChildText("a[href]")
			urlIndex := strings.Index(e.Text, url)
			if urlIndex <= len(e.Text) && urlIndex != -1 {
				//println(urlIndex, len(e.Text), e.Text)
				text := e.Text[0:urlIndex]
				textUrls = append(textUrls, message.TextUrl{text, url})
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

		var newTextUrls []message.TextUrl
		uniMap := make(map[string]bool)
		for _, v := range textUrls {
			if _, ok := uniMap[v.Url]; !ok {
				uniMap[v.Url] = true
				newTextUrls = append(newTextUrls, v)
			}
		}

		message.Push(message.Message{
			DailyTitle: title,
			TextUrls:   newTextUrls,
			Author:     author,
			PostUrl:    e.Request.URL.String(),
		})
	})

	c.OnHTML("div[class=\"title media-heading\"]", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, el *colly.HTMLElement) {
			title = el.Attr("title")
			link := el.Attr("href")
			//fmt.Printf("Link found: %q -> %s\n", el.Text, link)
			if strings.Contains(el.Text, "每日新闻") {
				_ = c.Visit(el.Request.AbsoluteURL(link))
				url = el.Request.AbsoluteURL(link)
			}
		})
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
	_ = c.Visit("https://gocn.vip/topics/node18")
}

func urlTextDeal(val string, res *[]message.TextUrl) {
	httpIndex := strings.Index(val, "http")
	if httpIndex <= len(val) && httpIndex != -1 {
		text := val[0:httpIndex]
		url := val[httpIndex:]
		*res = append(*res, message.TextUrl{text, url})
	}
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
