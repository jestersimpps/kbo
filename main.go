package main

import (
	"fmt"
	"net/http"
)

 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
