package main

import (
	"html/template"
	"io"
	"net/http"

	"./models"

	"github.com/h8gi/hello/controllers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=yagi dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&models.User{})

	t := &Template{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	e := echo.New()
	// register templates
	e.Renderer = t

	usersController := controllers.UsersController{
		DB: db,
	}

	// Middleware
	// remove trailing slash. /hello/ -> /hello
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login", usersController.Login)

	e.GET("/users", usersController.List)
	e.POST("/users", usersController.Create)

	e.GET("/users/:name", usersController.Get)
	e.PUT("/users/:name", usersController.Update)
	e.DELETE("/users/:name", usersController.Delete)

	e.GET("/restricted", controllers.Restricted, middleware.JWT([]byte("secret")))

	e.Logger.Fatal(e.Start(":1323"))
}
