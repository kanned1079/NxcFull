# NxcNetworks 接口文档

## 状态码说明

### 通用状态码

  - 200 `http.StatusOK` 一般用于表示请求成功 数据获取、更新、新建、下单结果成功的响应标志
  - 500 `http.StatusInternalServerError` 服务器内部错误 一般在json的序列化/反序列化失败、数据库、消息队列、Redis连接、查询、设置、文件操作错误等
  - 404 `http.StatusNotFound` 用于用户登录没有找到对应的用户、没有查到对应的订单、订阅计划、工单等
  - 401 `http.StatusUnauthorized` 用户没有认证 请先提供token以继续请求

### 支付状态码

  - 202 *WaitUserScanQrCode* `http.StatusAccepted` 预生成订单成功后的等待用户扫码
  - 102 *ScannedQrCodeWaitUserPay* `http.StatusProcessing` 用户扫描二维码后生成订单等待支付
  - 408 *TradeClosed* `http.StatusRequestTimeout` 未付款交易超时关闭或支付完成后全额退款
  - 200 *TradeSuccess* `http.StatusAccepted` 订单成功
  - 410 *TradeFinished* `http.StatusGone` 订单已经成功并且不可以退款
    - 200 *UpdateUserBalanceSuccess* `http.StatusOK` 更新用户和邀请人的信息都成功
    - 206 *UpdateUserBalanceAccept* `http.StatusPartialContent` 更新用户的成功 但是邀请人可能没有成功 请等待后续同步
    - 500 *UpdateUserBalanceFailed* `http.StatusInternalServerError` 都更新失败 建议客户后面联系客服处理


## 携带Token

> [!NOTE]
>
> 除了公用方法请求时不需要在`Header`中携带Token以外 其他所有请求（包括`websocket`）都需要携带Token以进行身份验证

### 封装Axios

- 封装`Axios`在请求拦截器中设置Token值 在每一次请求中携带该Token

  ```ts
  instance.interceptors.request.use(config => {
      const token = sessionStorage.getItem('token');
      console.log(token)
      if (token) {
          config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
  }, error => {
      return Promise.reject(error);
  });
  ```

- 设置响应拦截器 如果服务器中间件返回401(http.StatusUnauthed) 处理登出操作

  ```ts
  instance.interceptors.response.use(response => {
      return response;
  }, error => {
      if (error.response && error.response.status === 401) {
          let userInfoStore = useUserInfoStore();
          // token 过期
          console.error('Token expired or invalid. Please log in again.');
          // 登出操作
          userInfoStore.logout()
  
      }
      return Promise.reject(error);
  });
  ```
### 封装Websocket实例

- websocket请求 在创建ws请求实例的时候直接将token放置在`params`参数中

  ```ts
  const createWebSocket = () => {
    if (socket) socket.close(); // 先关闭已有连接
    // 创建新的ws连接实例 并传入需要的参数和Token
    socket = new WebSocket(`ws://localhost:8081/ws/user/v1/chat?token=${token.value}&user_id=${paramsData.value.userId}&ticket_id=${paramsData.value.ticketId}`);
  	// 打开链接
    socket.onopen = () => {
      console.log("WebSocket connection established");
    };
  	// 如果有新的消息 在这里完善具体业务流程
    socket.onmessage = (event: MessageEvent) => {
  		// ...
    };
  	// 如果ws连接以外关闭 执行重新连接 入如果还是连接失败 延迟5s再进行下一次尝试
    socket.onclose = () => {
      if (paramsData.value.status !== 204) {
        console.log("WebSocket connection closed. Reconnecting...");
        setTimeout(createWebSocket, 5000); // 每5秒尝试重新连接
      }
    };
  	// 在ws连接遇到错误时 处理错误
    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  };
  ```

  

## 通用

###  **`/api/admin/v1/login`**

#### **请求方式** `POST`

  - **功能** 用户或管理员登录

  - **是否需要携带Token** `否`
  
  - **请求参数**
  
    - `email (string)` 用户邮箱地址
    
    - `password (string)` 密码
    
    - `two_fa_code (string)` 两步验证码
    
    - `role (string)` 用户角色
    
      - 请求示例
    
        ```json
         {
           email: "kanned1079@gmail.com", 
           password: "MDAwMA==", 
           two_fa_code: "", 
           role: "user"
         }
        ```
  
  - **响应参数**
  
    - `code (int32)` 响应状态码
    
    - `isAuthed (bool)` 是否验证通过
    
    - `msg (string)` 提示消息
    
    - `token (string)` 用户Token
    
    - `user_data (Object{email: string, id: number, balance: number, ....})` 用户基本数据
    
      - 响应示例
    
        ```json
         {
             "code": 200,
             "isAuthed": true,
             "msg": "登录成功",
             "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
             "user_data": {
                 "balance": 22.770002,
                 "email": "kanned1079@gmail.com",
                 "group_id": 0,
                 "id": 2,
                 "invite_code": "z0dap3z1jkvuh02ja2",
                 "invite_user_id": "3",
                 "is_admin": false,
                 "is_staff": true,
                 "is_two_fa": false,
                 "status": true,
             }
         }
        ```

  

  ## Admin

###  **`/api/admin/v1/setting`**

#### 请求方式 `GET`

- **功能** 获取所有系统配置项目的键值 根据类别进行分类

- **是否需要携带Token** `YES`

- **请求参数**
  
  - 无
  
- **响应参数**
  
  - `code (int32)` 响应状态码
  
  - `msg (string)` 提示消息
  
  - `settings (Object{frontend: Object{...}, site: Object:{...}})` 系统配置项对象
  
    - 响应示例
  
      ```json
      {
         "code":200,
         "msg":"Query ok",
         "settings":{
            "frontend":{
               "frontend_background_url":"32rdwefdw1111",
               "frontend_theme":"defaultDay",
              ...
            },
            "my_app":{
               "android_download":"andver",
               "osx_download":"osxapp",
               "win_download":"winapp"
            },
           ...
         }
      }
      ```

#### 请求方式 `PUT`

- **功能** 修改一个配置项目的值 如果配置不存在于数据库中则创建

- **是否需要携带Token** `YES`

- **请求参数**

  - `category (string)` 配置参数的类别

  - `key (string)` 配置参数的键

  - `value (string | bool | number | null)` 配置参数具体的值 可以是`bool/number/string`

    - 请求示例

      ```json	
      {
        category: "site",
        key: "subscribe_url",
        value: "url1122"
      }
      ```

- **响应参数**

  - `code (int32)` 响应状态码

  - `msg` 提示信息

    - 响应示例

      ```json
      {
        "code":200,
        "msg":"配置项更新成功"
      }
      ```

###  **`/api/admin/v1/mail/test`**

#### 请求方式 `POST`

- **功能** 测试邮件配置 如果成功则返回使用的配置信息

- **是否需要携带Token** `YES`

- **请求参数** `Params参数`

  - `email (string)` 要发送到的邮箱地址

    - 请求示例

      ```json
      {
        email: "admin@ikanned.com"
      }
      ```

- **响应参数**

  - `code (int32)` 响应状态码

  - `info (Object{from: string, host: string, port: int32, ...})` 如果正确 则返回正确的邮件配置信息

  - `msg (string)` 提示信息

    - 响应示例

      ```json	
      {
          "code": 200,
          "info": {
              "from": "no-reply@ikanned.com",
              "port": 465,
              "use_ssl": true,
          },
          "msg": "发送测试邮件成功"
      }
      ```

      


### **`/api/admin/v1/notice`**

#### 请求方式 `GTE`

- **功能** 获取所有的通知信息 注意请求params参数中需要携带`is_user`参数

- **需要携带Token** `YES`

- **请求参数**
  - `page (number)` 分页页面号 *该参数用户请求时不可用*
  
  - `size (number)` 分页每一个页面条数 *该参数用户请求时不可用*
  
  - `is_user (boolean)` 是否为用户的请求
  
    - 请求示例
  
      ```json
      {
        page: 1,
      	size: 10,
      	is_user: false,
      }
      ```
  
- **响应参数**

  - `code (int32)` 响应状态码

  - `msg (string) ` 提示信息

  - `notices ([Object{id: number, content: string, img_url: string}, Object{...}, ...])` 具体的通知信息 对象数组 

    - 响应示例

      ```json
      {
          "code": 200,
          "msg": "获取成功",
          "notices": [
              {
                  "content": "NxcNetworks预祝全體用戶2025年新年快樂。",
                  "created_at": "2024-11-01T23:54:00+08:00",
                  "deleted_at": null,
                  "id": 18,
                  "img_url": "https://ikanned.com:2444/d/Upload/nxc-test/notice-2025.jpg",
                  "show": true,
                  "tags": "优惠券",
                  "title": "謹賀新年 2025",
                  "updated_at": "2024-11-12T19:10:44+08:00"
              },
              {
                  "content": "优惠券码：HappyWinter2024\n此优惠码使用截至日期2025-01-12，祝大家2025寒假快乐！",
                  "created_at": "2024-11-01T23:49:47+08:00",
                  "deleted_at": null,
                  "id": 17,
                  "img_url": "https://ikanned.com:2444/d/Upload/nxc-test/notice-winter.jpg",
                  "show": true,
                  "tags": "优惠券",
                  "title": "2024冬季优惠 40%off",
                  "updated_at": "2024-11-01T23:49:50+08:00"
              },
          ],
          "page_count": 2
      }
      ```


#### 请求方式 `POST`

- **功能** 添加一条新的通知

- **需要携带Token** `YES`

- **请求参数** `from-data`

  - `title (string)` 要新增公告的标题 它将作为在用户首页置顶的轮播图中的标题

  - `content (string)` 该公告的主要内容

  - `tags (string)` 公告左上角的标题

  - `img_url (string)` 如果公告有背景图片 在这里填写图片的直链

    - 请求示例

      ```json	
      {
        "title": "测试标题",
        "content": "这里是公告内容",
        "tags": "通知",
        "img_url": "http://xxx.xxxx.x.com/bg.img"
      }
      ```

- **响应参数**

  - `code (int32)` 响应状态码

  - `msg (string) ` 提示信息

    - 响应示例	

      ```json
      {
        "code": 200,
        "msg": "新建通知成功"
      }
      ```

#### x请求方式 `PUT`

#### **`/api/admin/v1/notice/ststus`**

##### 请求方式 `PUT`

#### 请求方式 `DELETE`

### **`/api/admin/v1/plan`**

#### **请求方式 `GET`**

#### **请求方式 `POST`**

#### **请求方式 `PUT`**

#### **请求方式 `DELETE`**

#### **`/api/admin/v1/plan/kv`**

##### **请求方式 `GET`**

#### **`/api/admin/v1/plan/sale`**

##### **请求方式 `PUT`**

#### **`/api/admin/v1/plan/renew`**

##### **请求方式 `PUT`**






```go	
adminAuthorized.GET("/plan", handler.HandleGetAllPlans)          // 获取所有的订阅
		adminAuthorized.POST("/plan", handler.HandleAddNewPlan)          // 添加新的订阅
		adminAuthorized.PUT("/plan", handler.HandleUpdatePlan)           // 修改订阅的细节
		adminAuthorized.DELETE("/plan", handler.HandleDeletePlan)        // 删除指定的订阅
		adminAuthorized.GET("/plan/kv", handler.HandleGetAllPlanKeyName) // 获取订阅计划的键值对
		adminAuthorized.PUT("plan/sale", handler.HandleUpdatePlanSale)   // 更新是否启用销售
		adminAuthorized.PUT("plan/renew", handler.HandleUpdatePlanRenew) // 更新是否允许续费
```





