package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/nurendraB/spiderNetscan/internal/scanner"
)

func main() {
	// Print ASCII banner
	printBanner()

	// Define flags for the tool
	portsFlag := flag.String("p", "", "Comma-separated list of ports to scan (e.g. 22,80,443)")
	subnetFlag := flag.String("s", "", "CIDR subnet to scan (e.g. 192.168.1.0/24)")
	cveFlag := flag.Bool("cve", false, "Check CVEs for open ports")
	onlineFlag := flag.Bool("online", false, "Fetch CVE data from online sources")
	apiKeyFlag := flag.String("api-key", "", "API key for online CVE sources")
	updateFlag := flag.Bool("update", false, "Update the tool to the latest version")

	// Parse the flags
	flag.Parse()

	// Handle --update flag
	if *updateFlag {
		err := updateTool()
		if err != nil {
			fmt.Println("Error updating tool:", err)
			return
		}
		return
	}

	// Validate and process flags
	if *portsFlag == "" || *subnetFlag == "" {
		fmt.Println("Error: Both -p (ports) and -s (subnet) are required.")
		return
	}

	ports := strings.Split(*portsFlag, ",")
	subnet := *subnetFlag

	// Fetch and check CVE data if needed
	if *cveFlag {
		if *onlineFlag {
			if *apiKeyFlag == "" {
				fmt.Println("Error: API key is required for online CVE check.")
				return
			}
			err := scanner.FetchOnlineCVEData(*apiKeyFlag, ports)
			if err != nil {
				fmt.Println("Error fetching CVE data:", err)
				return
			}
		} else {
			err := scanner.CheckOfflineCVE(ports, "data/cve_data.json")
			if err != nil {
				fmt.Println("Error checking CVE data:", err)
				return
			}
		}
	}

	// Proceed with scanning (port scanning or other operations can be added here)
	err := scanner.ScanPorts(subnet, ports)
	if err != nil {
		fmt.Println("Error scanning ports:", err)
	}
}

// printBanner prints the ASCII banner with tool information
func printBanner() {
	banner := `
           _     _              __     _                       
 ___ _ __ (_) __| | ___ _ __ /\ \ \___| |_ ___  ___ __ _ _ __  
/ __| '_ \| |/ __| |/ _ \ '__/  \/ / _ \ __/ __|/ __/ _\ | '_ \ 
\__ \ |_) | | (__| |  __/ | / /\  /  __/ |_\__ \ (_| (_| | | | |
|___/ .__/|_|\___|_|\___|_| \_\ \/ \___|\__|___/\___\__,_|_| |_|
    |_|                                                        

                        Developed by: @nurendraB (spiderinshell)
`
	fmt.Println(banner)
}

// updateTool updates the tool (can be implemented for git pull or other update logic)
func updateTool() error {
	fmt.Println("Updating spiderNetscan tool...")
	cmd := exec.Command("git", "pull")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to update the tool: %w", err)
	}
	fmt.Println("Update successful!")
	return nil
}
