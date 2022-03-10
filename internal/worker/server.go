package worker

import (
	"github.com/2rabs/nito-line-bot/internal/controller"
	"github.com/2rabs/nito-line-bot/internal/domain/model"
	"github.com/2rabs/nito-line-bot/internal/line"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	env *model.Env
	db  *gorm.DB
	bot *line.Bot
}

type ServerConfig struct {
	DB  *gorm.DB
	Bot *line.Bot
}

func NewServer(
	env *model.Env,
	config ServerConfig,
) (_ *Server, err error) {
	server := &Server{
		env: env,
		db:  config.DB,
		bot: config.Bot,
	}
	return server, nil
}

func (s Server) Install(
	env *model.Env,
	router *controller.Router,
) {
	// Root
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})

	// FIXME: 仮置き
	memberController := controller.NewMemberController(s.db, s.bot)
	router.HandleFunc("/api/message", memberController.MessageHandler)
	router.HandleFunc("/api/search-member", memberController.SearchMemberHandler)

	// Line のコールバック
	lineController := controller.NewLineController(s.bot)
	router.HandleFunc("/line/callback", lineController.CallbackHandler)
}
