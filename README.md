
# Simple Bank


4. In the 4th section, we will discuss several advanced backend topics such as managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP requests at the same time, embedding Swagger documentation as part of the backend service, partially updating a record using optional parameters, and writing structured logger HTTP middlewares and gRPC interceptors.

5. Then the 5th section will introduce you to asynchronous processing in Golang using background workers and Redis as its message queue. We'll also learn how to create and send emails to users via Gmail SMTP server. Along the way, we'll learn more about writing unit tests for our gRPC services that might involve mocking multiple dependencies at once.

6. The final section 6th concludes the course with lectures about how to improve the stability and security of the server. We'll keep updating dependency packages to the latest version, use Cookies to make the refresh token more secure, and learn how to gracefully shut down the server to protect the processing resources. As this part is still a work in progress, we will keep making and uploading new videos about new topics in the future. So please come back here to check them out from time to time.







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

 -->



<!-- 

# gRPC

1. 定义 proto 文件，并使用 protoc 生成 Go 代码
2. 实现 gRPC server，并测试，evans

    evans --host localhost --port 9090 -r repl
    show service
    call CreateUser 
    输入参数

3. grpc gateway
https://github.com/grpc-ecosystem/grpc-gateway


4. metadata 传递信息
grpc 为啥这么多子包


5. 参数校验（错误处理）

6. 可选参数（db 操作）
sql.NullString



50: 使用可选参数构建 gRPC 更新 API
51: 添加保护 gRPC API 的授权

 

52: 为 gRPC API 编写结构化日志

53: 如何在 Go 中编写 HTTP 日志记录器中间件

拦截器，打日志

 -->



<!-- 异步 [Asynq + Redis]

# Asynq + Redis


第5节: 使用后台工作者进行异步处理[ Asynq + Redis ]
54: 在 Go with Redis 和 Asynq 中实现后台 worker
55: 将异步 worker 集成到 Go web 服务器
56: 在 DB 事务中向 Redis 发送异步任务
57: 如何为 Go Asynq 工作者处理错误和打印日志
58: 稍微延迟一下可能对异步任务有好处
59: 如何通过 Gmail 发送电子邮件
60: 如何在 Go 中跳过测试，在 vscode 中配置测试标志
61: Go 中的电子邮件验证: 设计数据库并发送电子邮件
62: 在 Go 中实现电子邮件验证 API
63: 使用模拟 DB 和 Redis 对 gRPC API 进行单元测试
64: 如何测试需要身份验证的 gRPC API -->



<!-- 第6部分: 提高服务器的稳定性和安全性

1. 处理 db 错误，匹配 http 状态码（别啥都是 500）
2. docker-compose 启动不起来，那就多来几次


70: 在 Go 中实现以角色为基础的存取控制(RBAC) 

make new_migration name=add_role_to_users


-->

