package main

import (
	llm "apicore/projects/llm/entrypoints/controllers/http"
	"github.com/gofor-little/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func httpStart(port string) {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	llm.Start(app)
	app.Logger.Fatal(app.Start(port))
}

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	port := env.Get("PORT", ":3000")
	httpStart(port)
}
