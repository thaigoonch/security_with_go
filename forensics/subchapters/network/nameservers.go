package network

import (
	"fmt"
	"log"
	"net"
)

/*
"Looking up nameservers for a hostname

This program will find nameservers associated with a given hostname.
The primary function here is net.LookupNS():"
*/

func GetNameServers() {
	hostname := "devdungeon.com"

	fmt.Println("Looking up nameservers for " + hostname)

	nameservers, err := net.LookupNS(hostname)
	if err != nil {
		log.Fatal(err)
	}
	for _, nameserver := range nameservers {
		fmt.Println(nameserver.Host)
	}
}
