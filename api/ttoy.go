package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/my-Sakura/go-yuque-api/api"
)

func init() {
	token := "1574524831:AAEcjMbq_hKCyVlJtY9Qd4I29Wq0WLBOUSw"
	var bot, _ = tgbotapi.NewBotAPI(token)

	link := "go-telegram-bot.my-sakura.vercel.app/"

	bot.Debug = true

	_, err := bot.SetWebhook(tgbotapi.NewWebhook(link + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)

	var update tgbotapi.Update
	err := json.Unmarshal(bytes, &update)
	if err != nil {
		fmt.Println(err)
	}

	content := update.Data
	yuqueToken := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	namespace := "my-sakura/telegram"
	slug := "economistFifty"
	doc := api.GetDocumentInfo(yuqueToken, namespace, slug)
	id := doc.Data.Id
	api.UpdateDocument(yuqueToken, namespace, id, content)

	fmt.Println(update.Message.Text)
}
