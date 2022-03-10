package entity

import (
	"github.com/2rabs/nito-line-bot/internal/domain/model/member"
	"gorm.io/gorm"
	"time"
)

type Member struct {
	gorm.Model
	Id           int64
	LineUserId   string
	Nickname     string
	LastUpdateAt time.Time
	DeletedAt    *time.Time
}

func (m Member) ToMember() *member.Member {
	return member.NewMember(
		member.NewId(m.Id),
		member.NewLineUserId(m.LineUserId),
		member.NewNickname(m.Nickname),
	)
}
