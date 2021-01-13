package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FlashFeiFei/yuque/response"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegram struct {
	bot *tgbotapi.BotAPI
}

var t = new(telegram)

func init() {
	token := "1574524831:AAEcjMbq_hKCyVlJtY9Qd4I29Wq0WLBOUSw"
	t.bot, _ = tgbotapi.NewBotAPI(token)
	t.bot.Debug = true
}

func HelloServer(w http.ResponseWriter, req *http.Request) {

	content, _ := ioutil.ReadAll(req.Body)

	doc := response.ResponseDocDetailSerializer{}
	json.Unmarshal(content, &doc)

	text := doc.Data.Title + doc.Data.Body

	msg := tgbotapi.NewMessage(1200586530, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	if _, err := t.bot.Send(msg); err != nil {
		log.Println(err)
	}
}
