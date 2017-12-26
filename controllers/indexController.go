package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("访问了/index 页面")
	w.Write([]byte("<html></head><title>index/</title></head><body><h1>index.html</h1></body></html>"))
}
