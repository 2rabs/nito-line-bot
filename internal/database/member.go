package database

import (
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
