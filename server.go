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
	db, err := gorm.Open("postgres", "host=localhost user=yagi dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	// controller
	usersController := controllers.UsersController{
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

	e.GET("/user", usersController.List)
	e.POST("/user", usersController.Create)

	e.GET("/user/:name", usersController.Get)
	e.PUT("/user/:name", usersController.Update)
	e.DELETE("/user/:name", usersController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
