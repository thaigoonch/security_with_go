package steganography

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateHiddenArchive() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	// Open original file
	firstFile, err := os.Open(fmt.Sprintf("%s/test.jpg", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	// Second file
	secondFile, err := os.Open(fmt.Sprintf("%s/test.zip", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	// New file for output
	fileOut := fmt.Sprintf("%s/stego_image.jpg", dir)
	newFile, err := os.Create(fileOut)
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
	fmt.Printf("hidden archive created successfully: %s\n\n", fileOut)
}
