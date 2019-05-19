package main

import (
	"io"
	"net/http"
)

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		//数据库操作
		//逻辑处理
		//rest api json/xml 返回
		io.WriteString(writer, "hello world")
	})

	// 启动web服务器
	_ = http.ListenAndServe(":8080", nil)
}
