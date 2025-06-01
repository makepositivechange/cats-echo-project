package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/makepostivechange/cats-echo-project/models"
	"github.com/makepostivechange/cats-echo-project/requests"
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
	var cat_model models.CatInfo
	res := h.DB.Where("breed_name = ?", breed_name).First(&cat_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Breed not found in database")
	}
	return c.JSON(http.StatusOK, cat_model)
}

func (h *Handler) UpdateCatInfo(c echo.Context) error {
	var cat_model models.CatInfo
	payload := new(requests.UpdateCatInfo)
	err := (&echo.DefaultBinder{}).Bind(&payload, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}
	res := h.DB.Where("breed_name = ?", payload.CatBreed).First(&cat_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Breed not found in database")
	}
	if payload.CatTypeInfo != nil {
		cat_model.TypeInfo = payload.CatTypeInfo
		result := h.DB.Model(&cat_model).Updates(map[string]any{
			"type_info": payload.CatTypeInfo,
		})
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to store updated information")
		}

	}
	return c.JSON(http.StatusOK, cat_model)
}

func (h *Handler) AddNewCatToDB(c echo.Context) error {
	payload := new(requests.AddNewCat)
	err := (&echo.DefaultBinder{}).Bind(payload, c)
	if err != nil {
		log.Printf("Received an invalid request:%v", err)
		return c.JSON(http.StatusBadRequest, "invalid request received")
	}
	n_movie := models.CatInfo{
		BreedName:   payload.CatBreed,
		BreedOrigin: payload.CatOriginDescription,
		BreedType:   payload.CatType,
		BodyTypes:   payload.BodyType,
		CoatPattern: payload.CoatPattern,
		TypeInfo:    payload.CatTypeInfo,
	}
	result := h.DB.Create(&n_movie)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Could not add movie")
	}
	return c.JSON(http.StatusOK, n_movie)
}
