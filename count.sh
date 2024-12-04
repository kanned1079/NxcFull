#!/bin/bash

# 检查是否提供了文件后缀
if [ -z "$1" ]; then
    echo "Usage: $0 <file_extension>"
    exit 1
fi

extension=$1

# 初始化总计数
total_count=0

# 遍历当前目录及其子目录
for dir in $(find . -type d); do
    # 统计当前目录下匹配的文件数量
    count=$(find "$dir" -maxdepth 1 -type f -name "*.$extension" | wc -l)
    echo "Directory: $dir - $count .$extension files"
    total_count=$((total_count + count))
done

echo "Total .$extension files: $total_count"
