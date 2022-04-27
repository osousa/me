package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Home struct {
	name     string
	template *template.Template
}

type Blog struct {
	name     string
	template *template.Template
}

type PostControl struct {
	name     string
	template *template.Template
}

type Contact struct {
	name     string
	template *template.Template
}

type Admin struct {
	name     string
	template *template.Template
}

type About struct {
	name     string
	template *template.Template
}

type Controller interface {
	Execute(io.Writer, map[string]string)
}

func (c Blog) Execute(w io.Writer, fields_map map[string]string) {
	id, _ := strconv.Atoi(fields_map["page"])
	lpost := new(Post)
	//define this at settings level
	postNumber := 4

	err_lst, lst := GetList(lpost, id*(postNumber))

	if err_lst != nil {
		Error.Println(err_lst)
	}

	c.template.ExecuteTemplate(w, c.name, map[string][]interface{}{"Post": lst})
}

func (c PostControl) Execute(w io.Writer, fields_map map[string]string) {
	id, _ := strconv.Atoi(fields_map["id"])
	post := &Post{}
	GetById(post, id)
	c.template.ExecuteTemplate(w, c.name, map[string]interface{}{"post": struct {
		Post *Post
	}{Post: post}})
}

func (c Home) Execute(w io.Writer, fields_map map[string]string) {
	post := new(Post)
	GetById(post, 1)
	c.template.ExecuteTemplate(w, c.name, map[string]interface{}{"p": struct {
		Post *Post
	}{Post: post}})
}

func (c Contact) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, c.name, struct{ Title string }{Title: "ok"})
}

func (c About) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, c.name, struct{ Title string }{Title: "ok"})
}

func (c Admin) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, c.name, struct{ Title string }{Title: "ok"})
}

func NewController(typeCtrl interface{}, name string) interface{} {
	ptrType := reflect.ValueOf(typeCtrl)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New("").Funcs(template.FuncMap{
		"formatDate": func(datetime string) string {
			Error.Println(datetime)
			t, _ := time.Parse("2006-01-02 15:04:00", datetime)
			Error.Println(t)
			return t.Format("January 1, 2006")
		},
	}).ParseFiles(wd+"/web/html/"+name+".tpl", wd+"/web/html/base.tpl")
	if err != nil {
		log.Fatalf("Make sure your NewControl parameter equals a template name: %s", err.Error())
	}

	switch ptrType.Type().String() {
	case "main.Blog":
		return Blog{name, tpl}
	case "main.Home":
		return Home{name, tpl}
	case "main.Contact":
		return Contact{name, tpl}
	case "main.Admin":
		return Admin{name, tpl}
	case "main.PostControl":
		return PostControl{name, tpl}
	case "main.About":
		return About{name, tpl}
	default:
		panic(errors.New(fmt.Sprintf("Unknonwn Controller Type: %s", ptrType.Type().String())))
	}
}
