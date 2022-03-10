package line

import (
	"github.com/2rabs/nito-line-bot/internal/domain/model"
	"github.com/2rabs/nito-line-bot/internal/domain/model/value"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
)

type Bot struct {
	Client        *linebot.Client
	notifyGroupId *value.NotifyGroupId
}

func NewLineBot(appEnvironment *model.Env) (_ *Bot, err error) {
	client, err := linebot.New(
		appEnvironment.ChannelSecret,
		appEnvironment.ChannelToken,
	)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Client:        client,
		notifyGroupId: appEnvironment.NotifyGroupId,
	}
	return bot, nil
}

func (bot Bot) PushMessageForNotifyGroupId(content string) {
	if _, err := bot.Client.PushMessage(bot.notifyGroupId.Value, linebot.NewTextMessage(content)).Do(); err != nil {
		log.Print(err)
	}
}
