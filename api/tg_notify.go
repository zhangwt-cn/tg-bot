package api

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
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
}


func tgNotify(writer http.ResponseWriter,  request *http.Request)  {
	// json to Notify
	var notify Notify
	if err := json.NewDecoder(request.Body).Decode(&notify); err != nil {
		request.Body.Close()
		log.Panic(err)
	}
	// 创建bot
	bot, err := tgbotapi.NewBotAPI(notify.Token)   
	if err != nil {     
		log.Panic(err)   
	}
	fmt.Println(bot)
	// 发送消息    
	bot.Debug = true 
	log.Printf("Authorized on account %s", bot.Self.UserName)     
 	msg := tgbotapi.NewMessage(2109288988, notify.MsgText)   
	if _, err := bot.Send(msg); err != nil {     
  		log.Panic(err)    
 	}  
	// 返回请求结果
	var result Resp
	result.Code = "200"
	result.Msg = "消息已经发送：" +  notify.MsgText
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Panic(err)
	}
}