## vlumes readme
数据挂载卷文件夹. 挂载内容非上传到git

### docker-mysql
```bash
# 将mysql 的数据 挂载到本地
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /Users/qicong/disk-quaint/go-project/meetlife/vlumes/mysql:/var/lib/mysql -d mysql:8.0.19

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
```
