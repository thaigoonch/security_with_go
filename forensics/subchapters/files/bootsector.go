package files

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
"Reading the boot sector

This program will read the first 512 bytes of a disk and print the results
as decimal values, hex, and a string. The io.ReadFull() function is like
a normal read, but it ensures that the byte slice you provide with for
data is completely filled. It returns an error if there are not enough
bytes in the file to fill the byte slice.

A practical use for this is to check a machine's boot sector to see if it
has been modified. Rootkits and malware may hijack the boot process by
modifying the boot sector. You can manually inspect it for anything strange
or compare it to a known good version. Perhaps a backup image of the machine
or a fresh install can be compared to see if anything has changed.

Note that you can technically pass it any filename and not specifically a
disk, since everything in Linux is treated as a file. If you pass it the name
of the device directly, such as /dev/sda, it will read the first 512 bytes of
the disk, which is the boot sector. The primary disk device is typically /dev/sda,
but may also be /dev/sdb or /dev/sdc. Use mount or the df tools to get more
information about what your disks are named. You will need to run the application
with sudo in order to have the permission to read the disk device directly.

For more information on files, input, and output, look into the os, bufio, and
io packages, as demonstrated in the following code block:"
*/

func CheckBootSector() {
	// Ensure the program is run as root
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root. Please use sudo.")
		os.Exit(1)
	}

	// Note: I've added a few alternate disk names, but you still may need to add yours
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
