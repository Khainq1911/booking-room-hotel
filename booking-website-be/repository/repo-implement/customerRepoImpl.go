package repoimplement

import (
	"booking-website-be/database"
	"booking-website-be/model"
	repo "booking-website-be/repository"
	"context"
	"fmt"
)

type CustomerRepo struct {
	sql *database.Sql
}

func NewCustomerRepo(sql *database.Sql) repo.CustomerRepo {
	return &CustomerRepo{
		sql: sql,
	}
}

func (db *CustomerRepo) SaveAccountRepo(account model.SaveAccount, ctx context.Context) (model.SaveAccount, error) {
	query := "INSERT INTO Users (name, dob, phone, password, email) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.sql.Db.Exec(query, account.Name, account.Dob, account.Phone, account.Password, account.Email)
	if err != nil {
		fmt.Println(err)
		return model.SaveAccount{}, err
	}
	return account, nil
}

func (db *CustomerRepo) CheckSignInRepo(ctx context.Context, userName string) (model.SignIn, error) {
	user := model.SignIn{}
	query := "SELECT phone, password, role, user_id FROM Users WHERE phone = $1"

	if err := db.sql.Db.Get(&user, query, userName); err != nil {
		fmt.Println("error in select user signIn:", err)
		return model.SignIn{}, fmt.Errorf("error in select user signIn: %w", err)
	}

	return user, nil
}
