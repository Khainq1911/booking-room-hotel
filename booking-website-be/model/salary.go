package model

import "time"

type CreateSalary struct {
	//SalaryID    string    `json:"salary_id" db:"salary_id"`
	EmployeeID  string    `json:"employee_id" db:"employee_id"`
	BaseSalary  float64   `json:"base_salary" db:"base_salary" `
	Bonus       float64   `json:"bonus" db:"bonus"`
	Allowance   float64   `json:"allowance" db:"allowance"`
	Deduction   float64   `json:"deduction" db:"deduction"`
	NetSalary   float64   `json:"net_salary" db:"net_salary" `
	PaymentDate time.Time `json:"payment_date" db:"payment_date"`
	Note        string    `json:"note" db:"note"`
	CreateBy    string    `json:"createby" db:"createby"`
}
type UpdateSalary struct {
	//SalaryID    string    `json:"salary_id" db:"salary_id"`
	EmployeeID  string    `json:"employee_id" db:"employee_id"`
	BaseSalary  float64   `json:"base_salary" db:"base_salary" `
	Bonus       float64   `json:"bonus" db:"bonus"`
	Allowance   float64   `json:"allowance" db:"allowance"`
	Deduction   float64   `json:"deduction" db:"deduction"`
	NetSalary   float64   `json:"net_salary" db:"net_salary" `
	PaymentDate time.Time `json:"payment_date" db:"payment_date"`
	Note        string    `json:"note" db:"note"`
	UpdateBy    string    `json:"updateby" db:"updateby"`
}
type Salary struct {
	SalaryID    string    `json:"salary_id" db:"salary_id"`
	EmployeeID  string    `json:"employee_id" db:"employee_id"`
	BaseSalary  float64   `json:"base_salary" db:"base_salary" `
	Bonus       float64   `json:"bonus" db:"bonus"`
	Allowance   float64   `json:"allowance" db:"allowance"`
	Deduction   float64   `json:"deduction" db:"deduction"`
	NetSalary   float64   `json:"net_salary" db:"net_salary" `
	PaymentDate time.Time `json:"payment_date" db:"payment_date"`
	Note        string    `json:"note" db:"note"`
}
