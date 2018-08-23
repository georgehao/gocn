package db

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	write("GoCN每日新闻(2018-01-12)")
	write("GoCN每日新闻(2018-01-12)")
}

func TestRead(t *testing.T) {
	read()
	fmt.Println(dailyTotal)
}
