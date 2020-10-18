package main

import (
	"ddnsv6/service"
	"fmt"
	"ddnsv6/param"
	"runtime"
)



func main(){
	param.Parse()
	if(len(param.DnsPodToken)==0&&len(param.Cmd)==0){
		fmt.Printf("token null")
		return ;
	}
	if(len(param.Domain)==0&&len(param.Cmd)==0){
		fmt.Printf("domain null")
		return ;
	}
	fmt.Printf("dnspodtoken:%s\r\n",param.DnsPodToken)
	fmt.Printf("domain:%s\r\n",param.Domain)
	fmt.Printf("subdomain:%s\r\n",param.SubDomain)
	fmt.Printf("iptype:%d\r\n",param.Iptype)


	//运行服务
	if(runtime.GOOS=="windows") {
		service.StartWin();
	}
	//运行
	service.StartServer()
}





