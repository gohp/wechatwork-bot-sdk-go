package main

import (
	"fmt"
	"github.com/gohp/wechatwork-bot-sdk-go"
	"log"
)

func main() {
	bot := wechatwork_bot_sdk_go.NewBot("your-key-here")
	resp, err := bot.SendText(wechatwork_bot_sdk_go.TextContent{
		Content: "test",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp))
}
