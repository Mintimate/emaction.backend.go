# Build stage
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖
RUN apk add --no-cache gcc musl-dev

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=1 GOOS=linux go build -o emaction

# Final stage
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*

# 创建非 root 用户
RUN adduser -D -g '' emaction

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/emaction .
# 复制启动脚本
COPY entrypoint.sh .
# 创建配置目录
RUN mkdir -p ./config

# 给予启动脚本执行权限
RUN chmod +x ./entrypoint.sh

# 将所有文件的所有者更改为 appuser
RUN chown -R emaction:emaction /app

# 切换到非 root 用户
USER emaction

# 暴露端口
EXPOSE 8080

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

# 设置入口点和命令
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./emaction"]
