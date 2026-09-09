package main
import(
  "fmt"
  "net/http"
)
func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r "http.Resquest){
      fmt.Fprint(w,"Hello from Google app engine - Go!")
})
    http.ListenAndServer(":8000",nil)
}
  
                  
