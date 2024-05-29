package handler

import (
	"go-backend/service"
	"go-backend/views/home"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	GenerateService service.GenerateService
}

func (h HomeHandler) ShowHome(c echo.Context) error {
	name := "teste"
	return render(c, home.Hello(name))
}

func (h HomeHandler) Generate(c echo.Context) error {
	generatedText := h.GenerateService.Generate()
	return c.String(http.StatusOK, generatedText.Text)
}
