package routes

import (
	"fmt"
	"net/http"
	"simplaBank/services"
)

func Routes() {
	fmt.Println("Rotas aqui:")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/transaction", services.Transaction)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}
