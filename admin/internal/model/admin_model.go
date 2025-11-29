package model

import (


	"gorm.io/gorm"
)


type Admin struct {
    gorm.Model    
    Username string `gorm:"type:varchar(255);not null"`
    Password string `gorm:"type:varchar(255);not null"`
}