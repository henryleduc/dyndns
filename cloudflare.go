package cloudflare

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Client is the Cloudflare API client the allows you to access and update different records on the cloudflare API
type Client struct {
	apiEmail string
	apiKey   string
	zoneID   uuid.UUID
}

// DNSRecord is the Cloudflare API DNS Record data used by PutDnsRecord
type DNSRecord struct {
	RecordType string `json:"type"`
	Name       string
	Content    string
	TTL        int
	Proxied    bool
}

// NewClient will return a new cloudflare client
func NewClient(apiEmail, apiKey string, zoneID uuid.UUID) Client {
	return Client{
		apiEmail: apiEmail,
		apiKey:   apiKey,
		zoneID:   zoneID,
	}
}

// GetAllZones will return all zones for the given ZoneID on the client
func (c *Client) GetAllZones() (*http.Response, error) {
	httpClient := &http.Client{}

	reqURI := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", c.zoneID.String())
	req, err := http.NewRequest("GET", reqURI, nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetZone will return the details of a zone by it's id
func (c *Client) GetZone() (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", c.zoneID.String()), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PutDNSRecord will update a given dns record id and new dns record
func (c *Client) PutDNSRecord(recordID uuid.UUID, dnsRecord DNSRecord) (*http.Response, error) {
	httpClient := &http.Client{}

	dnsRecordJSON, err := json.Marshal(dnsRecord)
	if err != nil {
		log.Fatalf("failed to marhsal DNS Record to JSON: (%v)", dnsRecord)
		return nil, err
	}
	body := strings.NewReader(string(dnsRecordJSON))

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", c.zoneID.String(), recordID.String()), body)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetDNSRecord will get by a given dns record id
func (c *Client) GetDNSRecord(recordID uuid.UUID) (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", c.zoneID.String(), recordID.String()), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v) or DNS record ID (%v)", c.zoneID, recordID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) addDefaultHeaders(req *http.Request) {
	req.Header.Add("X-Auth-Email", c.apiEmail)
	req.Header.Add("X-Auth-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
}

func execRequest(httpClient *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("failed to execute request (%s): %v", resp.Status, err)
		return nil, err
	}

	if resp.StatusCode%100 != 2 {
		var body []byte
		log.Fatalf("request was not successful (%s): %v", resp.Status, body)
		return nil, err
	}

	return resp, err
}
