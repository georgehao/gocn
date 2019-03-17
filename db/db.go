package db

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var fd *os.File
var markdownFd *os.File
var dailyTotal []string
var bufferChain chan string
var markdownBufferChain chan string

// CheckSend check send to dingding if need
func CheckSend(msg string) bool {
	for _, v := range dailyTotal {
		if msg == v {
			return true
		}
	}
	return false
}

func writeMarkdown(msg string) {
	fmt.Fprintln(markdownFd, "\n"+msg)
	markdownFd.Sync()
}

func write(str string) {
	fmt.Fprintln(fd, str)
	dailyTotal = append(dailyTotal, str)
	fd.Sync()
}

func read() error {
	br := bufio.NewReader(fd)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		dailyTotal = append(dailyTotal, string(a))
	}

	return nil
}

func Push(str string) {
	bufferChain <- str
}

func PushMarkdown(msg string) {
	markdownBufferChain <- msg
}

func Run() {
	for {
		select {
		case m := <-bufferChain:
			write(m)
		case m := <-markdownBufferChain:
			writeMarkdown(m)
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

func init() {
	bufferChain = make(chan string, 1000)
	markdownBufferChain = make(chan string, 1000)

	var err error
	fd, err = os.OpenFile("./db.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}

	markdownFd, err = os.OpenFile("./daily/golang-daily.md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}

	read()
}
