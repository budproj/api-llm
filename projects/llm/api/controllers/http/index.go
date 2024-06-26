package http

import (
	"apicore/projects/llm/handler"
	service "apicore/projects/llm/services"

	"github.com/labstack/echo/v4"
)

func Start(app *echo.Echo) {
	generateService := service.GenerateService{}

	HomeHandler := handler.HomeHandler{
		GenerateService: generateService,
	}

	app.GET("/", HomeHandler.ShowHome)
	app.POST("/generate", HomeHandler.Generate)
	app.Static("/static", "assets")
}
