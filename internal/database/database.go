package database

import (
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/member"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDb(appEnv *model.AppEnvironment) *member.Member {
	pg := postgres.Open(appEnv.DatabaseUrl)
	db, err := gorm.Open(pg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	var memberEntity Member
	db.First(&memberEntity, 1)
	fmt.Println(memberEntity)
	return memberEntity.toMember()
}

func (m Member) toMember() *member.Member {
	return member.NewMember(
		member.NewId(m.Id),
		member.NewLineUserId(m.LineUserId),
		member.NewNickname(m.Nickname),
	)
}
