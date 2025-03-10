#!/bin/bash

# 检查是否传入了正确的参数
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <platform> <architecture>"
  exit 1
fi

# 获取平台名和架构
PLATFORM=$1
ARCHITECTURE=$2

# 微服务目录列表
SERVICES=("couponServices" "documentServices" "groupServices" "mailServices" "noticeServices" "subscribeServices"
          "systemServices" "userServices" "orderServices" "orderHandleServices" "keyServices"
          "ticketHandleServices" "logServices" "gateway")

# 创建一个 export 目录，用于存放打包后的 Docker 镜像
EXPORT_DIR="./export"
mkdir -p "$EXPORT_DIR"
if [ $? -ne 0 ]; then
  echo "Error: Failed to create export directory."
  exit 1
fi

# 遍历每个微服务目录并执行构建命令
for SERVICE in "${SERVICES[@]}"; do
  if [ -d "$SERVICE" ]; then
    echo "Entering directory: $SERVICE"
    cd "$SERVICE" || exit
    echo "Running: bash run_build.sh $PLATFORM $ARCHITECTURE"
    bash run_build.sh "$PLATFORM" "$ARCHITECTURE"

    # 查找并将构建的 .tar.gz 文件剪切到 export 目录
    DOCKER_TAR="./build/${SERVICE}-${PLATFORM}-${ARCHITECTURE}.tar.gz"
    if [ -f "$DOCKER_TAR" ]; then
      echo "Copying $DOCKER_TAR to export directory."
      mv "$DOCKER_TAR" "../$EXPORT_DIR/"
      if [ $? -ne 0 ]; then
        echo "Error: Failed to move $DOCKER_TAR to export directory."
        exit 1
      fi
    else
      echo "Docker tar.gz file for $SERVICE not found, skipping."
    fi

    cd .. || exit
  else
    echo "Directory $SERVICE does not exist, skipping."
  fi
done

# 复制compose
cp ./docker-compose.yaml ./export
cp ./manage_docker_images.sh ./export
cp ./README.md ./export

echo "Build process completed for all services. Docker tar.gz files are available in the $EXPORT_DIR directory."
