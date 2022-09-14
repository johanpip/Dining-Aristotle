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
			fmt.Println(i, " is eatingOWOWOWOWO ")
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
		} else {
			if forkLeft {
				chUsing[i] <- !forkLeft
				fmt.Println(i, "is putting single fork down")
			}

			if forkRight {
				if i == 0 {
					chUsing[4] <- !forkRight
				} else {
					chUsing[i-1] <- !forkRight
				}

				fmt.Println(i, "is putting single fork down")
			}

			time.Sleep(10 * time.Millisecond)
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
