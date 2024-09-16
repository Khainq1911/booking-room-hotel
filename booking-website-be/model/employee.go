package model

import "time"

type CreateEmp struct {
	FullName    string    `json:"full_name" db:"full_name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Address     string    `json:"address" db:"address"`
	Position    string    `json:"position" db:"position"`
	Salary      float32   `json:"salary" db:"salary"`
	HireDate    time.Time `json:"hire_date" db:"hire_date"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	IdDocument  string    `json:"id_document" db:"id_document"`
	Status      string    `json:"status" db:"status"`
	Note        string    `json:"note" db:"note"`
	CreateBy    string    `json:"createby" db:"createby"`
}

type Employee struct {
	EmployeeID  string     `json:"employee_id" db:"employee_id"`
	FullName    string     `json:"full_name" db:"full_name"`
	Email       string     `json:"email" db:"email"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	Address     string     `json:"address" db:"address"`
	Position    string     `json:"position" db:"position"`
	Salary      float64    `json:"salary" db:"salary"`
	HireDate    time.Time  `json:"hire_date" db:"hire_date"`
	DateOfBirth *time.Time `json:"date_of_birth" db:"date_of_birth"`
	IDDocument  string     `json:"id_document" db:"id_document"`
	Status      string     `json:"status" db:"status"`
	Note        string     `json:"note" db:"note"`
}

type UpdateEmp struct {
	FullName    string    `json:"full_name" db:"full_name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Address     string    `json:"address" db:"address"`
	Position    string    `json:"position" db:"position"`
	Salary      float32   `json:"salary" db:"salary"`
	HireDate    time.Time `json:"hire_date" db:"hire_date"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	IdDocument  string    `json:"id_document" db:"id_document"`
	Status      string    `json:"status" db:"status"`
	Note        string    `json:"note" db:"note"`
	UpdateBy    string    `json:"updateby" db:"updateby"`
}

type DeleteEmp struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
