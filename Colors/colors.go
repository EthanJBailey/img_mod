// Package Declaration
package Colors

// Imported required packages
import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // import this package to decode PNGs
	"os"
)

// Function Name: PrintPixels()
// Purpose: Function that prints the RGB color and address of each pixel of an image
// Receives: The path of the original image to be scanned
// Returns: Saves the output to a new file the local directory
func PrintPixels() {

	//Prompt user to enter image name
	fmt.Println("Enter the exact image name with extension: ")

	// var then variable name then variable type
	var imge string

	// Taking input from user
	fmt.Scanln(&imge)

	// Open the image  referenced
	reader, err := os.Open(imge)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Open output file for writing
	file, errs := os.Create("colors_pixel_counts.txt")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	fmt.Println("Processing Pixel Colors")

	// Prints out results of the decoder
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()

			values := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
			file.WriteString(values)

			values = ""
		}
	}

	// Prints out completetion statement to console
	fmt.Println("Saving pixel output to 'colors_pixel_counts.txt'" + "...done")
}
