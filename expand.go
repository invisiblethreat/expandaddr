package expandaddr

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

func ExpandAddrs(addrs string) ([]string, error) {
	// Discrete IP addresses, comma seperated.
	if strings.Contains(addrs, ",") {
		addrs = strings.ReplaceAll(addrs, " ", "")
		return expandCommaDelim(addrs), nil
		// CIDR that needs expanded
	} else if strings.Contains(addrs, "/") {
		return expandCIDR(addrs)
	}
	// Single address inserted into slice
	return []string{addrs}, nil

}

func expandCommaDelim(s string) []string {
	return strings.Split(s, ",")
}

func expandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); addrInc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	if len(ips) == 1 {
		return ips, nil
	}
	return ips[1 : len(ips)-1], nil
}
func addrInc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ExpandPorts(port string) ([]string, error) {
	var ports []string
	if strings.Contains(port, "-") {
		ranges := strings.Split(port, "-")
		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			return nil, err
		}
		if end > maxPort {
			end = maxPort
		}
		for i := start; i <= end; i++ {
			ports = append(ports, strconv.Itoa(i))
		}

	} else if strings.Contains(port, ",") {
		port = strings.ReplaceAll(port, " ", "")
		return expandCommaDelim(port), nil
	} else {
		ports = []string{port}
	}
	return ports, nil
}

func ExpandPortsString(port string) ([]string, error) {
	return ExpandPorts(port)

}

func ExpandPortsInt(port string) ([]int, error) {
	var ports []int
	if strings.Contains(port, "-") {
		ranges := strings.Split(port, "-")
		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			return nil, err
		}
		if end > maxPort {
			end = maxPort
		}
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}

	} else if strings.Contains(port, ",") {
		port = strings.ReplaceAll(port, " ", "")
		portStr := expandCommaDelim(port)

		for _, v := range portStr {
			casted, err := strconv.Atoi(v)
			if err != nil {
				continue
			} else {
				ports = append(ports, casted)
			}
		}
	} else {
		casted, err := strconv.Atoi(port)
		if err != nil {
			return ports, errors.New("No valid ports specified: " + err.Error())
		}
		ports = []int{casted}
	}
	return ports, nil
}
