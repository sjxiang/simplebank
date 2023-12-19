DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

# 创建 simple_bank 数据库
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

# 连接到 psql 客户端
psql:
	docker exec -it postgres psql -U root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank


mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0

# 打开 mysql 控制台
mysql_cli:
	docker exec -it mysql8 bash 
# mysql -uroot -psecret
# create database simple_bank;


migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
# statik -src=./doc/swagger -dest=./doc


# --{$plugin_name}_out={$out_dir} --{$plugin_name}_opt={$options}
# plugin_name
# 表示要执行的插件名；例如 "protoc-gen-go"，那么插件名为 "go"
# out_dir
# 表示插件生成代码的路径；如无特殊需求，一般指定为 "." 即可
# options
# 表示传递给插件的选项；通常会传递一些类似 "go module" 等信息


# 测试 gRPC
evans:
	evans --host localhost --port 9090 -r repl


redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

# 区分 target 和目录下文件名
.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis
