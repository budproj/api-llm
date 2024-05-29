package main

import (
	"go-backend/handler"
	"go-backend/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {

	generateService := service.GenerateService{}

	HomeHandler := handler.HomeHandler{
		GenerateService: generateService,
	}

	app := echo.New()
	app.Static("/static", "static")

	app.GET("/", HomeHandler.ShowHome)
	app.POST("/generate", HomeHandler.Generate)

	app.GET("/users/:id", getUser)

	app.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Logger.Fatal(app.Start(port))
}
