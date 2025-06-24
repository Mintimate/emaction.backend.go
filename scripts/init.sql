-- 创建数据库
CREATE DATABASE IF NOT EXISTS emaction CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE emaction;

-- 删除现有表（如果存在）以避免字段类型冲突
DROP TABLE IF EXISTS reactions;

-- 创建Emoji点击统计表（与 GORM 模型兼容）
CREATE TABLE reactions (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    target_id VARCHAR(255) NOT NULL COMMENT '目标ID',
    reaction_name VARCHAR(50) NOT NULL COMMENT '反应名称',
    count INT NOT NULL DEFAULT 0 COMMENT '反应计数',
    created_at BIGINT NOT NULL COMMENT '创建时间（毫秒时间戳）',
    updated_at BIGINT NOT NULL COMMENT '更新时间（毫秒时间戳）',
    INDEX idx_target_reaction (target_id, reaction_name) COMMENT '目标ID和反应名称索引',
    UNIQUE KEY uk_target_reaction (target_id, reaction_name) COMMENT '目标ID和反应名称唯一键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Emoji点击统计表';
