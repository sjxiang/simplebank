
# Simple Bank


<!-- 
# win 开发环境配置

wsl --install

$ sudo apt-get update
$ sudo apt-get install make
$ apt-get install -y git
$ sudo apt-get install protobuf-compiler
$ sudo apt-get automove protobuf-compiler


缝缝补补
https://github.com/docker/for-win/issues/13318


```powershell
cd C:\Users\相升杰\Downloads
Start-Process "Docker Desktop Installer.exe" -Verb RunAs -Wait -ArgumentList "install --installation-dir=C:\Docker\"
```

docker general 设置

-->


<!-- 
# Postgres  

1. 表结构设计
https://dbdiagram.io/d，拉个图倒是不赖


2. 使用 Docker + Postgres

镜像管理
    
    拉取镜像
    docker pull <image>:<tag>

容器运行
    创建并运行容器
	docker run \
    --rm \
	--name <container_name> \
    -e <environment_variable> \ 
	-p <host_ports:container_ports> \
    -d <image>:<tag>

容器管理
    查看正在运行的容器
    docker ps
        
    查看所有容器
    docker ps -a
        
    查看容器日志
    docker logs <container_name_or_id>

    进入容器的控制台
    docker exec -it <container_name_or_id> <command> [args]

    停止正在运行的容器
    docker stop <container_name_or_id>

    启动容器
    docker start <container_name_or_id>

    删除容器
    docker rm <container_name_or_id>


镜像管理
    删除镜像
    docker rmi <image_name_or_id> 



3. 数据库迁移

文档 https://github.com/golang-migrate/migrate


$ migrate create -ext sql -dir db/migration -seq $(name) 
# 扩展名，后缀
# 路径
# 有序，前缀
# 文件名，填补中间 

$ migrate -path db/migration -database "$(DB_URL)" -verbose down 1
# 回滚上一次 1

$ migrate -path db/migration -database "$(DB_URL)" -verbose up 1
# 升级 1




4. sqlc 代码生成
Create
insert new records to the database

Read
select to search for records int the database

Update
change some fields of the records in the database

Delete
remove records from the database

文档 https://docs.sqlc.dev/en/latest/index.html

sqlc init 
sqlc generate


5. 编写单元测试

6. 事务

e.g. 转账 
    
    - 创建交易记录 
    - 创建账户 1 流水记录
    - 创建账户 2 流水记录
    - 更新账户 1 余额
    - 更新账户 2 余额


7. 死锁

业务上禁用外键，约束太搞了

begin;
set transaction isolation level read committed;
show transaction isolation level;
select * from accounts;
update accounts set balance = balance - 10 where id = 1 returning *;
commit;

隔离级别，以及读异常（脏读、幻读 ...）


 -->



<!-- 

# web [Gin]

1. 设计 RESTful 风格的 HTTP API 
各种输入，参数绑定参数校验

2. viper 加载配置文件（.env）
3. 自定义参数验证器
4. 新增用户表
5. 处理数据库错误（约束）
6. 安全加密（bcrypt）
7. 身份认证，jwt、paseto
8. 中间件（拦截器）
9. 授权（不是你的账户，不能看；不是你的账户，不能转账）
10. 管理用户会话
无状态 -> 有状态

 -->




<!-- 

1. 打包镜像

docker build -t simplebank:latest .
docker images


2. docker 网络
docker inspect <container_name_or_id>
docker network ls
docker network inspect bridge

创建一个网络
    docker network create <bank-network>

容器加入网络
    docker network connect bank-network <container_name_or_id>


3. docker 卷
docker volume ls
docker volume rm <vloume_id>


docker-compose 
    容器内 IP 漂移，viper 还是要读取环境变量，覆盖 .env 文件，确保万无一失

 -->



<!-- # gRPC

1. 定义 proto 文件，并使用 protoc 生成 Go 代码
2. 实现 gRPC server，并测试，evans

    evans --host localhost --port 9090 -r repl
    show service
    call CreateUser 
    输入参数

3. gRPC gateway
https://github.com/grpc-ecosystem/grpc-gateway


4. metadata 传递信息
grpc 为啥这么多子包

5. 参数校验（错误处理）

6. 可选参数（db 操作）
sql.NullString

7. 认证鉴权
http gateway 统一模板，只能放 gRPC 操作 jwt，RBAC 鉴权


8. 拦截器，打日志
gRPC gateway 不会直接调用 gRPC client
zerolog


 -->


<!-- 


# 首选项 - 设置 - 搜索 protoc - vscode-proto3 扩展设置 - settings.json

```json
"protoc": {
    "options": [
        "--proto_path=proto",
        "--proto_path=idl"
    ]
},

相对于当前 workspace，
```


# 首选项 - 设置 - 搜索 go test flag - go 扩展 - settings.json

```json

"go.testFlags": [
    "-v",
    "-count=1"
]


-v 详情
-count=1 禁用测试缓存

```


# github 

.gitignore
追踪记录缓存，需要清除
git rm -r --cached . -->