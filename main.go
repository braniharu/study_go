package main

import (
	"fmt"
	"time"
)

type Flight interface {
	Fly()
}

type Bird struct {
	BirdName   string
	BirdLenght float64
}

type Airplane struct {
	AirplaneName string
}

func (bird Bird) Fly() {
	fmt.Println(bird.BirdName, "летит c размахом крыла", bird.BirdLenght)
}

func (airplane Airplane) Fly() {
	fmt.Println(airplane.AirplaneName, "летит")
}

func MakeFlight(flight Flight) {
	flight.Fly()
}

func main() {
	var bird [3]Bird
	bird[0] = Bird{BirdName: "Попугай", BirdLenght: 3.1}
	bird[1] = Bird{BirdName: "Утка", BirdLenght: 2.8}
	bird[2] = Bird{BirdName: "Ястреб", BirdLenght: 3.5}

	for i := range bird {
		go MakeFlight(bird[i])
	}

	var airplane []Airplane = make([]Airplane, 2)
	airplane[0] = Airplane{AirplaneName: "Boeing"}
	airplane[1] = Airplane{AirplaneName: "British Aerospace"}
	airplane = append(airplane, Airplane{AirplaneName: "\nLancair"})

	for i := range airplane {
		go MakeFlight(airplane[i])
	}

	duration := time.Second
	time.Sleep(duration)
}
