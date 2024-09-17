package model

import "time"

type CreateBooking struct {
	//BookingID     string `json:"booking_id" db:"booking_id"`
	CustomerID    string    `json:"customer_id" db:"customer_id"`
	RoomID        string    `json:"room_id" db:"room_id"`
	BookingDate   time.Time `json:"booking_date" db:"booking_date"`
	CheckInDate   time.Time `json:"check_in_date" db:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date" db:"check_out_date"`
	TotalPrice    float64   `json:"total_price" db:"total_price"`
	Status        string    `json:"status" db:"status"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	Note          string    `json:"note" db:"note"`
	EmployeeId    string    `json:"employee_id" db:"employee_id"`
	CreateBy      string    `json:"createby" db:"createby"`
}

type Booking struct {
	BookingID     string    `json:"booking_id" db:"booking_id"`
	CustomerID    string    `json:"customer_id" db:"customer_id"`
	RoomID        string    `json:"room_id" db:"room_id"`
	BookingDate   time.Time `json:"booking_date" db:"booking_date"`
	CheckInDate   time.Time `json:"check_in_date" db:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date" db:"check_out_date"`
	TotalPrice    float64   `json:"total_price" db:"total_price"`
	Status        string    `json:"status" db:"status"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	Note          string    `json:"note" db:"note"`
	EmployeeId    string    `json:"employee_id" db:"employee_id"`
}

type CancelBooking struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
