package models

import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Name string `json:"name", required`
	Email string `json:"email" gorm:"unique"`
}