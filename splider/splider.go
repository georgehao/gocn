package splider

import (
	"fmt"
	"github.com/georgehao/gocn/message"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

var c *colly.Collector

func Run() {
	// On every a element which has href attribute call callback
	c.OnHTML("div[class=\"aw-mod aw-question-detail aw-item\"]", func(e *colly.HTMLElement) {
		dailyTitle := e.ChildText("h1")
		var textUrls []message.TextUrl
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
		})

		message.Push(message.Message{
			DailyTitle: dailyTitle,
			TextUrls:   textUrls,
		})
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		if /*e.Text == ">" ||*/ strings.Contains(e.Text, "GoCN每日新闻") {
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://gocn.vip/explore/category-14")
}

func init() {
	// Instantiate default collector
	c = colly.NewCollector(
		colly.AllowedDomains("gocn.vip"),
		colly.MaxDepth(1),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "gocn.vip",
		Parallelism: 10,
		Delay:       1 * time.Second,
	})
}
