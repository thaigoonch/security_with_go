package steganography

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CreateZip() {
	// Get the path to ../files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "files")

	// Read the directory
	files, err := ioutil.ReadDir(dir)
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

	fmt.Printf("ZIP archive created successfully: %s\n\n", zipOut)
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
