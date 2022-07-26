package main
 
import (
   "fmt"
   "net/http"
    // 03.设置version
	"os"
)
 

// Job4 当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, req *http.Request) {
   w.WriteHeader(200)
   w.Write([]byte("healthz is ok\n"))
}
   
//w，给客户端回复数据
//req，读取客户端发送的数据


func homework02(w http.ResponseWriter, req *http.Request) {




	fmt.Println("req.Method = ", req.Method)         //请求方法
	fmt.Println("req.URL = ", req.URL)               //请求的地址
	fmt.Println("req.Header = ", req.Header)         //头部信息，map键值对
	fmt.Println("req.RemoteAddr = ", req.RemoteAddr) //客户端IP和端口

   // 02.将requst中的header 设置到 reponse中
      for k, v := range req.Header {
      for _, vv := range v {
         //fmt.Printf("Header key: %s, Header value: %s \n", k, v)
         w.Header().Set(k, vv)
      }
   }

   //w.Header().Set("name", "my name is smallsoup")
   w.WriteHeader(200)
   w.Write([]byte("HomeworkMoudle2\n"))
}


func get_envVersion(w http.ResponseWriter, req *http.Request) {
    // 03.设置version  v2022.7.23
   os.Setenv("VERSION", "v2022.7.23")
   version:= os.Getenv("VERSION")
   fmt.Println("req.Header = ", version)

	fmt.Println("req.Method = ", req.Method)         //请求方法
	fmt.Println("req.URL = ", req.URL)               //请求的地址
	//fmt.Println("req.Header = ", req.Header)         //头部信息，map键值对
	//fmt.Println("req.Body = ", req.Body)             //空
	fmt.Println("req.RemoteAddr = ", req.RemoteAddr) //客户端IP和端口

   w.WriteHeader(200)
   w.Write([]byte("check version!\n"))
}


func main() {
	//注册处理函数，用户连接时，自动调用指定的处理函数
   http.HandleFunc("/homework02",homework02)
   http.HandleFunc("/get_envVersion",get_envVersion)
   // Job4 当访问 localhost/healthz 时，应返回 200
   http.HandleFunc("/healthz",healthz)

	//监听绑定
	http.ListenAndServe(":8080", nil)
}

