package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/henryleduc/dyndns"
)

func main() {
	cloudflareEmail := os.Getenv("CLOUDFLARE_API_EMAIL")
	cloudflareKey := os.Getenv("CLOUDFLARE_API_KEY")
	cloudflareZoneID := os.Getenv("CLOUDFLARE_API_ZONEID")
	if cloudflareEmail == "" || cloudflareKey == "" || cloudflareZoneID == "" {
		log.Fatalf("environment variables CLOUDFLARE_API_EMAIL && CLOUDFLARE_API_KEY were not defined")
	}

	zoneID := uuid.MustParse(cloudflareZoneID)
	client := cloudflare.NewClient(cloudflareEmail, cloudflareKey, zoneID)

	_, err := client.GetAllZones()
	fmt.Println(err)

}
