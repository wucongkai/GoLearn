package main

import (
	"fmt"
	"log"
	"net/http"
)

//	定义一个空的结构体，实现方法ServeHTTP
//
// 第一个参数：ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
// 第二个参数：Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)

	log.Fatal(http.ListenAndServe(":9998", engine))
}
