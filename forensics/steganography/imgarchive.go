package steganography

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ArchiveImage() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "files")

	// Open original file
	firstFile, err := os.Open(fmt.Sprintf("%s/test.jpg", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	// Second file
	secondFile, err := os.Open("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	// New file for output
	newFile, err := os.Create("stego_image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	_, err = io.Copy(newFile, firstFile)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(newFile, secondFile)
	if err != nil {
		log.Fatal(err)
	}
}
