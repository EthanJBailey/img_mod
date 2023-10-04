package Text

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func PrintText() {
	const W = 500
	const H = 300

	//Prompt user to enter image URL
	fmt.Println("Enter the text that you want to color: ")

	// var then variable name then variable type
	var txt string

	// Taking input from user
	fmt.Scanln(&txt)

	// Create a temporary file and write the byte slice to it
	tempFile, err := ioutil.TempFile("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored(txt, W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	fmt.Println("Saving labeled image to colors_labeled.jpg ... done")
	dc.SavePNG("colors_labeled.jpg")
}
