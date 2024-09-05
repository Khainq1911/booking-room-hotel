package router

import (
	"booking-website-be/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo            *echo.Echo
	CustomerHandler handler.CustomerHandler
	AdminHandler    handler.AdminHandler
	AuthenHandler   handler.AuthenHandler
}

func (api *Api) SetupRouter() {

	//public routes
	api.Echo.POST("/register", api.AuthenHandler.SaveAccount)
	api.Echo.POST("/login", api.AuthenHandler.CheckLogin)
	api.Echo.GET("/rooms", api.CustomerHandler.GetAllRoom)
	api.Echo.GET("/room", api.CustomerHandler.SelectRoom)
	api.Echo.GET("/filter", api.CustomerHandler.FilterRoom)
	/* protected := api.Echo.Group("/") */
	/* protected.Use(middleware.AuthenticateMiddleware) */

	api.Echo.POST("/booking", api.CustomerHandler.BookingRoom)

	// admin routes
	api.Echo.POST("/admin/addRoom", api.AdminHandler.AddInforRoomHandler)
	api.Echo.GET("/admin/bookings_list", api.AdminHandler.GetBookingList)
	api.Echo.GET("/admin/bookings_list", api.AdminHandler.GetDetailBooking)
	api.Echo.PUT("/admin/bookings_list/cancel", api.AdminHandler.CancelBooking)
	api.Echo.DELETE("/admin/rooms/:room_id/delete", api.AdminHandler.DeleteRoom)
	api.Echo.PUT("/admin/rooms/:room_id/update", api.AdminHandler.UpdateRoom)
	api.Echo.PUT("/admin/bookings/:booking_id/update", api.AdminHandler.UpdateBooking)
}
