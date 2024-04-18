package initializers

import "github.com/BasantaBhusan/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Kyc{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.WorkingArea{})
	DB.AutoMigrate(&models.Activity{})
	DB.AutoMigrate(&models.ActivityItem{})
	DB.AutoMigrate(&models.Service{})

}
