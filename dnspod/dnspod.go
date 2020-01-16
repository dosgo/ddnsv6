package dnspod

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	params["format"]="json";
	params["login_token"]=dp.Token;
	var paramStr = url.Values{}
	for k, v := range params {
		paramStr.Add(k,v.(string))
	}
	client := http.DefaultClient
	//access_token在url中，内容在request body中
	fmt.Printf("pastr:%s\r\n",paramStr.Encode())
	resp, err := client.Post(api+cmd, "application/x-www-form-urlencoded", bytes.NewReader([]byte(paramStr.Encode())))
	if err != nil {
		return nil,err;
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err;
	}
	return result,nil
}

func (dp *DnsPod) Ddns(domain string,value string,sub_domain string,record_type string) (error) {
	res,err:=dp.GetRecord(domain,record_type,sub_domain);
	if(err!=nil){
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
		return errors.New(status["message"].(string))
	}
	return nil;
}

func (dp *DnsPod) Modify(domain string,value string,sub_domain string,record_type string) (error) {
	res,err:=dp.GetRecord(domain,record_type,sub_domain);
	if(err!=nil){
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
		return errors.New(status["message"].(string))
	}
	return nil;
}

