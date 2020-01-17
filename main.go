package main

import (
	"ddnsv6/dnspod"
	"ddnsv6/iptool"
	"flag"
	"fmt"
)
var dnsPodToken string;
var domain string;
var subDomain string;
var iptype int;
func init(){
	flag.StringVar(&dnsPodToken, "token", "", " dnspod token")
	flag.StringVar(&domain, "domain", "", " domain")
	flag.StringVar(&subDomain, "subdomain", "www", " subdomain")
	flag.IntVar(&iptype, "iptype", 6, "ip ver")
}

func main(){
	flag.Parse()
	if(len(dnsPodToken)==0){
		fmt.Printf("token null")
		return ;
	}
	if(len(domain)==0){
		fmt.Printf("domain null")
		return ;
	}

	if(iptype==4&&iptype==0){
		ipv4,errv4:=iptool.GetPublicIP(4);
		if(errv4==nil&&len(ipv4)>0){
			fmt.Printf("find ipv4 address")
			dnspodClent:=dnspod.DnsPod{Token:dnsPodToken}
			err:=dnspodClent.Modify(domain,ipv4,subDomain,"A");
			fmt.Printf("err:%#v\r\n",err)
		}
	}
	if(iptype==6||iptype==0){
		ipv6,errv6:=iptool.GetPublicIP(6);
		if(errv6==nil&&len(ipv6)>0){
			fmt.Printf("find ipv6 address")
			dnspodClent:=dnspod.DnsPod{Token:dnsPodToken}
			err:=dnspodClent.Modify(domain,ipv6,subDomain,"AAAA");
			fmt.Printf("err:%#v\r\n",err)
		}
	}

}


