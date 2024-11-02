package main

import (
	"fmt"
	"net/http"

	"ganja/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error %s", err))
	}
}
