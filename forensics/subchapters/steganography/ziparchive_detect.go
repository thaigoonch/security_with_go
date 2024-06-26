package steganography

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
"Detecting a ZIP archive in a JPEG image

If data is hidden using the technique from [CreateZip()],
it can be detected by searching for the ZIP file signature in the image.
A file may have a .jpg extension and still load properly in a photo viewer,
but it may still have a ZIP archive stored in the file.
The following program searches through a file and looks for a ZIP file signature.
We can run it against the file created in [CreateZip()]:"
*/

func DetectZip() {
	// Get the path to ../test_files
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentDir, "test_files")

	// Zip signature is "\x50\x4b\x03\x04"
	filename := fmt.Sprintf("%s/stego_image.jpg", dir)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	fileStat, _ := file.Stat()
	// 0 is being cast to an int64 to force i to be initialized as
	// int64 because filestat.Size() returns an int64 and must be
	// compared against the same type
	for i := int64(0); i < fileStat.Size(); i++ {
		myByte, err := bufferedReader.ReadByte()
		if err != nil {
			log.Fatal(err)
		}

		if myByte == '\x50' {
			// First byte match. Check the next 3 bytes
			byteSlice := make([]byte, 3)
			// Get bytes without advancing pointer with Peek
			byteSlice, err = bufferedReader.Peek(3)
			if err != nil {
				log.Fatal(err)
			}
			if bytes.Equal(byteSlice, []byte{'\x4b', '\x03', '\x04'}) {
				log.Printf("Found zip signature at byte %d.", i)
			}
		}
	}
}
