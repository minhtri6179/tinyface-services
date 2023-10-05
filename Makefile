DB_URL=postgresql://root:secret@localhost:5432/tinyface?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
redis:
	docker run --name redis -p 6379:6379 -d redis:6.0.9-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root tinyface

new_migration:
	migrate create -ext sql -dir scripts/migration/ -seq $(MESSAGE_NAME)

migrateup:
	migrate -path scripts/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path scripts/migration -database "${DB_URL}" -verbose down

dockerClear:
	docker volume rm $(docker volume ls -qf dangling=true) & docker rmi $(docker images -f "dangling=true" -q)

proto:
	protoc --proto_path=proto --go_out=proto/pb --go_opt=paths=source_relative \
	--go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
	proto/*.proto
	
.PHONY: postgres createdb migrateup migratedown new_migration dockerClear proto
 