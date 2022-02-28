package model

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/value"
	"os"
)

type AppEnvironment struct {
	Port          string
	DatabaseUrl   string
	ChannelSecret string
	ChannelToken  string
	NotifyGroupId *value.NotifyGroupId
}

func NewAppEnvironment() (*AppEnvironment, error) {
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

	appEnvironment := &AppEnvironment{
		Port:          port,
		DatabaseUrl:   databaseUrl,
		ChannelSecret: channelSecret,
		ChannelToken:  channelToken,
		NotifyGroupId: value.NewNotifyGroupId(notifyGroupId),
	}
	return appEnvironment, nil
}
