package main

import "fmt"

type Flight interface {
	Fly()
}

// werwe
type Bird struct {
}

type Airplane struct {
}

func (bird Bird) Fly() {
	fmt.Println("Птица летит")
}

func (airplane Airplane) Fly() {
	fmt.Println("Самолёт летит")
}

func MakeFlight(flight Flight) {
	flight.Fly()
}

func main() {
	bird := Bird{}
	MakeFlight(bird)

	airplane := Airplane{}
	MakeFlight(airplane)
}
