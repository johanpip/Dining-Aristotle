package main

import (
	"fmt"
	"time"
)

var chForks [5]chan bool
var chUsing [5]chan bool

//var chPhils [5]chan bool

func main() {

	for i := 0; i < 5; i++ {
		chForks[i] = make(chan bool)
		//chPhils[i] = make(chan bool)
		chUsing[i] = make(chan bool)
	}
	for i := 0; i < 5; i++ {
		go fork(i)
		//go philosopher(i)
	}
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 5; i++ {
		//go fork(i)
		go philosopher(i)
		//go fork(i)
		//chForks[i] <- true
	}

	time.Sleep(17000 * time.Millisecond)
}

func philosopher(i int) {
	var eaten int = 0

	for {
		var forkRight bool
		forkLeft := <-chForks[i]

		if i == 0 {
			forkRight = <-chForks[4]
		} else {
			forkRight = <-chForks[i-1]
		}

		if forkLeft && forkRight {
			fmt.Println(i, " is eating ")
			eaten++
			if eaten >= 3 {
				fmt.Println(i, " ate 3 times!")
			}

			time.Sleep(1000 * time.Millisecond)

			fmt.Println(i, " putting forks down to think ")
			chUsing[i] <- false
			if i == 0 {
				chUsing[4] <- false
			} else {
				chUsing[i-1] <- false
			}

			/*chForks[i] <- true
			if i == 0 {
				chForks[4] <- true
			} else {
				chForks[i-1] <- true
			}*/
		}
	}

}

func fork(i int) {
	chForks[i] <- true

	for {
		inUse := <-chUsing[i]

		if !inUse {
			fmt.Println("fork", i, "is free now")
			chForks[i] <- true
		}
	}

}
