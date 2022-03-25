package main

import (
	"time"
)

type Post struct {
	title string
	body  string
	date  string
}

type User struct {
	name       string
	pass       string
	experience []Experience
}

type Exp struct {
	description string
	position    string
	company     string
	from        time.Time
	term        string
	to          time.Time
}

type any = interface{}

type ok any

type Experience interface {
	AddExp(desc, position, company, term string) error
	GetExperienceAll() ([]Exp, error)
	GetExperience() ([]Exp, error)
}

func (u User) AddExp(desc, pos, comp, term string) error {

	//todo
	return nil
}

func (p Post) NewPost(title, body string) {
	//todo
}
