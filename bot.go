package wechatwork_bot_sdk_go

import (
	"encoding/json"
	"github.com/wzyonggege/goutils/httplib"
	"log"
)

const (
	baseUrl      = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	typeText     = "text"
	typeMarkdown = "markdown"
	typeImage    = "image"
	typeFile     = "file"
	typeNews     = "news"
)

type Bot struct {
	token string
}

type TextContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type MarkdownContent struct {
	Content string `json:"content"`
}

type NewsContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type FileContent struct {
	MediaId string `json:"media_id"`
}

type ImageContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type PostBody struct {
	MsgType  string          `json:"msgtype"`
	Text     TextContent     `json:"text,omitempty"`
	Image    ImageContent    `json:"image,omitempty"`
	Markdown MarkdownContent `json:"markdown,omitempty"`
	News     NewsContent     `json:"news,omitempty"`
	File     FileContent     `json:"file,omitempty"`
}

func NewBot(token string) *Bot {
	bot := &Bot{
		token: token,
	}
	return bot
}

func (bot *Bot) SendText(msg TextContent) ([]byte, error) {
	data := PostBody{
		MsgType: typeText,
		Text:    msg,
	}
	req, _ := json.Marshal(data)
	return bot.httpDo(req)
}

func (bot *Bot) SendMarkdown(msg MarkdownContent) ([]byte, error) {
	data := PostBody{
		MsgType:  typeMarkdown,
		Markdown: msg,
	}
	req, _ := json.Marshal(data)
	return bot.httpDo(req)
}

func (bot *Bot) SendImage(msg ImageContent) ([]byte, error) {
	data := PostBody{
		MsgType: typeImage,
		Image:   msg,
	}
	req, _ := json.Marshal(data)
	return bot.httpDo(req)
}

func (bot *Bot) SendFile(msg FileContent) ([]byte, error) {
	data := PostBody{
		MsgType: typeFile,
		File:    msg,
	}
	req, _ := json.Marshal(data)
	return bot.httpDo(req)
}

func (bot *Bot) SendNews(msg NewsContent) ([]byte, error) {
	data := PostBody{
		MsgType: typeNews,
		News:    msg,
	}
	req, _ := json.Marshal(data)
	return bot.httpDo(req)
}

func (bot *Bot) httpDo(data []byte) ([]byte, error) {
	resp, err := httplib.Post(baseUrl+bot.token).
		Body(data).
		Header("Content-Type", "application/json").
		Bytes()

	if err != nil {
		log.Fatal(err)
	}

	return resp, err
}
