package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	//创建gee实例
	r := gee.New()

	//使用GET方法添加路由
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	//启动web服务器
	r.Run(":9979")
}
