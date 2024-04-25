migrate/create:
	migrate create -ext=sql -dir=internal/infra/database/sql/migrations -seq init

migrate/up:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrate/down:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

graphql:
	go run github.com/99designs/gqlgen generate

grpc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

wire:
	wire ./cmd/ordersystem/

.PHONY: migrate/create migrate/up migrate/down graphql grpc wire