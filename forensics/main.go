package main

import (
	"fmt"

	"github.com/thaigoonch/securitywithgo/fileinfo"
	"github.com/thaigoonch/securitywithgo/filesort"
)

func main() {
	fmt.Println(">>> Running getFileInfo():")
	fileinfo.GetFileInfo()

	fmt.Println(">>> Running getLargeFiles():")
	filesort.GetLargeFiles()

	fmt.Println(">>> Running GetMostRecentlyModifiedFiles():")
	filesort.GetMostRecentlyModifiedFiles()
}
