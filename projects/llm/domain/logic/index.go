package logic

import (
	"github.com/labstack/echo/v4"
)

type LLMLogic struct {
}

func (_ LLMLogic) Get(c echo.Context) string {
	return "test"
}
