package ddns

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


type Cloudflare struct {
	Email string
	Apikey string
	Zoneid string
}

var  cloudflareApi string="https://api.cloudflare.com/client/v4/";




func (dp *Cloudflare)   Post(cmd string, params map[string]interface{})  ([]byte,error) {

	data,err:=json.Marshal(params)
	if err!=nil {
		return nil,err;
	}
	url := cloudflareApi+cmd;
	payload := strings.NewReader(string(data))

	req, err := http.NewRequest("POST", url, payload)
	if err!=nil {
		return nil,err;
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Email",dp.Email)
	req.Header.Add("X-Auth-Key",dp.Apikey)

	res, err := http.DefaultClient.Do(req)
	if err!=nil {
		return nil,err;
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}




func (dp *Cloudflare) Get(cmd string) ([]byte,error) {
	url := cloudflareApi+cmd;
	req, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		return nil,err;
	}
	req.Header.Add("X-Auth-Email",dp.Email)
	req.Header.Add("X-Auth-Key",dp.Apikey)
	req.Header.Add("Content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err!=nil {
		return nil,err;
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}



func (dp *Cloudflare)   Put(cmd string, params map[string]interface{})  ([]byte,error){

	data,err:=json.Marshal(params)
	if err!=nil {
		return nil,err;
	}

	url := cloudflareApi+cmd;
	payload := strings.NewReader(string(data))

	req, err := http.NewRequest("PUT", url, payload)
	if err!=nil {
		return nil,err;
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Email",dp.Email)
	req.Header.Add("X-Auth-Key",dp.Apikey)

	res, err := http.DefaultClient.Do(req)
	if err!=nil {
		return nil,err;
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}



func (dp *Cloudflare)  GetDomainID(domain string) string{
	result,err:=dp.Get("zones/"+dp.Zoneid+"/dns_records?name="+domain);
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

func (dp *Cloudflare)   Getuser() (string,error) {
	result,err:= dp.Get("user");
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
	domainid:=dp.GetDomainID(domain);
	var params=make(map[string]interface{})
	params["type"]=_type;
	params["name"]=domain;
	params["content"]=ip;
	params["proxied"]=false;
	params["ttl"]=ttl;

	var err error
	var result []byte
	if domainid=="" {
		result,err= dp.Post("zones/"+dp.Zoneid+"/dns_records",params);
	}else{
		result,err= dp.Put("zones/"+dp.Zoneid+"/dns_records/"+domainid,params);
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
