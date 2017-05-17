package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
}

// https://echo.labstack.com/guide/request
