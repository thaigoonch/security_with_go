package network

import (
	"fmt"
	"log"
	"net"
)

/*
"Looking up a hostname from an IP address

This program will take an IP address and figure out what the hostnames are.
The net.parseIP() function is used to validate the IP address provided,
and net.LookupAddr() does the real work of figuring out what the hostname is."
*/

func GetHostNameFromIP() {
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
