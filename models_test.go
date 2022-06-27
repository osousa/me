package main

import (
	mockdb "me/mocks"
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func randomAccount() User {
	return User{
		Id:         fake.Number(1, 1000),
		Name:       new(string),
		Pass:       new(string),
		Experience: nil,
	}
}

func TestGetById(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockDatabase(ctrl)
	store.EXPECT().
		GetById(gomock.Any(), gomock.Eq(account.Id)).Times(1).Return(nil)

	store.GetById(account, account.Id)
}
