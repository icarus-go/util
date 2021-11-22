package network

import (
	"errors"
	"net"
	"strings"
)

type _ip struct{}

var IP = new(_ip)

func (ip *_ip) Get() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := ip.getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}

func (*_ip) getIpFromAddr(addr net.Addr) net.IP {
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

//GetRequestIP 获取请求时使用的IP
func GetRequestIP(uri ...string) string {
	if len(uri) == 0 {
		uri = append(uri, "www.baidu.com:80")
	}

	// conn, err := net.Dial("udp", "www.baidu.com:80")
	conn, err := net.Dial("udp", uri[0])
	if err != nil {
		return ""
	}
	defer conn.Close()

	addr := strings.Split(conn.LocalAddr().String(), ":")

	return addr[0]
}
