package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/makepostivechange/cats-echo-project/models"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	health_check_struct := struct {
		Health bool `json:"health"`
	}{
		Health: true,
	}
	return c.JSON(http.StatusOK, health_check_struct)
}

// GetCats will return all the cat info information in the database
func (h *Handler) GetCats(c echo.Context) error {
	var cat_model []models.CatInfo
	res := h.DB.Find(&cat_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Could not find cats info from database")
	}
	return c.JSON(http.StatusOK, cat_model)
}

func (h *Handler) GetCat(c echo.Context) error {
	breed_name := c.Param("breed_name")
	var car_model models.CatInfo
	res := h.DB.Where("breed_name = ?", breed_name).First(&car_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Breed not found in database")
	}
	return c.JSON(http.StatusOK, car_model)
}
