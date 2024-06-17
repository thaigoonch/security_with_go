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
"Finding recently modified files

When examining a victim machine forensically, one of the first things you
can do is to look for files that have been recently altered. It could give
you clues as to where an attacker was looking, what settings they modified,
or what their motive was.

However, if an investigator is looking through an attacker's machine, then
the goal is slightly different. Recently accessed files may give clues as
to what tools they were using to attack where they might be hiding data,
or what software they use.

The following example will search a directory and subdirectories to find
all the files and sort them by the last modified time."
*/

func GetMostRecentlyModifiedFiles() {
	fileList := list.New()

	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")
	helpers.GetFilesInDirRecursivelyBySize(fileList, dir, "time")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(helpers.FileNode).Info.ModTime())
		fmt.Printf("%s\n", element.Value.(helpers.FileNode).FullPath)
	}
}
