package service

import (
	"ddnsv6/ddns"
	"ddnsv6/param"
	"fmt"
	"time"
)

func StartServer(){
	var ddnsClient ddns.DdnsClient;
	if param.DdnsType=="cloudflare"{
		ddnsClient=&ddns.Cloudflare{param.Email,param.Apikey,param.Zoneid}
	}
	if param.DdnsType=="dnspod"{
		ddnsClient=&ddns.DnsPod{Token:param.DnsPodToken}
	}

	if ddnsClient==nil {
		return ;
	}

	for {
		ddns.DdnsUpdate(ddnsClient,param.Iptype,param.Domain,param.SubDomain);
		time.Sleep(5*time.Minute);
	}
}

func StopServer() {
	fmt.Println("exit")
}
