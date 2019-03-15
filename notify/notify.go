package notify

import (
	"log"
	"net/http"
	"os"

	"github.com/Candy_Scrape/env"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Notifyer struct {
	line     *linebot.Client
	ID       string
	PUBLICID string
}

func NewNotifyer(config *env.Config) (*Notifyer, error) {
	n := Notifyer{}
	var err error
	n.line, err = linebot.New(config.Secret, config.Token)
	if err != nil {
		return &n, err
	}
	n.ID = config.ID
	n.PUBLICID = config.PUBLICID
	return &n, nil
}

func (n *Notifyer) Notify(message string) {
	log.Println(message)
	leftBtn := linebot.NewPostbackAction("はい", "予約依頼", "予約お願い！", "")
	rightBtn := linebot.NewPostbackAction("いいえ", "予約不要", "今回は回避！", "")
	pushMessage := linebot.NewConfirmTemplate(message, leftBtn, rightBtn)
	m := linebot.NewTemplateMessage("Sorry :(, please update your app.", pushMessage)
	_, err := n.line.PushMessage(n.ID, m).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (n *Notifyer) Wait() {
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := n.line.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			log.Println(event.Type)
			if event.Type == linebot.EventTypePostback {
				log.Println(event.Postback.Data)
				if event.Postback.Data == "予約依頼" {
					postMessage := linebot.NewTextMessage("やり方勉強中・・・、こんど教えてください！")
					_, err := n.line.PushMessage(n.ID, postMessage).Do()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
