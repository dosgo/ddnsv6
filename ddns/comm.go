package ddns

import (
	"ddnsv6/iptool"
	"fmt"
	"net"
	"strings"
)

type DdnsClient interface {
	Modify(domain string,value string,sub_domain string,record_type string) (error)
}


func  CheckIP(iptype int,caddr string,domain string,subDomain string)bool{
	var network="ip";
	switch iptype {
	case 4:
		network="ip4"
		break;
	case 6:
		network="ip6"
		break;
	}
	addr,err:=net.ResolveIPAddr(network,subDomain+"."+domain);
	if err!=nil {
		return false;
	}
	if iptype==6 {
		if strings.Index(addr.String(),":")!=-1&&caddr==addr.String() {
			return true;
		}
	}
	if iptype==4 {
		if strings.Index(addr.String(),":")==-1&&caddr==addr.String() {
			return true;
		}
	}
	return false;
}

func  DdnsUpdate(dc DdnsClient,iptype int,domain string,subDomain string){
	fmt.Printf("checkIpUpdateing....\r\n");
	if iptype==4||iptype==0 {
		ipv4,errv4:=iptool.GetPublicIP(4);
		if errv4==nil&&len(ipv4)>0 {
			fmt.Printf("find ipv4 address :%s\r\n",ipv4)
			if CheckIP(4,ipv4,domain,subDomain)==false {
				err := dc.Modify(domain, ipv4, subDomain, "A");
				fmt.Printf("err:%#v\r\n",err)
			}else{
				fmt.Printf("No changes\r\n")
			}
		}
	}
	if iptype==6||iptype==0 {
		ipv6,errv6:=iptool.GetPublicIP(6);
		if errv6==nil&&len(ipv6)>0 {
			fmt.Printf("find ipv6 address :%s\r\n",ipv6)
			//Modify  update
			if CheckIP(6,ipv6,domain,subDomain)==false {
				err := dc.Modify(domain, ipv6, subDomain, "AAAA");
				fmt.Printf("err:%#v\r\n", err)
			}else{
				fmt.Printf("No changes\r\n")
			}
		}
	}
}