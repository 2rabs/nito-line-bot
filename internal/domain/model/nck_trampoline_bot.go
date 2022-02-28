package model

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/value"
	"log"
	"net/http"
)

type NckTrampolineBot struct {
	client        *linebot.Client
	notifyGroupId *value.NotifyGroupId
}

func NewNckTrampolineBot(appEnvironment *AppEnvironment) (*NckTrampolineBot, error) {
	client, err := linebot.New(
		appEnvironment.ChannelSecret,
		appEnvironment.ChannelToken,
	)
	if err != nil {
		return nil, err
	}

	nckTrampolineBot := &NckTrampolineBot{
		client:        client,
		notifyGroupId: appEnvironment.NotifyGroupId,
	}
	return nckTrampolineBot, nil
}

func (bot NckTrampolineBot) PushMessageForNotifyGroupId(content string) {
	if _, err := bot.client.PushMessage(bot.notifyGroupId.Value, linebot.NewTextMessage(content)).Do(); err != nil {
		log.Print(err)
	}
}

func (bot NckTrampolineBot) CallbackHandler(w http.ResponseWriter, req *http.Request) {
	events, err := bot.client.ParseRequest(req)
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
			//	if _, err = bot.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			//		log.Print(err)
			//	}
			//case *linebot.StickerMessage:
			//	replyMessage := fmt.Sprintf(
			//		"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
			//	if _, err = bot.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			//		log.Print(err)
			//	}
			//}
		}
	}
}
