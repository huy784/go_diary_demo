# go\_diary 项目设计文档

## 项目概述

一个基于 Golang 的微服务架构日记后端项目，用于学习 CI/CD 相关内容。

- **项目名称**: go\_diary
- **阶段**: Demo
- **目标**: 实现日记的增删改查基础功能，配合 GitHub Actions 学习 CI/CD

## 服务架构

### 核心服务 (go-diary-core)

- 提供日记 CRUD API
- 用户认证（简化版）
- 数据持久化

### AI 总结服务 (go-diary-ai-summary)

- 对日记内容进行 AI 总结（暂不实现）

### 数据库服务 (go-diary-sqldata)

- 通过 Docker Compose 启动 PostgreSQL
- 存储日记数据

## 技术栈

| 组件     | 选型                      |
| ------ | ----------------------- |
| 语言     | Go 1.21+                |
| Web 框架 | Gin                     |
| ORM    | GORM                    |
| 数据库    | PostgreSQL              |
| 容器化    | Docker + Docker Compose |
| CI/CD  | GitHub Actions          |

## 目录结构

```
go_diary
├── core/                    # 核心服务源代码
│   ├── db/                  # 数据库相关
│   │   └── migrations/      # 数据库迁移脚本
│   ├── src/                 # 源代码目录
│   │   ├── handlers/        # HTTP 处理器层
│   │   ├── input_models/    # 输入模型（请求参数）
│   │   ├── models/          # 数据模型（数据库映射）
│   │   ├── services/        # 业务逻辑层（含接口定义）
│   │   ├── view_models/     # 视图模型（响应数据）
│   │   └── main.go          # 应用入口
│   ├── pkg/                 # 公共包
│   │   ├── config/          # 配置管理
│   │   ├── logger/          # 日志
│   │   └── response/        # 统一响应格式
│   └── go.mod
├── core_run/                # 核心服务运行配置（Dockerfile 等）
├── core_unittest/           # 核心服务单元测试
├── core_apitest/            # 核心服务 API 测试
├── ai/                      # AI 总结服务源代码（暂不实现）
├── ai_run/                  # AI 服务运行配置
├── ai_unittest/             # AI 服务单元测试
├── ai_apitest/              # AI 服务 API 测试
├── docker-compose.yml       # Docker Compose 配置
└── .gitignore
```

## 核心服务模块职责

### Handlers 层

- 处理 HTTP 请求
- 参数绑定与校验
- 调用 Service 层
- 返回统一响应格式

### Input Models 层

- 定义请求参数结构
- 包含参数校验标签

### View Models 层

- 定义响应数据结构
- 按需组装返回数据

### Models 层

- 数据模型定义
- 数据库表映射

### Services 层

- 业务逻辑处理
- 定义接口（i\_\*.go）便于测试和替换
- 提供接口实现
- 依赖注入（通过工厂或构造函数）
- 调用数据访问

## 接口设计

### Service 接口定义

```go
go_diary
├── core/                    # 核心服务源代码
│   ├── db/                  # 数据库相关
│   │   └── migrations/      # 数据库迁移脚本
│   ├── src/                 # 源代码目录
│   │   ├── handlers/        # HTTP 处理器层
│   │   ├── input_models/    # 输入模型（请求参数）
│   │   ├── models/          # 数据模型（数据库映射）
│   │   ├── services/        # 业务逻辑层（含接口定义）
│   │   ├── view_models/     # 视图模型（响应数据）
│   │   └── main.go          # 应用入口
│   ├── pkg/                 # 公共包
│   │   ├── config/          # 配置管理
│   │   ├── logger/          # 日志
│   │   └── response/        # 统一响应格式
│   └── go.mod
├── core_run/                # 核心服务运行配置（Dockerfile 等）
├── core_unittest/           # 核心服务单元测试
├── core_apitest/            # 核心服务 API 测试
├── ai/                      # AI 总结服务源代码（暂不实现）
├── ai_run/                  # AI 服务运行配置
├── ai_unittest/             # AI 服务单元测试
├── ai_apitest/              # AI 服务 API 测试
├── docker-compose.yml       # Docker Compose 配置
└── .gitignore
```

## API 接口设计

### 日记接口

| 方法     | 路径                  | 描述     |
| ------ | ------------------- | ------ |
| GET    | /api/v1/diaries     | 获取日记列表 |
| GET    | /api/v1/diaries/:id | 获取单条日记 |
| POST   | /api/v1/diaries     | 创建日记   |
| PUT    | /api/v1/diaries/:id | 更新日记   |
| DELETE | /api/v1/diaries/:id | 删除日记   |

### 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

## 模型设计

### Input Models

```go
// CreateDiaryInput 创建日记请求
type CreateDiaryInput struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

// UpdateDiaryInput 更新日记请求
type UpdateDiaryInput struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}
```

### Models

```go
// Diary 日记数据模型
type Diary struct {
    ID        int64     `gorm:"primaryKey;autoIncrement"`
    Title     string    `gorm:"size:255;not null"`
    Content   string    `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
```

### View Models

```go
// DiaryViewModel 日记视图模型
type DiaryViewModel struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## 数据库设计

### diaries 表

| 字段          | 类型           | 说明   |
| ----------- | ------------ | ---- |
| id          | bigint       | 主键   |
| title       | varchar(255) | 标题   |
| content     | text         | 内容   |
| created\_at | timestamp    | 创建时间 |
| updated\_at | timestamp    | 更新时间 |

## 依赖注入与测试

- Service 层定义接口，便于 Mock 测试
- 通过构造函数注入依赖（Repository、DatabaseProvider 等）
- 单元测试可使用 Mock 实现替代真实依赖

## CI/CD 流程

通过 GitHub Actions 实现：

1. 代码提交触发单元测试
2. 构建 Docker 镜像
3. 运行 API 测试
4. 推送镜像（可选）

