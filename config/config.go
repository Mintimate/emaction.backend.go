package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构
type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
	CORS     CORSConfig     `yaml:"cors"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Type       string `yaml:"type"` // mysql or sqlite
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	Charset    string `yaml:"charset"`
	SQLitePath string `yaml:"sqlite_path"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// CORSConfig CORS配置
type CORSConfig struct {
	AllowOrigins []string `yaml:"allowOrigins"`
	AllowMethods []string `yaml:"allowMethods"`
	AllowHeaders []string `yaml:"allowHeaders"`
}

// LoadConfig 从指定目录加载配置
func LoadConfig(configPath string) (Config, error) {
	var config Config

	// 构建配置文件路径
	configFile := fmt.Sprintf("%s/config.yaml", configPath)

	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		return config, fmt.Errorf("读取配置文件失败: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return config, nil
}