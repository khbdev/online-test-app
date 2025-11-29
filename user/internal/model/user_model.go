package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	FirstName       string      `gorm:"column:first_name;type:varchar(255)" json:"first_name"`
	LastName        string      `gorm:"column:last_name;type:varchar(255)" json:"last_name"`
	Phone           string      `gorm:"column:phone;type:varchar(100)" json:"phone"`
	Email           string      `gorm:"column:email;type:varchar(255)" json:"email"`
	TgUsername      string      `gorm:"column:tg_username;type:varchar(255)" json:"tg_username"`
	Bolimlar        string      `gorm:"column:bolimlar;type:json" json:"-"`
	Savollar        string      `gorm:"column:savollar;type:json" json:"-"`
	Javoblar        string      `gorm:"column:javoblar;type:json" json:"-"`
	TogriJavoblar   int         `gorm:"column:togri_javoblar" json:"togri_javoblar"`
	NatogriJavoblar int         `gorm:"column:natogri_javoblar" json:"natogri_javoblar"`
	ScorePercent    int         `gorm:"column:score_percent" json:"score_percent"`
	Description     string      `gorm:"column:description;type:text" json:"description"`
	
	BolimlarParsed interface{} `gorm:"-" json:"bolimlar,omitempty"`
	SavollarParsed interface{} `gorm:"-" json:"savollar,omitempty"`
	JavoblarParsed interface{} `gorm:"-" json:"javoblar,omitempty"`
}
