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

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "OK", 12}, title: "666"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}

/*
	fmt.Printf("#%v\n", t.Field(0))
	->
	#{User  main.User  0 [0] true}
*/
/*
	fmt.Printf("%#v\n", t.Field(0))
	->
	reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10b47e0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
*/
