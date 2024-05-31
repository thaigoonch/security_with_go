package filesort

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetMostRecentlyModifiedFiles() {
	fileList := list.New()

	// Get the path to ../files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "files")
	GetFilesInDirRecursivelyBySize(fileList, dir, "time")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(FileNode).Info.ModTime())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
