package initializers

import "github.com/BasantaBhusan/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
