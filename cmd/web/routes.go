package main

import (
	"github.com/bmizerany/pat"
	"github.com/welihenry/ws/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
