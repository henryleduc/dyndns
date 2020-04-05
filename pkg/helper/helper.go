package helper

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/henryleduc/dyndns/pkg/cloudflare"
	"github.com/henryleduc/dyndns/pkg/ip"
	"log"
)

var currentIP string

func UpdateAllARecords(client cloudflare.Client) error {
	dnsRecords, err := client.GetAllDNSRecords()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get all dns records from cloudflare: %v", err))
	}

	newIP, err := ip.GetIPv4()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get IPv4: %v", err))
	}
	if currentIP == newIP {
		return nil
	}
	currentIP = newIP
	log.Printf("Found new IP Address of %s", newIP)

	for _, dnsRecord := range dnsRecords {
		if dnsRecord.RecordType == "A" {
			dnsRecord.IP = newIP

			id := uuid.MustParse(dnsRecord.ID)
			dnsRecord.ID = "" // done so that the filed gets omitted as cloudflare does not accept the ID in the put request

			resp, err := client.PutDNSRecord(id, dnsRecord)
			if err != nil || resp.StatusCode != 200 {
				log.Printf("error failed to update record %s: %v", dnsRecord.Name, err)
			}
			log.Printf("Succesfully Updated Record %s to New IP %s", dnsRecord.Name, newIP)
		}
	}

	return nil
}
