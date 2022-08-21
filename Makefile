createdb:
	docker exec -it bank-account-db-1 createdb --username=postgres simple_bank

dropdb:
	docker exec -it bank-account-db-1 dropdb --username=postgres simple_bank

migratedbup:
	migrate -path ./db/migrations -database "postgresql://postgres:123%40123A@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedbdown:
	migrate -path ./db/migrations -database "postgresql://postgres:123%40123A@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedbawsup:
	migrate -path ./db/migrations -database "postgresql://postgres:m69Aj2ymrh08nWOVCqSJ@simple-bank.cnrlmwqijr7a.ap-south-1.rds.amazonaws.com:5432/simple_bank" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mockdb:
	mockgen -package mockdb  -destination db/mock/store.go bank_account/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=bank_account \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9999 -r repl

.PHONY: createdb dropdb migratedbup migratedbdown sqlc test server mockdb migratedbawsup proto