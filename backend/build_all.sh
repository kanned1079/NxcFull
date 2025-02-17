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

# 遍历每个微服务目录并执行构建命令
for SERVICE in "${SERVICES[@]}"; do
  if [ -d "$SERVICE" ]; then
    echo "Entering directory: $SERVICE"
    cd "$SERVICE" || exit
    echo "Running: bash run_build.sh $PLATFORM $ARCHITECTURE"
    bash run_build.sh "$PLATFORM" "$ARCHITECTURE"
    cd .. || exit
  else
    echo "Directory $SERVICE does not exist, skipping."
  fi
done

echo "Build process completed for all services."