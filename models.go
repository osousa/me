package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Id         int     `db:"id"`
	Name       *string `db:"username"`
	Pass       *string `db:"password"`
	Experience *[]Experience
	db         Database
}

type Post struct {
	Id       int     `db:"id"`
	Title    *string `db:"title"`
	Body     *string `db:"body"`
	Date     *string `db:"date"`
	Url      *string `db:"url"`
	Abstract *string `db:"abstract"`
	db       Database
}

type Experience struct {
	Id          int     `db:"id"`
	Description *string `db:"desc"`
	Position    *string `db:"position"`
	Company     *string `db:"company"`
	db          Database
	// From        string `db:"from"`
	// Term        string `db:"term"`
	// To          string `db:"to"`
}

type Exp interface {
	AddExp(desc, position, company, term string) error
	GetExperienceAll() ([]Experience, error)
	GetExperience() ([]Experience, error)
}

type Object interface {
	GetList(int, *[]interface{}) error
	GetById(int) error
	GetLast() error
	Save() error
}

func NewPost(id int, title, body, date, url, abstract string, db Database) *Post {
	_Title := &title
	_Body := &body
	_Date := &date
	_Url := &url
	_Abstract := &abstract
	return &Post{
		Id:       id,
		Title:    _Title,
		Body:     _Body,
		Date:     _Date,
		Url:      _Url,
		Abstract: _Abstract,
		db:       db,
	}
}

func NewExperience(id int, desc, pos, company string, db Database) *Experience {
	_Description := new(string)
	_Position := new(string)
	_Company := new(string)

	return &Experience{
		Id:          id,
		Description: _Description,
		Position:    _Position,
		Company:     _Company,
		db:          db,
	}
}

func NewUser(id int, username, password string, exp *[]Experience, db Database) *User {
	_username := new(string)
	_password := new(string)
	_username = &username
	_password = &password
	return &User{
		Id:         id,
		Name:       _username,
		Pass:       _password,
		Experience: exp,
		db:         db,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s, Password: %s}", u.Id, *u.Name, *u.Pass)
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (u *User) Save() error {
	var err error
	if len(*u.Pass) < 3 {
		return errors.New("Error! Password is too short")
	}
	if u.db == nil {
		err = DB.UpdateRow(u)
	} else {
		err = u.db.UpdateRow(u)
	}
	if err != nil {
		return err
	}
	return nil
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (e *Experience) Save() error {
	var err error
	if e.db == nil {
		err = DB.UpdateRow(e)
	} else {
		err = e.db.UpdateRow(e)
	}
	if err != nil {
		return err
	}
	return nil
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (e *Post) Save() error {
	var err error
	if e.db == nil {
		err = DB.UpdateRow(e)
	} else {
		err = e.db.UpdateRow(e)
	}
	if err != nil {
		return err
	}
	return nil
}

// GetById method changes the object pointed by e, and for that change to occur
// we must store an address at tmp, then passing it to the DB.GetById as change
// will happen in place , do not try to modify the original variable e . The id
// value should exist on a database as a primary key of the corresponding table
// with the same name as the struct (User in this case)
func (u *User) GetById(user_id int) error {
	tmp := NewUser(0, "", "", nil, nil)
	var err error
	if u.db == nil {
		err = DB.GetById(tmp, user_id)
	} else {
		err = u.db.GetById(tmp, user_id)
	}
	if err != nil {
		return err
	}

	*u = *tmp
	return nil
}

// GetById method changes the object pointed by e, and for that change to occur
// we must store an address at tmp, then passing it to the DB.GetById as change
// will happen in place , do not try to modify the original variable e . The id
// value should exist on a database as a primary key of the corresponding table
// with the same name as the struct (Experience in this case)
func (e *Post) GetById(id int) error {
	tmp := NewPost(0, "", "", "", "", "", nil)

	var err error
	if e.db == nil {
		err = DB.GetById(tmp, id)
	} else {
		err = e.db.GetById(tmp, id)
	}
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func (e *Post) GetLast() error {
	tmp := NewPost(0, "", "", "", "", "", nil)
	str, err := DB.RawQueryRow("select id from Post ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(*str)
	GetById(tmp, id)
	*e = *tmp
	return nil
}

func (e *Post) GetList(id int, list *[]interface{}) error {
	tmp := NewPost(0, "", "", "", "", "", nil)
	err := DB.GetList(tmp, list, id)
	for i, val := range *list {
		v := val.(reflect.Value)
		(*list)[i] = v.Interface().(Post)
	}

	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func (e *User) GetList(id int, list *[]interface{}) error {
	tmp := NewUser(0, "", "", nil, nil)

	err := DB.GetList(tmp, list, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func (e *Experience) GetList(id int, list *[]interface{}) error {
	tmp := NewExperience(0, "", "", "", nil)
	err := DB.GetList(tmp, list, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

// GetById method changes the object pointed by e, and for that change to occur
// we must store an address at tmp, then passing it to the DB.GetById as change
// will happen in place , do not try to modify the original variable e . The id
// value should exist on a database as a primary key of the corresponding table
// with the same name as the struct (Experience in this case)
func (e *Experience) GetById(id int) error {
	tmp := NewExperience(0, "", "", "", nil)

	err := DB.GetById(tmp, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

// Parameter "o" of Type Object accepts any variable that implements the Object
// interface, which happens to be the case for type User and Experience structs
// GetById accepts an o pointer and id int,the Object type variable will change
// in place, no error should return if Get is performed correctly
func GetById(o Object, id int) error {
	return o.GetById(id)
}

func GetLast(o Object) error {
	return o.GetLast()
}

// Parameter "o" of Type Object accepts any variable that implements the Object
// interface, which happens to be the case for type User and Experience structs
// GetById accepts an o pointer and id int,the Object type variable will change
// in place, no error should return if Get is performed correctly
func GetList(o Object, id int) (error, []interface{}) {
	list := make([]interface{}, 0)
	err := o.GetList(id, &list)
	if err != nil {
		return err, nil
	}
	return nil, list
}
