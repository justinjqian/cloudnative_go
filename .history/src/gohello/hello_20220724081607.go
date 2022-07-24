package main

import (
	"fmt"
	"net/http"
	// 02.For version
	"os"
	// 03.For version
	"log"
	"strings"
)

// 04. 当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("healthz is working\n"))
}

//w，给客户端回复数据
//req，读取客户端发送的数据

func homework02(w http.ResponseWriter, req *http.Request) {
	// 02.设置version  v2022.7.23 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "v2022.7.23")
	version := os.Getenv("VERSION")
	fmt.Println("req.Header = ", version)

	//

	fmt.Println("req.Method = ", req.Method) //请求方法
	fmt.Println("req.URL = ", req.URL)       //请求的地址
	fmt.Println("req.Header = ", req.Header) //头部信息，map键值对

	// 02.将requst中的header 设置到 reponse中
	for k, v := range req.Header {
		for _, vv := range v {
			//fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	// 04.记录日志并输出
	//fmt.Println("req.RemoteAddr = ", req.RemoteAddr) //客户端IP和端口
	clientip := getIP(req)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)

	w.WriteHeader(200)
	w.Write([]byte("<h1>Welcome to Cloud Native</h1>  Homework - Moudle 2\n"))
}

func getIP(req *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
   //ip := req.Header.Get("X-Forwarded-For")
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := req.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(req.RemoteAddr, ":")[0]
	}
	return ip
}

func main() {
	//注册处理函数，用户连接时，自动调用指定的处理函数
	http.HandleFunc("/homework02", homework02)
	// Job4 当访问 localhost/healthz 时，应返回 200
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/healthz", healthz)
	//监听绑定
	http.ListenAndServe(":8080", nil)
}
