SQLC_VER = 1.25.0
PG_PORT = 5432
PG_USER = postgres
PG_PASSWORD = qwerty
PG_DB = postgres

PROTO_DIR=./api
OUT_DIR=./pkg/proto
VALIDATE_DIR=./protoc-gen-validate
GOOGLEAPIS_DIR=./googleapis

.PHONY: run proto-get proto-install proto-gen migrate-up migrate-down migrate-create \
        sqlc test-db-up test-db-down test-db-exec test-db-delete buf-update

# ---------- Запуск сервиса ----------

run:
	GRPC_PORT=9007 \
	POSTGRES_HOST=localhost \
	POSTGRES_PORT=$(PG_PORT) \
	POSTGRES_USER=$(PG_USER) \
	POSTGRES_DB=$(PG_DB) \
	POSTGRES_PASSWORD=$(PG_PASSWORD) \
	POSTGRES_SSLMODE=disable \
	go run ./cmd/auth

# ---------- Buf / Protobuf ----------

buf-update:
	buf dep update # require VPN or access

proto-get:
	git clone https://github.com/envoyproxy/protoc-gen-validate.git $(VALIDATE_DIR)
	git clone https://github.com/googleapis/googleapis.git $(GOOGLEAPIS_DIR)

proto-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

proto-gen:
	protoc \
		-I $(PROTO_DIR) \
		-I $(VALIDATE_DIR) \
		-I $(GOOGLEAPIS_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUT_DIR) --grpc-gateway_opt=paths=source_relative \
		--validate_out=lang=go,paths=source_relative:$(OUT_DIR) \
		$(PROTO_DIR)/sync/final-boss/v1/auth.proto

# ---------- SQLC ----------

sqlc:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc:$(SQLC_VER) generate

# ---------- Postgres локально через docker ----------

test-db-up:
	docker run -e POSTGRES_USER=$(PG_USER) -e POSTGRES_PASSWORD=$(PG_PASSWORD) \
		-p $(PG_PORT):5432 --name mock-test-db -d \
		postgres:15 postgres -c log_statement=all || docker start mock-test-db
	docker exec mock-test-db timeout 20s bash -c "until pg_isready -d $(PG_DB) -U $(PG_USER); do sleep 0.5; done"
	sleep 0.5

test-db-down:
	docker stop mock-test-db || true

test-db-delete: test-db-down
	docker rm mock-test-db || true

test-db-exec:
	docker exec -it mock-test-db psql -U $(PG_USER)

# ---------- Миграции ----------

migrate-create:
	migrate create -ext sql -dir migrations/postgres mock

migrate-up:
	migrate -path ./migrations/postgres -database "postgres://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable" up

migrate-down:
	migrate -path ./migrations/postgres -database "postgres://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable" down -all

swagger-gen:
	protoc \
		-I $(PROTO_DIR) \
		-I $(VALIDATE_DIR) \
		-I $(GOOGLEAPIS_DIR) \
		--openapiv2_out=docs/swagger --openapiv2_opt logtostderr=true \
		$(PROTO_DIR)/sync/final-boss/v1/auth.proto
