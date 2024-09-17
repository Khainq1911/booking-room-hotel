package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"time"
)

type SalaryRepo interface {
	CreateSalary(ctx context.Context, model model.CreateSalary) error
	ViewListSalaryRepo(ctx context.Context) ([]model.Salary, error)
	ViewDetailSalaryRepo(ctx context.Context, salary_id string) ([]model.Salary, error)
	UpdateSalaryRepo(ctx context.Context, salary_id string, model model.UpdateSalary) error
}

type SalarySql struct {
	Sql *database.Sql
}

func NewSalaryRepo(sql *database.Sql) SalaryRepo {
	return &SalarySql{
		Sql: sql,
	}
}

func (db *SalarySql) CreateSalary(ctx context.Context, model model.CreateSalary) error {
	query := `insert into salary (employee_id, base_salary, bonus, allowance, deduction, net_salary, payment_date, note, createtime, createby) values ($1, $2, $3, $4, $5,$6, $7, $8, $9, $10)`

	current := time.Now()
	if _, err := db.Sql.Db.Exec(query, model.EmployeeID, model.BaseSalary, model.Bonus, model.Allowance, model.Deduction, model.NetSalary, model.PaymentDate, model.Note, current, model.CreateBy); err != nil {
		return err
	}
	return nil
}

func (db *SalarySql) ViewListSalaryRepo(ctx context.Context) ([]model.Salary, error) {
	data := []model.Salary{}

	query := `select salary_id, employee_id, base_salary, bonus, allowance, deduction, net_salary, payment_date, note from salary`

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Salary{}, err
	}

	return data, nil
}

func (db *SalarySql) ViewDetailSalaryRepo(ctx context.Context, salary_id string) ([]model.Salary, error) {
	data := []model.Salary{}

	query := `select salary_id, employee_id, base_salary, bonus, allowance, deduction, net_salary, payment_date, note
	from salary where salary_id = $1`
	if err := db.Sql.Db.Select(&data, query, salary_id); err != nil {
		return []model.Salary{}, err
	}

	return data, nil

}

func (db *SalarySql) UpdateSalaryRepo(ctx context.Context, salary_id string, model model.UpdateSalary) error {
	query := `
	update salary set
	employee_id = $1, base_salary = $2, bonus = $3, allowance = $4, deduction = $5, net_salary = $6, payment_date = $7, note = $8, updatetime = $9, updateby=$10
	where salary_id = $11
	`
	current := time.Now()
	if _, err := db.Sql.Db.Exec(query, model.EmployeeID, model.BaseSalary, model.Bonus, model.Allowance, model.Deduction, model.NetSalary, model.PaymentDate, model.Note, current, model.UpdateBy, salary_id); err != nil {
		return err
	}
	return nil
}
