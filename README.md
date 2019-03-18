# gocn

## 项目背景

由于想在公司内部传播golang, 某一天突发其想，为何不将gocn每日新闻每天定时发送到公司的技术群呢。于是在空闲时间码出这些代码

## 使用到开源工程

1. golang爬虫[colly](https://github.com/gocolly/colly)
2. 我最喜欢的配置文件[viper](https://github.com/spf13/viper)

## 相关目录的介绍

其实写的比较简单，欢迎吐槽

###  message

这里主要考虑可能后续会推送微信等工具。但是没有给出相应的interface, 如果有后续开发可能改动会较大

### db

主要为了存放那些已经爬取过的新闻

因为好像是侧边栏`相关问题`，在某些人评论后，每日新闻也会出现在那里，为了防止重复爬取相同的新闻，这里就用一个文本存放那些已经爬取过的新闻，当爬虫再次爬取到这些新闻时就忽略它

**后续优化：**

这里后续可以改成使用SQlite


## 配置文件splider.all

第一次抓取全部的时候，设置all为`true`, 其他时候都要设置为false

程序会判读当前页是否有下一页，如果有下一页会自动跳转到下一页进行爬取

### 如何启动

本程序也可以使用`crontab`方式启动，程序内没有封装定时任务的程序

## 有关具体的爬取内容

目前的内容不定时通过手动push到了[daily](./daily/golang-daily.md)


![](http://images.haohongfan.com/dingdinggocn.png)
