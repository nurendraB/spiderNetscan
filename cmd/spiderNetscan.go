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
	Version = "latest" // Default version, will be replaced during build
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
	sourceFlag := flag.String("source", "nvd", "CVE data source (nvd or mitre)")
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

	// Scan the ports and get the open ones
	openPorts, err := scanner.ScanPorts(subnet, ports)
	if err != nil {
		fmt.Println("Error scanning ports:", err)
		return
	}

	// If CVE flag is set, fetch and check CVE data
	if *cveFlag {
		if *onlineFlag {
			if *apiKeyFlag == "" {
				fmt.Println("Error: API key is required for online CVE check.")
				return
			}
			// Fetch and check CVE data from the specified online source
			err := scanner.FetchOnlineCVEData(*sourceFlag, *apiKeyFlag, ports, subnet)
			if err != nil {
				fmt.Println("Error fetching CVE data:", err)
				return
			}
		} else {
			// Check CVE data offline from local JSON file
			err := scanner.CheckOfflineCVE(openPorts, "data/cve_data.json", subnet)
			if err != nil {
				fmt.Println("Error checking CVE data:", err)
				return
			}
		}
	}

	// Display open ports only once
	if len(openPorts) > 0 {
		fmt.Printf("Port scan completed successfully. Open ports: %v\n", openPorts)
	} else {
		fmt.Println("No open ports found.")
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
	fmt.Println("\nUpdating spiderNetscan tool...")

	// Check if git is installed
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("❌ Git is not installed. Please install Git and try again.")
	}

	// Stash local changes to avoid merge issues
	exec.Command("git", "stash", "--include-untracked").Run()

	// Pull the latest code from the repo
	cmd := exec.Command("git", "pull")
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output))
	if err != nil {
		return fmt.Errorf("❌ Failed to pull updates: %w", err)
	}

	if strings.Contains(string(output), "Already up to date.") {
		fmt.Println("✅ No new updates available.")
		return nil
	}

	// Get the latest Git tag (version)
	versionBytes, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	version := "latest"
	if err == nil {
		version = strings.TrimSpace(string(versionBytes))
	}

	// Build the updated binary with embedded version
	cmdBuild := exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X main.Version=%s", version), "-o", "spiderNetscan", "cmd/spiderNetscan.go")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Run(); err != nil {
		return fmt.Errorf("❌ Failed to rebuild after update: %w", err)
	}

	// Move the updated binary to /usr/local/bin
	cmdInstall := exec.Command("sudo", "mv", "spiderNetscan", "/usr/local/bin/")
	cmdInstall.Stdout = os.Stdout
	cmdInstall.Stderr = os.Stderr
	if err := cmdInstall.Run(); err != nil {
		return fmt.Errorf("❌ Failed to install updated binary: %w", err)
	}

	// Restore stashed changes if any
	exec.Command("git", "stash", "pop").Run()

	Version = version
	fmt.Printf("✅ spiderNetscan updated to version %s successfully!\n", version)
	return nil
}
