package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	//router.NotFound
	//router.PanicHandler
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id")
		_, err := fmt.Fprintf(writer, text)
		if err != nil {
			return
		}
	})
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello World!")
	})

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
