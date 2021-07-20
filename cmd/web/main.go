package main

import (
	"fmt"
	"github.com/welihenry/ws/handlers"
	"net/http"
)

func main()  {
	fmt.Println("initializing application")

	mux:= Routes()

	go handlers.ListenToWsChan()

	http.ListenAndServe(":8080", mux)
}
