package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Controller struct {
	DB *gorm.DB
}

func (cn *Controller) CreateUser(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "create failed")
		return
	}
	cn.DB.Create(&user)
	return c.String(http.StatusCreated, "user created")
}

func (cn *Controller) GetUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	cn.DB.First(&user, "id = ?", id)
	if user.ID != 0 {
		return c.String(http.StatusOK, fmt.Sprintf("get %s", id))
	}
	return c.String(http.StatusNotFound, fmt.Sprintf("get: %s", id))
}

func (cn *Controller) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("put: %s", id))
}

func (cn *Controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("delete: %s", id))
}
