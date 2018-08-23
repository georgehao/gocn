package message

import "errors"

type TextUrl struct {
	Text string
	Url  string
}

type Message struct {
	TextUrls   []TextUrl
	DailyTitle string
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
	return Message{}, errors.New("error")
}

func init() {
	messageChain = make(chan Message, 10000)
}
