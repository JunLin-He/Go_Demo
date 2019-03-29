package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type", t.Name()) // Print the name of the type of o

	if k := t.Kind(); k != reflect.Struct { // Indicate that the import type is incorrect
		fmt.Println("Error")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)               // Get the field by index
		val := v.Field(i).Interface() // Get the value of the field
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "Ok", 12}
	Info(u)
}
