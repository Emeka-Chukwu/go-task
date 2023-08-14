DB_URL=postgresql://root:secret@localhost:5432/go-task?sslmode=disable
gql:
	printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	go mod tidy

init:
	go run github.com/99designs/gqlgen init
	go mod tidy

gen:
	go run github.com/99designs/gqlgen generate

docker:
	docker compose down && docker compose up --build

DB_URL=postgresql://root:secret@localhost:5432/go-task?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-task

dropdb:
	docker exec -it postgres dropdb go-task


migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...
run:
	go run main.go

mock:
	mockgen -package mockdb -destination internal/auths/repository/mock/store.go go-task/internal/auths/repository/postgres Authentication
	mockgen -package mockdb -destination internal/auths/usecase/mock/store.go go-task/internal/auths/usecase Authusecase 
	mockgen -package mockdb -destination internal/labels/repository/mock/store.go go-task/internal/labels/repository/postgres Label
	mockgen -package mockdb -destination internal/labels/usecase/mock/store.go go-task/internal/labels/usecase Labelusecase
	mockgen -package mockdb -destination internal/tasks/repository/mock/store.go go-task/internal/tasks/repository/postgres Task
	mockgen -package mockdb -destination internal/tasks/usecase/mock/store.go go-task/internal/tasks/usecase Taskusecase
	# mockgen -package mockdb -destination db/mock/store.go blog-api/db/sqlc Store

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema proto redis new_migratio



# make new_migration name=init