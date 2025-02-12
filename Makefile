DOCKER_COMPOSE_INMEMORY := docker-compose.inmemory.yml

DOCKER_COMPOSE_POSTGRES := docker-compose.postgres.yml

up-inmemory:
	docker compose -f docker-compose.inmemory.yml up --build

up-postgres:
	docker compose -f docker-compose.postgres.yml up --build

down:
	docker compose down

LOCAL_BIN:="$(CURDIR)/bin" 

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-urlshortener-api

generate-urlshortener-api:
	mkdir -p pkg/urlshortener_v3
	protoc --proto_path api/urlshortener_v3 \
	--go_out=pkg/urlshortener_v3 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/urlshortener_v3 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/urlshortener_v3/urlshortener.proto
