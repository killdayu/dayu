package parse

import "net"

func IP(IPs string) ([]net.IP, error) {
	_, ipNet, err := net.ParseCIDR(IPs)
	if err != nil {
		return nil, err
	}

	ipList := make([]net.IP, 0)
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ipCopy := make(net.IP, len(ip))
		copy(ipCopy, ip)
		ipList = append(ipList, ipCopy)
	}

	return ipList, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
