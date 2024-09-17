package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {

	fmt.Println()
	ClientToServer := make(chan string)
	ServerToClient := make(chan string)

	go client(ClientToServer, ServerToClient)
	go server(ClientToServer, ServerToClient)

	for {

	}

}

func client(ClientToServer chan string, ServerToClient chan string) {

	//Step 1:
	// Makes a random SEQ nr for the client and send it to server.
	ClientSeqNr := rand.Intn(8000) + 1
	fmt.Println("Sending", ClientSeqNr, "to server")
	ClientToServer <- convertIntToString(ClientSeqNr)

	// Step 3
	// Gets the servers ACK and SEQ
	ServerACKNr, ServerSeqNr, err := convertStringToTwoInt(<-ServerToClient)

	_ = err

	fmt.Println(ServerACKNr, ServerSeqNr)

	// the Client recives the ServerSEQ and +1. This is the new ClientACKNr
	ClientACKNr := ServerSeqNr + 1

	// The clients new SEQ Nr is the Servers ACK nr
	ClientSeqNr = ServerACKNr

	// Sends the new information to the server.
	ClientToServer <- convertIntToString(ClientACKNr) + " " + convertIntToString(ClientSeqNr)

}

func server(ClientToServer chan string, ServerToClient chan string) {
	// Step 2
	// Gets the client SEQ
	recivedSeqNr := <-ClientToServer
	AckNr, err := convertStringToInt(recivedSeqNr)
	AckNr = AckNr + 1
	_ = err

	// Makes Client SEQ
	NewSeqNr := rand.Intn(16000) + 8000 + 1

	fmt.Println(convertIntToString(AckNr) + " " + convertIntToString(NewSeqNr))

	// Sends the servers ACK and SEQ
	ServerToClient <- convertIntToString(AckNr) + " " + convertIntToString(NewSeqNr)

	fmt.Println("Sending ACKnr and SEQnr")

	//Recives ACK and SEQ From client. Establishes connection
	RecSeq, RecACK, err := convertStringToTwoInt(<-ClientToServer)

	fmt.Println("Connection established!")
	fmt.Println(RecSeq, RecACK)

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
