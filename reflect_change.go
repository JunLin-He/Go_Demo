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

func main() {
	u := User{007, "Alan", 22}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	// If the o is not pointer interface, we can't change the interface
	//Two condition: 1. changeable; 2. pointer interface
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("Invalid Field By Name")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("HAHA")
	}

}
