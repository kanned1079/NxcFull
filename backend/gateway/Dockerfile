# 使用轻量级的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 定义构建时参数
#ARG BINARY_NAME=services_platform_arch
ARG PLATFORM_DIR=linux-arm64

RUN echo "dir name is: ${PLATFORM_DIR}"

# 复制编译好的二进制文件到容器
COPY ./build/${PLATFORM_DIR}/exec /app/

# 复制配置文件到容器中
#COPY ./config/config.xml /app/config/config.xml

RUN ls -l .

# 暴露端口
#EXPOSE ${PORT}

# 启动应用程序
#CMD ["./${BINARY_NAME}"]
CMD ["./exec"]
