package config

import "github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Users{})
}
