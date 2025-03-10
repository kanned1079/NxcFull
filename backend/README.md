## 後端安裝流程

此處的安裝演示適用於在Linux環境下 如果使用其他操作系統 如FreeBSD(Unix) Windows 請使用其他的方法 因此在構建的時候請選擇合適的操作系統類型和架構

### 構建和打包
運行根目錄下的`build_all.sh`來構建所有微服務
```shell
chmod +x build_all.sh #赋予可执行权限
./build_all.sh linux amd64 #打包linux系统amd64架构
./build_all.sh darwin arm64 #打包macOS系统arm64架构
```
#### 支持的構建項目
- linux
    - amd64 `./build_all.sh linux amd64`
    - arm64 `./build_all.sh linux arm64`
- OSX(macOS X)
    - amd64 `./build_all.sh darwin amd64`
    - arm64 `./build_all.sh darwin arm64`
- FreeBSD
    - amd64 `./build_all.sh freebsd amd64`
    - arm64 `./build_all.sh freebsd arm64`

全部完成後可以在當前目錄下的`export/`文件夾中找到所有的打包後的為服務 以及用於導入到Docker的腳本和使用`docker-compose`運行的配置文件
> 如果系統為非Linux環境 在每一個為服務的 `build` 目錄下也可找到不同的可執行文件

### 上傳到服務器
建議打包為tar.gz後在服務器中進行解壓 可以使用不同的ssh連接軟件進行文件上傳
```shell
tar -zcvf backend.export.tar.gz ./export
```
在要運行的服務器上解壓
```shell
tar -zxvf backend.export.tar.gz
```

### 導入所有鏡像
#### 使用 `manage_docker_images.sh` 腳本
使用方法
```shell
chmod +x manage_docker_images.sh <import | delete>
```
- **import參數** 將所有該目錄下的docker鏡像進行導入
- **delete參數** 刪除所有導入的鏡像

### 運行所有服務
服務器上需要有`docker-compose`安裝 較為新版本的Linux發行版在安裝docker時默認已經安裝
#### 初次啟動所有服務
```shell
TAG=linux-amd64 docker-compose up -d #單獨安裝的docker-compose
TAG=linux-amd64 docker compose up -d #較爲新版本
```
#### 停止所有服務
```shell
docker compose stop
```
#### 再次開啟服務
```shell
docker compose start
```