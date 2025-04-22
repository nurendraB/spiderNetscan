package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts scans the given subnet and ports for open services and returns the open ports
func ScanPorts(subnet string, ports []string) ([]string, error) {
	openPorts := []string{}
	// Here, you would implement actual port scanning logic
	for _, port := range ports {
		address := net.JoinHostPort(subnet, port)
		_, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("Port %s is closed or unreachable\n", port)
		} else {
			fmt.Printf("Port %s is open\n", port)
			openPorts = append(openPorts, port)
		}
	}
	return openPorts, nil
}
