package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"log"
	"time"
)

type EmployeeRepo interface {
	CreateEmpRepo(ctx context.Context, employee model.CreateEmp) error
	ViewListEmpRepo(ctx context.Context) ([]model.Employee, error)
	ViewDetailEmpRepo(ctx context.Context, employeeId string) ([]model.Employee, error)
	UpdateEmpRepo(ctx context.Context, employee_id string, employee model.UpdateEmp) error
	DeleteEmpRepo(ctx context.Context, employee_id string, model model.DeleteEmp) error
	CheckLogin(ctx context.Context, phone_number string) ([]model.Login, error)
}

type EmployeeSql struct {
	Sql *database.Sql
}

func NewEmployeeRepo(sql *database.Sql) EmployeeRepo {
	return &EmployeeSql{
		Sql: sql,
	}
}

//employee

// create employee
func (db *EmployeeSql) CreateEmpRepo(ctx context.Context, employee model.CreateEmp) error {
	query := `  insert into employee (full_name, 
				email, 
				phone_number, 
				address, 
				position, 
				salary,
				hire_date, 
				date_of_birth,
				id_document,
				status, 
				note, 
				createtime, 
				createby) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	current := time.Now()

	if _, err := db.Sql.Db.Exec(query,
		employee.FullName,
		employee.Email,
		employee.PhoneNumber,
		employee.Address,
		employee.Position,
		employee.Salary,
		employee.HireDate,
		employee.DateOfBirth,
		employee.IdDocument,
		employee.Status,
		employee.Note,
		current,
		employee.CreateBy); err != nil {
		return err
	}
	return nil
}

// view list employee

func (db *EmployeeSql) ViewListEmpRepo(ctx context.Context) ([]model.Employee, error) {
	data := []model.Employee{}
	query := `select employee_id, full_name, email, phone_number, address, position, salary, hire_date, date_of_birth, id_document, status, note
	from employee where is_deleted = false`

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Employee{}, err
	}
	return data, nil
}

// view detail
func (db *EmployeeSql) ViewDetailEmpRepo(ctx context.Context, employeeId string) ([]model.Employee, error) {
	data := []model.Employee{}

	query := `select employee_id, full_name, email, phone_number, address, position, salary, hire_date, date_of_birth, id_document, status, note
	from employee where employee_id = $1 `

	if err := db.Sql.Db.Select(&data, query, employeeId); err != nil {
		return []model.Employee{}, err
	}

	return data, nil
}

// update
func (db *EmployeeSql) UpdateEmpRepo(ctx context.Context, employee_id string, employee model.UpdateEmp) error {
	query := `UPDATE employee
		SET full_name = $2, email = $3, phone_number = $4, address = $5, position = $6, salary = $7, hire_date = $8,
		    date_of_birth = $9, id_document = $10, status = $11, note = $12, updatetime = $13, updateby = $14
		WHERE employee_id = $1`

	current := time.Now()
	_, err := db.Sql.Db.Exec(query, employee_id, employee.FullName, employee.Email, employee.PhoneNumber, employee.Address, employee.Position, employee.Salary, employee.HireDate, employee.DateOfBirth, employee.IdDocument, employee.Status, employee.Note, current, employee.UpdateBy)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// delete
func (db *EmployeeSql) DeleteEmpRepo(ctx context.Context, employee_id string, model model.DeleteEmp) error {
	query :=
		`update employee
	set deleteby = $2, deletetime = $3, is_deleted = true where employee_id = $1`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, employee_id, model.DeleteBy, current); err != nil {
		return err
	}

	return nil
}

// login
func (db *EmployeeSql) CheckLogin(ctx context.Context, phone_number string) ([]model.Login, error) {
	data := []model.Login{}
	query := `select employee_id, password, is_admin from employee where phone_number = $1`

	if err := db.Sql.Db.Select(&data, query, phone_number); err != nil {
		return []model.Login{}, err
	}

	return data, nil
}
