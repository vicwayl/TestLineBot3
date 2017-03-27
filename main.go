package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	
	//if _, err := bot.PushMessage("U2c68fd429a99dceccc8956571baa7d00", linebot.NewTextMessage("hello")).Do(); err != nil {
	//	txt= txt + err.Error()
	//}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//var txt = Send(message.Text);
				//rand.Seed(99)
				answers := []string{"您好!我是鐵蛋機器人","歡迎來找我玩喔~","工程師哥哥真帥","我最喜歡導覽姐姐了<3","我是阿呆博士的機器人!!"}
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!"+txt+"  "+event.Source.UserID+"   "+event.ReplyToken)).Do(); err != nil {
				//	log.Print(err)
				//}
				var txt =answers[rand.Intn(len(answers))]
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(txt)).Do(); err != nil {
					log.Print(err)
				}
				
			}
		}
	//	if event.Type == linebot.EventTypeFollow {
	//		var text = "Hi!歡迎使用魔物獵人LINE@BOT\n"+
	//			"指令:\n"+
	//			"@魔物名稱\n"+
	//			"/功能\n"+
				
	//			"\n若不知道該如何下指令，請輸入/help查詢。"
	//		if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
	//			log.Print(err)
	//		}
//		}
	}
}
