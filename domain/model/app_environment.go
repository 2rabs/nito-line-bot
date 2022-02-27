package model

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppEnvironment struct {
	Port          string
	ChannelSecret string
	ChannelToken  string
	NotifyGroupId string
}

func NewAppEnvironment() (*AppEnvironment, error) {
	// 本番環境では .env を使用しないためエラーは無視
	_ = godotenv.Load(".env")

	channelSecret, isReadChannelSecret := os.LookupEnv("CHANNEL_SECRET")
	if !isReadChannelSecret {
		return nil, errors.New("CHANNEL_SECRET を読み込めませんでした")
	}

	channelToken, isReadChannelToken := os.LookupEnv("CHANNEL_TOKEN")
	if !isReadChannelToken {
		return nil, errors.New("CHANNEL_TOKEN を読み込めませんでした")
	}

	port, isReadPort := os.LookupEnv("PORT")
	if !isReadPort {
		return nil, errors.New("PORT を読み込めませんでした")
	}

	notifyGroupId, isReadNotifyGroupId := os.LookupEnv("NOTIFY_GROUP_ID")
	if !isReadNotifyGroupId {
		return nil, errors.New("NOTIFY_GROUP_ID を読み込めませんでした")
	}

	appEnvironment := &AppEnvironment{
		Port:          port,
		ChannelSecret: channelSecret,
		ChannelToken:  channelToken,
		NotifyGroupId: notifyGroupId,
	}
	return appEnvironment, nil
}
