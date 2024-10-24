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
    "subscribeServices"
    "systemServices"
    "userServices"
  )

  # 检查是否已有服务在运行
  if [ -f "$PID_FILE" ]; then
    echo "Services are already running. Stop them first."
    exit 1
  fi

  # 清空 PID 文件
  > "$PID_FILE"

  # 遍历启动每个服务
  for service in "${services[@]}"; do
    echo "Starting $service..."
    (cd "$service" && nohup go run command/server/main.go > /dev/null 2>&1 &)
    sleep 1  # 等待进程启动

    # 使用 ps 获取服务的 PID
    pid=$(ps -ef | grep "[g]o run command/server/main.go" | awk '{print $2}')
    if [ -n "$pid" ]; then
      echo "$service: $pid" >> "$PID_FILE"
    else
      echo "Failed to find PID for $service"
    fi
  done

  # 等待 10 秒启动 gateway 服务
  echo "Waiting 10 seconds to start gateway service..."
  sleep 10

  # 启动 gateway 服务
  echo "Starting gateway service..."
  (cd "gateway" && nohup go run command/server/main.go > /dev/null 2>&1 &)
  sleep 1  # 等待进程启动

  # 使用 ps 获取 gateway 的 PID
  pid=$(ps -ef | grep "[g]o run command/server/main.go" | awk '{print $2}')
  if [ -n "$pid" ]; then
    echo "gateway: $pid" >> "$PID_FILE"
  else
    echo "Failed to find PID for gateway"
  fi

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
      kill -TERM "$pid"
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