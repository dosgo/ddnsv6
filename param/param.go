package param

import "flag"
var DnsPodToken string;
var Domain string;
var SubDomain string;
var Cmd string;
var Iptype int;
func init(){
	flag.StringVar(&DnsPodToken, "token", "", " dnspod token")
	flag.StringVar(&Domain, "domain", "", " domain")
	flag.StringVar(&SubDomain, "subdomain", "www", " subdomain")
	flag.IntVar(&Iptype, "iptype", 6, "ip ver")
	flag.StringVar(&Cmd, "cmd", "", "cmd instart/uninstart")
}

func Parse(){
	flag.Parse()
}