package main

import (
	"fmt"
	"math/rand"
	"time"
)

var forkArry = []bool

func main() {
	ch1 := make(chan bool)
	forksch := make(chan int)
	forkArry = []bool{false, false, false, false, false}
	philosipher()
	philosipher()
	philosipher()
	philosipher()
	philosipher()

	fork(forkArry)
}

func philosipher() {
	//how many forks that can be grabbed
	eating := false
	forksAva := 0

	//indexes where philosipher has forks
	indexEating := -1
	indexEating2 := -1
	for {
		for i, v := range forkArry {

			if v == false && forksAva == 1 {
				forksch <- i
				indexEating2 = i
				eating = true
				break
			} else if v == false {
				forksch <- i
				indexEating = i
				forksAva++
			}

		}

		if indexEating != 1 && forksAva > 0 { //if only got one, put it back
			forksch <- indexEating
			indexEating = -1
		}

		if eating {
			fmt.Println("eating")
			forksch <- indexEating //put forks back
			forksch <- indexEating2
			indexEating = -1
			indexEating2 = -2
		} else {
			fmt.Println("thinking")
		}

		rand.Seed(time.Now().UnixNano()) //wait between 0 and 1 sec
		time.Sleep(time.Duration(rand.Intn(1000) * time.Millisecond))

	}
}

func fork(array []bool) {
	for {
		index := <-forksch
		changeArry(index, array)
	}
}

func changeArry(index int, array []bool) []bool {
	if array[index] {
		array[index] = false
	} else {
		array[index] = true
	}
	return array
}
