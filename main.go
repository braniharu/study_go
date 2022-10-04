package main

import (
	"errors"
	"fmt"
)

type Flight interface {
	Fly() error
}

func MakeFlight(flight Flight) error {
	if flight.Fly() == nil {
		return errors.New("функция принимает nil: ")
	}
	return errors.New("вызвана ошибка: ")
}

type Bird struct{}

type Airplane struct{}

func (bird Bird) Fly() error {
	return errors.New("ошибка полёта птицы")
}

func (airplane Airplane) Fly() error {
	return errors.New("ошибка полёта самолёта")
}

func main() {
	var parrot Bird
	var boeing Airplane

	errOne := MakeFlight(parrot)
	fmt.Println(errOne, parrot.Fly())
	if errOne != nil {
		fmt.Println(errOne, parrot.Fly())
		return
	}

	errTwo := MakeFlight(boeing)
	fmt.Println(errTwo, boeing.Fly())
	if errTwo != nil {
		fmt.Println(errTwo, boeing.Fly())
		return
	}
}
