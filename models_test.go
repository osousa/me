package main

import (
	mockdb "me/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockDatabase(ctrl)
	user := NewUser(0, "", "", nil, store)
	store.EXPECT().
		GetById(gomock.Any(), 0).Times(1).Return(errors.Errorf("Error"))

	err := user.GetById(0)
	if strings.Contains(err.Error(), "Errorooo") {
		t.Errorf("Abs(-1) = %d; want 1", err)
	}
}

func TestGetPostById(t *testing.T) {

	var PostTest = []struct {
		testname string
		id       int
		title    string
		body     string
		date     string
		url      string
		abstract string
		err      error
		db       bool
	}{
		{"Correct", 0, "Post title", "Post body", "2022-23-02 00:00:00", "post-title", "lorem ispum abstract", nil, true},
	}

	for _, tt := range PostTest {
		t.Run(tt.testname, func(t *testing.T) {
			var post *Post = nil
			if tt.db {
				ctrl := gomock.NewController(t)
				store := mockdb.NewMockDatabase(ctrl)
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, store)
				store.EXPECT().GetById(gomock.Any(), 0).Times(1).Return(nil)
			} else {
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, nil)
			}
			err := post.GetById(post.Id)
			if err != nil && !strings.Contains(tt.err.Error(), err.Error()) {
				t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
			}
		})
	}
}

func TestSaveUser(t *testing.T) {
	var UserTest = []struct {
		testname string
		username string
		password string
		err      error
		db       bool
	}{
		{"correct", "userone", "password", nil, true},
		{"error1", "userone", "", errors.New("Error! Password is too short"), false},
		{"error2", "", "password", errors.New("Error! Username is too short"), true},
	}

	for _, tt := range UserTest {
		t.Run(tt.testname, func(t *testing.T) {
			var user *User = nil
			if tt.db {
				ctrl := gomock.NewController(t)
				store := mockdb.NewMockDatabase(ctrl)
				user = NewUser(0, tt.username, tt.password, nil, store)
				store.EXPECT().UpdateRow(user).Return(nil)
			} else {
				user = NewUser(0, tt.username, tt.password, nil, nil)
			}
			err := user.Save()
			if err != nil && !strings.Contains(tt.err.Error(), err.Error()) {
				t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
			}
		})
	}
}

func TestGetList(t *testing.T) {
	//TODO
}
