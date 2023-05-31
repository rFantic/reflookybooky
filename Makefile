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

upProfile:
	docker compose --profile $(Profile) up -d --build --remove-orphans
