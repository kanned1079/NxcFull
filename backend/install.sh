#!/bin/bash

# 检查是否安装了 Golang
if ! command -v go &> /dev/null; then
    echo "Golang is not installed."
    exit 1
fi

# 获取 Golang 版本
go_version=$(go version)

# 提取版本号（从 "go version go1.x.x" 中提取 "1.x.x"）
version_number=$(echo $go_version | awk '{print $3}' | sed 's/go//')

# 检查版本是否大于或等于 1.22.2
required_version="1.22.2"

# 使用 `sort` 来比较版本号
if [[ $(echo -e "$version_number\n$required_version" | sort -V | head -n1) == "$required_version" ]]; then
    echo "Golang version is greater than or equal to 1.22.2. Version: $version_number"
else
    echo "Golang version is less than 1.22.2. Version: $version_number"
    exit 2
fi

# 检查是否传入了两个命令行参数
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <system> <architecture>"
  echo "Example: $0 linux amd64"
  exit 1
fi

# 获取命令行参数
SYSTEM=$1
ARCH=$2

# 定义子脚本路径
BUILD_SCRIPT="run_build.sh"

# 进入到 ./setup/envFreshInit 目录
echo "Entering directory: ./setup/envFreshInit"
cd ./setup/envFreshInit || exit

# 确保子脚本存在
if [ ! -f "$BUILD_SCRIPT" ]; then
  echo "Error: $BUILD_SCRIPT not found!"
  exit 1
fi

# 调用子脚本，传递参数
echo "Step 1: Building the binary for ${SYSTEM} ${ARCH}..."
bash "$BUILD_SCRIPT" "$SYSTEM" "$ARCH"
if [ $? -ne 0 ]; then
  echo "Error: Failed to build the binary in $BUILD_SCRIPT."
  exit 1
fi

# 返回上一级目录（即脚本最开始的目录）
cd - || exit

echo "Build completed successfully."