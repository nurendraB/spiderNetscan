package scanner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// CVEItem represents a CVE entry in JSON format
type CVEItem struct {
	CVEID       string `json:"CVE_id"`
	Description string `json:"description"`
	Port        string `json:"port"`
}

// CheckOfflineCVE checks for CVE data in the local JSON file
func CheckOfflineCVE(ports []string, filePath string) error {
	// Read the CVE data from the JSON file
	data, err := os.ReadFile(filePath) // Using os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return fmt.Errorf("failed to read CVE data: %w", err)
	}

	var cveData []CVEItem
	if err := json.Unmarshal(data, &cveData); err != nil {
		return fmt.Errorf("failed to parse CVE data: %w", err)
	}

	// Check if the ports are vulnerable based on the CVE data
	for _, cve := range cveData {
		for _, port := range ports {
			if cve.Port == port {
				fmt.Printf("CVE ID: %s | Description: %s | Port: %s\n", cve.CVEID, cve.Description, cve.Port)
			}
		}
	}

	return nil
}

// FetchOnlineCVEData fetches CVE data from an online source (like NVD)
func FetchOnlineCVEData(apiKey string, ports []string) error {
	// Example of using the NVD API (replace with a real URL)
	url := fmt.Sprintf("https://api.nvd.nist.gov/vuln/search?apiKey=%s&cpeName=*", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch CVE data from NVD: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Using io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var cveData []CVEItem
	if err := json.Unmarshal(body, &cveData); err != nil {
		return fmt.Errorf("failed to parse CVE data: %w", err)
	}

	// Filter and display relevant CVE data for the given ports
	for _, cve := range cveData {
		for _, port := range ports {
			if strings.Contains(cve.Port, port) {
				fmt.Printf("CVE ID: %s | Description: %s | Port: %s\n", cve.CVEID, cve.Description, cve.Port)
			}
		}
	}

	return nil
}
