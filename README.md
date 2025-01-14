# CPF e CNPJ Validator

Este projeto é uma ferramenta de linha de comando para validar números de CPF e CNPJ. Ele foi desenvolvido em Go e pode ser utilizado para verificar a validade desses documentos brasileiros.

## Funcionalidades

- Validação de CPF
- Validação de CNPJ

## Como Usar

### Pré-requisitos

- Go 1.23.4 ou superior

### Instalação

Clone o repositório:

```sh
git clone https://github.com/fabiosoliveira/cpf_cnpj_validator.git
cd cpf_cnpj_validator
```

## Compilação
Para compilar o projeto, execute:

```sh
go build -o validator cmd/cli/main.go
```

## Execução
Para validar um CPF:
```sh
./validator --cpf <cpf>
```

Para validar um CNPJ:
```sh
./validator --cnpj <cnpj>
```

## Exemplos
Validando um CPF:
```sh
./validator --cpf 12345678909
```

Validando um CNPJ:
```sh
./validator --cnpj 12345678000195
```

## Estrutura do Projeto
```sh
.
├── .gitignore
├── cmd/
│   └── cli/
│       └── main.go
├── go.mod
├── LICENSE
├── pkg/
│   └── validator/
│       ├── cnpj-validator.go
│       ├── cnpj-validator_test.go
│       ├── cpf-validator.go
│       ├── cpf-validator_test.go
├── README.md
└── validator.exe
```

## Testes
Para rodar os testes, utilize o comando:
```sh
go test ./pkg/validator/...
```

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

Desenvolvido por Fabio S. Oliveira