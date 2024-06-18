package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
"Getting file information"

"This program will print the information about a file, namely when it was
last modified, who owns it, how many bytes it is, and what its permissions are."
*/

func GetFileInfo() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err := os.Stat(fmt.Sprintf("%s/test.txt", dir))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
