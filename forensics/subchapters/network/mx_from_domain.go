package network

import (
	"fmt"
	"log"
	"net"
)

/*
"Looking up MX records

This program will take a domain name and return the MX records.
MX records, or mail exchanger records, are DNS records that point to the mail server.
For example, the MX server of https://www.devdungeon.com/ is mail.devdungeon.com. The
net.LookupMX() function performs this lookup and returns a slice of the net.MX structs:"
*/

func GetMXFromDomain() {
	domain := "devdungeon.com"

	fmt.Println("Looking up MX records for " + domain)

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, mxRecord := range mxRecords {
		fmt.Printf("Host: %s\tPreference: %d\n", mxRecord.Host, mxRecord.Pref)
	}
}
