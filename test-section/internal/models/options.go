package models

import "gorm.io/gorm"


type Option struct {
	gorm.Model
	QuestionID uint      `gorm:"column:question_id;not null"`
	Text       string    `gorm:"column:text;type:varchar(255);not null"`
	IsCorrect  bool      `gorm:"column:is_correct;type:boolean;default:false"`
	Question   *Question `gorm:"foreignKey:QuestionID"`
}