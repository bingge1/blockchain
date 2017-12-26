// userController
package controllers

import (
	"fmt"
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问了/user 页面")
	w.Write([]byte("<html></head><title>index/</title></head><body><h1>users.html</h1></body></html>"))
}
