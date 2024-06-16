package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"securitywithgo/forensics/files"
	"securitywithgo/forensics/network"
	"securitywithgo/forensics/steganography"
)

func main() {
	removeZipFiles()

	fmt.Println(">>> Running getFileInfo():")
	files.GetFileInfo()

	fmt.Println(">>> Running getLargeFiles():")
	files.GetLargeFiles()

	fmt.Println(">>> Running GetMostRecentlyModifiedFiles():")
	files.GetMostRecentlyModifiedFiles()

	fmt.Println(">>> Running CheckBootSector():")
	files.CheckBootSector()

	fmt.Println(">>> Running CreateImage():")
	steganography.CreateImage()

	fmt.Println(">>> Running HashFiles():")
	steganography.HashFiles()

	fmt.Println(">>> Running CreateZip():")
	steganography.CreateZip()

	fmt.Println(">>> Running CreateHiddenArchive():")
	steganography.CreateHiddenArchive()

	fmt.Println(">>> Running DetectZip():")
	steganography.DetectZip()

	fmt.Println(">>> Running GetHostNameFromIP():")
	network.GetHostNameFromIP()

	fmt.Println(">>> Running GetIPFromHostname():")
	network.GetIPFromHostname()

	fmt.Println(">>> Running GetMXFromDomain():")
	network.GetMXFromDomain()
}

func removeZipFiles() {
	// Define the directory
	dir := "test_files"

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the files and remove ZIP files
	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".zip" || filepath.Ext(file.Name()) == ".jpg") {
			filePath := filepath.Join(dir, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				log.Printf("Failed to remove %s: %v\n", filePath, err)
			} else {
				fmt.Printf("Removed %s\n", filePath)
			}
		}
	}
}
