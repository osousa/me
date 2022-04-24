package main

import (
	"reflect"
)

type User struct {
	Id         int     `db:"id"`
	Name       *string `db:"username"`
	Pass       *string `db:"password"`
	Experience *[]Experience
}

type Post struct {
	Id    int     `db:"id"`
	Title *string `db:"title"`
	Body  *string `db:"body"`
	Date  *string `db:"date"`
}

type Experience struct {
	Id          int     `db:"id"`
	Description *string `db:"desc"`
	Position    *string `db:"position"`
	Company     *string `db:"company"`
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
	Save() error
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (u *User) Save() error {
	err := DB.UpdateRow(u)
	if err != nil {
		return err
	}
	return nil
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (e *Experience) Save() error {
	err := DB.UpdateRow(e)
	if err != nil {
		return err
	}
	return nil
}

// To save a variable of type struct to the database , we use the save() method
// Notice that save wont work if called upon a struct pointer that has not been
// previously created using  Create()  or does not exist at all in our database
func (e *Post) Save() error {
	err := DB.UpdateRow(e)
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
	tmp := &User{0, new(string), new(string), new([]Experience)}

	err := DB.GetById(tmp, user_id)
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
	tmp := &Post{0, new(string), new(string), new(string)}

	err := DB.GetById(tmp, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func (e *Post) GetList(id int, list *[]interface{}) error {
	tmp := &Post{0, new(string), new(string), new(string)}
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
	tmp := &User{0, new(string), new(string), nil}

	err := DB.GetList(tmp, list, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func (e *Experience) GetList(id int, list *[]interface{}) error {
	tmp := &Experience{0, new(string), new(string), new(string)}
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
	tmp := &Experience{0, new(string), new(string), new(string)}

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
