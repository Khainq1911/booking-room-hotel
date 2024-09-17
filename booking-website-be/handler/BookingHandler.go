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

// view list booking
func (u *BookingHandler) ViewListBooking(ctx echo.Context) error {
	data, err := u.BookingRepo.ViewListBookingRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u *BookingHandler) ViewDetailBooking(ctx echo.Context) error {
	booking_id := ctx.Param("booking_id")

	data, err := u.BookingRepo.ViewDetailBookingRepo(ctx.Request().Context(), booking_id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

// cancel
func (u *BookingHandler) CancelBooking(ctx echo.Context) error {
	booking_id := ctx.Param("booking_id")
	req := model.CancelBooking{}
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.BookingRepo.CancelBookingRepo(ctx.Request().Context(), booking_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to cancel data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}
