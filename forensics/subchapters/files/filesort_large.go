package files

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"securitywithgo/forensics/helpers"
)

/*
"Finding the largest files

Large files are always prime suspects when investigating. Large database dumps,
password dumps, rainbow tables, credit card caches, stolen intellectual property,
and other data are often stored in one large archive that is easy to spot if
you have the right tools. Also, it would be helpful to find exceptionally large
image or video files that may have steganographically-hidden information inside...

This program will search in a directory and all subdirectories for all files and
sort them by file size."
*/

func GetLargeFiles() {
	fileList := list.New()

	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")
	helpers.GetFilesInDirRecursivelyBySize(fileList, dir, "size")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d ", element.Value.(helpers.FileNode).Info.Size())
		fmt.Printf("%s\n", element.Value.(helpers.FileNode).FullPath)
	}
}
