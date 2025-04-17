package entity

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID  int    `json:"user_id" gorm:""`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
	Content string `json:"content"`
}
