package orm

import (
	"fmt"
	"reflect"
	"strings"
)

type TestStructure struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func New(name string, age int) TestStructure {
	return TestStructure{
		Name: name,
		Age:  age,
	}
}

func (s *TestStructure) Inspect() {
	// Get the type and value
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// If it's a pointer, dereference it
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	// Loop through fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)        // metadata
		value := v.Field(i)        // actual value
		tag := field.Tag.Get("db") // struct tag

		fmt.Println("Field name:", field.Name)
		fmt.Println("Type:", field.Type)
		fmt.Println("DB tag:", tag)
		fmt.Println("Value:", value.Interface())
		fmt.Println("---")
	}

}

func BuildInsertQuery(model any) (string, []any, error) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	// Dereference pointer
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("BuildInsertQuery: expected struct")
	}

	table := strings.ToLower(t.Name()) + "s"

	var columns []string
	var placeholders []string
	var args []any

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip unexported fields
		if !value.CanInterface() {
			continue
		}

		// Read struct tag
		col := field.Tag.Get("db")
		if col == "" {
			col = strings.ToLower(field.Name)
		}

		// Only handle string, int, bool
		switch value.Kind() {
		case reflect.String, reflect.Int, reflect.Bool:
			columns = append(columns, col)
			placeholders = append(placeholders, "?")
			args = append(args, value.Interface())
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args, nil
}

func SaveUserToDatabase(s TestStructure) {
	sql, args, err := BuildInsertQuery(s)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(err)
}
