package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"emaction/config"
	"emaction/internal/database"
	"emaction/internal/model"

	"gorm.io/gorm"
)

var (
	dbInitOnce sync.Once
	dbInstance *gorm.DB
)

// initDatabase 初始化数据库连接（单例模式）
func initDatabase() *gorm.DB {
	dbInitOnce.Do(func() {
		configPath := os.Getenv("APP_CONFIG_PATH")
		if configPath == "" {
			configPath = "./config" // 默认当前目录
		}
		// 加载配置
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		// 初始化数据库连接
		dbInstance, err = database.InitDB(cfg.Database)
		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}
	})
	return dbInstance
}

// GetReactions 获取指定目标的所有 Emoji
func GetReactions(targetID string) ([]model.ReactionResponse, error) {
	// 获取共享的数据库连接
	db := initDatabase()

	var reactions []model.Reaction
	err := db.Where("target_id = ?", targetID).Find(&reactions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query reactions: %w", err)
	}

	// 转换为响应格式
	results := make([]model.ReactionResponse, 0, len(reactions))
	for _, reaction := range reactions {
		results = append(results, model.ReactionResponse{
			ReactionName: reaction.ReactionName,
			Count:        reaction.Count,
		})
	}

	return results, nil
}

// UpdateReaction 更新 Emoji 计数
func UpdateReaction(targetID, reactionName string, diff int) error {
	// 标准化 diff 值
	if diff > 0 {
		diff = 1
	} else if diff < 0 {
		diff = -1
	} else {
		return errors.New("invalid diff value")
	}

	// 获取共享的数据库连接
	db := initDatabase()

	var reaction model.Reaction
	err := db.Where("target_id = ? AND reaction_name = ?", targetID, reactionName).First(&reaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 记录不存在，创建新记录
			newCount := max(0, diff)
			now := time.Now().UnixMilli()
			newReaction := model.Reaction{
				TargetID:     targetID,
				ReactionName: reactionName,
				Count:        newCount,
				CreatedAt:    now,
				UpdatedAt:    now,
			}
			result := db.Create(&newReaction)
			if result.Error != nil {
				return fmt.Errorf("failed to create reaction: %w", result.Error)
			}
		} else {
			return fmt.Errorf("failed to query reaction: %w", err)
		}
	} else {
		// 记录存在，更新计数
		newCount := max(0, reaction.Count+diff)
		now := time.Now().UnixMilli()
		result := db.Model(&reaction).Updates(map[string]interface{}{
			"count":      newCount,
			"updated_at": now,
		})
		if result.Error != nil {
			return fmt.Errorf("failed to update reaction: %w", result.Error)
		}
	}

	return nil
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
