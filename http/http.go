package http

import (
	"bytes"
	_rsa "crypto/rsa"
	"encoding/base64"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"httpencrypt/rsa"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var client = http.Client{}
var publicKey *_rsa.PublicKey
var privateKey *_rsa.PrivateKey

func Start(port, publicKeyString, privateKeyString, token, targetUrl string, isClient, toGBK bool) error {
	targetUrl = strings.TrimRight(targetUrl, "/")
	publicKey, err := rsa.LoadPublicKey([]byte(publicKeyString))
	if err != nil {
		return err
	}
	privateKey, err := rsa.LoadPrivateKey([]byte(privateKeyString))
	if err != nil {
		return err
	}
	return http.ListenAndServe(port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//进行转发，把路径和内容转发
		bys, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("请求错误[%s]\n", r.RequestURI)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if !isClient {
			if r.Header.Get("xx___token") != token {
				fmt.Printf("认证失败token:%s,期望:%s\n", r.Header.Get("xx___token"), token)
				w.WriteHeader(http.StatusBadGateway)
				return
			}
		}
		if isClient {
			bys, err = rsa.EncryptPKCS1v15(publicKey, bys)
			if err != nil {
				fmt.Printf("加密失败:%s\n", err.Error())
				w.WriteHeader(http.StatusBadGateway)
				return
			}

			enc := base64.StdEncoding.EncodeToString(bys)
			fmt.Printf("提交的加密的内容:%s\n", enc)
		} else {
			enc := base64.StdEncoding.EncodeToString(bys)
			fmt.Printf("接受的加密的内容:%s\n", enc)
			bys, err = rsa.Decrypt(privateKey, bys)
			if err != nil {
				fmt.Printf("解密失败:%s\n", err.Error())
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			if toGBK {
				fmt.Printf("转换到gbk编码:%s", string(bys))
				bys, err = Utf8ToGbk(bys)
				if err != nil {
					fmt.Printf("编码转换失败[%s]\n", r.RequestURI)
					w.WriteHeader(http.StatusBadGateway)
					return
				}
			}
		}
		req, err := http.NewRequest(r.Method, targetUrl+r.URL.Path, bytes.NewReader(bys))
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		copyHeader(r.Header, req.Header)
		if !isClient {
			req.Header.Del("xx___token")
		} else {
			req.Header.Set("xx___token", token)
		}

		rsp, err := client.Do(req)
		if err != nil {
			fmt.Printf("请求%s失败,%s\n", req.URL, err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		defer rsp.Body.Close()
		bys, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Printf("请求%s失败,%s\n", req.URL, err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if isClient {
			bys, err = rsa.Decrypt(privateKey, bys)
			if err != nil {
				fmt.Printf("解密失败:%s\n", err.Error())
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			fmt.Printf("客户端返回内容:%s", string(bys))
		} else {
			bys, err = rsa.EncryptPKCS1v15(publicKey, bys)
			if err != nil {
				fmt.Printf("解密失败:%s\n", err.Error())
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			fmt.Printf("服务端返回内容:%s\n", string(bys))
		}
		copyHeader(rsp.Header, w.Header())
		w.Header().Set("Content-Length", strconv.FormatInt(int64(len(bys)), 10))
		w.WriteHeader(http.StatusOK)
		w.Write(bys)
	}))
}
func copyHeader(source, target http.Header) {
	for k, v := range source {
		for _, vv := range v {
			target.Add(k, vv)
		}
	}
}
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
