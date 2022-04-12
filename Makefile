postgres:
	docker run --name go-postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:alpine3.15

createdb:
	docker exec -it go-postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it go-postgres dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --destination db/mock/store.go --package mockdb github.com/renatoviolin/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock