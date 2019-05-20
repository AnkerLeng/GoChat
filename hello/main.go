package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// 将函数抽离出来
func userLogin(writer http.ResponseWriter, request *http.Request) {
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
	if (loginok) {
		// {"id":1,"token":"xxx"}
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "")
	} else {
		Resp(writer, 1, nil, "密码不正确")
	}

}

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // data 为null不显示 omitempty
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header为json 默认的text/html 所以特别指出返回的格式为
	//为application/json
	w.Header().Set("Content-Type", "application/json")
	//返回json ok

	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化为json字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	_, _ = w.Write(ret)
}

func RegisterView() {
	tpl, err := template.ParseGlob("hello/view/**/*")
	//如果报错就不要继续了
	if nil != err {
		log.Fatal(err.Error())

	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			_ = tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}

}

func main() {

	/*

		// 绑定请求和处理函数
		http.HandleFunc("/user/login", userLogin)

		// 1.提供静态资源目录支持
		//http.Handle("/",http.FileServer(http.Dir(".")))
		// 提供指定目录静态资源服务
		http.Handle("/asset/", http.FileServer(http.Dir(".")))
		//user.login.shtml
		http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
			//解析
			tpl, err := template.ParseFiles("hello/view/user/login.html")
			if nil != err {
				log.Fatal(err.Error())

			}
			_ = tpl.ExecuteTemplate(writer, "/user/login.shtml", nil)

		})

	*/
	RegisterView()
	// 启动web服务器
	_ = http.ListenAndServe(":8080", nil)
}
