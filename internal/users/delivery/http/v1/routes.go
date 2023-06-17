package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *usersHandlers) MapRoutes() {
	h.group.POST("/createuser", h.CreateUser())
	h.group.GET("/test", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "OKK")
	})
	h.group.Any("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
