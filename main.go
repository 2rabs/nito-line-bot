package main

import (
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/db"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"log"
	"net/http"
)

func main() {
	appEnv := loadAppEnv()

	db.Init(appEnv)
	defer db.Close()

	bot, err := model.NewNckTrampolineBot(appEnv)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, World")
	})
	http.HandleFunc("/callback", bot.CallbackHandler)

	//mem := db.GetMember(member.NewId(1))
	//bot.PushMessageForNotifyGroupId(
	//	fmt.Sprintf("MemberId: %d は %s さんです", mem.Id.Value, mem.Nickname.Value),
	//)

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
