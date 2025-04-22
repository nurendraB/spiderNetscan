package scanner

import (
	"fmt"
	"net"
	"time"
)

// formatAddress properly formats IP and port for IPv4 and IPv6
func formatAddress(ip string, port string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP != nil && parsedIP.To4() == nil {
		// IPv6
		return fmt.Sprintf("[%s]:%s", ip, port)
	}
	// IPv4 or invalid (let it fail gracefully)
	return fmt.Sprintf("%s:%s", ip, port)
}

// ScanPorts scans the given ports on the specified subnet
func ScanPorts(ip string, ports []string) error {
	for _, port := range ports {
		fmt.Printf("Scanning %s on port %s...\n", ip, port)
		address := formatAddress(ip, port)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			fmt.Printf("Port %s is closed or unreachable on %s\n", port, ip)
		} else {
			fmt.Printf("Port %s is open on %s\n", port, ip)
			conn.Close()
		}
	}
	return nil
}
