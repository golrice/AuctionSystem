# AuctionSystem

[![Go Report Card](https://goreportcard.com/badge/github.com/golrice/AuctionSystem)](https://goreportcard.com/report/github.com/golrice/AuctionSystem)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

基于Go语言开发的拍卖系统，支持实时竞价、拍卖品管理等功能。

## 功能特性

- 拍卖品管理（创建、查询、更新、删除）
- 实时竞价系统
- WebSocket实时通信
- 缓存优化（Redis）
- 数据持久化（MySQL）
- 用户认证与授权
- API文档自动生成

## 技术栈

- **语言**: Go 1.24.4
- **Web框架**: Gin
- **数据库**: MySQL (GORM)
- **缓存**: Redis
- **实时通信**: WebSocket (gorilla/websocket)
- **配置管理**: Viper
- **API文档**: Swagger
- **测试框架**: testify
- **部署**: Docker

## 目录结构

```
AuctionSystem/
├── api/              # API路由定义
├── bootstrap/        # 应用启动引导
├── cmd/              # 应用入口
├── docs/             # 文档
├── internal/         # 核心业务逻辑
│   ├── auction/      # 拍卖模块
│   ├── auth/         # 认证模块
│   ├── common/       # 公共模块
│   ├── user/         # 用户模块
│   └── ...
├── pkg/              # 公共包
├── testutil/         # 测试工具
└── ...
```

## 快速开始

### 环境要求

- Go 1.24+
- MySQL 5.7+
- Redis 5.0+

### 配置

1. 复制配置文件：
   ```bash
   cp .env.example .env
   ```

2. 修改 `.env` 文件中的配置项：
   ```bash
   # 数据库配置
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=
   DB_PASS=
   DB_NAME=
   
   # Redis配置
   REDIS_HOST=localhost
   REDIS_PORT=6379
   ```

### 运行

```bash
# 安装依赖
go mod tidy

# 启动服务
go run cmd/main.go
```

服务默认运行在 `http://localhost:8080`。

### Docker部署

```bash
# 构建镜像
docker build -t auctionsystem .

# 运行容器
docker run -p 8080:8080 auctionsystem
```

## API文档

启动服务后，访问 `http://localhost:8080/swagger/index.html` 查看API文档。

## 贡献

欢迎提交Issue和Pull Request。

## 许可证

本项目采用MIT许可证，详情请见 [LICENSE](LICENSE) 文件。