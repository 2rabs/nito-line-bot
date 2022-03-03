package model

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/value"
	"log"
)

type NckTrampolineBot struct {
	Client        *linebot.Client
	notifyGroupId *value.NotifyGroupId
}

var (
	bot *NckTrampolineBot
)

func InitBot(appEnvironment *AppEnvironment) error {
	client, err := linebot.New(
		appEnvironment.ChannelSecret,
		appEnvironment.ChannelToken,
	)
	if err != nil {
		return err
	}

	bot = &NckTrampolineBot{
		Client:        client,
		notifyGroupId: appEnvironment.NotifyGroupId,
	}
	return nil
}

func GetBot() *NckTrampolineBot {
	return bot
}

func (bot NckTrampolineBot) PushMessageForNotifyGroupId(content string) {
	if _, err := bot.Client.PushMessage(bot.notifyGroupId.Value, linebot.NewTextMessage(content)).Do(); err != nil {
		log.Print(err)
	}
}
