package nets

import (
	"errors"
	"net"
)

// 获取本机ip,根据传入的网关获取
func GetLocalIp(preferGateway string) (string, error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr = range addrs {
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过ipv6
			if ipNet.IP.To4() != nil {
				contains := ipNet.Contains(net.ParseIP(preferGateway))
				if contains {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", errors.New("未找到地址段")
}
