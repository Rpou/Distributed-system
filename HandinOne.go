package main

import {
	"fmt"
	"sync"
}

ch1 := make(chan bool)
forkArry := []bool{false,false,false,false,false}

func main() {
	philosipher()
	philosipher()
	philosipher()
	philosipher()
	philosipher()

	fork()
	fork()
	fork()
	fork()
	fork()
}

func philosipher(){
	//how many forks that can be grabbed
	forksAva := 0;
	//channel to communicate between forksava
	forksch := make(chan int)


	for i, v := range forkArry {
		forksAva = <-forksch
		if(v == false){
			forksch <- forksAva + 1
			forkArry[i] = true
		}
		ch <- false
	}

	if bla bla {
		fmt.Println("eating")
	} else {
		fmt.Println("thinking")
	}
}

func fork(){

	beingUsed := false
	
}

