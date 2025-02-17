#!/bin/bash

# 检查是否传入了两个命令行参数
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <system> <architecture>"
  echo "Example: $0 linux amd64"
  exit 1
fi

# 获取命令行参数
SYSTEM=$1
ARCH=$2

# 设置 GOOS 和 GOARCH 环境变量
export GOOS=${SYSTEM}
export GOARCH=${ARCH}

# 定义对应的目录和文件名
PLATFORM_DIR="${SYSTEM}-${ARCH}"
EXEC_FILE="./build/${PLATFORM_DIR}/exec"

# 定义 Makefile 路径
MAKEFILE="./Makefile"

# 打印当前路径
echo "Current directory: $(pwd)"

# 构建二进制文件
echo "Building the binary for ${SYSTEM} ${ARCH}..."
make -f ${MAKEFILE} build-${SYSTEM}-${ARCH}
if [ $? -ne 0 ]; then
  echo "Error: Failed to build the binary."
  exit 1
fi

# 查看是否成功生成二进制文件
echo "Checking if binary exists at: $EXEC_FILE"
if [ ! -f "$EXEC_FILE" ]; then
  echo "Error: Binary file 'exec' not found!"
  exit 1
fi

# 赋予执行权限
echo "Granting execute permissions to the binary..."
chmod +x "$EXEC_FILE"
if [ $? -ne 0 ]; then
  echo "Error: Failed to grant execute permissions to the binary."
  exit 1
fi

# 执行可执行文件
echo "Executing the binary..."
"$EXEC_FILE"
if [ $? -ne 0 ]; then
  echo "Error: Failed to execute the binary."
  exit 1
fi

echo "Success: Binary executed successfully."