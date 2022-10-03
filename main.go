package main

import (
	"errors"
)

type Bird struct {
	Flight
}

type Airplane struct {
	Flight
}

type Flight interface {
	Fly()
}

func (bird Bird) Fly() error {
	return errors.New("Ошибка полёта птицы")
}

func (airplane Airplane) Fly() error {
	return errors.New("Ошибка полёта самолёта")
}

func MakeFlight(flight Flight) error {
	flight.Fly()
	return errors.New("Вызвана ошибка")
}

func main() {
	//var bird Bird 
	var airplane Airplane
	MakeFlight(airplane.Flight)
	//errOne := MakeFlight(bird.Flight)
	//errTwo := MakeFlight(airplane.Flight)
	////if errOne != nil {
	////	fmt.Println("Птица: ", errOne)
	////	return
	////}
	//if errTwo != nil {
	//	fmt.Println(errTwo)
	//	return
	//}
}
