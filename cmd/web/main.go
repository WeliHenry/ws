package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("initializing application")

	mux:= Routes()

	http.ListenAndServe(":8080", mux)
}
