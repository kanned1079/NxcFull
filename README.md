## 这是一个测试的售卡系统

### 免责声明

1. 本软件是由个人独立开发，旨在为用户提供特定的功能和服务。由于个人开发资源和能力有限，软件可能存在一些不完善之处。
2. 对于因使用本软件而导致的任何设备损坏、数据丢失、利润损失、业务中断等后果，个人开发者不承担任何责任。
3. 个人开发者不保证软件中所包含的信息、功能和服务的准确性、完整性、及时性和可靠性。用户应自行核实和判断软件中的信息，并对自己的使用行为负责。

### 構建前端程序

前端部分分为两部分 `主要用户和管理员交互界面` `单独的用于配置微服务配置的交互界面` 因此需要在不同目录执行依赖安装并打包

#### 安裝前端項目依賴

項目使用`npm`管理項目依賴

```shell
make install || npm install # 安裝依賴
```

運行構建

```shell
# 需要運行環境 node(20.16+)
make build || npm run build # 使用vite進行打包
```

將在`nxc-frontend/`中生成`dist`構建文件夾 接下來參照`nxc-frontend/README.md`進行安裝和運行

### 構建后端程序

#### 統一構建

這是推薦的構建方式

在項目的`backend/`目錄下運行下面命令以構建所有後端程序  
```shell
#需要在執行構建的系統中有安裝docker(27.1+) golang(1.22.2+) make 軟件包
./build_all linux amd64 #構建Linux平台amd64系統架構的可執行文件和Docker鏡像
```

構建完成後再`backend/export/`下可以看到所有微服務的Docker鏡像包`*.tar.gz`以及
- `README.md` 在服務器環境中的安裝流程指南
- `docker-compose.yaml` 在服務端使用docker compose運行的配置文件
- `manage_docker_images.sh` 方便在服務端快速進行Docker鏡像導入和刪除的腳本

至此已經打包好後所有的後端程序 接下來參照`backend/README.md`進行安裝和運行

#### 單獨構建

所有的微服务位置为 `./backend/*`  

在每一個微服務中执行 `./run_build <操作系統> <系統架構>` 构建可执行文件
```shell
# 构建可执行文件
./run_build.sh linux amd64 # 构建linux平台amd64架构
./run_build.sh build-darwin-arm64 # 构建osx平台arm64架构
./run_build.sh build-windows-amd64 # 构建windows平台amd64架构
```

構建結束後可以在當前目錄中的`build`目錄中找到對應架構的可執行文件`exec`和Docker鏡像 

