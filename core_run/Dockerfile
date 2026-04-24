FROM golang:1.23-alpine AS builder
 
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 复制整个项目（包含 go.work 和所有模块）
COPY . ./

# 进入 core 目录下载依赖
WORKDIR /app/core
RUN go mod download

# 构建 core_run（它依赖 core 模块）
WORKDIR /app
RUN go build -o main ./core_run

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
