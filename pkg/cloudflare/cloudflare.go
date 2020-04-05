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

// GetAllZones will return all zones for the user
func (c *Client) GetAllZones() (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones", nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
	return resp, nil
}

// GetZone will return the details of a zone by it's id
func (c *Client) GetZone() (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", stripUUID(c.zoneID)), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
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

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", stripUUID(c.zoneID), stripUUID(recordID)), body)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
	return resp, nil
}

// GetDNSRecord will get by a given dns record id
func (c *Client) GetAllDNSRecords() (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", stripUUID(c.zoneID)), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v)", c.zoneID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
	return resp, nil
}

// GetDNSRecord will get by a given dns record id
func (c *Client) GetDNSRecord(recordID uuid.UUID) (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", stripUUID(c.zoneID), stripUUID(recordID)), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v) or DNS record ID (%v)", c.zoneID, recordID)
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
	return resp, nil
}

// GetUser will verify that the current client credentials are valid.
func (c *Client) GetUser() (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/user", nil)
	if err != nil {
		log.Fatal("failed to create request")
		return nil, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return nil, err
	}

	// TODO: return response body json rather than http.Response
	return resp, nil
}

func (c *Client) addDefaultHeaders(req *http.Request) {
	req.Header.Add("X-Auth-Email", c.apiEmail)
	req.Header.Add("X-Auth-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
}

func stripUUID(id uuid.UUID) string {
	return strings.ReplaceAll(id.String(), "-", "")
}

func execRequest(httpClient *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("failed to execute request %s (%s): %v", resp.Request.URL.RequestURI(), resp.Status, err)
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("failed to execute request %s (%s)", resp.Request.URL.RequestURI(), resp.Status)
		return nil, err
	}

	return resp, err
}