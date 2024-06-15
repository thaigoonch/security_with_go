package main

import (
	"fmt"

	"securitywithgo/forensics/bootsector"
	"securitywithgo/forensics/info"
	"securitywithgo/forensics/sort"
	"securitywithgo/forensics/steganography"
)

func main() {
	fmt.Println(">>> Running getFileInfo():")
	info.GetFileInfo()

	fmt.Println(">>> Running getLargeFiles():")
	sort.GetLargeFiles()

	fmt.Println(">>> Running GetMostRecentlyModifiedFiles():")
	sort.GetMostRecentlyModifiedFiles()

	fmt.Println(">>> Running CheckBootSector():")
	bootsector.CheckBootSector()

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
}
