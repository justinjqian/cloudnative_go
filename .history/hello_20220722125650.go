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

