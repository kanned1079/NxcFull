## 这是一个测试的售卡系统

### 免责声明

1. 本软件是由个人独立开发，旨在为用户提供特定的功能和服务。由于个人开发资源和能力有限，软件可能存在一些不完善之处。
2. 对于因使用本软件而导致的任何设备损坏、数据丢失、利润损失、业务中断等后果，个人开发者不承担任何责任。
3. 个人开发者不保证软件中所包含的信息、功能和服务的准确性、完整性、及时性和可靠性。用户应自行核实和判断软件中的信息，并对自己的使用行为负责。

### 系统主要功能

#### 前端功能

- 使用`Vue3`和`Naive-UI`构建，拥有美观的web界面的售卡平台，提供api接口可将服务绑定到各类第三方开发的应用软件、定制化系统、硬件设备等，全局为响应式布局 一套设计适用于移动端、平板端、桌面端，减小开发难度，支持国际化，内置多种不同的语言以方便跨境使用，支持深色模式，自动根据用户浏览器适配浅色、深色主题，在此基础上还提供不同的主题配色，设计有三种欢迎界面，用于介绍产品优势、功能、定价
- **用户侧功能** `多语言切换` `UI深浅色适配` `登录` `邮箱注册` `第三方单点登录/注册` `邮箱验证码认证` `查看通知` `阅读说明文档` `邀请用户` `更改密码` `是否允许发送邮件通知` `提交新工单` `工单中的即时通信实时与管理员联系` `查询用户购买的密钥信息` `换绑定序列号设备` `购买新订阅计划` `优惠券验证` `查看订单我的订单`
- **管理侧功能：** `登录` `设定管理员/员工` `收入查看` `7天内收入图标` `7天内访问人数图表` `服务端监控` `服务器基本信息查看` `队列服务监控` `配置系统的各项功能 如网站名、主题配色、邮箱验证、通知配置、邮箱通知地址配置、Brak即时通知api地址配置、专有app设置` `配置支付接口` `设置介绍页面数据` `权限组管理` `订阅计划管理` `优惠券管理` `用户管理` `公告管理` `工单管理`  `知识库管理`

#### 后端功能

- 使用微服务架构配合容器技术，将一整个后端程序拆分为多个单独的服务，前端所有请求连接到`gateway服务` 再通过RPC来调用对应的二级模块的功能 `userServices` `couponServices` `noticeServices` `ocumentServices` `orderServices` `subscribeServices` `mailServices` `systemServices` `expiryService`等微服务，微服务之间带有心跳和断线重连机制，单个服务掉线不影响其他服务，所有微服务的配置信息都来自于`Etcd` 其中的数据库、Redis、MQ连接信息都将由单独的一个程序推送
- 数据库使用`MySQL` 和 `Redis` 加入noSQL的Redis来作为MySQL的缓存 存放最近的用户信息和需要频繁读取的数据

### 技术棧

- 前端
    - Vite + Vue3 + TypeScript + pinia + NaiveUI + md-editor + i18next
- 后端
    - golang + gin + gorm + gRPC + Mysql + Redis + Etcd + RabbitMq

### 使用说明

- 前端部分分为两部分 `主要用户和管理员交互界面` `单独的用于配置微服务配置的交互界面` 因此需要在不同目录执行依赖安装并打包
```shell 
# 安装运行依赖
npm install
# 运行开发者预览
npm run dev -o
# 打包前端项目
npm run build
```

#### 后端

- 所有的微服务位置为 `./backend/` 
- 首先需要执行 `make` 构建可执行文件
```shell
# 构建可执行文件
# 支持架构 arm64 amd64 如果没有定义平台架构 将构架所有 下面是一些示例
make build-linux-amd64 # 构建linux平台amd64架构的可执行文件
make build-darwin-arm64 # 构建osx平台arm64架构的可执行文件
make build-windows-amd64 # 构建windows平台amd64架构的可执行文件
# 构建结束后可以直接运行
./build/<微服务名>-linux-amd64
```

- 打包服务到docker镜像文件
```shell
# 打包为镜像
docker build -t <用户名>/<镜像名>@<标签名>
# 使用docker run来运行
``` 