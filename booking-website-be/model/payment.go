package model

import "time"

type CreatePayment struct {
	//PaymentID     string    `json:"payment_id" db:"payment_id"`
	BookingID     string    `json:"booking_id" db:"booking_id"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
	Amount        float64   `json:"amount" db:"amount"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	Note          string    `json:"note" db:"note"`
	CreateBy      string    `json:"createby" db:"createby"`
}
type UpdatePayment struct {
	Amount        float64 `json:"amount" db:"amount"`
	PaymentMethod string  `json:"payment_method" db:"payment_method"`
	PaymentStatus string  `json:"payment_status" db:"payment_status"`
	Note          string  `json:"note" db:"note"`
	UpdateBy      string  `json:"updateby" db:"updateby"`
}

type Payment struct {
	PaymentID     string    `json:"payment_id" db:"payment_id"`
	BookingID     string    `json:"booking_id" db:"booking_id"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
	Amount        float64   `json:"amount" db:"amount"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	Note          string    `json:"note" db:"note"`
}
