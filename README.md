# go_diary

基于 Golang 的微服务架构日记后端项目。

## 项目结构

```
go_diary
├── core/                    # 核心服务源代码
│   ├── db/                  # 数据库相关
│   │   └── migrations/      # 数据库迁移脚本
│   ├── src/                 # 源代码目录
│   │   ├── handlers/        # HTTP 处理器层
│   │   ├── input_models/    # 输入模型（请求参数）
│   │   ├── middleware/      # 中间件
│   │   ├── models/          # 数据模型（数据库映射）
│   │   ├── services/        # 业务逻辑层（含接口定义）
│   │   ├── view_models/     # 视图模型（响应数据）
│   ├── pkg/                 # 公共包
│   │   ├── config/          # 配置管理
│   │   ├── logger/          # 日志
│   │   └── response/        # 统一响应格式
│   └── go.mod
├── core_run/                # 核心服务运行配置
│   ├── src/                 # 应用入口
│   │   └── main.go          # 主程序入口
│   ├── testing/             # HTTP 测试文件
│   └── Dockerfile           # Docker 构建文件
├── core_unittest/           # 核心服务单元测试
├── core_apitest/            # 核心服务 API 测试
├── ai/                      # AI 总结服务（待实现）
├── docker-compose.yml       # Docker Compose 配置
└── .env.example             # 环境变量示例
```

## 快速开始

### 1. 启动数据库

```bash
docker-compose up -d
```

### 2. 配置环境变量（可选）

复制 `.env.example` 为 `.env` 并根据需要修改：

```bash
cp .env.example .env
```

### 3. 本地运行核心服务

```bash
cd core
go mod tidy
cd ../core_run
go run src/main.go
```

服务启动后运行在 `http://localhost:8080`

### 4. Docker 方式运行

```bash
# 构建镜像
docker-compose build

# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps
```

## API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /:userIdentity/api/v1/diaries | 获取指定用户的日记列表 |
| GET | /:userIdentity/api/v1/diaries/:id | 获取指定用户的单条日记 |
| POST | /:userIdentity/api/v1/diaries | 为指定用户创建日记 |
| PUT | /:userIdentity/api/v1/diaries/:id | 更新指定用户的日记 |
| DELETE | /:userIdentity/api/v1/diaries/:id | 删除指定用户的日记 |

**注意**：`:userIdentity` 只能是 `admin1` 或 `admin2`（模拟用户）。

### 请求示例

创建日记（admin1 用户）：

```bash
curl -X POST http://localhost:8080/admin1/api/v1/diaries \
  -H "Content-Type: application/json" \
  -d '{"title": "我的日记", "content": "今天是个好日子"}'
```

获取日记列表（admin1 用户）：

```bash
curl http://localhost:8080/admin1/api/v1/diaries
```

## 技术栈

- Go 1.23+
- Gin Web 框架
- GORM ORM
- PostgreSQL 数据库
- Docker + Docker Compose
- GitHub Actions
- UUID 生成库（github.com/google/uuid）

## 端口说明

| 服务 | 宿主机端口 | 容器端口 | 说明 |
|------|-----------|---------|------|
| go-diary-core | 8080 | 8080 | 核心服务 API |
| postgres | 5432 | 5432 | PostgreSQL 数据库 |

左侧为宿主机访问端口，右侧为容器内部监听端口。