package main

import (
	"fmt"
	"sync"
)

ch1 := make(chan bool)
forkArry := []bool{false,false,false,false,false}
forksch := make(chan int)

func main() {
	philosipher()
	philosipher()
	philosipher()
	philosipher()
	philosipher()

	fork()
}

func philosipher(){
	//how many forks that can be grabbed
	eating := false
	forksAva := 0
	
	//indexes where philosipher has forks
	indexEating := -1
	indexEating2 := -1


	for i, v := range forkArry {
		forksAva = <-forksch
		if(v == false && forksAva == 1){
			forksch <- i
			indexEating2 = i
			break;
		} else if {
			forksch <- i
			indexEating = i
			forksAva++
		}
		
		if forksAva > 1 {
			eating = true
		} 	
		
	}

	if(indexEating != 1 && forksAva > 0){ //if only got one, put it back
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

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1000) * time.Millisecond))

	
}

func fork(){
	for {
		index := <- forksch
		if forkArry[index] {
			forkArry[index] = true;
		} else {
			forkArry[index] = false;
		}
	}
}

