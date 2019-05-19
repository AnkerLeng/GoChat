package main

import (
	"net/http"
)

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		//数据库操作
		//逻辑处理
		//rest api json/xml 返回

		// 1. 获取前端传递的参数
		// mobile, password
		// 解析参数
		// 如何获得参数
		// 解析参数

		_ = request.ParseForm()

		mobile := request.PostForm.Get("mobile")
		password := request.PostForm.Get("password")
		loginok := false
		if (mobile == "168000000000" && password == "123456") {
			loginok = true
		}
		//使用application/x-www-form-urlencoded POST数据
		//curl http://localhost:8080/user/login -X POST -d "mobile=186000000000&password=123456"
		str := `{"code":0,"data":{"id":1,"token":"test"}}`
		if (!loginok) {
			//返回失败的json
			str = `{"code":-1,"msg":"密码不正确"}`
		}
		//设置header为json 默认的text/html 所以特别指出返回的格式为
		//为application/json
		writer.Header().Set("Content-Type", "application/json")
		//返回json ok
		_, _ = writer.Write([]byte(str))

	})

	// 启动web服务器
	_ = http.ListenAndServe(":8080", nil)
}
