package main

import (
	"booking-website-be/database"
	"booking-website-be/handler"
	"booking-website-be/repository"

	"booking-website-be/router"

	"github.com/labstack/echo/v4"
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

	CustomerDb := handler.CustomerHandler{
		Repo: repository.NewCustomerRepo(sql),
	}
	AdminDb := handler.AdminHandler{
		Repo: repository.NewAdminRepo(sql),
	}

	api := router.Api{
		Echo:    e,
		CustomerHandler: CustomerDb,
		AdminHandler: AdminDb,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1912"))
}
