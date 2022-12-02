package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Set the target URL and the text to search for
	url := "https://github.com/BlakePeavy/GoFindText"
	searchText := "I made this in an effort to brush up on Golang."

	// Send an HTTP GET request to the target URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the response body to a string
	html := string(body)

	// Check if the search text is found in the HTML
	if strings.Contains(html, searchText) {
		fmt.Println("The text string was found in the HTML")
	} else {
		fmt.Println("Did not find the text string in the HTML")
	}
}
