package main

import (
	"fmt"

	"securitywithgo/forensics/bootsector"
	"securitywithgo/forensics/image"
	"securitywithgo/forensics/info"
	"securitywithgo/forensics/sort"
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

	fmt.Println(">>> Running GenImage():")
	image.GenImage()
}
