package handler

import (
	"fmt"
	"go-backend/entrypoint/templates/html/home"
	service "go-backend/services"
	model "go-backend/services/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	GenerateService service.GenerateService
}

func (h HomeHandler) ShowHome(c echo.Context) error {
	name := "2"

	return render(c, home.Hello(name))
}

func (h HomeHandler) Generate(c echo.Context) error {
	fmt.Println(c)
	summarized := new(model.SummarizeKeyResultInput)
	if err := c.Bind(summarized); err != nil {
		return err
	}
	// return c.JSON(http.StatusCreated, summarized))
	generatedText := h.GenerateService.Generate(summarized)
	return c.String(http.StatusOK, generatedText)
}
