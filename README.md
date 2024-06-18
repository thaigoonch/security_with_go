# Security with Go

This project contains various subchapters and examples from the "Forensics" chapter of the book “Security with Go” by John Daniel Leon. 
Below is the order to explore the code.

## Exploring the Code

### 1. Files Subchapter

Begin with the "Files" subchapter, which covers file-related operations:

- [fileinfo.go](forensics/subchapters/files/fileinfo.go): This script demonstrates how to retrieve and display various metadata about a file, such as its name, size, and permissions
- [filesort_large.go](forensics/subchapters/files/filesort_large.go): This script simply sorts files by size, useful in forensics when trying to spot suspicious large files
- [filesort_modified.go](forensics/subchapters/files/filesort_modified.go): This script sorts files in a directory by their last modification date
- [bootsector.go](forensics/subchapters/files/bootsector.go): This script displays info about the boot sector of a disk, which is crucial for understanding the structure and integrity of storage media

### 2. Steganography Subchapter

Next, move on to the "Steganography" subchapter to learn about hiding information within other files:

- [hashfiles.go](forensics/subchapters/steganography/hashfiles.go): This script generates hash values for files, which can be used for verifying file integrity
- [image.go](forensics/subchapters/steganography/image.go): This script simply generates a colored-pixel-arranged JPG image for use in further programs
- [hiddenarchive.go](forensics/subchapters/steganography/hiddenarchive.go): This script creates a ZIP archive out of files in the `forensics/test_files` directory
- [ziparchive_detect.go](forensics/subchapters/steganography/ziparchive_detect.go): This script scans files for hidden ZIP archives

### 3. Network Subchapter

Finally, explore the "Network" subchapter to understand network-related functionalities:

- [hostname_from_ip.go](forensics/subchapters/network/hostname_from_ip.go): This script converts IP addresses to their corresponding hostnames
- [ip_from_hostname.go](forensics/subchapters/network/ip_from_hostname.go): This script resolves domain names to their IP addresses
- [mx_from_domain.go](forensics/subchapters/network/mx_from_domain.go): This script fetches MX records for a given domain, which are used in email routing
- [nameservers.go](forensics/subchapters/network/nameservers.go): This script retrieves the nameserver records for a domain

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
