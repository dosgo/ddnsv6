package iptool

import "net"
import "errors"
/*
*获取本机公网IP
*ipver 4 ipv4  6 ipv6
 */
func GetPublicIP(ipver int) (ip string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	//取IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			//ipv4跳过ipv6
			if(ipver==4&&ipNet.IP.To16() != nil){
				continue;
			}
			//ipv6跳过ipv4
			if(ipver==6&&ipNet.IP.To4() != nil){
				continue;
			}
			if(IsPublicIP(ipNet.IP)){
				ip = ipNet.IP.String() // 192.168.1.1
				return ;
			}
		}
	}
	err = errors.New("no public ip")
	return
}


//IsPublicIP 判断是否公网IP,支持IPv4,IPv6
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	// IPv4私有地址空间
	// A类：10.0.0.0到10.255.255.255
	// B类：172.16.0.0到172.31.255.255
	// C类：192.168.0.0到192.168.255.255
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		case ip4[0] == 169 && ip4[1] == 254:
			return false
		default:
			return true
		}
	}
	// IPv6私有地址空间：以前缀FEC0::/10开头
	if ip6 := ip.To16(); ip6 != nil {
		if ip6[0] == 15 && ip6[1] == 14 && ip6[2] <= 12 {
			return false
		}
		return true
	}
	return false
}
