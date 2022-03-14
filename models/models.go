package models

import(
	"time"
)

type Post struct {
	var title string
	var body string
	var date Duration
}

type Experience struct {
	var description string
	var position string
	var company string
	var from Time
	var term string
	var to Time
}


func (p Post) NewPost(title, body string) {
	
}