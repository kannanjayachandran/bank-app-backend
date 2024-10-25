postgres:
	docker run --name bank_app_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpassword -d postgres:17-alpine

createdb:
	docker exec -it bank_app_db createdb --username=root --owner=root simpleBank

dropdb:
	docker exec -it bank_app_db dropdb simpleBank

migrateup:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simpleBank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simpleBank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc