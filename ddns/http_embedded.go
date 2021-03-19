// +build mipsle  mips  mips64  mips64le

package ddns

import (
	"io/ioutil"
	"net/url"
	"crypto/tls"
	"net"
	"strings"
)

func  post(_url string,contentType string,data []byte,headers map[string]string  ) ([]byte, error) {
	return _req("POST",_url,contentType,data,headers);
}

func  _req(method  string,_url string,contentType string,data []byte,headers map[string]string  ) ([]byte, error) {
	uInfo,err:=url.Parse(_url);
	if err!=nil {
		return nil,err;
	}
	var host=uInfo.Host;
	var path="/";
	var query="";
	var port="";
	if strings.HasPrefix(_url,"http://"){
		port="80"
	}else{
		port="443"
	}
	if uInfo.Path!="" {
		path=uInfo.Path
	}
	if uInfo.Query()!=nil {
		query="?"+uInfo.Query().Encode()
	}
	if uInfo.Port()!="" {
		port=uInfo.Port()
	}




	// building POST-request:
	request:=method+" "+path+query+" HTTP/1.1\n";
	request+="Host: "+host+"\n";
	request+="Content-type: "+contentType+"\n";
	//request+="Content-length: "+strconv.Itoa(len(data))+"\n";
	request+="Connection: close\n";
	request+="\n";
	request+=string(data)+"\n";



	var fp net.Conn;
	if strings.HasPrefix(_url,"http://"){
		fp, err = net.Dial("tcp", host+":"+port)
	}else{
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		fp, err = tls.Dial("tcp", host+":"+port, tlsConf)
	}
	if err!=nil {
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

func  put(_url string,contentType string,data []byte,headers map[string]string  ) ([]byte, error) {
	return _req("PUT",_url,contentType,data,headers);
}

func  get(_url string,headers map[string]string) ([]byte,error) {
	uInfo,err:=url.Parse(_url);
	if err!=nil {
		return nil,err;
	}
	var host=uInfo.Host;
	var path="/";
	var query="";
	var port="";
	if strings.HasPrefix(_url,"http://"){
		port="80"
	}else{
		port="443"
	}
	if uInfo.Path!="" {
		path=uInfo.Path
	}
	if uInfo.Query()!=nil {
		query="?"+uInfo.Query().Encode()
	}
	if uInfo.Port()!="" {
		port=uInfo.Port()
	}

	// building GET-request:
	request:="GET "+path+query+" HTTP/1.1\n";
	request+="Host: "+host+"\n";
	request+="Connection: close\n";
	request+="\n";



	var fp net.Conn;
	if strings.HasPrefix(_url,"http://"){
		fp, err = net.Dial("tcp", host+":"+port)
	}else{
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}
		fp, err = tls.Dial("tcp", host+":"+port, tlsConf)
	}
	if err!=nil {
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