# Build stage
# 构建阶段
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
# 运行阶段
FROM alpine:3.18
WORKDIR /app
# 可执行文件
COPY --from=builder /app/main .
# 配置文件
COPY app.env .
# 
COPY start.sh .
COPY wait-for.sh .
# 迁移
COPY db/migration ./db/migration

EXPOSE 8080 9090
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
