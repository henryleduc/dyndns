package main

import (
	"github.com/henryleduc/dyndns/pkg/cloudflare"
	"github.com/henryleduc/dyndns/pkg/helper"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const checkPeriod = 2 * time.Minute

func main() {
	cloudflareEmail := os.Getenv("CLOUDFLARE_API_EMAIL")
	cloudflareKey := os.Getenv("CLOUDFLARE_API_KEY")
	cloudflareZoneID := os.Getenv("CLOUDFLARE_API_ZONEID")

	if cloudflareEmail == "" || cloudflareKey == "" || cloudflareZoneID == "" {
		log.Fatalf("environment variables not defined the following environments variables are required: " +
			"CLOUDFLARE_API_EMAIL and CLOUDFLARE_API_KEY")
	}

	zoneID := uuid.MustParse(cloudflareZoneID)
	client, err := cloudflare.NewClient(cloudflareEmail, cloudflareKey, zoneID)
	if err != nil {
		log.Fatalf("failed to create cloudflare API client: %v", err)
	}

	for {
		err = helper.UpdateAllARecords(client)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(checkPeriod)
	}
}
