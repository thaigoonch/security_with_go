package steganography

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

/*
"Creating a ZIP archive

This program will create a ZIP archive, so we have an archive to use
with our steganography experiments.
The Go standard library has a zip package, but it also supports
TAR archives with the tar package."

Note: I had to use chatGPT to generate this code, because the book did not
include the correct program for this part.
*/

func CreateZip() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new ZIP file
	zipOut := fmt.Sprintf("%s/test.zip", dir)
	zipFile, err := os.Create(zipOut)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add files to the ZIP archive
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(dir, file.Name())
		if err := addFileToZip(zipWriter, filePath, file.Name()); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("ZIP archive created successfully: %s\n", zipOut)
}

// addFileToZip adds a file to the ZIP archive
func addFileToZip(zipWriter *zip.Writer, filePath, filename string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new file in the ZIP archive
	zipFile, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	// Copy the file content to the ZIP archive
	_, err = io.Copy(zipFile, file)
	if err != nil {
		return err
	}

	return nil
}
