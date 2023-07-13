package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:9000/test-multipartform"
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	frontWriter, err := multipartWriter.CreateFormField("front")
	if err != nil {
		panic(err)
	}
	file, err := os.Open("front.png")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(frontWriter, file)
	if err != nil {
		panic(err)
	}
	// Close the multipart writer to finalize the form
	err = multipartWriter.Close()
	if err != nil {
		panic(err)
	}

	// Create a POST request with the multipart form data
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code %d", response.StatusCode)
	}

	// Process the response as needed
	fmt.Println("Upload successful")
}
