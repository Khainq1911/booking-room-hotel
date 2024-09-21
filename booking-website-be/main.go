package main

import (
	"booking-website-be/database"
	"booking-website-be/handler"
	"booking-website-be/repository"

	"booking-website-be/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	sql := &database.Sql{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
		Dbname:   "booking-room-hotel",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.CORS())
	
	TypeRoomDb := handler.TypeRoomHandler{
		TypeRoomRepo: repository.NewTypeRoomRepo(sql),
	}

	AccountHandler := handler.AccountHandler{
		Repo: repository.NewAccountRepo(sql),
	}

	RoomDb := handler.RoomHandler{
		RoomRepo: repository.NewRoomRepo(sql),
	}

	bookingDb := handler.BookingHandler{
		BookingRepo: repository.NewBookingRepo(sql),
	}

	salaryDb := handler.SalaryHandler{
		Repo: repository.NewSalaryRepo(sql),
	}

	employeeDb := handler.EmployeeHandler{
		EmployeeRepo: repository.NewEmployeeRepo(sql),
	}

	paymentDb := handler.PaymentHandler{
		Repo: repository.NewPaymentRepo(sql),
	}
	api := router.Api{
		Echo:            e,
		AccountHandler:  AccountHandler,
		TypeRoomHandler: TypeRoomDb,
		RoomHandler:     RoomDb,
		BookingHandler:  bookingDb,
		SalaryHandler:   salaryDb,
		EmployeeHandler: employeeDb,
		PaymentHandler:  paymentDb,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1912"))
}
