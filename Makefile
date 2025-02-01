include .env

PROTO_DIR=api/proto
OUT_DIR=pkg/proto

proto:
	protoc --proto_path=$(PROTO_DIR) \
	       --go_out=$(OUT_DIR) \
	       --go-grpc_out=$(OUT_DIR) \
	       --validate_out="lang=go:$(OUT_DIR)" \
	       $(PROTO_DIR)/*.proto

create-migration:
	goose -dir ./database/postgres/migrations create example_migration sql

migrate:
	@source .env && goose -dir ./database/postgres/migrations postgres "$${DATABASE_URL}" up
down:
	@source .env && goose -dir ./database/postgres/migrations postgres "$${DATABASE_URL}" down

docker:
	docker-compose -f docker-compose-local.yaml up -d

sqlc:
	sqlc generate -f ./database/sqlc.yaml
