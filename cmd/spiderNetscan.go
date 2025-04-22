package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nurendraB/spiderNetscan/internal/scanner"
)

var (
	Version = "dev" // Default version, will be replaced during build
)

func main() {
	// Print ASCII banner with version
	printBanner()

	// Define flags for the tool
	portsFlag := flag.String("p", "", "Comma-separated list of ports to scan (e.g. 22,80,443)")
	subnetFlag := flag.String("s", "", "CIDR subnet to scan (e.g. 192.168.1.0/24)")
	cveFlag := flag.Bool("cve", false, "Check CVEs for open ports")
	onlineFlag := flag.Bool("online", false, "Fetch CVE data from online sources")
	apiKeyFlag := flag.String("api-key", "", "API key for online CVE sources")
	updateFlag := flag.Bool("update", false, "Update the tool to the latest version")
	versionFlag := flag.Bool("version", false, "Show version of the tool")

	// Parse the flags
	flag.Parse()

	// Handle --version flag
	if *versionFlag {
		fmt.Printf("spiderNetscan version: %s\n", Version)
		os.Exit(0)
	}

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

// printBanner prints the ASCII banner with tool information and version
func printBanner() {
	banner := fmt.Sprintf(`
           _     _              __     _                       
 ___ _ __ (_) __| | ___ _ __ /\ \ \___| |_ ___  ___ __ _ _ __  
/ __| '_ \| |/ __| |/ _ \ '__/  \/ / _ \ __/ __|/ __/ _\ | '_ \ 
\__ \ |_) | | (__| |  __/ | / /\  /  __/ |_\__ \ (_| (_| | | | |
|___/ .__/|_|\___|_|\___|_| \_\ \/ \___|\__|___/\___\__,_|_| |_|
    |_|                                                        

                        Developed by: @nurendraB (spiderinshell)
                        
                        Version: %s
`, Version)
	fmt.Println(banner)
}

// updateTool updates the tool (can be implemented for git pull or other update logic)
func updateTool() error {
	fmt.Println("Updating spiderNetscan tool...")

	// Pull latest code
	cmd := exec.Command("git", "pull")
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		return fmt.Errorf("failed to pull updates: %w", err)
	}

	if strings.Contains(string(output), "Already up to date.") {
		fmt.Println("No new updates available.")
		return nil
	}

	// Get the latest Git tag as version (fallback to "latest" if no tag exists)
	versionBytes, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	version := "latest"
	if err == nil {
		version = strings.TrimSpace(string(versionBytes))
	}

	// Rebuild the binary with updated version
	cmdBuild := exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X main.Version=%s", version), "-o", "spiderNetscan", "cmd/spiderNetscan.go")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Run(); err != nil {
		return fmt.Errorf("failed to rebuild after update: %w", err)
	}

	// Move the binary to /usr/local/bin for global use
	cmdInstall := exec.Command("sudo", "mv", "spiderNetscan", "/usr/local/bin/")
	cmdInstall.Stdout = os.Stdout
	cmdInstall.Stderr = os.Stderr
	if err := cmdInstall.Run(); err != nil {
		return fmt.Errorf("failed to install updated binary: %w", err)
	}

	fmt.Println("Update and rebuild successful!")
	return nil
}
