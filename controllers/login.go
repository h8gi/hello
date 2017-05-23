package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/h8gi/hello/models"
	"github.com/labstack/echo"
)

func (cntrl *Controller) ShowLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func (cntrl *Controller) Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	user := new(models.User)
	if cntrl.DB.First(&user, "name = ?", name).RecordNotFound() {
		return echo.ErrUnauthorized
	}
	if *user.Password != password {
		return echo.ErrUnauthorized
	}
	//  create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = *user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour * 24)
	// cookie.Secure = true
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
