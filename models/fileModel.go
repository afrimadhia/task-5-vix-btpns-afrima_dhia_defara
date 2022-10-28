package models

import "gorm.io/gorm"

type FilePhoto struct {
	gorm.Model
	PhotoName string `gorm:"type:varchar(255)"`
	PhotoPath string `gorm:"type:varchar(255)"`
}
