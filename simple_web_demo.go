package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, the road is very long!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("127.0.0.1:8080/hello is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}

