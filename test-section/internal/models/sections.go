package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Name      string     `gorm:"column:name;type:varchar(255);not null"`
	Questions []Question `gorm:"foreignKey:SectionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}