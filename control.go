package main

import (
	"errors"
	"html/template"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
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

type Controller interface {
	Execute(io.Writer, map[string]string)
}

func (c Blog) Execute(w io.Writer, fields_map map[string]string) {
	id, _ := strconv.Atoi(fields_map["page"])
	lpost := new(Post)
	//define this at settings level
	postNumber := 5

	err_lst, lst := GetList(lpost, id*(postNumber))

	if err_lst != nil {
		Error.Println(err_lst)
	}

	for _, val := range lst {
		Info.Println(val)
	}

	//post := &Post{}
	//GetById(post, id)
	//Debug.Println(post)
	c.template.ExecuteTemplate(w, "blog", map[string][]interface{}{"Post": lst})
}

func (c PostControl) Execute(w io.Writer, fields_map map[string]string) {
	id, _ := strconv.Atoi(fields_map["id"])
	post := &Post{}
	GetById(post, id)
	Debug.Println(post)
	c.template.ExecuteTemplate(w, "blog", map[string]*Post{"Post": post})
}

func (c Home) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, "layout", struct{ Title string }{Title: "ok"})
}

func (c Contact) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, "layout", struct{ Title string }{Title: "ok"})
}

func (c Admin) Execute(w io.Writer, fields_map map[string]string) {
	c.template.ExecuteTemplate(w, "layout", struct{ Title string }{Title: "ok"})
}

func NewController(typeCtrl interface{}, name string) interface{} {
	ptrType := reflect.ValueOf(typeCtrl)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tpl, err := template.ParseFiles(wd + "/html/" + name + ".tpl")
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

	default:
		panic(errors.New("Unknonwn Controller Type"))
	}
}
