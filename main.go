package main
import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>Bienvenue dans Golang !</h1>")
	
}
func main() {
	http.handlerFunc("/",handlerFunc)
	fmt.Println("Starting the server on : 3000...")
	http.ListenAndServe(":3000",nil)
}

