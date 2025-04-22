package scanner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CVEItem represents a CVE entry in JSON format
type CVEItem struct {
	CVEID       string `json:"CVE_id"`
	Description string `json:"description"`
	Port        string `json:"port"`
}

// CheckOfflineCVE checks for CVE data in the local JSON file
func CheckOfflineCVE(ports []string, filePath, subnet string) error {
	// Scan ports to check which are open
	openPorts, err := ScanPorts(subnet, ports)
	if err != nil {
		return fmt.Errorf("failed to scan ports: %w", err)
	}

	// Read the CVE data from the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read CVE data: %w", err)
	}

	var cveData []CVEItem
	if err := json.Unmarshal(data, &cveData); err != nil {
		return fmt.Errorf("failed to parse CVE data: %w", err)
	}

	// Check if the ports are vulnerable based on the CVE data
	for _, cve := range cveData {
		for _, openPort := range openPorts {
			if cve.Port == openPort {
				fmt.Printf("CVE ID: %s | Description: %s | Port: %s\n", cve.CVEID, cve.Description, cve.Port)
			}
		}
	}

	return nil
}

// FetchOnlineCVEData fetches CVE data from either NVD or MITRE
func FetchOnlineCVEData(source, apiKey string, ports []string, subnet string) error {
	switch source {
	case "nvd":
		// Fetch from NVD
		url := fmt.Sprintf("https://api.nvd.nist.gov/vuln/search?apiKey=%s&cpeName=*", apiKey)
		return fetchCVEFromURL(url, ports, subnet)
	case "mitre":
		// Fetch from MITRE
		url := fmt.Sprintf("https://cve.mitre.org/api/cve-search?apiKey=%s", apiKey)
		return fetchCVEFromURL(url, ports, subnet)
	default:
		return fmt.Errorf("unsupported CVE source: %s", source)
	}
}

// fetchCVEFromURL fetches CVE data from a given URL
func fetchCVEFromURL(url string, ports []string, subnet string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch CVE data from URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var cveData []CVEItem
	if err := json.Unmarshal(body, &cveData); err != nil {
		return fmt.Errorf("failed to parse CVE data: %w", err)
	}

	// Scan ports to check which are open
	openPorts, err := ScanPorts(subnet, ports)
	if err != nil {
		return fmt.Errorf("failed to scan ports: %w", err)
	}

	// Filter and display relevant CVE data for the open ports
	for _, cve := range cveData {
		for _, openPort := range openPorts {
			if cve.Port == openPort {
				fmt.Printf("CVE ID: %s | Description: %s | Port: %s\n", cve.CVEID, cve.Description, cve.Port)
			}
		}
	}

	return nil
}
