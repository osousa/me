package main

import (
	mockdb "me/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestGetByIdUser(t *testing.T) {
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

func TestGetByIdPost(t *testing.T) {

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
		{
			"Correct", 0,
			"Post title",
			"Post body",
			"2022-23-02 00:00:00",
			"post-title",
			"lorem ispum abstract",
			nil,
			true,
		},
		{
			"Error",
			-1,
			"Post title",
			"Post body",
			"2022-23-02 00:00:00",
			"post-title",
			"lorem ispum abstract",
			errors.New("cannot be negative"),
			false,
		},
	}

	for _, tt := range PostTest {
		t.Run(tt.testname, func(t *testing.T) {
			var post *Post = nil
			if tt.db {
				ctrl := gomock.NewController(t)
				store := mockdb.NewMockDatabase(ctrl)
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, store)
				store.EXPECT().GetById(gomock.Any(), 0).Times(1).Return(tt.err)
			} else {
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, nil)
			}
			err := post.GetById(post.Id)
			if err != nil && !strings.Contains(err.Error(), tt.err.Error()) {
				t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
			}
		})
	}
}
func TestGetLastPost(t *testing.T) {
	var PostTest = []struct {
		testname string
		id       int
		title    string
		body     string
		date     string
		url      string
		abstract string
		err_gbi  error
		err_rqr  error
		db       bool
	}{
		{"Correct", 0, "", "", "", "", "", nil, nil, true},
		{"Error", 0, "", "", "", "", "", nil, errors.New("cannot connect to DB"), true},
		{"Error2", 0, "", "", "", "", "", errors.New("Cannot connect to DB"), nil, true},
	}

	for _, tt := range PostTest {
		t.Run(tt.testname, func(t *testing.T) {
			var post *Post = nil
			if tt.db {
				ctrl := gomock.NewController(t)
				store := mockdb.NewMockDatabase(ctrl)
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, store)
				stringret := "0"
				if (tt.err_rqr) == nil {
					store.EXPECT().GetById(gomock.Any(), 0).Times(1).Return(tt.err_gbi)
				}
				store.EXPECT().RawQueryRow(gomock.Any()).Times(1).Return(&stringret, tt.err_rqr)
			} else {
				post = NewPost(tt.id, tt.title, tt.body, tt.date, tt.url, tt.abstract, nil)
			}
			err := post.GetLast()
			if err != nil {
				if tt.err_rqr != nil {
					if !strings.Contains(err.Error(), tt.err_rqr.Error()) {
						t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
					}
				}
				if tt.err_gbi != nil {
					if !strings.Contains(err.Error(), tt.err_gbi.Error()) {
						t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
					}
				}

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
		{
			"correct",
			"userone",
			"password",
			nil,
			true,
		},
		{
			"error1",
			"userone",
			"",
			errors.New("Error! Password is too short"),
			false,
		},
		{
			"error2",
			"",
			"password",
			errors.New("Error! Username is too short"),
			true,
		},
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
			if err != nil && !strings.Contains(err.Error(), tt.err.Error()) {
				t.Errorf("Error in %v; Test failed: \"%v\"", tt.testname, err)
			}
		})
	}
}

func TestGetList(t *testing.T) {
	//TODO
}
