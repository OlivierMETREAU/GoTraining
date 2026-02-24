package orm

import (
	"fmt"
	"reflect"
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
