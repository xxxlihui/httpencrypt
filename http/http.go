package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Start(port, key, token, targetUrl string, isClient bool) error {
	http.Handle("", http.NotFoundHandler())
	return http.ListenAndServe(port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//进行转发，把路径和内容转发
		bys, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("请求错误[%s]", r.RequestURI)
			return
		}
		http.NewRequest(r.Method, r.RequestURI)
	}))
}
func copyHeader(header http.Header) http.Header {
	t := &http.Header{}
	for k, v := range header {
		for _, vv := range v {
			t.Add(k, vv)
		}
	}
	return *t
}
