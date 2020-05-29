package main

import (
	"ddnsv6/dnspod"
	"ddnsv6/service"
	"flag"
	"fmt"
	"github.com/chai2010/winsvc"
	"time"
)
var dnsPodToken string;
var domain string;
var subDomain string;
var cmd string;
var iptype int;
func init(){
	flag.StringVar(&dnsPodToken, "token", "", " dnspod token")
	flag.StringVar(&domain, "domain", "", " domain")
	flag.StringVar(&subDomain, "subdomain", "www", " subdomain")
	flag.IntVar(&iptype, "iptype", 6, "ip ver")
	flag.StringVar(&cmd, "cmd", "", "cmd instart/uninstart")
}

func main(){
	flag.Parse()
	if(len(dnsPodToken)==0&&len(cmd)==0){
		fmt.Printf("token null")
		return ;
	}
	if(len(domain)==0&&len(cmd)==0){
		fmt.Printf("domain null")
		return ;
	}
	fmt.Printf("dnspodtoken:%s\r\n",dnsPodToken)
	fmt.Printf("domain:%s\r\n",domain)
	fmt.Printf("subdomain:%s\r\n",subDomain)
	fmt.Printf("iptype:%d\r\n",iptype)


	//运行服务
	if !winsvc.IsAnInteractiveSession() {
		if err := winsvc.RunAsService("ddnsv6", StartServer, StopServer, false); err != nil {
			fmt.Printf("service run err:%s\r\n",err.Error())
		}
		return
	}else{
		if(len(cmd)>0){
			switch (cmd) {
			case "install":
				err:=service.Install();
				if(err==nil){
					fmt.Printf("install success\r\n")
				}else{
					fmt.Printf("install fail err:"+err.Error()+"\r\n")
				}
				break;
			case "uninstall":
				err:=service.UnInstall();
				if(err==nil){
					fmt.Printf("uninstall success\r\n")
				}else{
					fmt.Printf("uninstall fail err:"+err.Error()+"\r\n")
				}
				break;
			}
		}
		return ;
	}
	//运行
	StartServer()
}


func StartServer(){
	dnspodClent:=dnspod.DnsPod{Token:dnsPodToken}
	for {
		dnspodClent.DdnsUpdate(iptype,domain,subDomain);
		time.Sleep(5*time.Minute);
	}
}

func StopServer() {
	fmt.Println("exit")
}




