package network

import (
	"fmt"
	"log"
	"net"
)

func GetIPFromHostname() {
	hostname := "www.example.com"

	fmt.Println("Looking up IP addresses for hostname: " + hostname)

	ips, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
