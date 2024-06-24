package main

import (
	"go-backend/handler"
	service "go-backend/services"

	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	generateService := service.GenerateService{}

	HomeHandler := handler.HomeHandler{
		GenerateService: generateService,
	}

	app := echo.New()

	app.GET("/", HomeHandler.ShowHome)
	app.POST("/generate", HomeHandler.Generate)
	app.Static("/static", "assets")

	port := env.Get("PORT", ":3000")

	app.Logger.Fatal(app.Start(port))
}
