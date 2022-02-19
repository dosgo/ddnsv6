//go:build !windows

package ddns

import (
	"bytes"
	"io/ioutil"

	"tinygo.org/x/drivers/net/http"
)

func post(_url string, contentType string, data []byte, headers map[string]string) ([]byte, error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}
	req, err := http.NewRequest("POST", _url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func put(_url string, contentType string, data []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("PUT", _url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
