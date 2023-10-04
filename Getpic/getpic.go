package Getpic

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadPicture() {
	//Prompt user to enter image URL
	fmt.Println("Enter the image URL of image: ")

	// var then variable name then variable type
	var imageUrl string

	// Taking input from user
	fmt.Scanln(&imageUrl)

	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Create a new file to save the image
	outputFile, err := os.Create("colors.jpg")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Downloading " + imageUrl + "...done")

	fmt.Println("Saving 'colors.jpg' to local file...done")
}
