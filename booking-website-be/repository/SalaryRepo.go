package repository

import "booking-website-be/database"

type SalaryRepo interface {
}

type SalarySql struct {
	Sql *database.Sql
}

func NewSalaryRepo(sql *database.Sql) SalaryRepo {
	return &SalarySql{
		Sql: sql,
	}
}
