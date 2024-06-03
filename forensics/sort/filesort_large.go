package sort

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetLargeFiles() {
	fileList := list.New()

	// Get the path to ../files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "files")

	GetFilesInDirRecursivelyBySize(fileList, dir, "size")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d ", element.Value.(FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(FileNode).FullPath)
	}
}
