package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter card name (type 'quit' to exit): ")

		scanner.Scan()

		// Retrieve the text that the user entered
		userInput := scanner.Text()

		// Trim any leading or trailing whitespace
		userInput = strings.TrimSpace(userInput)

		if userInput == "quit" {
			fmt.Println("Exiting the program...")
			break
		}

		getCard(userInput)
	}
}

func getCard(cardname string) {
	url := fmt.Sprintf("https://api.scryfall.com/cards/named?fuzzy=%s", cardname)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving card: %s", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %s", err)
		return
	}

	fmt.Println("Response Body:", prettyJSON.String())
}
