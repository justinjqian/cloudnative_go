package main
 
import (
   "fmt"
   "net/http"
//   "context"
//	"io"
//	"log"
//	"os"
//	"os/signal"
)
 
//func HelloHandler(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Hello World")
//}

func HandleConn(w http.ResponseWriter, req *http.Request) {

   
//w，给客户端回复数据
//req，读取客户端发送的数据
func HandleConn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req.Method = ", req.Method)         //请求方法
	fmt.Println("req.URL = ", req.URL)               //请求的地址
	fmt.Println("req.Header = ", req.Header)         //头部信息，map键值对
	fmt.Println("req.Body = ", req.Body)             //空
	fmt.Println("req.RemoteAddr = ", req.RemoteAddr) //客户端IP和端口
   w.Header().Set("name", "my name is smallsoup")
   w.WriteHeader(202)
   w.Write([]byte("hello world\n"))
}



func main() {
	//注册处理函数，用户连接时，自动调用指定的处理函数
	http.HandleFunc("/go.html", HandleConn)
   http.HandleFunc("healthz",healthz)

	//监听绑定
	http.ListenAndServe(":8000", nil)
}

 
//func main () {
//   http.HandleFunc("/", HelloHandler)
//   http.ListenAndServe(":8845", nil)
//}