package main

import (
	"fmt"
	"github.com/henryleduc/dyndns/pkg/cloudflare"
	"github.com/henryleduc/dyndns/pkg/ip"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	cloudflareEmail := os.Getenv("CLOUDFLARE_API_EMAIL")
	cloudflareKey := os.Getenv("CLOUDFLARE_API_KEY")
	cloudflareZoneID := os.Getenv("CLOUDFLARE_API_ZONEID")
	domain := os.Getenv("ROOT_DOMAIN")
	subdomains := os.Getenv("SUBDOMAINS")

	if cloudflareEmail == "" || cloudflareKey == "" || cloudflareZoneID == "" || domain == "" || subdomains == "" {
		log.Fatalf("environment variables not defined the following environments variables are required: " +
			"CLOUDFLARE_API_EMAIL and CLOUDFLARE_API_KEY and ROOT_DOMAIN and SUBDOMAINS")
	}

	zoneID := uuid.MustParse(cloudflareZoneID)
	client := cloudflare.NewClient(cloudflareEmail, cloudflareKey, zoneID)

	_, _ = client.GetAllDNSRecords()
	//fmt.Println(err)
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Request.URL.RequestURI())
	//fmt.Println(resp.Status)
	//fmt.Println(string(body))

	ipAdr, err := ip.GetIPv4()
	fmt.Println(ipAdr)
	fmt.Println(err)
}
