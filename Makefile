.PHONY: os gqlgen build lint test updb downdb dropdb

os:
	@echo ${OSFLAG}

genGql:
	go run github.com/99designs/gqlgen

genEnt:
	go generate ./ent

genProtoc:
	protoc --go_out=pb --go-grpc_out=pb \
	--go_opt=paths=import \
	--go-grpc_opt=paths=import \
	./proto/*.proto

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(OSFLAG) GOARCH=$(GOARCH) go build -ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT=$(COMMIT) -X main.DATE=$(DATE) -w -s" -v -o server cmd/main.go

api: build
	./server api

db:
	docker compose up -d db

run:
	docker compose up -d

teardown:
	docker compose down -v --remove-orphans
	docker compose rm --force --stop -v