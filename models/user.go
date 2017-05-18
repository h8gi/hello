package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     *string `json:"name" form:"name" gorm:"not null;unique"`
	Password *string `json:"password" form:"password" gorm:"not null"`
	Email    *string `json:"email" form:"email" gorm:"not null"`
}

// https://echo.labstack.com/guide/request
