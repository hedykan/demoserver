package httpserver

import (
	"fmt"
	"net/http"

	httphelper "github.com/hedykan/httpHelper"
)

func Serve() {
	port := ":10001"
	mux := http.NewServeMux()
	getHanderArr := httphelper.HandleArr{}
	postHanderArr := httphelper.HandleArr{}

	httphelper.SetMuxHandle(mux, getHanderArr)
	httphelper.SetMuxHandle(mux, postHanderArr)

	fmt.Println("lisent", port)
	http.ListenAndServe(port, mux)
}
