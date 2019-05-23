package main

import (
	"net"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func isIntranet(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}

	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

func getIps() []string {
	var ips = []string{"localhost", "127.0.0.1"}
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipStr := ipnet.IP.String(); isIntranet(ipStr) && ipnet.IP.To4() != nil {
					ips = append(ips, ipStr)
				}
			}
		}
	}

	return ips
}

func showAvailableIps(port int) {
	ips := getIps()
	color.Magenta("Available on:")
	for _, v := range ips {
		color.HiGreen("  http://%s:%d\n", v, port)
	}
}
