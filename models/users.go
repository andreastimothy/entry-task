package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id" gorm:"primaryKey"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Photo      string `json:"photo"`
}
