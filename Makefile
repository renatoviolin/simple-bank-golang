postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:alpine3.15

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	ENV_MODE=DEV go run main.go

server-prod:
	ENV_MODE=PROD go run main.go

mock:
	mockgen --destination db/mock/store.go --package mockdb github.com/renatoviolin/simplebank/db/sqlc Store

server-container:
	docker rm simplebank --force && docker run --name go-simplebank --network bank-network -p 8000:8000 -e DB_SOURCE="postgresql://postgres:secret@go-postgres:5432/simple_bank?sslmode=disable" -d simplebank


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 server-prod