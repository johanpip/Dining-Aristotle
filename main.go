//The phils will constantly check whether their neighboaring forks are avalible, if they are, they will eat, but if they aren't they will not pick up any forks.
//A phil might eat more than 3 times however, which courses other phils to have to wait, but deadlocks will never happen, as it is random which phils will get to eat next, if there are multiple phils with two forks next to them
//The program will finish once every phil has eaten at least 3 times. At that point, some phils might've eaten way more than 3 times, but eventually they will all eat 3 times, and the program will end.
//So the runtime is all up to the "luck of the draw" ;)

package main

import (
	"fmt"
	"time"
)

var chForks [5]chan bool
var chUsing [5]chan bool
var finished int = 0

//var chPhils [5]chan bool

func main() {

	for i := 0; i < 5; i++ {
		chForks[i] = make(chan bool)
		//chPhils[i] = make(chan bool)
		chUsing[i] = make(chan bool)
	}
	for i := 0; i < 5; i++ {
		go fork(i)
		go philosopher(i)
	}

	for true {
		if finished == 5 {
			return
		}
	}
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
			chUsing[i] <- true
			if i == 0 {
				chUsing[4] <- true
			} else {
				chUsing[i-1] <- true
			}
			fmt.Println(i, " is eating")
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

		} else {
			if forkLeft {
				chUsing[i] <- !forkLeft
				fmt.Println(i, " only had 1 fork ", i, " is putting left fork down")
			} else if forkRight {
				if i == 0 {
					chUsing[4] <- !forkRight
				} else {
					chUsing[i-1] <- !forkRight
				}

				fmt.Println(i, " only had 1 fork ", i, " is putting right fork down")
			} else {
				fmt.Println(i, " could not find any forks :(")
			}
		}

		if eaten == 3 {
			finished++
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
		} else {
			chForks[i] <- false
		}
	}

}
