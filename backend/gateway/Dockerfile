# 使用轻量级 Alpine 镜像
FROM alpine:latest

# 设置时区（可选）
RUN apk add --no-cache tzdata

# 设置工作目录
WORKDIR /app

# 将本地编译的可执行文件复制到镜像中
COPY ./build/gateway_linux_arm64 .

# 复制配置文件
COPY ./config/config.yaml /app/config/config.yaml

# 暴露服务端口
EXPOSE 8081

# 运行可执行文件
CMD ["./gateway_linux_arm64"]