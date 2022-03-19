package main

import(
	"time"
)

type Post struct {
	title string
	body string
	date string
}

type Experience struct {
	description string
	position string
	company string
	from time.Time
	term string
	to time.Time
}


func (p Post) NewPost(title, body string) {
	
}