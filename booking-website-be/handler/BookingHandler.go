package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	BookingRepo repository.BookingRepo
}

func (u *BookingHandler) CreateBooking(ctx echo.Context) error {
	req := model.CreateBooking{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.BookingRepo.CreateBookingRepo(ctx.Request().Context(), req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create booking",
		})
	}

	return ctx.JSON(http.StatusOK, model.ResWithOutData{
		StatusCode: http.StatusOK,
		Message:    "successful",
	})
}
