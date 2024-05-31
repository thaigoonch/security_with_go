package filesort

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode, sortBy string) {
	if fileList.Len() == 0 {
		// If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}

	for element := fileList.Front(); element != nil; element = element.Next() {
		switch sortBy {
		case "size":
			if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
				fileList.InsertBefore(fileNode, element)
				return
			}
		case "time":
			if fileNode.Info.ModTime().Before(element.Value.(FileNode).Info.ModTime()) {
				fileList.InsertBefore(fileNode, element)
				return
			}
		default:
			// Handle unsupported sort criteria
			return
		}
	}

	fileList.PushBack(fileNode)
}

func GetFilesInDirRecursivelyBySize(fileList *list.List, path string, sortMethod string) {
	dirFiles, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
		return
	}

	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		info, err := dirFile.Info()
		if err != nil {
			fmt.Println("Error getting file info: " + err.Error())
			continue
		}
		if dirFile.IsDir() {
			GetFilesInDirRecursivelyBySize(fileList, fullpath, sortMethod)
		} else if info.Mode().IsRegular() {
			insertSorted(fileList, FileNode{FullPath: fullpath, Info: info}, sortMethod)
		}
	}
}
