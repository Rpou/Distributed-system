package main

// This program works by having 5 go-rutines for both philosiphers and forks.
// Each philosipher tries to take one fork first, and then requests the next fork.
// If the philosipher fails to take the first fork (it takes too long time to request it)
// the philosipher will try again.
// If the philosipher successfully takes the first fork, he tries to request the next fork.
// If this takes too long, then he will drop both forks, and try again.

// The program will not reach a deadlock, because if philosipher 1 tries to access the fork array, and philosipher 2
// is already holding onto that fork, philosipher 1 will not do anything with the fork. the fork array will ignore the request.
// Another reason it wont deadlock, is because there is a set amount of time, to take each fork.
// If the fork has not been taken within that timeframe, the philosipher will stop the request.

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	forkArry   = []int{0, 0, 0, 0, 0} // array to keep track of the forks and who holds onto a fork.
	TimesEaten = []int{0, 0, 0, 0, 0} // array to see, how many times each philosipher has eaten.
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go fork(0, ch1)
	go fork(1, ch2)
	go fork(2, ch3)
	go fork(3, ch4)
	go fork(4, ch5)

	go philosipher(1, ch1, ch2)
	go philosipher(2, ch2, ch3)
	go philosipher(3, ch3, ch4)
	go philosipher(4, ch4, ch5)
	go philosipher(5, ch5, ch1)

	for {
		//change the number in "Status" to make them able to eat more than 3 times.
		if Status(3) == 1 {
			break
		}
	}

}

func philosipher(number int, ch1 chan int, ch2 chan int) {
	for {
		eating := false

		startTime := time.Now()                 // Record the start time
		timeoutDuration := 1 * time.Millisecond // 1 millisecond timeout

		// Checks if philosipher can take fork 1 within the 1 millisec timeframe
		for time.Since(startTime) < timeoutDuration {
			ch1 <- number
			if forkArry[number-1] == number { // Acquired the first fork
				break
			}

		}

		if forkArry[number-1] == number { //tjek om array for opdateret fra fork??? Check if the philosipher has fork 1.
			startTime = time.Now()
			// Checks if philosipher can take fork 2 within the 1 millisec timeframe
			for time.Since(startTime) < timeoutDuration {
				ch2 <- number
				if number != 5 {
					if forkArry[number] == number {
						eating = true
						TimesEaten[number-1]++
						break
					}

				} else {
					if forkArry[0] == number {
						eating = true
						TimesEaten[number-1]++
						break
					}
				}

			}
		} else { // If it has not gotten the first fork, then it will retry taking the fork again. (running the first for-loop again)
			fmt.Println(number, "thinkin")
			continue
		}

		if eating {
			fmt.Println(number, "spiser, nam nam...")
			ch1 <- number // returns both forks
			ch2 <- number
		} else {
			fmt.Println(number, "thinkin")
			ch1 <- number
		}

	}
}

func fork(number int, ch chan int) {
	for {
		index := <-ch
		if forkArry[number] == 0 {
			forkArry[number] = index
		} else if forkArry[number] == index {
			forkArry[number] = 0
		}
	}
}

func Status(limit int) int {
	for {
		time.Sleep(time.Duration(rand.Float64() * 100000))
		for i := 0; i < 100; i++ {
			//fmt.Println(TimesEaten)
			if TimesEaten[0] > limit && TimesEaten[1] > limit && TimesEaten[2] > limit && TimesEaten[3] > limit && TimesEaten[4] > limit {
				fmt.Println(TimesEaten)
				return 1
			}
		}
	}
}
