package param

import "flag"
var DnsPodToken string;
var Domain string;
var SubDomain string;
var Cmd string;
var Iptype int;
var DdnsType string
var Email string;
var Apikey string;
var Zoneid string;

func init(){
	flag.StringVar(&DnsPodToken, "token", "", " dnspod token")
	flag.StringVar(&Domain, "domain", "", " domain")
	flag.StringVar(&SubDomain, "subdomain", "www", " subdomain")
	flag.IntVar(&Iptype, "iptype", 6, "ip ver")
	flag.StringVar(&Cmd, "cmd", "", "cmd instart/uninstart")
	flag.StringVar(&DdnsType, "ddnstype", "cloudflare", "ddns type dnspod or cloudflare")
	flag.StringVar(&Email, "email", "", "cloudflare email")
	flag.StringVar(&Apikey, "apikey", "", "cloudflare apikey")
	flag.StringVar(&Zoneid, "zoneid", "", "cloudflare zoneid")
}

func Parse(){
	flag.Parse()
}