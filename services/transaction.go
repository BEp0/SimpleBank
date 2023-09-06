package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simplaBank/model"
)

func Transaction(writer http.ResponseWriter, request *http.Request) {
	users := []model.User{
		model.CreateUser("Bernardo",
			"bernardo@gmail.com",
			"111.111.111-11",
			model.CPF,
			"11"),
		model.CreateUser("Ivan",
			"ivan@gmail.com",
			"222.222.222-22",
			model.CPF,
			"22"),
	}
	fmt.Println("Subiu o /", users)

	jsonResp, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Deu merda")
	}
	_, erro := writer.Write(jsonResp)
	if erro != nil {
		return
	}
	return
}
