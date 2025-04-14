postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=mydatabase -p 5432:5432 -d postgres

createdb:
	docker exec -it postgres12 createdb --username=user --owner=user simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://user:admin@postgres:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://user:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go version
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test




