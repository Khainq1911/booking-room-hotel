package repository

import (
	"booking-website-be/model"
	"context"
)

type CustomerRepo interface {
	SaveAccountRepo(account model.SaveAccount, ctx context.Context) (model.SaveAccount, error)
	CheckSignInRepo(ctx context.Context, userName string) (model.SignIn, error)
}
