package handler

import (
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
	s := new(model.SummarizeKeyResultInput)
	if err := c.Bind(s); err != nil {
		return err
	}

	OKR := model.SummarizeKeyResultInput{
		Objective:   s.Objective,
		Cycle:       s.Cycle,
		Title:       s.Title,
		Description: s.Description,
		Goal:        s.Goal,
		Format:      s.Format,
		Owner:       s.Owner,
		Comments:    s.Comments,
		Checklist:   s.Checklist,
		CheckIns:    s.CheckIns,
	}

	// return c.JSON(http.StatusCreated, summarized))
	generatedText := h.GenerateService.Generate(OKR)
	return c.String(http.StatusOK, generatedText)
}
