package main

import (
	"net/http"
	"simplaBank/routes"
)

func main() {
	routes.Routes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
