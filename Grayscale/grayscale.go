// Package declaration
package Grayscale

// Imported required packages
import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

// Function Name: GraySclae()
// Purpose: Function that takes an image an converts each pixel to grayscale 
// Receives: The address of the original image from user
// Returns: Saves the gray scale image to the local directory
func GrayScale() {
	// Prompt user to enter image URL
	fmt.Println("Enter exact name of the image to grayscale: ")

	// Declare variable to hold image URL
	var imageUrl string

	// Taking input from user
	fmt.Scanln(&imageUrl)

	// Open the original image
	reader, err := os.Open(imageUrl)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	fmt.Println("Grayscaling Image...Processing Grayscale...done")

	// Save the grayscale image
	grayFile, err := os.Create("colors_gray_scale.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer grayFile.Close()
	png.Encode(grayFile, grayImg)

	fmt.Println("Saving grayscaled image to 'colors_gray_scale.jpg'" + "... done.")
}
