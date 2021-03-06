package model

import (
	"errors"
	"github.com/2rabs/nito-line-bot/internal/domain/model/value"
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	Port              string
	DatabaseUrl       string
	ChannelSecret     string
	ChannelToken      string
	NotifyGroupId     *value.NotifyGroupId
	TempNotifyGroupId *value.NotifyGroupId
}

func NewEnv() (*Env, error) {
	// 本番環境では .env を使用しないためエラーは無視
	_ = godotenv.Load(".env")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		return nil, errors.New("PORT を読み込めませんでした")
	}

	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, errors.New("DATABASE_URL を読み込めませんでした")
	}

	channelSecret, ok := os.LookupEnv("CHANNEL_SECRET")
	if !ok {
		return nil, errors.New("CHANNEL_SECRET を読み込めませんでした")
	}

	channelToken, ok := os.LookupEnv("CHANNEL_TOKEN")
	if !ok {
		return nil, errors.New("CHANNEL_TOKEN を読み込めませんでした")
	}

	notifyGroupId, ok := os.LookupEnv("NOTIFY_GROUP_ID")
	if !ok {
		return nil, errors.New("NOTIFY_GROUP_ID を読み込めませんでした")
	}

	tempNotifyGroupId, ok := os.LookupEnv("TEMP_NOTIFY_GROUP_ID")
	if !ok {
		return nil, errors.New("TEMP_NOTIFY_GROUP_ID を読み込めませんでした")
	}

	env := &Env{
		Port:              port,
		DatabaseUrl:       databaseUrl,
		ChannelSecret:     channelSecret,
		ChannelToken:      channelToken,
		NotifyGroupId:     value.NewNotifyGroupId(notifyGroupId),
		TempNotifyGroupId: value.NewNotifyGroupId(tempNotifyGroupId),
	}
	return env, nil
}
