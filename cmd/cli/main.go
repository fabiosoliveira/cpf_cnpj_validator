package main

import (
	"fmt"
	"os"

	"github.com/fabiosoliveira/cpf_cnpj_validator/pkg/validator"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		errorMessage()
	}

	_validator := args[0]
	_document := args[1]

	switch _validator {
	case "--cpf":
		if !validator.IsValidCPF(&_document) {
			fmt.Println("Invalid CPF")
			os.Exit(1)
		}
		fmt.Println("Valid CPF")
	case "--cnpj":
		if !validator.IsValidCNPJ(&_document) {
			fmt.Println("Invalid CNPJ")
			os.Exit(1)
		}
		fmt.Println("Valid CNPJ")
	default:
		errorMessage()
	}
}

func errorMessage() {
	fmt.Println("Usage: validator --cpf <cpf>")
	fmt.Println("Usage: validator --cnpj <cnpj>")
	os.Exit(1)
}
