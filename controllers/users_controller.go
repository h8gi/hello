package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) Create(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "create failed")
		return
	}
	if err = uc.DB.Create(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "create failed")
		return
	}
	return c.String(http.StatusCreated, "user created")
}

func (uc *UserController) Get(c echo.Context) error {
	user := new(models.User)
	name := c.Param("name")
	if uc.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, fmt.Sprintf("get: %s", name))
	}
	return c.String(http.StatusOK, fmt.Sprintf("get %s", name))
}

func (uc *UserController) Update(c echo.Context) (err error) {
	user := new(models.User)
	name := c.Param("name")
	if uc.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, fmt.Sprintf("not found"))
	}
	if err = c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "bad params")
		return err
	}
	if err = uc.DB.Save(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "bad params")
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("put: %s", name))
}

func (uc *UserController) Delete(c echo.Context) (err error) {
	user := new(models.User)
	name := c.Param("name")
	if uc.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, "not found")
	}
	if err = uc.DB.Delete(user).Error; err != nil {
		c.String(http.StatusInternalServerError, "???")
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("delete: %s", name))
}
