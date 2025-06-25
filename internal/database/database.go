package database

import (
	"fmt"
	"sync"
	"time"

	"emaction/config"
	"emaction/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB 初始化数据库连接（使用单例模式）
func InitDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	var initErr error

	once.Do(func() {
		var dialector gorm.Dialector

		switch cfg.Type {
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				cfg.Username,
				cfg.Password,
				cfg.Host,
				cfg.Port,
				cfg.Database,
				cfg.Charset,
			)
			dialector = mysql.Open(dsn)
		case "sqlite":
			dialector = sqlite.Open(cfg.SQLitePath)
		default:
			initErr = fmt.Errorf("unsupported database type: %s", cfg.Type)
			return
		}

		var err error
		db, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // 设置为静默模式，生产环境建议使用
		})
		if err != nil {
			initErr = fmt.Errorf("failed to connect database: %w", err)
			return
		}

		// 配置连接池
		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("failed to get underlying sql.DB: %w", err)
			return
		}

		// 设置连接池参数以支持高并发
		sqlDB.SetMaxOpenConns(100)                  // 最大打开连接数
		sqlDB.SetMaxIdleConns(10)                   // 最大空闲连接数
		sqlDB.SetConnMaxLifetime(300 * time.Second) // 连接最大生存时间

		// 自动迁移模型
		err = db.AutoMigrate(&model.Reaction{})
		if err != nil {
			initErr = fmt.Errorf("failed to migrate database: %w", err)
			return
		}
	})

	if initErr != nil {
		return nil, initErr
	}

	return db, nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
