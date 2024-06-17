# Security with Go

This project contains various subchapters and examples from the "Forensics" chapter of the book “Security with Go” by John Daniel Leon. 
Below is the order to explore the code.

## Exploring the Code

### 1. Files Subchapter

Begin with the "Files" subchapter, which covers file-related operations:

- [fileinfo.go](forensics/subchapters/files/fileinfo.go)
- [filesort_large.go](forensics/subchapters/files/filesort_large.go)
- [filesort_modified.go](forensics/subchapters/files/filesort_modified.go)
- [bootsector.go](forensics/subchapters/files/bootsector.go)

### 2. Steganography Subchapter

Next, move on to the "Steganography" subchapter to learn about hiding information within other files:

- [hashfiles.go](forensics/subchapters/steganography/hashfiles.go)
- [image.go](forensics/subchapters/steganography/image.go)
- [ziparchive.go](forensics/subchapters/steganography/ziparchive.go)
- [hiddenarchive.go](forensics/subchapters/steganography/hiddenarchive.go)
- [ziparchive_detect.go](forensics/subchapters/steganography/ziparchive_detect.go)

### 3. Network Subchapter

Finally, explore the "Network" subchapter to understand network-related functionalities:

- [hostname_from_ip.go](forensics/subchapters/network/hostname_from_ip.go)
- [ip_from_hostname.go](forensics/subchapters/network/ip_from_hostname.go)
- [mx_from_domain.go](forensics/subchapters/network/mx_from_domain.go)
- [nameservers.go](forensics/subchapters/network/nameservers.go)

## To run
```
cd forensics && \
sudo go run main.go
```

## Definitions
- **Digital forensics**: gathering info on electronic data
   - E.g. gathering info on an attacking IP, capturing network communication
- **Steganography**: the hiding of archives inside images
- **Hashing**: using a one-way mathematical function to convert data into a unique string of text that can't be decoded or reversed
   - E.g. MD5, SHA1
