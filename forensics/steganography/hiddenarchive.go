package steganography

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
"Creating a steganographic image archive

Now that we have an image and a ZIP archive, we can combine them together
to "hide" the archive within the image.
This is probably the most primitive form of steganography.
A more advanced way would be to split up the file byte by byte,
store the information in the low bits of the image, use a special program
to extract the data from the image, and then reconstruct the original data.
This example is nice because we can easily test and verify if it still
loads as an image and still behaves like a ZIP archive.

The following example will take a JPEG image and a ZIP archive
and combine them to create a hidden archive.
The file will retain the .jpg extension and will still function
and look like a normal image.
However, the file also still works as a ZIP archive.
You can unzip the .jpg file and the archived files will be extracted:"
*/

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
