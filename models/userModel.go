package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email       string `form:"email"`
	Password    string `form:"password"`
	Username    string `form:"username"`
	NamaLengkap string `form:"nama_lengkap"`
}
