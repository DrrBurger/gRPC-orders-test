# установливает зависимости
deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc

# билдит бинарный файл
build:
	go build -o gRPC-orders-test -v

# запускает сервер grpc
run_serv:
	go run cmd/server/main.go

# запускает код клиента
run_cli:
	go run cmd/client/main.go

# генерирует код из .proto файла
gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/order_service/order_service.proto

# очищяет и удаляет бинарный файл
clean:
	go clean
	rm -f gRPC-orders-test

# запускает все тесты
test:
	go test -v ./...

.PHONY: test deps gen_proto build test run_serv run_cli