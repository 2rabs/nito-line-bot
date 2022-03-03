package db

import (
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/db/entity"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/member"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init(appEnv *model.AppEnvironment) {
	pg := postgres.Open(appEnv.DatabaseUrl)
	db, err = gorm.Open(pg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	log.Printf("データベースと接続しました")
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	database, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := database.Close(); err != nil {
		panic(err)
	}
	log.Printf("データベースの接続を切断しました")
}

func GetMember(id member.Id) *member.Member {
	var memberEntity entity.Member
	db.First(&memberEntity, id.Value)
	fmt.Println(memberEntity)
	return memberEntity.ToMember()
}
