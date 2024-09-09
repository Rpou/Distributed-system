package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	forkArry   = []int{0, 0, 0, 0, 0}
	TimesEaten = []int{0, 0, 0, 0, 0}
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

	//go Status() //viser

	//select {} // prevent main from exiting

	for {
		if Status(1000000000) == 1 {
			break
		}
	}

}

func philosipher(number int, ch1 chan int, ch2 chan int) {
	for {

		//time.Sleep(time.Duration(rand.Int31()))
		//time.Sleep(time.Duration(rand.Float32()))

		eating := false

		startTime := time.Now()                   // Record the start time
		timeoutDuration := 200 * time.Millisecond // 1 second timeout

		// tjekker om den kan få gaffel 1 inden for 1 sec
		for time.Since(startTime) < timeoutDuration {
			ch1 <- number
			if forkArry[number-1] == number { // Acquired the first fork
				break
			}

		}

		if forkArry[number-1] == number { //tjek om array for opdateret fra fork???
			startTime = time.Now()
			// tjekker om den kan få gaffel 2 inden for 1 sec
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
		} else { // hvis den ikke har fået fat i gaffel 1, så continue (kør for-loop fra starten igen.)
			fmt.Println(number, "thinkin")
			//time.Sleep(time.Duration(rand.Float64() * 10000)) //mængden af tid de venter med at
			continue
		}

		if eating {
			fmt.Println(number, "spiser, nam nam...")
			ch1 <- number //returner begge gafler
			ch2 <- number
		} else {
			fmt.Println(number, "thinkin")
			ch1 <- number

		}

		//time.Sleep(time.Duration(rand.Float64() * 10000)) //mængden af tid de venter med at
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
			fmt.Println(TimesEaten)
			fmt.Println(forkArry)
			if TimesEaten[0] > limit && TimesEaten[1] > limit && TimesEaten[2] > limit && TimesEaten[3] > limit && TimesEaten[4] > limit {
				return 1
			}
		}
	}
}
