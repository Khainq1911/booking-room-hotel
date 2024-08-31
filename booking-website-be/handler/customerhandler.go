package handler

import (
	"booking-website-be/model"
	repo "booking-website-be/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	Repo repo.CustomerRepo
}

func (u CustomerHandler) GetAllRoom(ctx echo.Context) error {

	data, err := u.Repo.GetAllRoomRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "loi khi lay du lieu",
		})
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{
		"message": "succesful",
		"data":    data,
	})
}

func (u CustomerHandler) SelectRoom(ctx echo.Context) error {

	room_id, err := strconv.Atoi(ctx.QueryParam("room_id"))
	if err != nil {
		fmt.Println("failed to get room_id")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get room_id",
		})
	}

	data, err := u.Repo.SelectRoomRepo(ctx.Request().Context(), room_id)
	if err != nil {
		fmt.Println("failed to get room information")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get room information",
		})
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{
		"message": "succesful",
		"data":    data,
	})
}

func (u CustomerHandler) BookingRoom(ctx echo.Context) error {
	data := model.BookingRoom{}

	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
			Data:       nil,
		})
	}

	_, err := u.Repo.BookingRoomRepo(ctx.Request().Context(), data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to booking room",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Booking successfully",
		Data:       data,
	})
}
