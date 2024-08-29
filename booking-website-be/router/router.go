package router

import (
	"booking-website-be/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo            *echo.Echo
	CustomerHandler handler.CustomerHandler
	AdminHandler    handler.AdminHandler
}

func (api *Api) SetupRouter() {
	api.Echo.POST("/signUp", api.CustomerHandler.SaveAccount)
	api.Echo.POST("/signIn", api.CustomerHandler.CheckLogin)
	/* protected := api.Echo.Group("/") */
	/* protected.Use(middleware.AuthenticateMiddleware) */
	api.Echo.GET("/", api.CustomerHandler.GetAllRoom)
	api.Echo.GET("/room", api.CustomerHandler.SelectRoom)
	api.Echo.POST("/booking", api.CustomerHandler.BookingRoom)

	api.Echo.POST("/admin/addRoom", api.AdminHandler.AddInforRoomHandler)
	api.Echo.GET("/admin/bookings_list", api.AdminHandler.GetBookingList)
	api.Echo.GET("/admin/bookings_list", api.AdminHandler.GetDetailBooking)
	api.Echo.PUT("/admin/bookings_list/cancel", api.AdminHandler.CancelBooking)
}
