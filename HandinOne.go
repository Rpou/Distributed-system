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
	TimesEaten = []int{0, 0, 0, 0, 0} // array to see, how many times each philosipher has eaten.
)

func main() {
	fork1 := make(chan struct{}, 1)
	fork2 := make(chan struct{}, 1)
	fork3 := make(chan struct{}, 1)
	fork4 := make(chan struct{}, 1)
	fork5 := make(chan struct{}, 1)

	// Initially, each fork is available by sending a struct{} to the channel
	fork1 <- struct{}{}
	fork2 <- struct{}{}
	fork3 <- struct{}{}
	fork4 <- struct{}{}
	fork5 <- struct{}{}

	// Start goroutines for philosophers
	go philosipher(1, fork1, fork2)
	go philosipher(2, fork2, fork3)
	go philosipher(3, fork3, fork4)
	go philosipher(4, fork4, fork5)
	go philosipher(5, fork5, fork1)
	for {
		//change the number in "Status" to make them able to eat more than 3 times.
		if Status(100000000) == 1 {
			break
		}
	}

}

func philosipher(number int, leftFork, rightFork chan struct{}) {
	for {
		eating := false
		//startTime := time.Now()
		timeDuration := time.Duration(rand.Float32() * 10000)

		// Try to pick up the left fork within the timeout duration
		select {
		case <-leftFork:
			// If successful, try to pick up the right fork within the timeout
			select {
			case <-rightFork:
				// Got both forks, start eating
				eating = true
				TimesEaten[number-1]++
				fmt.Println("Philosopher", number, "is eating...")

				time.Sleep(timeDuration) // Simulate eating time

				// Release both forks
				leftFork <- struct{}{}
				rightFork <- struct{}{}
			default:
				// Failed to pick up right fork, release left fork
				leftFork <- struct{}{}
			}
		default:
			// Timeout occurred, philosopher keeps thinking
		}

		if !eating {
			fmt.Println("Philosopher", number, "is thinking...")
			time.Sleep(timeDuration) // Simulate thinking time before retrying
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
