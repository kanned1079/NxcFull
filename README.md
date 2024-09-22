## 这是一个测试的售卡系统

### 免责声明

1. 本软件是由个人独立开发，旨在为用户提供特定的功能和服务。由于个人开发资源和能力有限，软件可能存在一些不完善之处。
2. 对于因使用本软件而导致的任何设备损坏、数据丢失、利润损失、业务中断等后果，个人开发者不承担任何责任。
3. 个人开发者不保证软件中所包含的信息、功能和服务的准确性、完整性、及时性和可靠性。用户应自行核实和判断软件中的信息，并对自己的使用行为负责。

### 技术棧

- 前端
    - Vue3 + TypeScript + pinia + NaiveUI + md-editor + i18next
- 后端
    - golang + gin + gorm + gRPC + Mysql + Redis + RabbitMq

### 使用说明

```shell 
  # 安装运行依赖
  npm install
  # 运行开发者预览
  npm run dev -o
  ## 打包前端项目
  npm run build
```

#### 后端

```shell
  # 构建可执行文件
  go build -o server
  # 运行
  ./server
```