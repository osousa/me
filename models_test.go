package main

import (
	mockdb "me/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	//fake "github.com/brianvoe/gofakeit/v6"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockDatabase(ctrl)
	user := NewUser(0, "", "", nil, store)
	store.EXPECT().
		GetById(gomock.Any(), 0).Times(1).Return(nil)

	user.GetById(0)
}

func TestGetPostById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockDatabase(ctrl)
	post := NewPost(0, "", "", "", "", "", nil, store)
	store.EXPECT().
		GetById(gomock.Any(), 0).Times(1).Return(nil)

	post.GetById(0)
}

func TestGetList(t *testing.T) {
	//TODO
}
