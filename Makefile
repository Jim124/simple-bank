postgres:
		docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d -p 5432:5432 postgres:17-alpine
createdb:
		docker exec -it postgres17 createdb --username=root --owner=root simple_bank
dropdb:
		docker exec -it postgres17 dropdb simple_bank
migrateup:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
		sqlc generate
test:
		go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
