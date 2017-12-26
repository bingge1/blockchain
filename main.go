package main

import (
	"fmt"
	"net/http"
	"router"
)

func main() {
	app := http.NewServeMux()
	router.Init()
	http.ListenAndServe(":18888", app)
}
func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html></head><title>go的第一章</title></head><body><h1>server is start on localhost:18888！</h1></body></html>"))
	fmt.Printf("")
	fmt.Printf("")
}
