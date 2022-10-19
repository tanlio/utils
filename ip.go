package utils

import (
	"net"
	"os"
)

func GetLocalIP() string {
	var localIP string

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			localIP = ipv4.String()
			break
		}
	}

	return localIP
}
