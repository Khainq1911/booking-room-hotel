package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"time"
)

type BookingRepo interface {
	CreateBookingRepo(ctx context.Context, booking model.CreateBooking) error
}

type BookingSql struct {
	Sql *database.Sql
}

func NewBookingRepo(sql *database.Sql) BookingRepo {
	return &BookingSql{
		Sql: sql,
	}
}

func (db *BookingSql) CreateBookingRepo(ctx context.Context, booking model.CreateBooking) error {
	query := `insert into booking (customer_id, room_id, booking_date, 
	check_in_date, check_out_date, total_price, status, payment_status, note,
	employee_id, createtime, createby) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, booking.CustomerID, booking.RoomID, booking.BookingDate, booking.CheckInDate,
		booking.CheckOutDate, booking.TotalPrice, booking.Status, booking.PaymentStatus, booking.Note, booking.EmployeeId, current, booking.CreateBy); err != nil {
		return err
	}

	return nil
}
