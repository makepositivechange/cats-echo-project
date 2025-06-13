package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/makepostivechange/cats-echo-project/models"
	"github.com/makepostivechange/cats-echo-project/requests"
	"github.com/makepostivechange/cats-echo-project/response"
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
		response := response.Response{
			Code:    http.StatusNotFound,
			Message: "Could not find cats info from database",
			Error:   res.Error,
		}
		return c.JSON(http.StatusNotFound, response)
	}
	res_success := response.Response{
		Code:    http.StatusOK,
		Message: "success",
	}
	return c.JSON(http.StatusOK, res_success)
}

func (h *Handler) GetCat(c echo.Context) error {
	breed_name := c.Param("breed_name")
	var cat_model models.CatInfo
	res := h.DB.Where("breed_name = ?", breed_name).First(&cat_model)
	if res.Error != nil {
		response := response.Response{
			Code:    http.StatusNotFound,
			Message: "Breed not found in database",
		}
		return c.JSON(http.StatusNotFound, response)
	}
	return c.JSON(http.StatusOK, cat_model)
}

func (h *Handler) UpdateCatInfo(c echo.Context) error {
	breed_name := c.Param("breed_name")
	var cat_model models.CatInfo
	payload := new(requests.UpdateCatTypeInfo)
	err := (&echo.DefaultBinder{}).Bind(&payload, c)
	if err != nil {
		response_request := response.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
		}
		return c.JSON(http.StatusBadRequest, response_request)
	}
	res := h.DB.Where("breed_name = ?", breed_name).Find(&cat_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Breed not found in database")
	}
	if payload.CatTypeInfo != nil {
		cat_model.TypeInfo = payload.CatTypeInfo
		result := h.DB.Save(&cat_model)
		if result.Error != nil {
			response_error := response.Response{
				Code:    http.StatusInternalServerError,
				Message: "Failed to store updated information",
			}
			return c.JSON(http.StatusInternalServerError, response_error)
		}
	}
	return c.JSON(http.StatusOK, cat_model)
}

func (h *Handler) AddNewCatToDB(c echo.Context) error {
	payload := new(requests.AddNewCat)
	err := (&echo.DefaultBinder{}).Bind(payload, c)
	if err != nil {
		log.Printf("Received an invalid request:%v", err)
		response := response.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid request received",
			Error:   err,
		}
		return c.JSON(http.StatusBadRequest, response)
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
		response_err := response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Could not add movie",
			Error:   result.Error,
		}
		return c.JSON(http.StatusInternalServerError, response_err)
	}
	return c.JSON(http.StatusOK, n_movie)
}

func (h *Handler) RemoveCatFromDB(c echo.Context) error {
	breed_name := c.Param("breed_name")
	res := h.DB.Where("breed_name = ?", breed_name).Delete(&models.CatInfo{})
	if res.Error != nil {
		res := response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Breed not found in database",
			Error:   res.Error,
		}
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, "Successful delete")
}
