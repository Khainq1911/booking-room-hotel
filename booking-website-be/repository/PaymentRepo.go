package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"time"
)

type PaymentRepo interface {
	CreatePaymentRepo(ctx context.Context, payment model.CreatePayment) error
	ViewListPaymentRepo(ctx context.Context) ([]model.Payment, error)
	ViewDetailPaymentRepo(ctx context.Context, payment_id string) ([]model.Payment, error)
	UpdatePaymentRepo(ctx context.Context, payment_id string, payment model.UpdatePayment) error
}

type PaymentSql struct {
	Sql *database.Sql
}

func NewPaymentRepo(sql *database.Sql) PaymentRepo {
	return &PaymentSql{
		Sql: sql,
	}
}

func (db *PaymentSql) CreatePaymentRepo(ctx context.Context, payment model.CreatePayment) error {
	query := `insert into payment (booking_id, payment_date, 
	amount, payment_method, payment_status, note, createtime, createby) values ($1, $2, $3, $4, $5, $6, $7, $8)`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, payment.BookingID, payment.PaymentDate, payment.Amount, payment.PaymentMethod, payment.PaymentStatus, payment.Note, current, payment.CreateBy); err != nil {
		return err
	}

	return nil
}

func (db PaymentSql) ViewListPaymentRepo(ctx context.Context) ([]model.Payment, error) {
	data := []model.Payment{}

	query := `select payment_id, booking_id, payment_date, 
	amount, payment_method, payment_status, note from payment`

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Payment{}, err
	}

	return data, nil
}

func (db PaymentSql) ViewDetailPaymentRepo(ctx context.Context, payment_id string) ([]model.Payment, error) {
	data := []model.Payment{}

	query := `select payment_id, booking_id, payment_date, 
	amount, payment_method, payment_status, note from payment where payment_id = $1`
	if err := db.Sql.Db.Select(&data, query, payment_id); err != nil {
		return []model.Payment{}, err
	}

	return data, nil

}

func (db *PaymentSql) UpdatePaymentRepo(ctx context.Context, payment_id string, payment model.UpdatePayment) error {
	query := ` UPDATE Payment
		SET amount = $1,
		    payment_method = $2,
		    payment_status = $3,
		    note = $4,
			updatetime = $5,
			updateby = $6

		WHERE payment_id = $7;`

	current := time.Now()
	if _, err := db.Sql.Db.Exec(query, payment.Amount, payment.PaymentMethod, payment.PaymentStatus, payment.Note, current, payment.UpdateBy, payment_id); err != nil {
		return err
	}

	return nil
}
