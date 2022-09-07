package main

import (
	"fmt"
	"time"
)

var chForks [5]chan bool
var chPhils [5]chan bool
var test chan bool

func main() {

	for i := 0; i < 5; i++ {
		chForks[i] = make(chan bool)
		chPhils[i] = make(chan bool)
	}
	for i := 0; i < 5; i++ {
		go philosopher(i)
		//go fork(i)
		chForks[i] <- true
	}

	time.Sleep(30000 * time.Millisecond)
}

func philosopher(i int) {

	for {
		var forkRight bool
		forkLeft := <-chForks[i]

		if i == 0 {
			forkRight = <-chForks[4]
		} else {
			forkRight = <-chForks[i-1]
		}

		if forkLeft && forkRight {
			chForks[i] <- false
			if i == 0 {
				chForks[4] <- false
			} else {
				chForks[i-1] <- false
			}
			fmt.Print("Imma eating namanamanm ")
			fmt.Print(i)
			time.Sleep(1000 * time.Millisecond)
		}
	}

}

//func fork(int i) {
//	for {

//	}
//}
