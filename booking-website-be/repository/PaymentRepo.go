package repository

import "booking-website-be/database"

type PaymentRepo interface {
}

type PaymentSql struct {
	Sql *database.Sql
}

func NewPaymentRepo(sql *database.Sql) PaymentRepo {
	return &PaymentSql{
		Sql: sql,
	}
}
