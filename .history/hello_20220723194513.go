package main
 
import (
   "fmt"
   "net/http"
	"os"
	"strings"
)
 

func getEnvs() {
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		}else{
			println(string(parts[0]),string(parts[1]))
		}
	}
}

func healthz(w http.ResponseWriter, req *http.Request) {
   w.WriteHeader(200)
   w.Write([]byte("healthz is ok\n"))
}
   
//w，给客户端回复数据
//req，读取客户端发送的数据


func test1(w http.ResponseWriter, req *http.Request) {
   getEnvs()
   VERSION:= os.Getenv("Version")
	fmt.Println("req.Method = ", req.Method)         //请求方法
	fmt.Println("req.URL = ", req.URL)               //请求的地址
	//fmt.Println("req.Header = ", req.Header)         //头部信息，map键值对
   fmt.Print
	fmt.Println("req.Body = ", req.Body)             //空
	fmt.Println("req.RemoteAddr = ", req.RemoteAddr) //客户端IP和端口
   w.Header().Set("name", "my name is smallsoup")
   w.WriteHeader(202)
   w.Write([]byte("hello world\n"))
}


func main() {
	//注册处理函数，用户连接时，自动调用指定的处理函数
   http.HandleFunc("/test1",test1)
   http.HandleFunc("/healthz",healthz)

	//监听绑定
	http.ListenAndServe(":8000", nil)
}

 
//func main () {
//   http.HandleFunc("/", HelloHandler)
//   http.ListenAndServe(":8845", nil)
//}