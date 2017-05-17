package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	_, err := gorm.Open("postgres", "host=localhost user=yagi dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		fmt.Print(err.Error())
	}

	e := echo.New()
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func saveUser(c echo.Context) error {
	return c.String(http.StatusOK, "save user")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("get: %s", id))
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("put: %s", id))
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("delete: %s", id))
}
