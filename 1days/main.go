package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", handlerGet)
	r.POST("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "POST METHOD")
		if err != nil {
			w.WriteHeader(500)
		}
	})
	r.Run(":8081")
}

func handlerGet(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "home url = %q", req.URL.Path)
	if err != nil {
		w.WriteHeader(500)
	}
}
