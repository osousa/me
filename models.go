package main

type Post struct {
	title string
	body  string
	date  string
}

type User struct {
	Id         int     `db:"id"`
	Name       *string `db:"username"`
	Pass       *string `db:"password"`
	Experience *[]Experience
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
	GetById(int) error
	Save() error
}

func (u *User) Save() error {
	DB.InsertRow(u)
	return nil
}

func (e *Experience) Save() error {
	return nil
}

func (u *User) GetById(user_id int) error {
	tmp := &User{0, new(string), new(string), new([]Experience)}

	err := DB.GetById(tmp, user_id)
	if err != nil {
		return err
	}

	*u = *tmp
	return nil
}

func (e *Experience) GetById(id int) error {
	tmp := &Experience{0, new(string), new(string), new(string)}

	err := DB.GetById(tmp, id)
	if err != nil {
		return err
	}
	*e = *tmp
	return nil
}

func GetById(o Object, id int) error {
	return o.GetById(id)
}
