package model

import "time"

type CreateCus struct {
	FullName         string    `json:"full_name" db:"full_name"`
	Email            string    `json:"email" db:"email"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	Address          string    `json:"address" db:"address"`
	Nationality      string    `json:"nationality" db:"nationality"`
	DateOfBirth      time.Time `json:"date_of_birth" db:"date_of_birth"`
	IDDocument       string    `json:"id_document" db:"id_document"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Note             string    `json:"note" db:"note"`
	CreateBy         string    `json:"createby" db:"createby"`
}

type Customer struct {
	CustomerId       string    `json:"customer_id" db:"customer_id"`
	FullName         string    `json:"full_name" db:"full_name"`
	Email            string    `json:"email" db:"email"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	Address          string    `json:"address" db:"address"`
	Nationality      string    `json:"nationality" db:"nationality"`
	DateOfBirth      time.Time `json:"date_of_birth" db:"date_of_birth"`
	IDDocument       string    `json:"id_document" db:"id_document"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Note             string    `json:"note" db:"note"`
}

type UpdateCus struct {
	FullName         string    `json:"full_name" db:"full_name"`
	Email            string    `json:"email" db:"email"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	Address          string    `json:"address" db:"address"`
	Nationality      string    `json:"nationality" db:"nationality"`
	DateOfBirth      time.Time `json:"date_of_birth" db:"date_of_birth"`
	IDDocument       string    `json:"id_document" db:"id_document"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Note             string    `json:"note" db:"note"`
	UpdateBy         string    `json:"updateby" db:"updateby"`
}

type DeleteCus struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
