package steganography

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
)

/*
"Generating an image with random noise

This program will create a JPEG image with every pixel set to a random color.
It is a simple program so we have just a jpeg image available to work with.
The Go standard library comes with jpeg, gif, and png packages.
The interface to all different image types is the same, so swapping from
a jpeg to a gif or png package is very easy:"
*/

func CreateImage() {
	// 100x200 pixels
	myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	for p := 0; p < 100*200; p++ {
		pixelOffset := 4 * p
		myImage.Pix[0+pixelOffset] = uint8(rand.Intn(256)) // Red
		myImage.Pix[1+pixelOffset] = uint8(rand.Intn(256)) // Green
		myImage.Pix[2+pixelOffset] = uint8(rand.Intn(256)) // Blue
		myImage.Pix[3+pixelOffset] = 255                   // Alpha
	}

	imgName := "test_files/test.jpg"
	outputFile, err := os.Create(imgName)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(outputFile, myImage, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image generated as %s\n", imgName)
}
