## 简介
memo save & search, 备忘录 保存和搜索小项目. 自己运行记录一些事情

> 目前参考了七米大佬的项目, 简单搭建了 gin / gorm 的项目:
https://github.com/Q1mi/bubble

### docker 启动 mysql, 并挂在到 宿主机

#### 第一次构建启动
```bash
### 将mysql 的数据 挂载到本地
# TODO 注意 挂载宿主机路径需要修改
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/youName/vlumes/mysql:/var/lib/mysql -d mysql:8.0.19

### 查看mysql 容器id
docker ps
---
CONTAINER ID   IMAGE                  COMMAND                  CREATED         STATUS         PORTS                                NAMES
1b511e4641cd   mysql:8.0.19           "docker-entrypoint.s…"   5 minutes ago   Up 5 minutes   33060/tcp, 0.0.0.0:13306->3306/tcp   mysql8019

### 进入mysql 容器
docker exec -it 4a74c75be74b /bin/bash

### 进入容器之后, 登录mysql, 创建表
root@1b511e4641cd:/# mysql -uroot -p

# 执行数据库脚本
CREATE DATABASE IF NOT EXISTS mosch;
# 修改连接权限
alter user 'root'@'%'  IDENTIFIED WITH  mysql_native_password BY 'root1234';

### 构建应用镜像
docker build . -t mosch_app

### 运行项目
# 前台运行
docker run --link=mysql8019:mysql8019 -p 8888:8888 mosch_app
# 后台运行
docker run -d --link=mysql8019:mysql8019 -p 8888:8888 mosch_app
```

#### 第二次构建部署

**本地测试**
```bash
# TODO 注意 挂载宿主机路径需要修改, 最好和第一次一样, 防止数据丢失
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/youName/vlumes/mysql:/var/lib/mysql -d mysql:8.0.19

# 本地启动1
cd mosch
make run
```

**docker 部署测试**
```bash
### 启动 Mysql
# TODO 注意 挂载宿主机路径需要修改, 最好和第一次一样, 防止数据丢失
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/youName/vlumes/mysql:/var/lib/mysql -d mysql:8.0.19

### 构建应用镜像
docker build . -t mosch_app

### 运行项目
# 前台运行
docker run --link=mysql8019:mysql8019 -p 8888:8888 mosch_app
# 后台运行
docker run -d --link=mysql8019:mysql8019 -p 8888:8888 mosch_app
```
