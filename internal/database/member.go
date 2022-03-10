package database

import (
	"fmt"
	"github.com/2rabs/nito-line-bot/internal/database/entity"
	"github.com/2rabs/nito-line-bot/internal/domain/model/member"
	"gorm.io/gorm"
)

func GetMember(db *gorm.DB, id member.Id) *member.Member {
	var memberEntity entity.Member
	db.First(&memberEntity, id.Value)
	fmt.Println(memberEntity)
	return memberEntity.ToMember()
}
