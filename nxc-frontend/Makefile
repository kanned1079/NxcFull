# 设置项目名称和目录
PROJECT_NAME := vue-vite-project
BUILD_DIR := dist
SRC_DIR := src

# Node.js 路径配置
NODE := node
NPM := npm
VITE := $(NODE) node_modules/.bin/vite

# 默认目标：开发环境
.PHONY: dev
dev:
	@echo "Starting development server..."
	$(VITE) serve

# 构建项目
.PHONY: build
build:
	@echo "Building project for production..."
	$(NPM) run build

# 清除构建目录
.PHONY: clean
clean:
	@echo "Cleaning the build directory..."
	rm -rf $(BUILD_DIR)

# 安装依赖
.PHONY: install
install:
	@echo "Installing dependencies..."
	$(NPM) install

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	$(NPM) run test

# 启动生产环境
.PHONY: start
start:
	@echo "Starting production server..."
	$(VITE) preview

# 更新依赖
.PHONY: update
update:
	@echo "Updating dependencies..."
	$(NPM) update