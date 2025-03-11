#!/bin/bash

APPNAME="keyservices"
EXPORTPORT=50008

# 检查是否传入了两个命令行参数
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <system> <architecture>"
  echo "Example: $0 linux amd64"
  exit 1
fi

# 获取命令行参数
SYSTEM=$1
ARCH=$2

# 定义对应的目录和文件名
PLATFORM_DIR="${SYSTEM}-${ARCH}"
EXEC_FILE="./build/${PLATFORM_DIR}/exec"
DOCKER_TAR="./build/${APPNAME}-${PLATFORM_DIR}.tar.gz"

# 删除已存在的 docker 包和可执行文件目录（如果存在）
if [ -f "$DOCKER_TAR" ]; then
  echo "Step 0: Removing existing Docker tar.gz file..."
  rm -f "$DOCKER_TAR"
  if [ $? -ne 0 ]; then
    echo "Error: Failed to remove the existing Docker tar.gz file."
    exit 1
  fi
fi

if [ -d "./build/${PLATFORM_DIR}" ]; then
  echo "Step 0: Removing existing ${PLATFORM_DIR} directory..."
  rm -rf "./build/${PLATFORM_DIR}"
  if [ $? -ne 0 ]; then
    echo "Error: Failed to remove the existing directory."
    exit 1
  fi
fi

# 定义编译命令
BUILD_CMD="make build-${SYSTEM}-${ARCH}"

# 第一步：构建二进制文件
echo "Step 1: Building the binary for ${SYSTEM} ${ARCH}..."
$BUILD_CMD
if [ $? -ne 0 ]; then
  echo "Error: Failed to build the binary."
  exit 1
fi

# 确保二进制文件生成
if [ ! -f "$EXEC_FILE" ]; then
  echo "Error: Binary file 'exec' not found!"
  exit 1
fi

echo "Exec file name: exec"

# 第二步：构建 Docker 镜像
echo "Step 2: Building the Docker image with tag ${APPNAME}:${SYSTEM}-${ARCH}..."

# 使用 docker buildx 来支持跨架构构建
docker buildx build --platform linux/${ARCH} --build-arg PLATFORM_DIR=${PLATFORM_DIR} --build-arg PORT=${EXPORTPORT} -t ${APPNAME}:${SYSTEM}-${ARCH} .

if [ $? -ne 0 ]; then
  echo "Error: Failed to build the Docker image."
  exit 1
fi

# 第三步：将 Docker 镜像保存为 tar.gz 文件
echo "Step 3: Saving the Docker image as tar.gz..."
docker save ${APPNAME}:${SYSTEM}-${ARCH} | gzip > "$DOCKER_TAR"
if [ $? -ne 0 ]; then
  echo "Error: Failed to save the Docker image as tar.gz."
  exit 1
fi

# 成功
echo "Successfully built and packaged the Docker image as $DOCKER_TAR"

# 解压并加载 Docker 镜像（可选步骤）
# docker load < "$DOCKER_TAR"