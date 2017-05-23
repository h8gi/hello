package controllers

import (
	"fmt"
	"net/http"

	"github.com/h8gi/hello/models"
	"github.com/labstack/echo"
)

func (cntrl *Controller) List(c echo.Context) (err error) {
	users := make([]models.User, 10)
	cntrl.DB.Find(&users)
	return c.Render(http.StatusOK, "users-list", users)
}

func (cntrl *Controller) CreateUser(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "create failed")
		return
	}
	if err = cntrl.DB.Create(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "create failed")
		return
	}
	return c.String(http.StatusCreated, "user created")
}

func (cntrl *Controller) GetUser(c echo.Context) error {
	user := new(models.User)
	name := c.Param("name")
	if cntrl.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, fmt.Sprintf("get: %s", name))
	}
	return c.String(http.StatusOK, fmt.Sprintf("%s!", user))
}

func (cntrl *Controller) UpdateUser(c echo.Context) (err error) {
	user := new(models.User)
	name := c.Param("name")
	if cntrl.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, fmt.Sprintf("not found"))
	}
	if err = c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "bad params")
		return err
	}
	if err = cntrl.DB.Save(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "bad params")
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("put: %s", name))
}

func (cntrl *Controller) DeleteUser(c echo.Context) (err error) {
	user := new(models.User)
	name := c.Param("name")
	if cntrl.DB.First(&user, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, "not found")
	}
	if err = cntrl.DB.Delete(user).Error; err != nil {
		c.String(http.StatusInternalServerError, "???")
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("delete: %s", name))
}
