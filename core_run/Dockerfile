FROM golang:1.23-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 复制整个项目（包含 go.work 和所有模块）
COPY . ./

# 进入 core 目录下载依赖
WORKDIR /app/core
RUN go mod download

# 进入 core_run 目录下载依赖并构建
WORKDIR /app/core_run
RUN go mod download
RUN go build -o main ./src

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/core_run/main .

EXPOSE 8080

CMD ["./main"]