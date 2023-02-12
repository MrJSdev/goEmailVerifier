package main

import (
	"log"
	"net"
	"strings"
)

func IsEmailExist(domain string) bool {
	var hasMx, hasSPF, hasDMARC bool

	maxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error in max record fetch: %v\n", err)
	}

	if len(maxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error in lookup of txt records: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error in lookup of dmarch records: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
		}
	}

	return hasMx && hasSPF && hasDMARC
}
