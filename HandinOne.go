package main

import {
	"fmt"
	"sync"
}

ch1 := make(chan bool)

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


	if bla bla {
		fmt.Println("eating")
	} else {
		fmt.Println("thinking")
	}
}

func fork(){
	beingUsed := false
}

