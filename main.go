package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// TODO: Move this to a utils file
func getAddress() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address:")

	address, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	address = address[:len(address)-1]

	return address
}

func main() {

	address := getAddress()
	fmt.Println("Greetings from: ", address)

}
