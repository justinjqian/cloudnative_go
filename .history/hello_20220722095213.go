//package main
//
//import (
//	"fmt"
// )

//func main() {
//	fmt.Println("Hello,playgroud")
//}



package main
 
import (
   "fmt"
   "net/http"
)
 
func HelloHandler(w http.ResponseWriter, r *http.Request) {cd 
   fmt.Fprintf(w, "Hello World")
}
 
func main () {
   http.HandleFunc("/", HelloHandler)
   http.ListenAndServe(":8000", nil)

————————————————
版权声明：本文为CSDN博主「kevin_tech」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/kevin_tech/article/details/104100835