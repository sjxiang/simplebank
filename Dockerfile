# 构建
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# 运行
FROM alpine:3.18
WORKDIR /app
# 拷贝文件（可执行文件、配置文件、脚本 ...）
COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration

EXPOSE 8080 9090
ENTRYPOINT ["/app/main"]