package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Set the target URL and the text to search for
	url := "https://www.example.com"
	searchText := "Hello, world!"

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
		fmt.Println("Found the text in the HTML")
	} else {
		fmt.Println("Did not find the text in the HTML")
	}
}
