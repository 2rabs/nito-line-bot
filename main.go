package main

import (
	"fmt"
	"github.com/2rabs/nito-line-bot/internal/cmdconfig"
	"github.com/2rabs/nito-line-bot/internal/controller"
	"github.com/2rabs/nito-line-bot/internal/domain/model"
	"github.com/2rabs/nito-line-bot/internal/line"
	"github.com/2rabs/nito-line-bot/internal/worker"
	"log"
	"net/http"
)

func main() {
	env := loadAppEnv()

	db, err := cmdconfig.OpenDB(env)
	if err != nil {
		log.Fatalf("%v", err)
	}

	bot, err := line.NewLineBot(env)
	if err != nil {
		log.Fatalf("%v", err)
	}

	server, err := worker.NewServer(env, worker.ServerConfig{
		DB:  db,
		Bot: bot,
	})

	router := controller.NewRouter()

	server.Install(env, router)

	if err := http.ListenAndServe(":"+env.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func loadAppEnv() *model.Env {
	appEnv, err := model.NewEnv()
	if err != nil {
		panic(fmt.Sprintf("環境変数を読み込めませんでした: %v", err))
	}
	log.Printf("環境変数を読み込みました")
	return appEnv
}
