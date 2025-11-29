package models

import "gorm.io/gorm"




type Question struct {
	gorm.Model
	SectionID uint     `gorm:"column:section_id;not null"`
	Text      string   `gorm:"column:text;type:text;not null"`
	Section   *Section `gorm:"foreignKey:SectionID"`

		Options []Option `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}