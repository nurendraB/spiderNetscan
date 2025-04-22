package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts scans the given subnet and ports for open services
func ScanPorts(subnet string, ports []string) error {
	// Here, you would implement actual port scanning logic, this is just a mock
	for _, port := range ports {
		address := fmt.Sprintf("%s:%s", subnet, port)
		_, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("Port %s is closed or unreachable\n", port)
		} else {
			fmt.Printf("Port %s is open\n", port)
		}
	}
	return nil
}
