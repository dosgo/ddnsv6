package main

import (
	"ddnsv6/param"
	"ddnsv6/service"
	"fmt"
	"runtime"
)



func main(){
	param.Parse()
	if len(param.Cmd)==0 {
		if len(param.Domain) == 0 && len(param.Cmd) == 0 {
			fmt.Printf("domain null")
			return;
		}

		if param.DdnsType=="dnspod" {
			if len(param.DnsPodToken) == 0  {
				fmt.Printf("token null")
				return;
			}
		}else{
			if len(param.Email) == 0  {
				fmt.Printf("cloudflare email null")
				return;
			}
			if len(param.Apikey) == 0  {
				fmt.Printf("cloudflare apikey null")
				return;
			}
			if len(param.Zoneid) == 0  {
				fmt.Printf("cloudflare zoneid null")
				return;
			}
		}
	}
	fmt.Printf("dnspodtoken:%s\r\n",param.DnsPodToken)
	fmt.Printf("domain:%s\r\n",param.Domain)
	fmt.Printf("subdomain:%s\r\n",param.SubDomain)
	fmt.Printf("iptype:%d\r\n",param.Iptype)

	//运行服务
	if runtime.GOOS=="windows" {
		service.StartWin();
	}else {
		//运行
		service.StartServer()
	}
}





