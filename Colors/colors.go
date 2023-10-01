package colors

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func ColorsPrint() {

	//Prompt user to enter image name
	fmt.Println("Enter the exact image name with extension: ")

	// var then variable name then variable type
	var imge string

	// Taking input from user
	fmt.Scanln(&imge)

	// Open the image
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

	// Prints out results of the decoder
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
