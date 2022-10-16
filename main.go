package main

import (
	"fmt"
	"sync"
	"time"
)

type Flight interface {
	Fly()
}

type Bird struct{}
type Airplane struct{}

var counter = 0
var mutex sync.Mutex

func (bird Bird) Fly() {
	for {
		mutex.Lock()
		counter++
		fmt.Println("птица летит", counter)

		if counter >= 25 {
			counter = 0
		}
		mutex.Unlock()

		break
	}
}

func (airplane Airplane) Fly() {
	for {
		mutex.Lock()
		counter++
		fmt.Println("самолёт летит", counter)

		if counter >= 25 {
			counter = 0
		}
		mutex.Unlock()

		break
	}
}

func MakeFlight(flight Flight) {
	flight.Fly()
}

func main() {
	var bird [1000]Bird
	var airplane = make([]Airplane, 1000)

	for i := range bird {
		go MakeFlight(bird[i])
	}

	for j := range airplane {
		go MakeFlight(airplane[j])
	}

	time.Sleep(time.Second * 2)
}
