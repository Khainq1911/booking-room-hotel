package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"time"
)

type BookingRepo interface {
	CreateBookingRepo(ctx context.Context, booking model.CreateBooking) error
	ViewListBookingRepo(ctx context.Context) ([]model.Booking, error)
	ViewDetailBookingRepo(ctx context.Context, booking_id string) ([]model.Booking, error)
	CancelBookingRepo(ctx context.Context, booking_id string, model model.CancelBooking) error
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

// view list booking
func (db *BookingSql) ViewListBookingRepo(ctx context.Context) ([]model.Booking, error) {
	data := []model.Booking{}

	query := `select booking_id, customer_id, employee_id, room_id, booking_date, check_in_date, check_out_date, total_price, status, payment_status, note from booking where is_canceled = false`
	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Booking{}, err
	}

	return data, nil
}

// view detail booking
func (db *BookingSql) ViewDetailBookingRepo(ctx context.Context, booking_id string) ([]model.Booking, error) {
	data := []model.Booking{}

	query := `select booking_id, customer_id, employee_id, room_id, booking_date, check_in_date, check_out_date, total_price, status, payment_status, note 
	from booking where is_canceled = false and booking_id = $1`
	if err := db.Sql.Db.Select(&data, query, booking_id); err != nil {
		return []model.Booking{}, err
	}

	return data, nil

}

// cancel
func (db *BookingSql) CancelBookingRepo(ctx context.Context, booking_id string, model model.CancelBooking) error {
	query := `update booking
	set is_canceled = true , deletetime = $2, deleteby= $3 where booking_id = $1`
	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, booking_id, current, model.DeleteBy); err != nil {
		return err
	}

	return nil
}
