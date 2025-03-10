#!/bin/bash

# 检查是否传入了正确的参数
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <import|delete>"
  exit 1
fi

ACTION=$1
EXPORT_DIR="."  # 存放 tar.gz 文件的目录
ARCHITECTURE="linux-amd64"    # 默认架构信息

# 检查指定目录是否存在
if [ ! -d "$EXPORT_DIR" ]; then
  echo "Error: Directory $EXPORT_DIR does not exist."
  exit 1
fi

# 将镜像名称转换为小写
convert_to_docker_name() {
  local name=$1
  # 将镜像名称转换为小写
  echo "$name" | tr '[:upper:]' '[:lower:]'
}

# 从 tar.gz 文件名提取镜像名称，去掉后缀（如 -linux-amd64）
extract_image_name() {
  local tar_file=$1
  local base_name=$(basename "$tar_file" .tar.gz)  # 去掉 .tar.gz 后缀
  echo "$base_name" | sed 's/-linux-amd64//g'  # 去掉 -linux-amd64 后缀
}

# 导入所有 tar.gz 文件到 Docker
import_images() {
  echo "Starting import of Docker images from $EXPORT_DIR..."
  for tar in "$EXPORT_DIR"/*.tar.gz; do
    if [ -f "$tar" ]; then
      # 获取镜像名称，去掉 -linux-amd64 后缀并转换为小写
      IMAGE_NAME=$(extract_image_name "$tar")
      IMAGE_NAME=$(convert_to_docker_name "$IMAGE_NAME")
      # 添加架构信息
      FULL_IMAGE_NAME="${IMAGE_NAME}:${ARCHITECTURE}"
      echo "Loading Docker image from $tar as $FULL_IMAGE_NAME..."
      docker load < "$tar"
      if [ $? -eq 0 ]; then
        echo "Successfully loaded $tar as image $FULL_IMAGE_NAME"
      else
        echo "Failed to load $tar"
      fi
    else
      echo "No tar.gz files found in $EXPORT_DIR."
      break
    fi
  done
}

# 删除 Docker 中已导入的 tar.gz 镜像
delete_images() {
  echo "Starting deletion of Docker images imported from tar.gz files in $EXPORT_DIR..."
  for tar in "$EXPORT_DIR"/*.tar.gz; do
    if [ -f "$tar" ]; then
      # 获取镜像名称，去掉 -linux-amd64 后缀并转换为小写
      IMAGE_NAME=$(extract_image_name "$tar")
      IMAGE_NAME=$(convert_to_docker_name "$IMAGE_NAME")
      # 添加架构信息
      FULL_IMAGE_NAME="${IMAGE_NAME}:${ARCHITECTURE}"
      echo "Removing Docker image $FULL_IMAGE_NAME..."
      docker rmi "$FULL_IMAGE_NAME"
      if [ $? -eq 0 ]; then
        echo "Successfully removed image $FULL_IMAGE_NAME"
      else
        echo "Failed to remove image $FULL_IMAGE_NAME or image not found."
      fi
    else
      echo "No tar.gz files found to delete in $EXPORT_DIR."
      break
    fi
  done
}

# 根据传入的参数执行相应的操作
case "$ACTION" in
  import)
    import_images
    ;;
  delete)
    delete_images
    ;;
  *)
    echo "Invalid action. Usage: $0 <import|delete>"
    exit 1
    ;;
esac

echo "Operation $ACTION completed."