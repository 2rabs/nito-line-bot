package controller

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/tkchry/nck-trampoline-bot/internal/line"
	"net/http"
)

type LineController struct {
	bot *line.Bot
}

func NewLineController(bot *line.Bot) *LineController {
	return &LineController{
		bot: bot,
	}
}

func (c LineController) CallbackHandler(w http.ResponseWriter, req *http.Request) {
	events, err := c.bot.Client.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			//switch message := event.Message.(type) {
			//case *linebot.TextMessage:
			//	for _, mentione := range message.Mention.Mentionees {
			//		log.Printf("mentioneUserId: %s", mentione.UserID)
			//	}
			//}
		}
	}
}
