package router

import (
	"booking-website-be/handler"
	"booking-website-be/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo            *echo.Echo
	AccountHandler  handler.AccountHandler
	TypeRoomHandler handler.TypeRoomHandler
	RoomHandler     handler.RoomHandler
	BookingHandler  handler.BookingHandler
	EmployeeHandler handler.EmployeeHandler
	SalaryHandler   handler.SalaryHandler
	PaymentHandler  handler.PaymentHandler
}

func (api *Api) SetupRouter() {
	adminGroup := api.Echo.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("Ftghghttfhgt44"),
	}))
	adminGroup.Use(middleware.AdminMiddleware)
	//customer routes
	api.Echo.POST("/customer/create", api.AccountHandler.CreateCustomer)
	adminGroup.GET("/customers", api.AccountHandler.ViewCusList)
	api.Echo.GET("/customers/:customer_id", api.AccountHandler.ViewCusDetail)
	api.Echo.PUT("/customers/update/:customer_id", api.AccountHandler.UpdateCus)
	api.Echo.PUT("/customers/delete/:customer_id", api.AccountHandler.DeleteCus)

	//typeroom routes
	api.Echo.POST("typeRoom/add", api.TypeRoomHandler.AddTypeRoom)
	api.Echo.GET("typeRoom", api.TypeRoomHandler.ViewTypeRoom)
	api.Echo.GET("typeRoom/:type_id", api.TypeRoomHandler.ViewDetailTypeRoom)
	api.Echo.PUT("/typeRoom/:type_id/update", api.TypeRoomHandler.UpdateTypeRoom)
	api.Echo.PUT("/typeRoom/:type_id/delete", api.TypeRoomHandler.DeleteTypeRoom)
	api.Echo.GET("/typeRoom/filter", api.TypeRoomHandler.FilterTypeRoom)

	//employee routes
	api.Echo.POST("employee/create", api.EmployeeHandler.CreateEmployee)
	api.Echo.GET("/employee", api.EmployeeHandler.ViewListEmp)
	api.Echo.GET("/employee/:employee_id", api.EmployeeHandler.ViewDetailEmp)
	api.Echo.PUT("/employee/:employee_id/update", api.EmployeeHandler.UpdateEmp)
	api.Echo.PUT("/employee/:employee_id/delete", api.EmployeeHandler.DeleteEmp)

	//rooms routes
	api.Echo.GET("rooms", api.RoomHandler.ViewListRoom)
	api.Echo.GET("rooms/:room_id", api.RoomHandler.ViewDetailRoom)
	api.Echo.POST("/rooms/add", api.RoomHandler.AddRoom)
	api.Echo.PUT("/rooms/:room_id/update", api.RoomHandler.UpdateRoom)
	api.Echo.PUT("/rooms/:room_id/delete", api.RoomHandler.DeleteRoom)

	//booking
	api.Echo.POST("/booking/create", api.BookingHandler.CreateBooking)
	api.Echo.GET("/booking", api.BookingHandler.ViewListBooking)
	api.Echo.GET("/booking/:booking_id", api.BookingHandler.ViewDetailBooking)
	api.Echo.PUT("/booking/:booking_id/cancel", api.BookingHandler.CancelBooking)

	//salary
	api.Echo.POST("/salary/create", api.SalaryHandler.CreateSalary)
	api.Echo.GET("/salary", api.SalaryHandler.ViewListSalary)
	api.Echo.GET("/salary/:salary_id", api.SalaryHandler.ViewDetailSalary)
	api.Echo.PUT("/salary/:salary_id/update", api.SalaryHandler.UpdateSalary)

	//payment
	api.Echo.POST("/payment/create", api.PaymentHandler.CreatePayment)
	api.Echo.GET("/payment", api.PaymentHandler.ViewListPayment)
	api.Echo.GET("/payment/:payment_id", api.PaymentHandler.ViewDetailPayment)
	api.Echo.PUT("/payment/:payment_id/update", api.PaymentHandler.UpdatePayment)

	api.Echo.POST("/login", api.EmployeeHandler.CheckLogin)
}
