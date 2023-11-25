
# Simple Bank



2. In the 2nd section, you will learn how to build a set of RESTful HTTP APIs using Gin - one of the most popular Golang frameworks for building web services. This includes everything from loading app configs, mocking DB for more robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.  

3. In the 3rd section, you will learn how to build your app with Docker and deploy it to a production Kubernetes cluster on AWS. The lectures are very detailed with a step-by-step guide, from how to build a minimal docker image, set up a free-tier AWS account, create a production database, store and retrieve production secrets, create a Kubernetes cluster with EKS, use GitHub Action to automatically build and deploy the image to the EKS cluster, buy a domain name and route the traffics to the service, secure the connection with HTTPS and auto-renew TLS certificate from Let's Encrypt.

4. In the 4th section, we will discuss several advanced backend topics such as managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP requests at the same time, embedding Swagger documentation as part of the backend service, partially updating a record using optional parameters, and writing structured logger HTTP middlewares and gRPC interceptors.

5. Then the 5th section will introduce you to asynchronous processing in Golang using background workers and Redis as its message queue. We'll also learn how to create and send emails to users via Gmail SMTP server. Along the way, we'll learn more about writing unit tests for our gRPC services that might involve mocking multiple dependencies at once.

6. The final section 6th concludes the course with lectures about how to improve the stability and security of the server. We'll keep updating dependency packages to the latest version, use Cookies to make the refresh token more secure, and learn how to gracefully shut down the server to protect the processing resources. As this part is still a work in progress, we will keep making and uploading new videos about new topics in the future. So please come back here to check them out from time to time.

This course is designed with a lot of details, so that everyone, even with very little programming experience can understand and do it by themselves. I strongly believe that after the course, you would be able to work much more confidently and effectively on your projects.







<!-- # win 开发环境配置

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


<!-- # Postgres  

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

    进入容器控制台，运行 cmd
    docker exec -it <container_name_or_id> <command> [args]

    停止正在运行的容器
    docker stop <container_name_or_id>

    启动容器
    docker start <container_name_or_id>

    删除容器
    docker rm <container_name_or_id>


3. 数据库迁移脚本

文档 https://github.com/golang-migrate/migrate

migrate create -ext sql -dir db/migration -seq $(name)
migrate up   
migrate down


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




<!-- # web [Gin]

1. 



 -->


<!-- 11: 使用 Gin 在 Go 中实现 RESTful HTTP API
12: 在 Go with Viper 中从文件和环境变量加载配置
13: 在 Go 中测试 HTTPAPI 并实现100% 覆盖率的模拟数据库
14: 使用自定义参数验证器实现转账 API
15: 在 PostgreSQL 中添加具有唯一和外键约束的用户表
16: 如何正确处理 Golang 的数据库错误
17: 如何安全地存储密码? 在 Go with Bcrypt 中散列密码！
18: 如何使用自定义 gomock 匹配器编写更强大的单元测试
19: 为什么 PASETO 比 JWT 更适合基于令牌的身份验证？
20: 如何在 Golang 创建和验证 JWT & PASETO 令牌
21: 在 Go 中实现返回 PASETO 或 JWT 访问令牌的登录用户 API
22: 使用金在 Golang 实现身份验证中间件和授权规则 -->


<!-- 第3节: 将应用程序部署到生产环境[ Kubernetes + AWS ]
23: 使用多阶段 Dockerfile 构建最小的 Golang Docker 映像
24: 如何使用 docker 网络连接两个独立的容器
25: 如何使用 wait-for.sh 编写 docker-撰写文件和控制服务启动订单
26: 如何创建免费层 AWS 帐户
27: 使用 Github 操作自动构建和推送 Docker 映像到 AWS ECR
28: 如何在 AWS RDS 上创建产品数据库
29: 使用 AWS 秘密管理器存储和检索生产秘密
30: Kubernetes 架构及如何在 AWS 上创建 EKS 集群
31: 如何使用 kubectl & k9s 连接到 AWS EKS 上的 kubernetes 集群
32: 如何在 AWS EKS 上将 Web 应用部署到 Kubernetes 集群
33: 注册一个域名并使用 Route53建立 A 记录
34: 如何使用 inress 将交通路由到 Kubernetes 的不同服务
35: 利用 Let’s Encrypt 在 Kubernetes 自动发布 TLS 证书
36: 使用 Github Action 自动部署到 Kubernetes -->



<!-- # session + gRPC

第4节: 高级后端主题[会话 + gRPC ]
37: 如何使用刷新令牌管理用户会话-Golang
38: 从 DBML 生成 DB 文档页面和模式 SQL 转储
39: gRPC 入门
40: 定义 gRPCAPI 并使用 Protobuf 生成 Go 代码
41: 如何运行 golang gRPC 服务器并调用其 API
42: 实现 gRPCAPI 在 Go 中创建和登录用户
43: gRPC 网关: 编写代码一次，同时服务 gRPC 和 HTTP 请求
44: 如何从 gRPC 元数据中提取信息
45: 从 Go 服务器自动生成和服务 Swagger 文档
46: 在 Golang 后端服务器的二进制文件中嵌入静态前端文件
47: 验证 gRPC 参数并发送人机友好响应
48: 在 Golang 代码中直接运行数据库迁移
49: 使用 SQLC 可空参数部分更新 DB 记录
50: 使用可选参数构建 gRPC 更新 API
51: 添加保护 gRPC API 的授权
52: 为 gRPC API 编写结构化日志
53: 如何在 Go 中编写 HTTP 日志记录器中间件 -->

<!-- 异步 [Asynq + Redis]

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
65: Go 和 Postgres 的配置 sqlc 版本2
66: 将 DB 驱动程序从 lib/pq 切换到 pgx
67: 如何使用 PGX 驱动程序处理数据库错误
68: Docker 撰写: 端口 + 卷映射
69: 如何在 Go 中安装和使用二进制包
70: 在 Go 中实现以角色为基础的存取控制(RBAC) -->

