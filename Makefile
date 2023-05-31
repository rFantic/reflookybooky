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

up:
	docker compose --profile graphql up -d --build --remove-orphans

down:
	docker compose --profile graphql down --remove-orphans

reUp: down up

upDebug:
	docker compose -f docker-compose.debug.yml --profile graphql up -d --build --remove-orphans

downDebug:
	docker compose -f docker-compose.debug.yml --profile graphql down --remove-orphans

reUpDebug: downDebug upDebug