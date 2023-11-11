package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labtrilhadev/go-tutorials/di-fx/internal/usecase"
)

func RegisterAnyController(e *echo.Echo, c AnyController) {
	e.GET("/", c.Hello)
}

type AnyController struct {
	AnyUseCase usecase.AnyUseCase
}

func NewAnyController(anyUseCase usecase.AnyUseCase) AnyController {
	return AnyController{anyUseCase}
}

func (c *AnyController) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, c.AnyUseCase.Perform(ctx.Request().Context()))
}
