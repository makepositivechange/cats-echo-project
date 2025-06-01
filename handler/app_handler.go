package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	health_check_struct := struct {
		Health bool `json:"health"`
	}{
		Health: true,
	}
	return c.JSON(http.StatusOK, health_check_struct)
}
