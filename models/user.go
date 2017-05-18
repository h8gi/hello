package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     *string `json:"name" form:"name" gorm:"not null;unique"`
	Password *string `json:"password" form:"password" gorm:"not null"`
	Email    *string `json:"email" form:"email" gorm:"not null"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s: %s", *u.Name, *u.Email)
}

// https://echo.labstack.com/guide/request
