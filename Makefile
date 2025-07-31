# Caminho para o arquivo principal com as anotações Swagger
SWAGGER_MAIN=internal/handler/http/handler.go

# Nome do binário (ajuste se quiser gerar executável)
BINARY_NAME=cashly

.PHONY: build run swagger clean

## Compila o projeto
build:
	go build -o $(BINARY_NAME) ./cmd/server

## Roda o servidor
run:
	go run $(SWAGGER_MAIN)

## Gera a documentação Swagger
swagger:
	swag init --generalInfo $(SWAGGER_MAIN)

## Limpa arquivos gerados
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf docs
