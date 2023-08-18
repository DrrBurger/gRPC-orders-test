## deps: Установливает зависимости
deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc

## build: Билдит бинарный файл
build:
	go build -o gRPC-orders-test -v

## run_serv: Запускает сервер grpc
run_serv:
	go run cmd/server/main.go

## run_cli: Запускает код клиента
run_cli:
	go run cmd/client/main.go

## gen_proto: Генерирует код из .proto файла
gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/order_service/order_service.proto

## clean: Очищяет и удаляет бинарный файл
clean:
	go clean
	rm -f gRPC-orders-test

## test: Запускает все тесты
test:
	go test -v ./...

## fmt: Форматирование кода для соответствия стандартному стилю Go
fmt:
	go fmt ./...

## vet: Статический анализ кода на поиск подозрительных конструкций
vet:
	go vet ./...

help: Makefile
	@echo " Choose a command run in "gRPC-orders-test":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test deps gen_proto build test run_serv run_cli fmt vet help