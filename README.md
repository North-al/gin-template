# gin-template

gin-template 是一个基于 `Go + Gin` 编写的gin模板，采用 **整洁架构（Clean Architecture）** 设计，支持 **多环境配置、数据库迁移、日志管理、中间件扩展** 等功能。

## 🚀 项目架构

本项目遵循 **整洁架构** 设计，核心思想如下：
- `internal/biz` 负责业务逻辑，包含 **领域对象（Entity）、仓储接口（Repository）、业务服务（Service）**。
- `internal/data` 负责 **数据库操作和数据持久化**，并提供数据库迁移和初始化方案。
- `internal/server` 负责 **HTTP 服务器及中间件管理**。
- `cmd/` 作为应用程序的 **入口**，用于加载配置、初始化依赖并启动服务器。

---

## 📂 目录结构
```plaintext
GO-GATEWAY
├── api/                # API 定义（如 Protobuf 文件）
├── cmd/                # 应用入口（Main 目录）
├── config/             # 配置文件
│   ├── dev/            # 开发环境配置
│   ├── prod/           # 生产环境配置
│   └── test/           # 测试环境配置
├── internal/           # 内部应用逻辑（遵循 Clean Architecture）
│   ├── biz/            # 业务逻辑层（核心领域）
│   │   ├── entity/     # 领域实体（业务对象）
│   │   ├── repository/ # 仓储接口（定义数据访问方法）
│   │   └── service/    # 业务服务（业务实现）
│   ├── data/           # 数据层（数据库访问）
│   │   ├── migrations/ # 数据库迁移 SQL
│   │   └── seed/       # 预设种子数据
│   ├── pkg/            # 通用工具（可复用模块）
│   │   ├── logger/     # 日志管理
│   │   └── utils/      # 工具类
│   ├── server/         # 服务器启动相关
│   │   └── middleware/ # Gin 中间件
├── docs/               # API 文档（Swagger 生成）
└── Makefile            # 构建 & 运行命令
```


## 🔧 安装 & 运行

### **1️⃣ 克隆项目**
```bash
git clone https://github.com/North-al/gin-template.git
cd gin-template
```

### 2️⃣ 配置环境
- 复制 `config/dev/config.yaml.example` 为 `config/dev/config.yaml` 并修改相应参数：
```bash
cp config/dev/config.yaml.example config/dev/config.yaml
```
### 3️⃣ 运行服务
使用 Makefile 快速启动：
```bash
make run
```
或者手动执行：
```bash
go run ./cmd
```

### 4️⃣ 数据库迁移
如果项目使用数据库（如 MySQL、PostgreSQL），可以执行数据库迁移：

```bash
make migrate
```
或：
```bash
go run internal/data/migrations/migrate.go
```

# 📖 开发指南
## 代码规范
- 使用 Go 1.20+ 版本。
- 遵循 Gin 作为 Web 框架，保持 RESTful API 设计。
- internal/biz 作为业务核心，不直接依赖数据库实现。
- internal/data 仅负责数据存取，实现 biz/repository 接口。
## 日志管理
- 统一使用 internal/pkg/logger 进行日志管理。
- 日志级别可以在 config 目录中配置。
