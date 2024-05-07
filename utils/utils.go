package utils

import (
	"fmt"
	"net"
	"time"
)

// Check whether the service running at the given ip and port is running
func IsServerUp(ip string, port uint16) bool {
	server := fmt.Sprintf("%s:%d", ip, port)
	_, err := net.DialTimeout("tcp", server, 5*time.Second)
	return err == nil
}
