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
    # 接下来操作
    go run quickStartServices.go
else
    echo "Golang version is less than 1.22.2. Version: $version_number"
    exit 2
fi