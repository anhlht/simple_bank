createdb:
	docker exec -it bank-account-db-1 createdb --username=postgres simple_bank

dropdb:
	docker exec -it bank-account-db-1 dropdb --username=postgres simple_bank

migratedbup:
	migrate -path ./db/migrations -database "postgresql://postgres:123%40123A@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedbdown:
	migrate -path ./db/migrations -database "postgresql://postgres:123%40123A@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb migratedbup migratedbdown sqlc test server