# go_diary

基于 Golang 的微服务架构日记后端项目。

## 项目结构

```
go_diary
├── core/                    # 核心服务
│   ├── db/migrations/       # 数据库迁移脚本
│   ├── src/                 # 源代码
│   │   ├── handlers/        # HTTP 处理器
│   │   ├── input_models/    # 输入模型
│   │   ├── models/          # 数据模型
│   │   ├── services/        # 业务逻辑
│   │   ├── view_models/     # 视图模型
│   │   └── main.go
│   ├── pkg/                 # 公共包
│   └── go.mod
├── core_run/                # 核心服务 Docker 配置
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

### 2. 配置环境变量

复制 `.env.example` 为 `.env` 并根据需要修改：

```bash
cp .env.example .env
```

### 3. 运行核心服务

```bash
cd core
go mod tidy
go run src/main.go
```

## API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /api/v1/diaries | 获取日记列表 |
| GET | /api/v1/diaries/:id | 获取单条日记 |
| POST | /api/v1/diaries | 创建日记 |
| PUT | /api/v1/diaries/:id | 更新日记 |
| DELETE | /api/v1/diaries/:id | 删除日记 |

## 技术栈

- Go 1.21+
- Gin Web 框架
- GORM ORM
- PostgreSQL 数据库
- Docker + Docker Compose
