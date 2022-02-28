package main

import (
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"log"
	"net/http"
)

func main() {
	appEnv := loadAppEnv()

	bot, err := model.NewNckTrampolineBot(appEnv)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, World")
	})
	http.HandleFunc("/callback", bot.CallbackHandler)

	if err := http.ListenAndServe(":"+appEnv.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func loadAppEnv() *model.AppEnvironment {
	appEnv, err := model.NewAppEnvironment()
	if err != nil {
		panic(fmt.Sprintf("環境変数を読み込めませんでした: %v", err))
	}
	log.Printf("環境変数を読み込みました")
	return appEnv
}
