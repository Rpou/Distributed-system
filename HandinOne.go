package main

import (
	"fmt"
	"sync"
)


var forkArry []bool;

func main() {
	ch1 := make(chan bool)
	forkArry = []bool{false,false,false,false,false}

	go philosipher()
	go philosipher()
	go philosipher()
	go philosipher()
	go philosipher()

	go fork()
	go fork()
	go fork()
	go fork()
	go fork()

	for i, v := range 100000{
		
	}

}

func philosipher(){
	eating := false
	//how many forks that can be grabbed
	forksAva := 0;
	//channel to communicate between forksava
	forksch := make(chan int)

	// go through array to see which forks are avalible 
	for i, v := range forkArry {
		//sends updated amount of avalible forks to the var
		forksAva = <-forksch

		//check if fork avalible
		if(v == false){
			// if fork ava, then update the forksava var.
			forksch <- forksAva + 1
			forkArry[i] = true
			
		}
		
		if(forksAva >= 2){
			ch1 <- true;
			forksch <- 0;
		}
	}

	if eating {
		fmt.Println("eating")
	} else {
		fmt.Println("thinking")
	}
}

func fork(){
	var beingUsed bool
	beingUsed = <-ch1
}

func changeArray(index int, boolean bool){
	forkArry[index] = boolean; 
}
