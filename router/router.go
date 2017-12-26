package router

import (
	"fmt"
	"net/http"

	controller "github.com/bingge/blockchain/controllers"
)

func Init(mux *http.ServeMux) {
	fmt.Printf("init all router mapping")
	mux.HandleFunc("/index", controller.Index)
	mux.HandleFunc("/user", controller.UserIndex)
	mux.HandleFunc("/", controller.Index)
}
