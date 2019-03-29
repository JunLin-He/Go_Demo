package main

import (
	"fmt"
)

type empty interface {
}

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
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
