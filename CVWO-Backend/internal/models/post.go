package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"` //Can have multiple categories
	User    User   `json:"owner"`
	Datetime string `json:"date_time"`
	Score    uint   `json:"score"`
}