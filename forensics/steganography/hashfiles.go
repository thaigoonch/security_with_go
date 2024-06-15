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

func HashFiles() {
	// Get the path to ../files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "files")

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
