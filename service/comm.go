package service

import (
	"ddnsv6/dnspod"
	"ddnsv6/param"
	"fmt"
	"time"
)

func StartServer(){
	dnspodClent:=dnspod.DnsPod{Token:param.DnsPodToken}
	for {
		dnspodClent.DdnsUpdate(param.Iptype,param.Domain,param.SubDomain);
		time.Sleep(5*time.Minute);
	}
}

func StopServer() {
	fmt.Println("exit")
}
