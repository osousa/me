package main

import (
	//"fmt"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type db struct {
	db        *sql.DB
	connected bool
	name      string
}

func (d *db) setConnectedState(state bool) {
	d.connected = state
}

func (d db) GetConnState() bool {
	return d.connected
}

func ConnectDB(name string) (*db, error) {
	dbase, err := sql.Open("mysql", "osousa:password@tcp(127.0.0.1:3306)/osousa")
	if err != nil {
		panic("Unable to connect to db")
	}
	db := &db{db: dbase, connected: false, name: "osousa"}

	db.setConnectedState(true)
	return db, nil
}

// Sets struct fields' values, given the mysql type, field name
// and arg to set the field pointer ptr. The ptr value must be
// the address of the field pointer, thus deferencing twice
// Example:
// in_type: VARCHAR
// field: Name (struct field)
// arg: John
// ptr: 0x0000958
func SetElem(in_type, field string, arg, ptr interface{}) error {
	arg1 := reflect.ValueOf(arg).Elem()
	ptr1 := reflect.Indirect(reflect.Indirect(ptr.(reflect.Value)))

	// Return if arg is nil, nothing to set on ptr value
	// All vars / struct fields should be pointers
	if arg1.Interface() == nil {
		return nil
	}

	switch in_type {
	case "VARCHAR":
		ptr1.SetString(string(arg1.Interface().([]byte)))
	case "TEXT":
		ptr1.SetString(string(arg1.Interface().([]byte)))
	case "INT":
		ptr1.SetInt(int64(arg1.Interface().(int64)))
	default:
		return errors.New(fmt.Sprintf("Database Type unknown:%s\n", in_type))
	}
	return nil
}

// To extract the values of each struct's field , we must
// provide the following function with the struct, and a
// reflect.StructField parameter as the latter contains
// information regarding the type. There is no way around
// to extract the concrete type of each field, you must use
// reflect.StructField
// Add types as needed
func structToStuctFieldString(structure interface{}, strField reflect.StructField) string {
	ptr1 := reflect.ValueOf(structure)
	ret := new(string)
	switch strField.Type.String() {
	case "int":
		*ret = strconv.Itoa(ptr1.Elem().FieldByName(strField.Name).Interface().(int))
	case "*string":
		ret = ptr1.Elem().FieldByName(strField.Name).Interface().(*string)
	}
	return *ret
}

// structure parameter must be an address pointing to a struct type
// and its fields should be pointers, otherwise it'll throw an error
// Make sure the reflect object is settable, no CanSet() is checked
// It should work with any struct with an Id associated. We can't get
// column names from *Row, only *Rows, thus we use Query. Structs must
// have the following tag (the tag should correspond to a database col):
// type Abc struct{
//		var Example *string `db: "example_column"`
// }
func (d *db) GetById(structure interface{}, id int) error {
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()

	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	columns, fields, _ := structToSlices(structure)

	row, err := d.db.Query("SELECT "+strings.Join(columns[:], ", ")+" FROM "+struct_name+" WHERE ID = ?", strconv.Itoa(id))
	defer row.Close()
	if err != nil || err == sql.ErrNoRows {
		panic(err.Error())
	}

	colTypes, err := row.ColumnTypes()
	values := make([]interface{}, len(columns))
	scan_args := make([]interface{}, len(columns))
	for i := range values {
		scan_args[i] = &values[i]
	}
	if row.Next() {
		err = row.Scan(scan_args...)
		if err != nil {
			panic(err.Error())
		}
	} else {
		return fmt.Errorf("%s object with Id %d does not exist", struct_name, id)
	}
	for i, arg := range scan_args {
		err := SetElem(colTypes[i].DatabaseTypeName(), fields[i], arg, structPtr.Elem().FieldByName(fields[i]).Addr())
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func structToSlices(structure interface{}) ([]string, []string, []reflect.StructField) {
	structPtr := reflect.TypeOf(structure)
	columns := []string{}
	fields := []string{}
	values := make([]reflect.StructField, 0)
	for i := 0; i < structPtr.Elem().NumField(); i++ {
		col := structPtr.Elem().Field(i).Tag.Get("db")
		if col != "" {
			//fmt.Println(structPtr.Field(i))
			fields = append(fields, structPtr.Elem().Field(i).Name)
			values = append(values, structPtr.Elem().Field(i))
			columns = append(columns, "`"+col+"`")
		}
	}

	return columns, fields, values
	//vals := make(map[string]string)
	//for i := 0; i < structPtr.Elem().NumField(); i++ {
	//	col := structPtr.Elem().Type().Field(i).Tag.Get("db")
	//	if col != "" {
	//		field := structPtr.Elem().Type().Field(i).Name
	//		vals[field] = "`" + col + "`"
	//	}
	//}
}

func structFieldFromTag(sFieldSlice []reflect.StructField, tag_key, tag_val string) reflect.StructField {
	var alias reflect.StructField
	for i := 0; i < len(sFieldSlice); i++ {
		field := sFieldSlice[i]
		if field.Tag.Get(tag_key) == tag_val {
			alias = field
		} else {
			continue
		}
	}
	return alias
}

func interfaceSlice(strlst []string) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(strlst))
	for i, d := range strlst {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func (db *db) InsertRow(structure interface{}) error {
	structPtr := reflect.ValueOf(structure)
	struct_name := structPtr.Type().Elem().Name()
	columns, _, vals := structToSlices(structure)
	vals_str := make([]string, 0)
	if structPtr.Type().Kind() != reflect.Ptr {
		return errors.New("You must Dereference Struct")
	}

	params := func(columns []string) string {
		x := make([]string, 0)
		for i := 0; i < len(columns); i++ {
			x = append(x, columns[i]+"=?")
		}
		return strings.Join(x[:], ", ")
	}

	for i := 0; i < len(columns); i++ {
		vals_str = append(vals_str, structToStuctFieldString(structure, vals[i]))
	}
	values := append(interfaceSlice(vals_str), structPtr.Elem().FieldByName(structFieldFromTag(vals, "db", "id").Name).Interface())
	res, err := db.db.Exec("UPDATE "+struct_name+" SET "+params(columns)+" WHERE id = ?", values...)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)
	return nil
}
