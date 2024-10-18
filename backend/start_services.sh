#!/bin/bash

# 存储进程 ID 的文件
PID_FILE="services_pids.txt"

# 启动服务函数
start_services() {
  # 定义服务目录的数组
  services=(
    "couponServices"
    "documentServices"
    "groupServices"
    "mailServices"
    "noticeServices"
    "orderServices"
    "subscribeServices"
    "systemServices"
    "userServices"
  )

  # 清空 PID 文件
  > "$PID_FILE"

  # 遍历启动每个服务
  for service in "${services[@]}"; do
    echo "Starting $service..."
    (cd "$service" && go run command/server/main.go &) 
    # 保存每个服务的进程 ID
    echo "$service: $!" >> "$PID_FILE"
  done

  # 等待 5 秒启动 gateway 服务
  echo "Waiting 5 seconds to start gateway service..."
  sleep 5

  # 启动 gateway 服务
  echo "Starting gateway service..."
  (cd "gateway" && go run command/server/main.go &)
  # 保存 gateway 的进程 ID
  echo "gateway: $!" >> "$PID_FILE"

  echo "All services started."
}

# 停止服务函数
stop_services() {
  if [ ! -f "$PID_FILE" ]; then
    echo "No running services found."
    exit 1
  fi

  # 遍历 PID 文件并停止进程
  while IFS= read -r line; do
    service=$(echo "$line" | cut -d':' -f1)
    pid=$(echo "$line" | cut -d':' -f2)

    if kill -0 "$pid" 2>/dev/null; then
      echo "Stopping $service with PID $pid..."
      kill "$pid"
    else
      echo "$service is not running."
    fi
  done < "$PID_FILE"

  # 删除 PID 文件
  rm -f "$PID_FILE"

  echo "All services stopped."
}

# 检查参数
if [ "$1" == "stop" ]; then
  stop_services
else
  start_services
fi
