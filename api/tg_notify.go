package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const(
	MARKDOWN = "Markdown"
	MARKDOWN_V2 = "MarkdownV2"
	JSON = `{"chat_id":%v,"text":"%s", "parse_mode":"%s"}`
)

// 返回实体
type Resp struct {
	Code	string `json:"code"`
	Msg		string `json:"msg"`
}

// 通知实体类
type Notify struct {
	Token     string  `json:"token"`
	MsgText   string  `json:"msgText"`
	ChatId    int64  `json:"chatId"`
	ParseMode string `json:"parseMode"`
}


func TgNotify(writer http.ResponseWriter,  request *http.Request)  {
	// json to Notify
	var notify Notify
	if err := json.NewDecoder(request.Body).Decode(&notify); err != nil {
		request.Body.Close()
		log.Panic(err)
	}
	// 创建bot
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", notify.Token)
	if notify.ParseMode == "" {
		notify.ParseMode = MARKDOWN
	}
	data := fmt.Sprintf(JSON, notify.ChatId, notify.MsgText, notify.ParseMode)
	resp, _ := http.Post(url, "application/json", strings.NewReader(data))
	log.Printf("msg send return : %v", resp)
	// 返回请求结果
	var result Resp
	result.Code = "200"
	result.Msg = "success"
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Panic(err)
	}
}