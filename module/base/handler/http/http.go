package http

import (
	"lolipad/boilerplate/middleware"
	"lolipad/boilerplate/module/base/usecase"

	"github.com/labstack/echo/v4" //nolint
)

// BaseHandler  represent the httphandler for Base
type BaseHandler struct {
	baseUsecase usecase.Usecase
}

// NewBaseHandler will initialize the contact/ resources endpoint
func NewBaseHandler(e *echo.Group, us usecase.Usecase) {
	handler := &BaseHandler{
		baseUsecase: us,
	}

	router := e.Group("/v1", middleware.Authorization)

	router.GET("/base", handler.GetBase)
}

func (h *BaseHandler) GetBase(c echo.Context) error {
	return c.JSON(200, "Base")
}
