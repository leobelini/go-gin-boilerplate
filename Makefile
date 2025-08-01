# Caminho para o arquivo principal com as anotações Swagger
SWAGGER_MAIN=internal/handler/http/handler.go

# Nome do binário (ajuste se quiser gerar executável)
BINARY_NAME_SERVER=cashly-server
BINARY_NAME_WORKERS=cashly-workers

.PHONY: build run swagger clean

## Compila o projeto
build:
	go build -o $(BINARY_NAME_SERVER) ./cmd/server
	go build -o $(BINARY_NAME_WORKERS) ./cmd/workers

## Roda o servidor
run-server:
	go run $(SWAGGER_MAIN)

## Roda os workers
run-workers:
	go run ./cmd/workers

## Gera a documentação Swagger
swagger:
	swag init --generalInfo $(SWAGGER_MAIN)

## Limpa arquivos gerados
clean:
	go clean
	rm -f $(BINARY_NAME_SERVER)
	rm -f $(BINARY_NAME_WORKERS)
	rm -rf docs
