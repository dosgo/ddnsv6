package main

import (
	"ddnsv6/dnspod"
	"flag"
	"fmt"
)
var dnsPodToken string;
func init(){
	flag.StringVar(&dnsPodToken, "token", "", " dnspod token")
}

func main(){
	flag.Parse()
	if(len(dnsPodToken)==0){
		fmt.Printf("token null")
		return ;
	}
	dnspodClent:=dnspod.DnsPod{Token:dnsPodToken}
	err:=dnspodClent.Modify("xxx.com","2408:805f:e220:","ddd6","AAAA");

	fmt.Printf("err:%#v\r\n",err)

}

