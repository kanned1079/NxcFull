# 使用 Nginx 官方镜像作为基础镜像
FROM nginx:alpine

# 设置工作目录
WORKDIR /usr/share/nginx/html

# 将打包好的前端项目拷贝到 Nginx 的默认静态资源目录
COPY ./dist /usr/share/nginx/html

# 复制自定义的 Nginx 配置文件
COPY ./runtime/nginx.conf /etc/nginx/nginx.conf

# 暴露 8080 端口
EXPOSE 8080

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]