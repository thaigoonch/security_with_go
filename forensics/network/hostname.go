package network

import (
	"fmt"
	"log"
	"net"
)

func GetHostNameFromIP() {
	// Hard-coded IP address example
	ipAddress := "8.8.8.8"

	// Parse the IP for validation
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		log.Fatal("Valid IP not detected. Value provided: " + ipAddress)
	}

	fmt.Println("Looking up hostnames for IP address: " + ipAddress)
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostname := range hostnames {
		fmt.Println(hostname)
	}
}
