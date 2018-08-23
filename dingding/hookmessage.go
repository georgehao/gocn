package dingding

//Package dinghook 详情参见 https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1
const (
	// MsgTypeText text 类型
	MsgTypeText = "text"
	// MsgTypeLink link 类型
	MsgTypeLink = "link"
	// MsgTypeMarkdown markdown 类型
	MsgTypeMarkdown = "markdown"
)

// Message 普通消息
type Message struct {
	Content   string `validate:"required"`
	AtPersion []string
	AtAll     bool
}

// Link 链接消息
type Link struct {
	Content    string `json:"text" validate:"required"`       // 要发送的消息， 必填
	Title      string `json:"title" validate:"required"`      // 标题， 必填
	ContentURL string `json:"messageUrl" validate:"required"` // 点击消息跳转的URL 必填
	PictureURL string `json:"picUrl"`                         // 图片 url
}

// Markdown markdown 类型
type Markdown struct {
	Content string `json:"text" validate:"required"`  // 要发送的消息， 必填
	Title   string `json:"title" validate:"required"` // 标题， 必填
}

// SimpleMessage push message
type SimpleMessage struct {
	Content string
	Title   string
}
