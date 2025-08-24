package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil) // nil to use default mux

}
func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the HomePage!</h1>"))
	fmt.Fprintf(w, "This is the home page")
	fmt.Println("Request received at console")
}
