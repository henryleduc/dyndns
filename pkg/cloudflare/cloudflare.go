package cloudflare

import (
	"encoding/json"
	"errors"
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
	ID         string `json:"id,omitempty"`
	RecordType string `json:"type"`
	Name       string `json:"name"`
	IP         string `json:"content"`
	TTL        int    `json:"ttl"`
	Proxied    bool   `json:"proxied"`
}

type dnsRecordResponse struct {
	Result DNSRecord
}
type allDNSRecordResponse struct {
	Result []DNSRecord
}

// NewClient will return a new cloudflare client
func NewClient(apiEmail, apiKey string, zoneID uuid.UUID) (Client, error) {
	client := Client{
		apiEmail: apiEmail,
		apiKey:   apiKey,
		zoneID:   zoneID,
	}

	resp, err := client.GetZone()
	if err != nil {
		log.Fatal("failed to find zone given for client")
		return Client{}, err
	}

	// TODO: Refactor once GetZone has been changed to return zone data...
	if resp.StatusCode != 200 {
		return Client{}, errors.New("failed to get zone when creating client, X-API-Key, X-API-Email or ZoneID may be incorrect")
	}

	return client, nil
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

// GetAllDNSRecords will return all the dns records for the current zone
func (c *Client) GetAllDNSRecords() ([]DNSRecord, error) {
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

	jsonBody := &allDNSRecordResponse{}
	err = json.NewDecoder(resp.Body).Decode(jsonBody)
	if err != nil {
		log.Fatalf("failed to unmarshal json response")
		return []DNSRecord{}, err
	}

	if len(jsonBody.Result) <= 0 {
		return []DNSRecord{}, errors.New("error no dns records found for zone")
	}

	return jsonBody.Result, nil
}

// GetDNSRecord will get by a given dns record id
func (c *Client) GetDNSRecord(recordID uuid.UUID) (DNSRecord, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", stripUUID(c.zoneID), stripUUID(recordID)), nil)
	if err != nil {
		log.Fatalf("failed to create request, possibly invalid zoneID (%v) or DNS record ID (%v)", c.zoneID, recordID)
		return DNSRecord{}, err
	}

	c.addDefaultHeaders(req)
	resp, err := execRequest(httpClient, req)
	if err != nil {
		return DNSRecord{}, err
	}

	jsonBody := &dnsRecordResponse{}
	err = json.NewDecoder(resp.Body).Decode(jsonBody)
	if err != nil {
		log.Fatalf("failed to unmarshal json response")
		return DNSRecord{}, err
	}
	return jsonBody.Result, nil
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
