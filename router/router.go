package router

import (
	"blockchain/controllers"
	"fmt"
	"net/http"
)

func Init(mux http.ServeMux) {
	mux.HandleFunc("/index", indexController.Index)
}
