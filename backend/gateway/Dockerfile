# 使用 Go 官方镜像
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件到工作目录中
COPY ../../go.mod ../../go.sum ./

# 下载依赖
RUN go mod download

# 复制项目的所有文件到工作目录中
COPY . .

# 将配置文件也复制到容器中
COPY config/config.yaml /app/config/config.yaml

# 编译应用程序
RUN go build -o gateway ./command/server/main.go

# 暴露端口（假设服务在8080端口上运行）
EXPOSE 50006

# 启动微服务
CMD ["./gateway"]