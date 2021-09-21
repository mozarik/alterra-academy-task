package controllers

import (
	"net/http"
	"strconv"
	"tugas-mvc/models"
	"tugas-mvc/repository"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	userData, err := repository.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.DefaultResponse{
		Code: http.StatusOK,
		Data: userData,
	})
}

func AddUser(c echo.Context) error {
	var user models.AddUserRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request body",
		})
	}

	userData, err := repository.AddUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.DefaultResponse{
		Code: http.StatusCreated,
		Data: userData,
	})
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request body",
		})
	}

	userData, err := repository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.DefaultResponse{
			Code:    http.StatusNotFound,
			Message: "Cannot found user with that id",
		})
	}

	return c.JSON(http.StatusOK, models.DefaultResponse{
		Code: http.StatusOK,
		Data: userData,
	})
}

func UpdateUser(c echo.Context) error {
	// Get the user data with the id first and change the state
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request body",
		})
	}

	userState, err := repository.GetUser(id)

	// Change the user state based on the request
	var userNewState models.UpdateUserRequest
	err = c.Bind(&userNewState)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request body",
		})
	}

	userState.Name = userNewState.Name
	userState.Email = userNewState.Email
	userState.DOB = userNewState.DOB

	// Update the user
	userData, err := repository.UpdateUser(userState)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.DefaultResponse{
		Code: http.StatusOK,
		Data: userData,
	})
}
