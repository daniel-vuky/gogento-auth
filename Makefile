DB_URL=postgresql://root:secret@localhost:5432/gogento?sslmode=disable&x-migrations-table=schema_migrations_auth
init_postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:16-alpine
start_postgres:
	docker start postgres16
create_db:
	docker exec -it postgres16 createdb --username=root --owner=root gogento
migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/daniel-vuky/gogento-auth/db/sqlc Store
test:
	go test -v -cover -short ./...

.PHONY: init_postgres start_postgres stop_postgres create_db drop_db migrate_up migrate_down sqlc mock test
