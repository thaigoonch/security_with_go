package network

import (
	"fmt"
	"log"
	"net"
)

/*
"Looking up IP addresses from a hostname

The following example takes a hostname and returns the IP address.
It is very similar to [GetHostNameFromIP()], but it is in reverse.
The net.LookupHost() function does the heavy lifting:"
*/

func GetIPFromHostname() {
	hostname := "devdungeon.com"

	fmt.Println("Looking up IP addresses for hostname: " + hostname)

	ips, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
