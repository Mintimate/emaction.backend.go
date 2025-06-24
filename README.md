# Emaction Backend (Go 版本)

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Framework-Gin-green.svg)](https://gin-gonic.com)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Emaction 的 Go 语言版本后端服务，提供 emoji reaction 统计功能。这是原 JavaScript 版本 [emaction.backend](https://github.com/emaction/emaction.backend) 的 Go 重构版本，具有更高的性能和更好的类型安全性。

## 📋 功能特性

- ✨ **Emoji Reaction 统计** - 记录和统计各种 emoji reaction 的点击次数
- 🔥 **高性能** - 基于 Go 和 Gin 框架，提供优秀的并发性能
- 🗄️ **MySQL 数据库** - 使用 GORM 进行数据持久化
- 🌐 **CORS 支持** - 内置跨域资源共享配置
- 📊 **RESTful API** - 简洁易用的 API 接口设计
- ⚡ **快速部署** - 支持一键 Docker 部署

## 🚀 快速开始

### 前置要求

- Go 1.21 或更高版本
- MySQL 5.7+ 或 MariaDB 10.3+

### 安装步骤

1. **克隆项目**
   ```bash
   git clone https://github.com/your-username/emaction.backend.go.git
   cd emaction.backend.go
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **配置数据库**
   
   修改 `config/config.yaml` 文件：
   ```yaml
   database:
     host: "localhost"
     port: 3306
     username: "your_username"
     password: "your_password"
     database: "emaction"
     charset: "utf8mb4"
   ```

4. **初始化数据库**
   ```bash
   mysql -u your_username -p < scripts/init.sql
   ```

5. **运行服务**
   ```bash
   go run main.go
   ```

服务将在 `http://localhost:8080` 启动。

## 📚 API 文档

### 1. 获取 Reactions

获取特定 `targetId` 的所有 reaction 统计。

**接口地址：** `GET /reactions`

**请求参数：**

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| targetId | String | 是 | 目标 ID |

**请求示例：**
```bash
curl "http://localhost:8080/reactions?targetId=article-123"
```

**响应格式：**
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "reactionsGot": [
      {
        "reaction_name": "thumbs-up",
        "count": 42
      },
      {
        "reaction_name": "heart",
        "count": 28
      }
    ]
  }
}
```

### 2. 更新 Reaction

新增或更新一个 reaction 的计数。

**接口地址：** `PATCH /reaction`

**请求参数：**

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| targetId | String | 是 | 目标 ID |
| reaction_name | String | 是 | reaction 名称 |
| diff | Int | 是 | 数量变动，只接受 1 或 -1 |

**请求示例：**
```bash
curl -X PATCH "http://localhost:8080/reaction?targetId=article-123&reaction_name=thumbs-up&diff=1"
```

**响应格式：**
```json
{
  "code": 0,
  "msg": "success"
}
```

## 🏗️ 项目结构

```
.
├── main.go                 # 主程序入口
├── config/
│   ├── config.go          # 配置文件加载器
│   └── config.yaml        # 配置文件
├── internal/
│   ├── controller/        # 控制器层
│   │   └── reaction.go
│   ├── database/          # 数据库连接
│   │   └── database.go
│   ├── dto/              # 数据传输对象
│   │   └── reaction.go
│   ├── model/            # 数据模型
│   │   └── models.go
│   ├── service/          # 业务逻辑层
│   │   └── reaction.go
│   └── until/            # 工具函数
│       └── response.go
├── scripts/
│   └── init.sql          # 数据库初始化脚本
├── go.mod              # Go 模块文件
└── go.sum              # Go 依赖校验文件
```

## 🛠️ 技术栈

- **语言：** Go 1.21+
- **Web 框架：** [Gin](https://gin-gonic.com/)
- **ORM：** [GORM](https://gorm.io/)
- **数据库：** MySQL / MariaDB
- **配置：** YAML

## 📄 许可证

本项目使用 MIT 许可证。详情请查看 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- 原项目：[emaction/emaction.backend](https://github.com/emaction/emaction.backend)
- 感谢所有贡献者的支持
