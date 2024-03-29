package iptool

import (
	"net"
)
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
				ip = ipNet.IP.String()
				return
			}
		}
	}
	err = errors.New("no public ip")
	return
}

func isPrivateIPv4(ip net.IP) bool {
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	return private24BitBlock.Contains(ip) ||
		private20BitBlock.Contains(ip) ||
		private16BitBlock.Contains(ip) ||
		ip.IsLoopback() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast()
}

func isPrivateIPv6(ip net.IP) bool {
	_, block, _ := net.ParseCIDR("fc00::/7")
	return block.Contains(ip) || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
}



//IsPublicIP 判断是否公网IP,支持IPv4,IPv6
func IsPublicIP(ip net.IP) bool {

	// IPv4私有地址空间
	// A类：10.0.0.0到10.255.255.255
	// B类：172.16.0.0到172.31.255.255
	// C类：192.168.0.0到192.168.255.255
	if ip4 := ip.To4(); ip4 != nil {
		return !isPrivateIPv4(ip);
	}
	// IPv6私有地址空间：以前缀FEC0::/10开头
	if ip6 := ip.To16(); ip6 != nil {
		return !isPrivateIPv6(ip);
	}
	return false
}
