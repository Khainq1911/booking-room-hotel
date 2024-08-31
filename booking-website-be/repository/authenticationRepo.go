package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
)

type AuthenticationRepo interface {
	SaveAccountRepo(account model.SaveAccount, ctx context.Context) (model.SaveAccount, error)
	CheckSignInRepo(ctx context.Context, userName string) (model.SignIn, error)
}

type AuthenticationSql struct {
	Sql *database.Sql
}

func NewAuthenticationRepo(sql *database.Sql) AuthenticationRepo {
	return &AuthenticationSql{
		Sql: sql,
	}
}

func (db *AuthenticationSql) SaveAccountRepo(account model.SaveAccount, ctx context.Context) (model.SaveAccount, error) {
	query := "INSERT INTO Users (name, dob, phone, password, email) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.Sql.Db.Exec(query, account.Name, account.Dob, account.Phone, account.Password, account.Email)
	if err != nil {
		fmt.Println(err)
		return model.SaveAccount{}, err
	}
	return account, nil
}

func (db *AuthenticationSql) CheckSignInRepo(ctx context.Context, userName string) (model.SignIn, error) {
	user := model.SignIn{}
	query := "SELECT phone, password, role, user_id FROM Users WHERE phone = $1"

	if err := db.Sql.Db.Get(&user, query, userName); err != nil {
		fmt.Println("error in select user signIn:", err)
		return model.SignIn{}, fmt.Errorf("error in select user signIn: %w", err)
	}

	return user, nil
}
