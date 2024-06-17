package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"securitywithgo/forensics/subchapters/files"
	"securitywithgo/forensics/subchapters/network"
	"securitywithgo/forensics/subchapters/steganography"
)

func main() {
	clearFiles()

	fmt.Println("\n\n>>> Running files.GetFileInfo():")
	files.GetFileInfo()

	fmt.Println("\n\n>>> Running files.GetLargeFiles():")
	files.GetLargeFiles()

	fmt.Println("\n\n>>> Running files.GetMostRecentlyModifiedFiles():")
	files.GetMostRecentlyModifiedFiles()

	fmt.Println("\n\n>>> Running files.CheckBootSector():")
	files.CheckBootSector()

	fmt.Println("\n\n>>> Running steganography.CreateImage():")
	steganography.CreateImage()

	fmt.Println("\n\n>>> Running steganography.HashFiles():")
	steganography.HashFiles()

	fmt.Println("\n\n>>> Running steganography.CreateZip():")
	steganography.CreateZip()

	fmt.Println("\n\n>>> Running CreateHiddenArchive():")
	steganography.CreateHiddenArchive()

	fmt.Println("\n\n>>> Running steganography.DetectZip():")
	steganography.DetectZip()

	fmt.Println("\n\n>>> Running network.GetHostNameFromIP():")
	network.GetHostNameFromIP()

	fmt.Println("\n\n>>> Running network.GetIPFromHostname():")
	network.GetIPFromHostname()

	fmt.Println("\n\n>>> Running network.GetMXFromDomain():")
	network.GetMXFromDomain()

	fmt.Println("\n\n>>> Running network.GetNameServers():")
	network.GetNameServers()
}

func clearFiles() {
	dir := "test_files"

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the files and remove ZIP/JPG files
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
