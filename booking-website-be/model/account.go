package model

import "time"

type SignIn struct {
	User_id  int    `json:"user_id" db:"user_id"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type Request struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

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
	CreateTime  time.Time `json:"createtime" db:"createtime"`
	CreateBy    string    `json:"createby" db:"createby"`
}
