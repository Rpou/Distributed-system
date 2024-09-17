package main

import (
	"fmt"
)

func main() {

	fmt.Scan()

	go client()
	go server()

	for {

	}

}
