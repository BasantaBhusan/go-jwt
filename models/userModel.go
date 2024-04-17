package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"unique; not null" json:"email" binding:"required"`
	Password       string `gorm:"not null" json:"password" binding:"required"`
	ProfilePicture string `json:"profile_picture"`
	Role           string `gorm:"default:USER" json:"role"`
	IsKyc          bool   `gorm:"default:false" json:"is_kyc"`
	Kyc            *Kyc   `gorm:"foreignKey:UserID"`
}
