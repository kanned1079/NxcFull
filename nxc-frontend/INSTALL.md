## 前端安裝流程

首先將前端打包後的`dist`文件夾傳輸到服務器上

### 安裝Nginx

Nginx是一個高性能的Web服務器 當然使用Caddy作為服務器也是一個不錯的選擇 這裡以Nginx為例

```shell
sudo dnf install nginx -y #redhat系
sudo apt install nginx -y #debian系
# 不同的操作系統安裝方式可能有些許差別
```

### 開始安裝

#### 新建一個文件夾用於存放靜態文件

```shell
sudo mkdir -p /var/www/nxc 
sudo cp -r ~/dist . #複製到此處
```

#### 新建一個Nginx配置文件

假設服務運行在27003端口上 下面給出一個例子

```
server {
    listen 27003 ssl http2;
    server_name ikanned.com www.ikanned.com;

    # SSL 配置
    ssl_certificate     /etc/nginx/ikanned.com_nginx/ikanned.com_bundle.crt;
    ssl_certificate_key /etc/nginx/ikanned.com_nginx/ikanned.com.key;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers off;

    # 静态文件根目录
    root /var/www/nxc/dist;
    index index.html index.htm;

    # 代理 API 请求到后端服务
    location /api/ {
        proxy_pass https://localhost:8081;  # 后端服务在 HTTPS 上运行
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_read_timeout 90;
        proxy_connect_timeout 90;
        proxy_send_timeout 90;
    }

    # 代理 WebSocket 请求到后端
    location /api/ws/ {
        proxy_pass https://localhost:8081;  # WebSocket 服务在 HTTPS 上运行
        proxy_http_version 1.1;  # WebSocket 使用 HTTP/1.1
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_read_timeout 3600s;  # 设置长时间超时以支持 WebSocket 长连接
    }

    # 处理单页应用的路径，转发所有请求到 index.html
    location / {
        try_files $uri $uri/ /index.html;
 
```

### 啟動前端服務

```shell
sudo nginx -t #測試Nginx配置是否通過
# 如果出現錯誤可查看Nginx日誌
sudo journalctl -xeu nginx.service

sudo systemctl restart nginx.service #重啟Nginx服務
sudo systemctl enable nginx.service #開機時自動啟動
```


