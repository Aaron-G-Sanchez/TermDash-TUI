package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Function to test the cloudflare tunnel set up
func testTunnel() string {
	// Make a get request to the designated endpoint
	resp, err := http.Get("https://termdash.aarongsanchez.me/hello-world")
	if err != nil {
		log.Fatalf("Error fetching: %v", err)
	}

	defer resp.Body.Close()
	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing body: %v", err)
	}
	return string(body)
}

func main() {

	call := testTunnel()

	// Print body and print human readable body.
	fmt.Println(call) // Should print {msg: Hello World}
}
