package models

import "gorm.io/gorm"

type FilePhoto struct {
	gorm.Model
	PhotoName string
	PhotoPath string
}
