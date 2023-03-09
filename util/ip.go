package util

import "net"

// GetIpFromAddr 获取当前IP
func GetIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

// GetLocalIPList 获取当前IP列表
func GetLocalIPList() (ipList []net.IP) {
	faces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if face.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrList, err := face.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrList {
			ip := GetIpFromAddr(addr)
			if ip == nil {
				continue
			}
			ipList = append(ipList, ip)
		}
	}
	return
}
