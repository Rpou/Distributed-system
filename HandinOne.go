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
	
	//indexes where philosipher has forks
	indexFork := -1
	indexFork2 := -1
	for {
		for i, v := range forkArry {
			forksAva = <-forksch
			if(v == false && indexFork != -1){
				forksch <- i
				indexFork2 = i
				forksAva++
				eating = true
				break;
			} else if false {
				forksch <- i
				indexFork = i
			}
			
	
			
		}
	
		if(indexFork != -1 && indexFork2 =< 0){ //if only got one, put it back
			forksch <- indexFork
			indexFork = -1
		}

		if eating {
			fmt.Println("eating")
			forksch <- indexFork //put forks back
			forksch <- indexFork2
			indexFork = -1
			indexFork2 = -2
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
