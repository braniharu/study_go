package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

const N = 20

func birdFy() {
	var counter int
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()

			if counter >= 5 {
				counter = 0
				wg.Done()
			}
		}()

		fmt.Println("Птица летит", counter)
	}

	wg.Wait()
}

func airplaneFly() {
	var counter int
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()

			if counter >= 5 {
				counter = 0
				wg.Done()
			}
		}()

		fmt.Println("Самолёт летит", counter)
	}

	wg.Wait()
}

func main() {

	go birdFy()
	go airplaneFly()

	time.Sleep(3 * time.Second)
}
