package model

import (
	"errors"
	"strings"
)

const (
	CPF  = "CPF"
	CNPJ = "CNPJ"
)

type User struct {
	Name         string
	Email        string
	Document     string
	DocumentType string
	password     string
}

func CreateUser(name string,
	email string,
	document string,
	documentType string,
	password string) User {

	docType, _ := getDocumentType(documentType)

	return User{
		Name:         name,
		Email:        email,
		Document:     document,
		DocumentType: docType,
		password:     password,
	}
}

func getDocumentType(documentType string) (string, error) {
	docType := strings.ToUpper(documentType)
	isNotValidType := docType != CPF || docType != CNPJ
	if isNotValidType {
		return CPF, errors.New("ERROR, DocumentType isn't CPF or CNPJ type")
	}
	return docType, nil
}
