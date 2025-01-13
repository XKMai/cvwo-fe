package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Content  string `json:"content"`
	User    User   `json:"owner"`
	Post	Post   `json:"post"`
	Datetime string `json:"date_time"`
	Score    uint   `json:"score"`
}