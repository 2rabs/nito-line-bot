package cmdconfig

import (
	"github.com/2rabs/nito-line-bot/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenDB(appEnv *model.Env) (_ *gorm.DB, err error) {
	db, err := gorm.Open(
		postgres.Open(appEnv.DatabaseUrl),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	log.Printf("database open finished")
	return db, nil
}
