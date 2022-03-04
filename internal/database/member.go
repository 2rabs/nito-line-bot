package database

import (
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/database/entity"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/member"
	"gorm.io/gorm"
)

func GetMember(db *gorm.DB, id member.Id) *member.Member {
	var memberEntity entity.Member
	db.First(&memberEntity, id.Value)
	fmt.Println(memberEntity)
	return memberEntity.ToMember()
}
