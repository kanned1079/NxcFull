# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
# 二进制文件名
BINARY_NAME=noticeServices
MAIN_FILE=./command/server/main.go
# 输出二进制文件的目录
BUILD_DIR=./build

# 如果make后没有传递任何参数 直接运行build-all
all: build-all

# 确保构建目录存在
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# 为Linux x86_64架构编译可执行文件
build-linux-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_amd64 $(MAIN_FILE)

# 为Linux arm64架构编译可执行文件
build-linux-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_arm64 $(MAIN_FILE)

# 为macOS amd64架构编译可执行文件
build-darwin-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_amd64 $(MAIN_FILE)

# 为macOS arm64架构编译可执行文件
build-darwin-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_arm64 $(MAIN_FILE)

# 为Windows amd64架构编译可执行文件
build-windows-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_windows_amd64.exe $(MAIN_FILE)

# 为Windows arm64架构编译可执行文件
build-windows-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_windows_arm64.exe $(MAIN_FILE)

# 为FreeBSD amd64架构编译可执行文件
build-freebsd-amd64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_freebsd_amd64 $(MAIN_FILE)

# 为FreeBSD arm64架构编译可执行文件
build-freebsd-arm64: $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_freebsd_arm64 $(MAIN_FILE)

# 同时编译x86_64和arm64
build-all: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-windows-arm64 build-freebsd-amd64 build-freebsd-arm64

# 编译完成后运行项目
run:
	./$(BUILD_DIR)/$(BINARY_NAME)

# 运行所有单元测试
test:
	$(GOTEST) -v ./...

# Clean the build
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	$(GOGET) -u ./...

# Lint the project
lint:
	golangci-lint run

# Format the code
fmt:
	$(GOCMD) fmt ./...

# Vendor dependencies
vendor:
	$(GOCMD) mod vendor

# Tidy go.mod and go.sum
tidy:
	$(GOCMD) mod tidy

.PHONY: all build build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-windows-arm64 build-freebsd-amd64 build-freebsd-arm64 build-all clean run test deps lint fmt vendor tidy