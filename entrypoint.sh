#!/bin/sh

# 检查配置文件目录是否存在，如果不存在则创建
if [ ! -d "/app/config" ]; then
  mkdir -p /app/config
  echo "Created config directory"
fi

# 检查配置文件是否存在，如果不存在则创建默认配置
if [ ! -f "/app/config/config.yaml" ]; then
  echo "Config file not found, creating default config..."
  cat > /app/config/config.yaml << EOF
# 数据库配置(支持mysql和sqlite)
database:
  # 支持mysql或sqlite
  type: "sqlite"
  # MySQL情况下数据库地址
  host: "localhost"
  # MySQL情况下数据库端口
  port: 3306
  # MySQL情况下数据库用户名
  username: "root"
  # MySQL情况下数据库密码
  password: "HelloWorld"
  # MySQL情况下数据库名
  database: "emaction"
  # MySQL情况下数据库编码
  charset: "utf8mb4"
  # SQLite情况下数据库路径
  sqlite_path: "./emaction.db"

# 跨域
cors:
  allowOrigins: ["*"]
  allowMethods: ["GET", "PATCH", "OPTIONS"]
  allowHeaders: ["Content-Type", "Authorization"]
EOF
  echo "Default config file created"
fi

# 执行主应用程序
exec "$@"
