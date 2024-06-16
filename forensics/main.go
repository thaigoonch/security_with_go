package main

import (
	"fmt"

	"securitywithgo/forensics/files"
	"securitywithgo/forensics/network"
	"securitywithgo/forensics/steganography"
)

func main() {
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
}
