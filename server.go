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
	cntrl := controllers.Controller{
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

	e.POST("/api/users", cntrl.CreateUser)
	e.GET("/api/users/:id", cntrl.GetUser)
	e.PUT("/api/users/:id", cntrl.UpdateUser)
	e.DELETE("/api/users/:id", cntrl.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))

}
