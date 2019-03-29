package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main() {
	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a == nil)
}

func Disconnect(usb USB) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknowed device.")
	}
}
