/*
 * @Author: Flyinsky w2084151024@gmail.com
 * @Description: None
 */
package entity

import "gorm.io/gorm"

type Virtual struct {
	gorm.Model
	ID          uint    `json:"id"`
	ProfileId   uint    `json:"profileId" gorm:"column:profile_id"`
	UserID      uint    `json:"user_id" gorm:"column:user_id"`
	Difficulty  string  `json:"difficulty"`
	User        User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Seconds     int     `json:"seconds"`
	Score       int     `json:"score"`
	Description string  `json:"description"`
	Profile     Profile `json:"profile" gorm:"foreignKey:ProfileId;references:ID"`
}

type VirtualQuestion struct {
	gorm.Model
	ID                uint    `json:"id"`
	VirtualId         uint    `json:"virtualId" gorm:"column:virtual_id"`
	Question          string  `json:"question"`
	Answer            string  `json:"answer"`
	QuestionAudioPath string  `json:"question_audio_path"`
	AnswerAudioPath   string  `json:"answer_audio_path"`
	Virtual           Virtual `json:"virtual" gorm:"foreignKey:VirtualId;references:ID"`
}
