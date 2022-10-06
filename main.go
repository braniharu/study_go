package main

import (
	"fmt"
	"time"
)

var Bird = [3]string{"Попугай", "Утка", "Ястреб"}
var Airplane = []string{"Boeing", "British Aerospace", "EADS Socata", "Lancair"}

func MakeFlightBird(Bird [3]string) {
	for i := range Bird {
		Bird[i] = Bird[i] + " летит"
		fmt.Println(Bird[i])
	}
}

func MakeFlightAirplane(Airplane []string) {
	for i := range Airplane {
		Airplane[i] = Airplane[i] + " летит"
		fmt.Println(Airplane[i])
	}
}

func main() {
	//for i := range Bird {
	//	go MakeFlightBird(Bird)
	//}
	go MakeFlightBird(Bird)
	go MakeFlightAirplane(Airplane)

	duration := time.Second
	time.Sleep(duration)
}
