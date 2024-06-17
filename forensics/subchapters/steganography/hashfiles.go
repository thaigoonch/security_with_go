package steganography

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/* This code was included in the book under the section
"Detecting a ZIP archive in a JPEG image", I believe as a mistake.

From what I can understand without any context from the book, this program
calculates the MD5, SHA-1, SHA-256, and SHA-512 hashes for the files.
*/

func HashFiles() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Get the full file path
		filename := filepath.Join(dir, file.Name())

		// Get bytes from file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// Hash the file and output results
		fmt.Printf("File: %s\n", filename)
		fmt.Printf("Md5: %x\n", md5.Sum(data))
		fmt.Printf("Sha1: %x\n", sha1.Sum(data))
		fmt.Printf("Sha256: %x\n", sha256.Sum256(data))
		fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
	}
}
