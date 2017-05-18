package main

import (
	"net/http"

	"./controllers"
	"./models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=yagihiroki dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	// controller
	user := controllers.UserController{
		DB: db,
	}

	e := echo.New()
	// Middleware
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/user", user.Create)
	e.GET("/user/:name", user.Get)
	e.PUT("/user/:name", user.Update)
	e.DELETE("/user/:name", user.Delete)

	e.Logger.Fatal(e.Start(":1323"))

}
