postgres:
	docker run --name simple-bank-postgresql -p 5436:5432 -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it simple-bank-postgresql createdb --username=postgres --owner=postgres db_simple_bank

dropdb:
	docker exec -it simple-bank-postgresql dropdb --username=postgres db_simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5436/db_simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5436/db_simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc