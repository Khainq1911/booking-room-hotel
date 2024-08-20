package main

import (
	"booking-website-be/database"
	"booking-website-be/handler"
	repoimplement "booking-website-be/repository/repo-implement"
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

	db := handler.CustomerHandler{
		Repo: repoimplement.NewCustomerRepo(sql),
	}

	api := router.Api{
		Echo:    e,
		Handler: db,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1912"))
}
