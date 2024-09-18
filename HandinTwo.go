package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println()
	//ClientToServer := make(chan string)
	//ServerToClient := make(chan string)

	ClientToMiddleman := make(chan string)
	MiddlemanToClient := make(chan string)
	ServerToMiddleman := make(chan string)
	MiddlemanToServer := make(chan string)

	go client(ClientToMiddleman, MiddlemanToClient)
	go server(ServerToMiddleman, MiddlemanToServer)
	go middleman(ClientToMiddleman, MiddlemanToClient, ServerToMiddleman, MiddlemanToServer)
	for {

	}

}

func middleman(ClientToMiddleman chan string, MiddlemanToClient chan string, ServerToMiddleman chan string, MiddlemanToServer chan string) {
	for {
		select {

		// Client sends a message to middleman, and middleman sends to server
		case msgFromClientToServer := <-ClientToMiddleman:
			// 10 percent chance func returns false, and won't send message
			if triesToSendMessagesCouldFail() {
				triesToSendMessagesCouldDelay()
				MiddlemanToServer <- msgFromClientToServer
			}

		// Server sends a message to middleman, and middleman sends to client
		case msgFromServerToMiddleman := <-ServerToMiddleman:
			// 10 percent chance func returns false, and won't send message
			if triesToSendMessagesCouldFail() {
				triesToSendMessagesCouldDelay()
				MiddlemanToClient <- msgFromServerToMiddleman
			}
		}
	}
}

func client(ClientToMiddleman chan string, MiddlemanToClient chan string) {

	//Step 1:
	// Makes a random SEQ nr for the client and send it to middleman.
	ClientSeqNr := rand.Intn(8000) + 1
	fmt.Println("Sending SEQ:", ClientSeqNr, "to server")
	ClientToMiddleman <- convertIntToString(ClientSeqNr)

	// Step 3
	// Gets the servers ACK and SEQ
	ServerACKNr, ServerSeqNr, err := convertStringToTwoInt(<-MiddlemanToClient)

	_ = err

	fmt.Println("Recieved:", ServerACKNr, ServerSeqNr, "from the server")

	// the Client recives the ServerSEQ and +1. This is the new ClientACKNr
	ClientACKNr := ServerSeqNr + 1

	// The clients new SEQ Nr is the Servers ACK nr
	ClientSeqNr = ServerACKNr

	// Sends the new information to the server.
	fmt.Println("Tries to send:", ClientACKNr, ClientSeqNr, "to the server")
	ClientToMiddleman <- convertIntToString(ClientACKNr) + " " + convertIntToString(ClientSeqNr)

	for i := 0; i < 10; i++ {
		RandMessageForServer := rand.Intn(313311122)
		ClientToMiddleman <- convertIntToString(RandMessageForServer)
	}

}

func server(ServerToMiddleman chan string, MiddlemanToServer chan string) {
	// Step 2
	// Gets the client SEQ and makes it the servers ACK
	recivedSeqNr := <-MiddlemanToServer
	fmt.Println("Server Recived:", recivedSeqNr, "from the client")

	AckNr, err := convertStringToInt(recivedSeqNr)
	AckNr = AckNr + 1
	_ = err

	// Makes Client SEQ
	NewSeqNr := rand.Intn(16000) + 8000 + 1

	fmt.Println("Trying to send ACKnr and SEQnr: " + convertIntToString(AckNr) + " " + convertIntToString(NewSeqNr))

	// Sends the servers ACK and SEQ
	ServerToMiddleman <- convertIntToString(AckNr) + " " + convertIntToString(NewSeqNr)

	//Recives ACK and SEQ From client. Establishes connection
	RecSeq, RecACK, err := convertStringToTwoInt(<-MiddlemanToServer)
	fmt.Println("Server Recived:", RecSeq, RecACK, "from the client")
	fmt.Println("-----------------------------------------------------------------------------------")

	fmt.Println("Connection established!")
	fmt.Println(RecSeq, RecACK)

	arry := make([]int, 10)
	fmt.Println("Now sending messages!")
	for i := 0; i < len(arry); i++ {
		Message, err := convertStringToInt(<-MiddlemanToServer)
		_ = err
		arry[i] = Message
		fmt.Println(arry[i])
	}
	fmt.Println("Messages finished")
}

func triesToSendMessagesCouldFail() bool {

	//10% chance it returns false
	return rand.Float32() < 0.9
}

func triesToSendMessagesCouldDelay() {

	//10% chance it returns false
	if rand.Float32() < 0.9 {

	} else { // pause i 1 sec (delay)
		time.Sleep(3 * time.Second)
	}
}

func convertStringToTwoInt(a string) (int, int, error) {
	parts := strings.Split(a, " ")

	num1, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return 0, 0, err1
	}

	num2, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		return 0, 0, err2
	}

	return num1, num2, nil
}

func convertStringToInt(a string) (int, error) {

	num1, err1 := strconv.Atoi(a)
	if err1 != nil {
		return 0, err1
	}
	return num1, nil

}

func convertIntToString(a int) string {

	return strconv.Itoa(a)

}
