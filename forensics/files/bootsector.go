package files

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CheckBootSector() {
	// Ensure the program is run as root
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root. Please use sudo.")
		os.Exit(1)
	}

	paths := []string{"/dev/sda", "/dev/sdb", "/dev/sdc", "/dev/nvme0n1", "/dev/nvme0n1p1"}
	var file *os.File
	var err error

	for _, path := range paths {
		file, err = os.Open(path)
		if err == nil {
			// Successfully opened the file, break out of the loop
			break
		}
	}

	if err != nil {
		log.Fatal("Error: Unable to open any of the devices.")
	}
	defer file.Close()

	// The file.Read() function will read a tiny file into a large
	// byte slice, but io.ReadFull() will return an
	// error if the file is smaller than the byte slice.
	byteSlice := make([]byte, 512)
	// ReadFull will error if 512 bytes not available to read
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal("Error reading 512 bytes from file. " + err.Error())
	}

	log.Printf("Bytes read: %d\n\n", numBytesRead)
	log.Printf("Data as decimal:\n%d\n\n", byteSlice)
	log.Printf("Data as hex:\n%x\n\n", byteSlice)
	log.Printf("Data as string:\n%s\n\n", byteSlice)
}
