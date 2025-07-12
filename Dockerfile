# 1. 构建阶段（使用 Go 官方镜像）
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o app ./cmd/main.go

# 2. 运行阶段（使用精简基础镜像）
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
