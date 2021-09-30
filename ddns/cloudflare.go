package ddns

import (
	"encoding/json"
	"errors"
	"fmt"
)


type Cloudflare struct {
	Email string
	Apikey string
	Zoneid string
}

var  cloudflareApi string="https://api.cloudflare.com/client/v4/";




func (dp *Cloudflare)   post(cmd string, params map[string]interface{})  ([]byte,error) {
	data,err:=json.Marshal(params)
	if err!=nil {
		return nil,err;
	}
	var headers=make(map[string]string)
	headers["X-Auth-Email"]=dp.Email;
	headers["X-Auth-Key"]=dp.Apikey;
	return post(cloudflareApi+cmd,"application/json",data,headers)
}




func (dp *Cloudflare) get(cmd string) ([]byte,error) {
	var headers=make(map[string]string)
	headers["X-Auth-Email"]=dp.Email;
	headers["X-Auth-Key"]=dp.Apikey;
	headers["Content-type"]="application/json";
	return get(cloudflareApi+cmd,headers);
}



func (dp *Cloudflare)   put(cmd string, params map[string]interface{})  ([]byte,error){
	data,err:=json.Marshal(params)
	if err!=nil {
		return nil,err;
	}
	var headers=make(map[string]string)
	headers["X-Auth-Email"]=dp.Email;
	headers["X-Auth-Key"]=dp.Apikey;
	return put(cloudflareApi+cmd,"application/json",data,headers)
}



func (dp *Cloudflare)  getDomainID(domain string) string{
	result,err:=dp.get("zones/"+dp.Zoneid+"/dns_records?name="+domain);
	if err==nil {
		//复用了map
		res:= make(map[string]interface{})
		err = json.Unmarshal(result, &res)

		if err!=nil {
			return "";
		}
		if res["success"].(bool) {
			resData:=res["result"].([]interface{})
			for i := range resData {
				item:=resData[i].(map[string]interface{})
				if item["name"].(string)==domain {
					return item["id"].(string)
				}
			}
		} else {
			return "";
		}
	}
	return ""
}

func (dp *Cloudflare)   getuser() (string,error) {
	result,err:= dp.get("user");
	if err!=nil {
		return "",err;
	}
	//复用了map
	res:= make(map[string]interface{})
	err = json.Unmarshal(result, &res)
	fmt.Printf("user:%+v\r\n",res);
	return "",err;
}

func (dp *Cloudflare)   updateDNS(domain string,ip string,_type string,ttl int) error{
	domainid:=dp.getDomainID(domain);
	var params=make(map[string]interface{})
	params["type"]=_type;
	params["name"]=domain;
	params["content"]=ip;
	params["proxied"]=false;
	params["ttl"]=ttl;

	var err error
	var result []byte
	if domainid=="" {
		result,err= dp.post("zones/"+dp.Zoneid+"/dns_records",params);
	}else{
		result,err= dp.put("zones/"+dp.Zoneid+"/dns_records/"+domainid,params);
	}
	if err!=nil {
		return err;
	}
	//复用了map
	res:= make(map[string]interface{})
	err = json.Unmarshal(result, &res)
	if err!=nil {
		return err;
	}
	if !res["success"].(bool) {
		return errors.New("update err")
	}
	return nil;
}

func (dp *Cloudflare) Modify(domain string,value string,sub_domain string,record_type string) (error) {
	return dp.updateDNS(sub_domain+"."+domain, value, record_type,120);
}
