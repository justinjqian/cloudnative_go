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
func healthz(w http.ResponseWriter, req *http.Request) {
   w.WriteHeader(200)
   w.Write([]byte("healthz is OK\n"))
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