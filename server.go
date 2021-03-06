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
	cntrl := controllers.Controller{
		DB: db,
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	e := echo.New()
	// register templates
	e.Renderer = t
	// e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

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

	e.GET("/login", cntrl.ShowLogin)
	e.POST("/login", cntrl.Login)

	e.GET("/users", cntrl.List)
	e.POST("/users", cntrl.CreateUser)

	e.GET("/users/:name", cntrl.GetUser)
	e.PUT("/users/:name", cntrl.UpdateUser)
	e.DELETE("/users/:name", cntrl.DeleteUser)

	e.GET("/restricted", controllers.Restricted, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:token",
	}))

	// e.Logger.Fatal(e.StartAutoTLS("localhost.com:443"))
	e.Logger.Fatal(e.Start(":1323"))
}
