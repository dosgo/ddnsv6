package dnspod

import (
	"crypto/tls"
	"ddnsv6/iptool"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"strconv"
	"strings"
)


type DnsPod struct {
	Token   string
}

var  api string="https://dnsapi.cn/";

func (dp *DnsPod) GetRecord(domain string,record_type string,sub_domain string) ( map[string]interface{}, error){
	var params=make(map[string]interface{})
	params["domain"]=domain;
	params["sub_domain"]=sub_domain;
	params["record_type"]=record_type;
	result,err:=dp.post("Record.List",params)
	//复用了map
	res:= make(map[string]interface{})
	err = json.Unmarshal(result, &res)
	return res,err;
}



func (dp *DnsPod) post(cmd string, params map[string]interface{}) ([]byte, error) {
	return dp.postv2(cmd,params);
}

/*Not using the http library is to reduce the size, because it needs to run in embedded
*/
func (dp *DnsPod) postv2( cmd string, params map[string]interface{}) ([]byte, error){
	params["format"]="json";
	params["login_token"]=dp.Token;
	urlStr:=api+cmd
	uInfo,err:=url.Parse(urlStr);
	if err!=nil {
		return nil,err;
	}
	var host=uInfo.Host;
	var path="/";
	var query="";
	var port="";
	if strings.HasPrefix(urlStr,"http://"){
		port="80"
	}else{
		port="443"
	}
	if(uInfo.Path!=""){
		path=uInfo.Path
	}
	if(uInfo.Query()!=nil){
		query="?"+uInfo.Query().Encode()
	}
	if(uInfo.Port()!=""){
		port=uInfo.Port()
	}

	// making string from $data
	var paramStr = url.Values{}
	for k, v := range params {
		paramStr.Add(k,v.(string))
	}


	// building POST-request:
	request:="POST "+path+query+" HTTP/1.1\n";
	request+="Host: "+host+"\n";
	request+="Content-type: application/x-www-form-urlencoded\n";
	request+="Content-length: "+strconv.Itoa(len(paramStr.Encode()))+"\n";
	request+="Connection: close\n";
	request+="\n";
	request+=paramStr.Encode()+"\n";



	var fp net.Conn;
	if strings.HasPrefix(urlStr,"http://"){
		fp, err = net.Dial("tcp", host+":"+port)
	}else{
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		fp, err = tls.Dial("tcp", host+":"+port, tlsConf)
	}
	if(err!=nil){
		return nil,err;
	}
	fp.Write([]byte(request))
	var result []byte;
	result,err=ioutil.ReadAll(fp)
	if err!=nil {
		return nil,err;
	}else{
		res := strings.Split(string(result),"\r\n\r\n")
		var TransferEncoding="";
		headers:=strings.Split(res[0],"\n");
		for _, v:= range headers {
		   if(strings.HasPrefix(strings.ToLower(v),strings.ToLower("Transfer-Encoding"))){
			  _header:=strings.Split(v,":")
			   TransferEncoding=_header[1];
		   }
		}
		if strings.Index(TransferEncoding,"chunked")!=-1 {
			bodys:=strings.Split(res[1],"\n")
			var i=0;
			var body="";
			for _, v:= range bodys {
				if(i==0||i==len(bodys)-1){
					i++
					continue;
				}
				body=body+v;
				i++;
			}
			return []byte(body),err;

		}else{
			return []byte(res[1]),err;
		}
	}
}

func (dp *DnsPod) Ddns(domain string,value string,sub_domain string,record_type string) (error) {
	res,err:=dp.GetRecord(domain,record_type,sub_domain);
	if(err!=nil){
		fmt.Printf("GetRecord error\r\n")
		return err;
	}
	status:=res["status"].(map[string]interface{})
	if(status["code"].(string)=="1"){
		records:=res["records"].([]interface{})
		var record_id="";
		var record_line_id="";
		var oldValue="";
		for _, record := range records {
			recordInfo:=record.(map[string]interface{});
			if(recordInfo["name"].(string)==sub_domain){
				record_id=recordInfo["id"].(string)
				record_line_id=recordInfo["line_id"].(string)
				oldValue=recordInfo["value"].(string)
			}
		}
		if(record_id==""){
			return errors.New("Record not found")
		}

		if(oldValue!=value){
			var params=make(map[string]interface{})
			params["domain"]=domain;
			params["sub_domain"]=sub_domain;
			params["record_id"]=record_id;
			params["record_line_id"]=record_line_id;
			params["value"]=value;
			fmt.Printf("params:%#v\r\n",params)
			result,err:=dp.post("Record.Ddns",params)
			if(err!=nil){
				return err
			}
			//复用了map
			res1:= make(map[string]interface{})
			err = json.Unmarshal(result, &res1)
			if(err!=nil){
				return err
			}
			status1:=res1["status"].(map[string]interface{})
			if(status1["code"].(string)=="1"){
				return nil;
			}else{
				return errors.New(status1["message"].(string))
			}

		}else{
			return errors.New("No changes in records")
		}
		return nil;
	}else{
		fmt.Printf("GetRecord2 error\r\n")
		return errors.New(status["message"].(string))
	}
	return nil;
}

func (dp *DnsPod) Modify(domain string,value string,sub_domain string,record_type string) (error) {
	res,err:=dp.GetRecord(domain,record_type,sub_domain);
	if(err!=nil){
		fmt.Printf("Modify:GetRecord error\r\n")
		return err;
	}
	status:=res["status"].(map[string]interface{})
	if(status["code"].(string)=="1"){
		records:=res["records"].([]interface{})
		var record_id="";
		var record_line_id="";
		var oldValue="";
		for _, record := range records {
			recordInfo:=record.(map[string]interface{});
			if(recordInfo["name"].(string)==sub_domain){
				record_id=recordInfo["id"].(string)
				record_line_id=recordInfo["line_id"].(string)
				oldValue=recordInfo["value"].(string)
			}
		}
		if(record_id==""){
			return errors.New("Record not found")
		}

		if(oldValue!=value){
			var params=make(map[string]interface{})
			params["domain"]=domain;
			params["sub_domain"]=sub_domain;
			params["record_id"]=record_id;
			params["record_line_id"]=record_line_id;
			params["record_type"]=record_type
			params["value"]=value;
			fmt.Printf("params:%#v\r\n",params)
			result,err:=dp.post("Record.Modify",params)
			if(err!=nil){
				return err
			}
			//复用了map
			res1:= make(map[string]interface{})
			err = json.Unmarshal(result, &res1)
			if(err!=nil){
				return err
			}
			status1:=res1["status"].(map[string]interface{})
			if(status1["code"].(string)=="1"){
				return nil;
			}else{
				return errors.New(status1["message"].(string))
			}

		}else{
			return errors.New("No changes in records")
		}
		return nil;
	}else{
		fmt.Printf("Modify:GetRecord2 error\r\n")
		return errors.New(status["message"].(string))
	}
	return nil;
}


func (dp *DnsPod) DdnsUpdate(iptype int,domain string,subDomain string){
	fmt.Printf("checkIpUpdateing....\r\n");
	if(iptype==4||iptype==0){
		ipv4,errv4:=iptool.GetPublicIP(4);
		if(errv4==nil&&len(ipv4)>0){
			fmt.Printf("find ipv4 address :%s\r\n",ipv4)
			if(dp.CheckIP(4,ipv4,domain,subDomain)==false) {
				err := dp.Modify(domain, ipv4, subDomain, "A");
				fmt.Printf("err:%#v\r\n",err)
			}else{
				fmt.Printf("No changes\r\n")
			}
		}
	}
	if(iptype==6||iptype==0){
		ipv6,errv6:=iptool.GetPublicIP(6);
		if(errv6==nil&&len(ipv6)>0){
			fmt.Printf("find ipv6 address :%s\r\n",ipv6)
			//Modify  update
			if(dp.CheckIP(6,ipv6,domain,subDomain)==false) {
				err := dp.Modify(domain, ipv6, subDomain, "AAAA");
				fmt.Printf("err:%#v\r\n", err)
			}else{
				fmt.Printf("No changes\r\n")
			}
		}
	}
}


func (dp *DnsPod) CheckIP(iptype int,caddr string,domain string,subDomain string)bool{
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
	if(err!=nil){
		return false;
	}
	if(iptype==6){
		if(strings.Index(addr.String(),":")!=-1&&caddr==addr.String()){
			return true;
		}
	}
	if(iptype==4){
		if(strings.Index(addr.String(),":")==-1&&caddr==addr.String()){
			return true;
		}
	}
	return false;
}
